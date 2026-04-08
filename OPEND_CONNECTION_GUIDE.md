# FutuAPI4Go - Futu OpenD Connection Guide

## Prerequisites Checklist

### Step 1: Verify Futu OpenD is Running
```bash
netstat -ano | findstr "11111"
```
Expected output:
```
TCP    127.0.0.1:11111        0.0.0.0:0              LISTENING       <PID>
```

### Step 2: Check Futu OpenD Status
1. Open Futu OpenD application
2. Verify you see:
   - "Connected" status
   - Logged into Futu account
   - Port 11111 configured
   - No RSA key required (or properly configured)

### Step 3: Common Issues

#### Issue A: OpenD Not Logged In
- **Symptom**: Port listening but no response
- **Fix**: Login to Futu account in OpenD first

#### Issue B: RSA Key Mismatch  
- **Symptom**: "packet too large" or connection rejected
- **Fix**: Either configure RSA in both OpenD and client, or disable completely

#### Issue C: Wrong Port
- **Symptom**: Connection refused
- **Fix**: Check OpenD settings for correct port

## Quick Test Commands

### Test 1: Basic Connectivity
```bash
# Test TCP connection
telnet 127.0.0.1 11111
# or
Test-NetConnection -ComputerName 127.0.0.1 -Port 11111
```

### Test 2: Run SDK Test
```bash
cd D:\gitee\futuapi4go
go run ./cmd/examples/test_real_opend/
```

### Test 3: Diagnostic Tool
```bash
go run ./cmd/examples/diag_packet/
```

## Current Status

- **Project Structure**: Complete (Go standard layout)
- **Protobuf Files**: Regenerated with correct paths
- **Examples**: 6 comprehensive examples created
- **Compilation**: Successful
- **OpenD Connection**: Needs verification

## Next Steps

1. **Verify OpenD is ready**:
   - Check OpenD shows "Connected" status
   - Confirm logged into account
   - Verify port 11111

2. **Test connection**:
   ```bash
   go run ./cmd/examples/test_real_opend/
   ```

3. **If still fails**:
   - Check OpenD logs
   - Try restarting OpenD
   - Verify no firewall blocking

## Connection Parameters

The SDK sends:
```
ProtoID: 1001 (InitConnect)
clientVer: 10100
clientID: "futuapi4go"  
packetEncAlgo: -1 (no encryption)
recvNotify: true
```

Expected response:
```
ProtoID: 1001
retType: 0 (success)
serverVer: <version>
connID: <connection ID>
connAESKey: <16 char key>
keepAliveInterval: 30
```
