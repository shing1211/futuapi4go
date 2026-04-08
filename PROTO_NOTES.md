# Protocol Implementation Notes / 協議實現說明

## Current Status / 當前狀態

### ✅ Working APIs / 正常工作的 API
- **InitConnect (1001)**: Connection handshake works correctly
- **KeepAlive (1002)**: Heartbeat mechanism functional
- **GetGlobalState (1004)**: Returns data, but structure parsing needs verification
- **Subscribe (3001)**: Subscription API works, returns RetType=0

### ⚠️ APIs with Proto Mismatch / 協議不匹配的 API
The following APIs return "解析protobuf协议失败" (Failed to parse protobuf protocol):
- GetBasicQot (2101)
- GetKL (2102)
- GetStaticInfo (2201)
- GetSubInfo (3002)
- GetOrderBook (2106) - Returns "未知的协议ID" (Unknown protocol ID)

### 🔍 Root Cause / 根本原因
The proto definition files may not exactly match the specific Futu OpenD build/version in use. 
The protocol structure returned by OpenD differs from what the proto files define.

## Verification Results / 驗證結果

### Subscribe Works / 訂閱 API 正常工作
```
Request: 23 bytes
Response: 6 bytes, RetType=0 (Success)
```
This confirms:
- ✅ TCP connection layer is correct
- ✅ Header encoding (44 bytes) is correct
- ✅ SHA1 hash calculation is correct
- ✅ Protobuf serialization works for some APIs

### GetGlobalState Response / GetGlobalState 響應
```
Raw bytes (14 bytes): 08 00 12 00 18 00 22 06 08 XX XX XX XX XX
```
The response structure is much smaller than the proto definition suggests, indicating
OpenD may be sending a different or simplified format.

## Next Steps / 下一步

1. Verify Futu OpenD version matches the proto files (10.2.6208)
2. If version mismatch, obtain correct proto files for your OpenD version
3. Consider using the official Futu Python SDK as reference for correct proto structure
4. Test with different OpenD configurations

## Working Test Command / 測試命令

```bash
# Connection and Subscribe test (should work):
go run ./cmd/examples/debug_test/

# The Subscribe API consistently returns success:
# ✅ Success (RetType=0)
```

## Reference / 參考
- Protocol Spec: https://openapi.futunn.com/futu-api-doc/en/ftapi/protocol.html
- Header Size: 44 bytes
- Endianness: Little-endian
- SHA1: Required for body hash
