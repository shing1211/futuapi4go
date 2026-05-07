# Futu OpenAPI Protocol Reference (Proto)

---

# Basic Functions

## Protocol ID Table

| Protocol ID | Protobuf File | Description |
|-------------|---------------|-------------|
| 1001 | InitConnect | Connection Initialization |
| 1002 | GetGlobalState | Get Global Status |
| 1003 | Notify | Event Notification Callback |
| 1004 | KeepAlive | Heartbeat Keep Alive |

## InitConnect

**Protocol ID**: 1001

InitConnect.proto:
```protobuf
message C2S
{
    required int32 clientVer = 1; //Client version number. clientVer = number before "." * 100 + number after ".". For example: clientVer = 1 * 100 + 1 = 101 for version 1.1 , and clientVer = 2 * 100 + 21 = 221 for version 2.21.
    required string clientID = 2; //The unique identifier for client, no specific generation rules, the client can guarantee the uniqueness
    optional bool recvNotify = 3; //Whether this connection receives notifications of market status or events that transaction needs to be re-unlocked. If True, OpenD will push these notifications to this connection, otherwise false means not receiving or pushing
    optional int32 packetEncAlgo = 4; //Specify the packet encryption algorithm, refer to the enumeration definition of Common.PacketEncAlgo
    optional int32 pushProtoFmt = 5; //Specify the push protocol format on this connection
}
```

### Connection Encryption

- If OpenD is configured with encryption, InitConnect must use RSA public key encryption to initialize the connection
- Subsequent protocols use AES encrypted communication with the random key returned by InitConnect
- RSA key of OpenD is 1024-bit, filling method is PKCS1, public key encryption, private key decryption

### Push Protocol Format

- The pushProtoFmt field in the initial connection protocol specifies the format of the data pushed on the connection
- Format types: Protobuf (0) or Json (1)

## KeepAlive

**Protocol ID**: 1004

```protobuf
syntax = "proto2";
package KeepAlive;
option java_package = "com.futu.openapi.pb";
option go_package = "github.com/futuopen/ftapi4go/pb/keepalive";

import "Common.proto";

message C2S
{
    required int64 time = 1; //Greenwich timestamp when the client sends the packet, in seconds
}

message S2C
{
    required int64 time = 1; //Greenwich timestamp when the server returned the packet, in seconds
}

message Request
{
    required C2S c2s = 1;
}

message Response
{
    required int32 retType = 1 [default = -400];
    optional string retMsg = 2;
    optional int32 errCode = 3;
    optional S2C s2c = 4;
}
```

- Send heartbeat keep alive protocol according to the interval returned by initialization protocol

---

# General Definitions

## Interface Result

**RetType:**
```protobuf
enum RetType
{
  RetType_Succeed = 0;   //Success
  RetType_Failed = -1;   //Failed
  RetType_TimeOut = -100; //Timeout
  RetType_Unknown = -400; //Unknown result
}
```

## Protocol Format

**ProtoFmt:**
```protobuf
enum ProtoFmt
{
  ProtoFmt_Protobuf = 0; //Google Protobuf
  ProtoFmt_Json = 1;     //Json
}
```

## Packet Encryption Algorithm

**PacketEncAlgo:**
```protobuf
enum PacketEncAlgo
{
  PacketEncAlgo_FTAES_ECB = 0; //AES ECB mode encryption modified by Futu
  PacketEncAlgo_None = -1;     //No encryption
  PacketEncAlgo_AES_ECB = 1;  //Standard AES ECB mode encryption
  PacketEncAlgo_AES_CBC = 2;   //Standard AES CBC mode encryption
}
```

## Program Status Type

**ProgramStatusType:**
```protobuf
enum ProgramStatusType
{
  ProgramStatusType_None = 0;
  ProgramStatusType_Loaded = 1;       //The necessary modules have been loaded
  ProgramStatusType_Loging = 2;        //Logging in
  ProgramStatusType_NeedPicVerifyCode = 3;  //Need a graphic verification code
  ProgramStatusType_NeedPhoneVerifyCode = 4; //Need phone verification code
  ProgramStatusType_LoginFailed = 5;    //Login failed
  ProgramStatusType_ForceUpdate = 6;   //The client version is too low
  ProgramStatusType_NessaryDataPreparing = 7; //Pulling necessary information
  ProgramStatusType_NessaryDataMissing = 8;   //Missing necessary information
  ProgramStatusType_UnAgreeDisclaimer = 9;     //Disclaimer is not agreed
  ProgramStatusType_Ready = 10;         //Ready to use
  ProgramStatusType_ForceLogout = 11;   //OpenD was forced to log out
  ProgramStatusType_DisclaimerPullFailed = 12; //Failed to get disclaimers
}
```

