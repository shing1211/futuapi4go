# WebSocket Transport (Experimental)

> **Status**: This package is experimental. The WebSocket transport is implemented but not yet wired into the main `Client`. API surface may change.

The `internal/ws` package provides a WebSocket-based transport for Futu OpenD as an alternative to the default TCP transport. It uses the same Futu binary protocol but over WebSocket (`ws://`) or secure WebSocket (`wss://`).

## Usage

```go
import "github.com/shing1211/futuapi4go/internal/ws"

// Connect via WebSocket
conn, err := ws.ConnectWS(ctx, "localhost:11111", 30*time.Second)
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

// Set push notification handler
conn.SetPushHandler(func(pkt *futuapi.Packet) {
    // Handle push notifications
})

// Write a packet
err = conn.WritePacket(protoID, serialNo, body)

// Read a response
resp, err := conn.ReadResponse(serialNo, 30*time.Second)
```

## Secure Connection (WSS)

```go
conn, err := ws.ConnectWSS(ctx, "localhost:11111", 30*time.Second)
```

## Limitations

- **Not integrated with main Client**: The main `client.Client` always uses the TCP transport. WebSocket is available for direct use with `pkg/qot`, `pkg/trd`, and `pkg/sys` packages by constructing requests manually.
- **Manual packet handling**: You must manage serial numbers and dispatching yourself when using `WSConn` directly.
- **No auto-reconnect**: Unlike the TCP client, the WebSocket transport does not implement automatic reconnection.

## Roadmap

- [ ] Wire WebSocket transport into `internal/client.Client` with `WithTransport("ws")` option
- [ ] Add auto-reconnection support
- [ ] TLS/SSL support via `wss://`
- [ ] Integration with the public `client.Client` API
