package events

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// EventType lists standard system event topics.
type EventType string

const (
	EventCommitCreated          EventType = "CommitCreated"
	EventBuildStarted           EventType = "BuildStarted"
	EventBuildCompleted         EventType = "BuildCompleted"
	EventDeploymentStarted      EventType = "DeploymentStarted"
	EventDeploymentSucceeded    EventType = "DeploymentSucceeded"
	EventPerformanceDropped     EventType = "PerformanceDropped"
	EventIncidentDetected       EventType = "IncidentDetected"
	EventRootCauseFound         EventType = "RootCauseFound"
	EventRegressionTestGen      EventType = "RegressionTestGenerated"
	EventKnowledgeUpdated       EventType = "KnowledgeUpdated"
	EventDeploymentRolledBack   EventType = "DeploymentRolledBack"
)

// Event encapsulates a state change in GabraOS.
type Event struct {
	EventID          string                 `json:"eventId"`
	EventType        EventType              `json:"eventType"`
	Timestamp        time.Time              `json:"timestamp"`
	SourceArtifactID string                 `json:"sourceArtifactId"`
	Payload          map[string]interface{} `json:"payload"`
}

// EventHandler represents an async event subscriber callback.
type EventHandler func(ctx context.Context, evt Event) error

// EventBus interface defines pub/sub capabilities.
type EventBus interface {
	Publish(ctx context.Context, evt Event) error
	Subscribe(eventType EventType, handler EventHandler) string
	Unsubscribe(subscriptionID string)
}

// MemoryEventBus provides a thread-safe in-memory pub-sub event bus.
type MemoryEventBus struct {
	mu           sync.RWMutex
	subscribers  map[EventType]map[string]EventHandler
	eventHistory []Event
}

// NewMemoryEventBus instantiates a MemoryEventBus.
func NewMemoryEventBus() *MemoryEventBus {
	return &MemoryEventBus{
		subscribers:  make(map[EventType]map[string]EventHandler),
		eventHistory: make([]Event, 0),
	}
}

// Publish dispatches an event asynchronously to subscribers.
func (b *MemoryEventBus) Publish(ctx context.Context, evt Event) error {
	if evt.EventID == "" {
		evt.EventID = fmt.Sprintf("evt_%s", uuid.New().String()[:8])
	}
	if evt.Timestamp.IsZero() {
		evt.Timestamp = time.Now()
	}

	b.mu.Lock()
	b.eventHistory = append(b.eventHistory, evt)
	handlersCopy := make([]EventHandler, 0)
	if subs, exists := b.subscribers[evt.EventType]; exists {
		for _, h := range subs {
			handlersCopy = append(handlersCopy, h)
		}
	}
	b.mu.Unlock()

	for _, handler := range handlersCopy {
		go func(h EventHandler) {
			_ = h(ctx, evt)
		}(handler)
	}

	return nil
}

// Subscribe registers an event handler for a specific event topic.
func (b *MemoryEventBus) Subscribe(eventType EventType, handler EventHandler) string {
	b.mu.Lock()
	defer b.mu.Unlock()

	subID := fmt.Sprintf("sub_%s", uuid.New().String()[:8])
	if _, exists := b.subscribers[eventType]; !exists {
		b.subscribers[eventType] = make(map[string]EventHandler)
	}
	b.subscribers[eventType][subID] = handler
	return subID
}

// Unsubscribe removes a subscription.
func (b *MemoryEventBus) Unsubscribe(subscriptionID string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for evtType, subs := range b.subscribers {
		if _, exists := subs[subscriptionID]; exists {
			delete(b.subscribers[evtType], subscriptionID)
			break
		}
	}
}

// GetHistory returns dispatched events.
func (b *MemoryEventBus) GetHistory() []Event {
	b.mu.RLock()
	defer b.mu.RUnlock()
	copied := make([]Event, len(b.eventHistory))
	copy(copied, b.eventHistory)
	return copied
}