## OpenD Event Notification Type

**GtwEventType:**
```protobuf
enum GtwEventType
{
  GtwEventType_None = 0;
  GtwEventType_LocalCfgLoadFailed = 1;   //Load local configuration failed
  GtwEventType_APISvrRunFailed = 2;      //Server start failed
  GtwEventType_ForceUpdate = 3;          //The client version is too low
  GtwEventType_LoginFailed = 4;          //Login failed
  GtwEventType_UnAgreeDisclaimer = 5;    //Disclaimer is not agreed
  GtwEventType_NetCfgMissing = 6;        //Missing necessary network configuration
  GtwEventType_KickedOut = 7;            //Account is logged in elsewhere
  GtwEventType_LoginPwdChanged = 8;      //Login password has been changed
  GtwEventType_BanLogin = 9;             //User is forbidden to log in
  GtwEventType_NeedPicVerifyCode = 10;   //Need graphic verification code
  GtwEventType_NeedPhoneVerifyCode = 11; //Need phone verification code
  GtwEventType_AppDataNotExist = 12;     //Program's own data does not exist
  GtwEventType_NessaryDataMissing = 13;  //Missing necessary data
  GtwEventType_TradePwdChanged = 14;     //Trading password has been changed
  GtwEventType_EnableDeviceLock = 15;    //Enable device lock
}
```

## System Notification Type

**NotifyType:**
```protobuf
enum NotifyType
{
  NotifyType_None = 0;
  NotifyType_GtwEvent = 1;    //OpenD running event notification
  NotifyType_ProgramStatus = 2; //Program status
  NotifyType_ConnStatus = 3;    //Connection status
  NotifyType_QotRight = 4;      //Quotes authority
  NotifyType_APILevel = 5;      //User level (deprecated)
  NotifyType_APIQuota = 6;      //API Quota
}
```

## Package Unique Identifier

**PacketID:**
```protobuf
message PacketID
{
  required uint64 connID = 1;  //The current TCP connection ID, unique identifier returned by InitConnect
  required uint32 serialNo = 2; //Increment serial number
}
```

## Program Status

**ProgramStatus:**
```protobuf
message ProgramStatus
{
  required ProgramStatusType type = 1;     //Current status
  optional string strExtDesc = 2;          //Additional description
}
```

---

# Protocol Introduction

## Protocol Request Process

1. Create a connection
2. Initialize the connection
3. Request data or receive pushed data
4. Send KeepAlive protocol periodically to keep connected

## Protocol Design

The protocol data includes the protocol header and the protocol body.

### Protocol Header

```c
struct APIProtoHeader
{
    u8_t  szHeaderFlag[2];   // Packet header start flag, fixed as "FT"
    u32_t nProtoID;          // Protocol ID
    u8_t  nProtoFmtType;     // Protocol type, 0 for Protobuf, 1 for Json
    u8_t  nProtoVer;         // Protocol version, currently 0
    u32_t nSerialNo;         // Packet serial number, must be incremented
    u32_t nBodyLen;          // Body length
    u8_t  arrBodySHA1[20];   // SHA1 hash of packet body
    u8_t  arrReserved[8];    // Reserved 8-byte extension
};
```

| Field | Description |
|-------|-------------|
| szHeaderFlag | Packet header start flag, fixed as "FT" |
| nProtoID | Protocol ID |
| nProtoFmtType | Protocol type, 0 for Protobuf, 1 for Json |
| nProtoVer | Protocol version, used for iterative compatibility |
| nSerialNo | Packet serial number, for request/response matching |
| nBodyLen | Body length |
| arrBodySHA1 | SHA1 hash value for data integrity |
| arrReserved | Reserved 8-byte extension |

**Notes:**
- OpenD internal processing uses Protobuf, so recommend using Protobuf format
- Binary stream uses little-endian byte order

### Protocol Body

#### Protobuf Request Structure

```protobuf
message C2S
{
    required int64 req = 1;
}

message Request
{
    required C2S c2s = 1;
}
```

#### Protobuf Response Structure

