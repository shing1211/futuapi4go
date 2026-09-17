# Configuration Reference

## Client options — public `client` package

| Option | Default | Purpose |
|--------|---------|---------|
| `WithDialTimeout(d)` | `10s` | TCP dial timeout |
| `WithAPISetTimeout(d)` | `30s` | Per-request response timeout |
| `WithKeepAliveInterval(d)` | `30s` | Keep-alive ping interval |
| `WithReconnectInterval(d)` | `3s` | Base reconnect backoff |
| `WithMaxRetries(n)` | `3` | Max reconnect attempts |
| `WithLogLevel(level)` | info | `0`=info, `1`=warn, `2`=error, `3`=silent |
| `WithLogLevelName(name)` | — | `debug` \| `info` \| `warn` \| `error` \| `silent` |
| `WithSlogLogger(l)` | — | Route SDK logs through a `*slog.Logger` |
| `WithRSAPublicKey(pem)` | — | RSA public key PEM for `InitConnect` |
| `WithRSAPrivateKey(pem)` | — | RSA private key PEM (encryption mode) |
| `WithEncryption(enable)` | `false` | Enable FTAES session encryption |
| `WithOnStateChange(fn)` | — | Connection state-change callback |
| `WithEnvConfig()` | — | Read the environment variables below |

Client methods (chainable): `WithTradeEnv(env)`, `WithTradeMarket(mkt)`,
`WithContext(ctx)`.

```go
cli := client.New(
    client.WithAPISetTimeout(15*time.Second),
    client.WithSlogLogger(slog.Default()),
).WithTradeEnv(constant.TrdEnv_Simulate)
```

## Environment variables

| Variable | Read by | Meaning |
|----------|---------|---------|
| `FUTU_OPEND_ADDR` | `futuapi.NewClientFromEnv()` | OpenD `host:port` (default `127.0.0.1:11111`) |
| `FUTU_RSA_PUBLIC_KEY` | `client.WithEnvConfig()` | RSA public key PEM (file path or inline) |
| `FUTU_RSA_PRIVATE_KEY` | `client.WithEnvConfig()` | RSA private key PEM (file path or inline) |
| `FUTU_ENCRYPT` | `client.WithEnvConfig()` | `"1"`/`"true"` enables encryption |
| `FUTU_LOG_LEVEL` | `client.WithEnvConfig()` | `0`=info … `3`=silent |

> There is **no** `FUTU_TRD_ENV`. Set the trading environment with
> `WithTradeEnv(...)`; it defaults to simulate.

## Internal-only options

The internal client (`internal/client`) additionally provides `WithTLS`,
`WithRateLimiter`, `WithRetryConfig`, `WithBreaker`, `WithReconnectBackoff`,
`WithWSSecretKey`, `WithPushHandler`, and `WithLogger`. These are **not**
re-exported on the public `client` package — see
[ADR-0007](adr/0007-public-client-wrapper.md). Rate limiting, the circuit
breaker, and retry are opt-in and off by default.

## Defaults

Dial `10s` · API `30s` · keep-alive `30s` · reconnect base `3s` · max retries
`3`. Override any of them with the options listed above.
