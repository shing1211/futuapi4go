# 📚 User Guide: Futu API for Golang (Updated)
## Introduction

This guide is designed for quantitative traders, bot developers, and end-users who wish to integrate `futuapi4go` into their applications. This SDK provides a robust, high-performance interface to the Futu OpenD protocol, allowing you to access market data, manage trading accounts, and run algorithmic strategies efficiently.

## 🚀 Getting Started (Quick Guide)
*(Please refer to README.md for the full Quick Start sequence)*

### Installation
```bash
go get github.com/shing1211/futuapi4go@latest
```

### Basic Usage Flow
The core workflow involves initialization, connection, API calls, and cleanup:

1.  **Initialize:** `cli := client.New()`
2.  **Connect:** `cli.Connect("host:port")` (Use environment variables for production!)
3.  **Execute:** Call wrapper functions like `GetQuote`, `PlaceOrder`.
4.  **Cleanup:** Always call `defer cli.Close()`.

## 🛡️ Operational Stability & Reliability Improvements (v1.0.0+)
The SDK has undergone significant architectural upgrades to improve reliability:

*   **Input Validation (New):** All public methods now validate their input parameters *before* making a network request. This prevents API calls from being sent with malformed data, leading to faster and clearer error feedback when inputs are incorrect.
*   **Architecture Overhaul:** The underlying system is more resilient. We have decoupled the network handling (`APIConnector` interface), which improves stability and allows for advanced testing capabilities.

## ⚙️ Production Deployment Guide (CRITICAL!)
To ensure maximum security, especially when running automated trading bots:

### **Security First: Configuration Management**
***DO NOT HARDCODE SECRETS!*** Never include private keys or passwords in your source code. Always load sensitive parameters from the operating system's environment variables (`os.Getenv()`). This is mandatory for secure deployment.

### Connection Settings
Use `cli.Connect()` with your OpenD address, which defaults to `127.0.0.1:11111`. You can override this by passing a custom string address.

## 📊 Key Functional Areas Overview
The library supports four main categories of functionality:

| Feature Area | Primary Functions | Description |
| :--- | :--- | :--- |
| **Market Data** (Read-Only) | `GetQuote`, `GetKLines`, `Subscribe` | Fetching real-time and historical market data. |
| **Trading Execution** | `PlaceOrder`, `ModifyOrder`, `UnlockTrade` | Full lifecycle management of orders and positions. |
| **System Info** | `GetUserInfo`, `GetCapitalFlow` | Retrieving user credentials, account balances, and metadata. |
| **Real-time Updates** | `SetQotPushHandler`, `SetTrdPushHandler` | Handling asynchronous data streams for live updates. |

## ❓ FAQ & Troubleshooting
*   **Connection refused:** Ensure Futu OpenD is running on the correct host and port (default: 127.0.0.1:11111).
*   **Timeout errors:** Check network latency or increase API timeouts via configuration options.

## 📋 Market Constants Quick Reference
| Category | Examples | Description |
| :--- | :--- | :--- |
| **Stock Markets** | `client.Market_HK_Security`, `client.Market_US_Security` | Identifies the exchange/market code. |
| **K-Line Types** | `client.KLType_Day`, `client.KLType_1Min` | Defines the time aggregation for historical data. |
| **Side (Trade Direction)** | `client.Side_Buy`, `client.Side_Sell` | Specifies buy or sell orders. |

*(For a complete list of all market constants, see the full API Reference.)*