```protobuf
message S2C
{
    required int64 data = 1;
}

message Response
{
    required int32 retType = 1 [default = -400]; //RetType, result of return
    optional string retMsg = 2;
    optional int32 errCode = 3;
    optional S2C s2c = 4;
}
```

| Field | Description |
|-------|-------------|
| c2s | Request parameter structure |
| req | Request parameters, actually defined according to the protocol |
| retType | Request result |
| retMsg | The reason for the failed request |
| errCode | The corresponding error code for failed request |
| s2c | Response data structure |
| data | Response data |

**Notes:**
- Enumeration value field definition uses signed integer
- Enumerations are generally defined in Common.proto, Qot_Common.proto, Trd_Common.proto
- Price, percentage and other data are transmitted in floating point type - need to be rounded (default 3 decimal places)

## Encrypted Communication Process

1. If OpenD is configured with encryption, InitConnect must use RSA public key encryption
2. Other subsequent protocols use AES encrypted communication with the random key returned by InitConnect
3. RSA encryption and decryption is only used for InitConnect requests

### RSA Encryption Rules

- For 1024-bit key, maximum length of single encryption string is (key_size)/8-11 = 100 bytes
- Divide plaintext into segments of up to 100 bytes for encryption

### AES Encryption

- The encryption key is returned by the InitConnect protocol
- Uses ECB encryption mode of AES by default
- Source data length must be integer multiple of 16 - needs padding with '0'

---

# Proto ID Reference

## Basic Protocols

| proto_id | Name | Description |
|----------|------|-------------|
| 1001 | InitConnect | Connection Initialization |
| 1002 | GetGlobalState | Get Global Status |
| 1003 | Notify | Event Notification Callback |
| 1004 | KeepAlive | Heartbeat Keep Alive |

## Quote Protocols (3xxx)

| proto_id | Name | Description |
|----------|------|-------------|
| 3001 | QOT_SUB | Subscribe |
| 3002 | QOT_REG_QOT_PUSH | Register Quote Push |
| 3003 | QOT_GET_SUB_INFO | Get Subscription Info |
| 3004 | QOT_GET_BASIC_QOT | Get Basic Quote |
| 3005 | QOT_UPDATE_BASIC_QOT | Update Basic Quote |
| 3006 | QOT_GET_KL | Get K-line |
| 3007 | QOT_UPDATE_KL | Update K-line |
| 3008 | QOT_GET_RT | Get Real-time Data |
| 3009 | QOT_UPDATE_RT | Update Real-time Data |
| 3010 | QOT_GET_TICKER | Get Ticker |
| 3011 | QOT_UPDATE_TICKER | Update Ticker |
| 3012 | QOT_GET_ORDER_BOOK | Get Order Book |
| 3013 | QOT_UPDATE_ORDER_BOOK | Update Order Book |
| 3014 | QOT_GET_BROKER | Get Broker |
| 3015 | QOT_UPDATE_BROKER | Update Broker |
| 3100 | QOT_GET_HISTORY_KL | Get History K-line |
| 3101 | QOT_GET_HISTORY_KL_POINTS | Get History K-line Points |
| 3200 | QOT_GET_TRADE_DATE | Get Trade Date |
| 3203 | QOT_GET_SECURITY_SNAPSHOT | Get Security Snapshot |
| 3204 | QOT_GET_PLATE_SET | Get Plate Set |
| 3205 | QOT_GET_PLATE_SECURITY | Get Plate Security |
| 3209 | QOT_GET_OPTION_CHAIN | Get Option Chain |

## Trade Protocols (2xxx)

| proto_id | Name | Description |
|----------|------|-------------|
| 2001 | TRD_GET_ACC_LIST | Get Account List |
| 2005 | TRD_UNLOCK_TRADE | Unlock Trade |
| 2101 | TRD_GET_FUNDS | Get Funds |
| 2102 | TRD_GET_POSITION_LIST | Get Position List |
| 2201 | TRD_GET_ORDER_LIST | Get Order List |
| 2202 | TRD_PLACE_ORDER | Place Order |
| 2205 | TRD_MODIFY_ORDER | Modify Order |
| 2211 | TRD_GET_ORDER_FILL_LIST | Get Order Fill List |
| 2221 | TRD_GET_HISTORY_ORDER_LIST | Get History Order List |
| 2222 | TRD_GET_HISTORY_ORDER_FILL_LIST | Get History Order Fill List |

---

*Source: https://openapi.futunn.com/futu-api-doc/en/*