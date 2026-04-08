# Protocol Implementation Notes

## Current Status

### Working APIs
- **InitConnect (1001)**: Connection handshake works correctly
- **KeepAlive (1002)**: Heartbeat mechanism functional
- **GetGlobalState (1004)**: Returns data, but structure parsing needs verification
- **Subscribe (3001)**: Subscription API works, returns RetType=0

### APIs with Proto Mismatch
The following APIs return "Failed to parse protobuf protocol":
- GetBasicQot (2101)
- GetKL (2102)
- GetStaticInfo (2201)
- GetSubInfo (3002)
- GetOrderBook (2106) - Returns "Unknown protocol ID"

### Root Cause
The proto definition files may not exactly match the specific Futu OpenD build/version in use. 
The protocol structure returned by OpenD differs from what the proto files define.

## Verification Results

### Subscribe Works
```
Request: 23 bytes
Response: 6 bytes, RetType=0 (Success)
```
This confirms:
- TCP connection layer is correct
- Header encoding (44 bytes) is correct
- SHA1 hash calculation is correct
- Protobuf serialization works for some APIs

### GetGlobalState Response
```
Raw bytes (14 bytes): 08 00 12 00 18 00 22 06 08 XX XX XX XX XX
```
The response structure is much smaller than the proto definition suggests, indicating
OpenD may be sending a different or simplified format.

## Next Steps

1. Verify Futu OpenD version matches the proto files (10.2.6208)
2. If version mismatch, obtain correct proto files for your OpenD version
3. Consider using the official Futu Python SDK as reference for correct proto structure
4. Test with different OpenD configurations

## Working Test Command

```bash
# Connection and Subscribe test (should work):
go run ./cmd/examples/debug_test/

# The Subscribe API consistently returns success:
# Success (RetType=0)
```

## Reference
- Protocol Spec: https://openapi.futunn.com/futu-api-doc/en/ftapi/protocol.html
- Header Size: 44 bytes
- Endianness: Little-endian
- SHA1: Required for body hash
