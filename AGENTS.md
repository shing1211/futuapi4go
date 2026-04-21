# AGENTS.md — futuapi4go Operational Guide

## 🧠 High-Signal Architectural Constraints (READ FIRST)
*   **Language:** Go (Golang).
*   **Architecture Flow:** The library uses a layered structure: `Application` $\rightarrow$ `client/Client` (Public API Wrapper) $\rightarrow$ `pkg/*` (Business Logic) $\rightarrow$ `internal/client/Conn` (Raw I/O) $\rightarrow$ `Futu OpenD Gateway`.
*   **Protobuf Reliance:** All communication relies on Protobuf; the SDK does not use JSON by default.
*   **Core Libraries:** The connection pool (`internal/client/pool.go`) and low-level connection logic (`internal/client/conn.go`) are critical for stability.

## ⚙️ Developer Workflow & Execution Commands

### Building & Testing
When making changes, always validate against the full test suite:
```bash
go build ./...  # Compiles all packages and examples
go test ./...   # Runs unit tests across the entire codebase (critical verification step)
```

### Code Review Focus Areas
*   **Concurrency:** Scrutinize mutex usage in connection handlers (`internal/client/conn.go`) for race condition prevention, especially when managing `connected` or `connActive` state flags.
*   **Pool Safety:** Verify that resource acquisition and release logic in the connection pool is thread-safe under heavy load.
*   **Context Usage (New Standard):** All public API methods **must** be updated to accept a `context.Context`. This is essential for production readiness to handle operation cancellation gracefully.

## ⚠️ Gotchas & Production Constraints

*   **Pool Management:** Always ensure an acquired client from the pool (`Pool.Get()`) is explicitly returned via `Pool.Put(client)` when finished, preventing connection starvation.
*   **Asynchronous Updates (Real-Time):** Real-time data arrives via push notifications. You must set a handler using `cli.SetPushHandler()` to receive these asynchronous packets after connecting.
*   **API Timeouts:** Utilize the granular API request timeout mechanism rather than relying solely on global connection timeouts for reliable operation of long-running queries.

## 🔗 Key Files & Entry Points
- **Main Client:** `client/client.go` — The primary public interface and wrapper.
- **Connection Pool:** `internal/client/pool.go` — Resource management logic.
- **Low-Level I/O:** `internal/client/conn.go` — Raw packet handling, serialization, and reading from the socket.
- **Examples:** `cmd/examples/*` - Provides concrete usage patterns for trading, market data, and push subscriptions.