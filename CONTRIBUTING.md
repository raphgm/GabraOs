# Contributing to GabraOS

First off, thank you for considering contributing to GabraOS! GabraOS is an open-source project building the **open standard for Autonomous Engineering**.

---

## Code of Conduct

We expect all contributors to adhere to an open, inclusive, and respectful community standard.

---

## How to Contribute

### 1. Architectural Proposals (RFCs & ADRs)
Before writing major code features, please submit a Request for Comments (RFC) or Architecture Decision Record (ADR) in `docs/architecture/rfc/`.

### 2. Developing Locally
- Fork the repository and clone locally.
- Run `go test -v ./...` to verify all Go packages compile and pass unit tests.
- Run `./bin/gabra status` to verify CLI command execution.

### 3. Pull Request Guidelines
- Follow standard Go formatting (`gofmt -w .`).
- Ensure all public functions have clear docstrings.
- Add test coverage for new event types, artifacts, or agent lifecycle stages.

---

Thank you for shaping the future of Autonomous Engineering!
