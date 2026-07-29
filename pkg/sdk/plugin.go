package sdk

import (
	"context"
	"fmt"

	"github.com/gabraos/gabraos/pkg/artifacts"
)

// PluginMetadata defines metadata for community and core plugins.
type PluginMetadata struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Supported   []string `json:"supported"`
}

// Plugin interface defines the contract for GabraOS extensions.
type Plugin interface {
	Metadata() PluginMetadata
	Init(ctx context.Context, config map[string]interface{}) error
	CollectArtifacts(ctx context.Context) ([]*artifacts.Artifact, error)
	ExecuteAction(ctx context.Context, action string, params map[string]interface{}) (map[string]interface{}, error)
}

// PluginRegistry manages loaded plugins.
type PluginRegistry struct {
	plugins map[string]Plugin
}

// NewPluginRegistry initializes a plugin registry.
func NewPluginRegistry() *PluginRegistry {
	return &PluginRegistry{
		plugins: make(map[string]Plugin),
	}
}

// Register registers a plugin instance into the ecosystem registry.
func (r *PluginRegistry) Register(p Plugin) error {
	meta := p.Metadata()
	if meta.Name == "" {
		return fmt.Errorf("plugin name cannot be empty")
	}
	r.plugins[meta.Name] = p
	return nil
}

// ListPlugins returns metadata for all registered ecosystem plugins.
func (r *PluginRegistry) ListPlugins() []PluginMetadata {
	list := make([]PluginMetadata, 0, len(r.plugins))
	for _, p := range r.plugins {
		list = append(list, p.Metadata())
	}
	return list
}
