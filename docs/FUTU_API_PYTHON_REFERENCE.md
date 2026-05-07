# Futu OpenAPI Documentation (Python)

---
# Introduction

## Overview

Futu API provides wide varieties of market data and trading services for your programmed trading to meet the needs of every developer's programmed trading and help your Quant dreams.

*Futu API* consists of *OpenD* and *API SDK*:
* *OpenD* is the gateway program of Futu API, running on your local computer or cloud server
* *API SDK* includes Python, Java, C#, C++, JavaScript

## Account Types
- **Futu ID** - your user account for Futubull
- **Universal Account** - trade across HK, US, A-shares markets

## Functionality
- **Quotation** - real-time quotes, K-lines, order book, ticker data
- **Trading** - place orders, cancel orders, modify orders, query positions

## Features
- Multi-platform: Windows, MacOS, CentOS, Ubuntu
- Multi-language: Python, Java, C#, C++, JavaScript
- Low latency: 0.0014s fastest order
- Free trading via API

---

# API Categories

## Quote API (Market Data)

### Market Snapshot
- `get_market_snapshot` - Get latest quotes for multiple stocks
- `get_stock_quote` - Get real-time quotes

### K-line Data  
- `get_cur_kline` - Get current K-line
- `get_history_kline` - Get historical K-line
- `request_history_kline` - Paginated historical K-line

### Order Book
- `get_order_book` - Get bid/ask order book

### Ticker
- `get_ticker` - Get tick-by-tick trades

### Time Sharing
- `get_stock_time_sharing` - Get intraday time-sharing data

### Market State
- `get_market_state` - Query market open/close status

### Capital Flow
- `get_capital_flow` - Get capital inflow/outflow
- `get_capital_distribution` - Get large/medium/small order distribution

### Plates
- `get_plate_set` - Get plate list
- `get_plate_stock` - Get stocks in a plate

### Stock Filter
- `stock_filter` - Filter stocks by criteria

### Subscription
- `subscribe` - Subscribe to real-time data
- `unsubscribe` - Cancel subscription

## Trade API (Trading)

### Order Operations
- `place_order` - Place an order
- `cancel_order` - Cancel an order
- `modify_order` - Modify an order
- `order_list_query` - Query order list
- `order_fill_query` - Query order fills

### Position & Funds
- `position_list_query` - Query positions
- `account_info_query` - Query account info

### Trading Environment
- `TrdEnv.SIMULATE` - Paper trading
- `TrdEnv.REAL` - Live trading

---

# OpenD

## Installation
- Download from Futu website
- Supports: Windows, MacOS, CentOS, Ubuntu

## Configuration
- IP: 127.0.0.1 (local) or 0.0.0.0 (all network)
- Port: 11111 (default)
- Log level: debug, info, error, fatal

## Connection
```python
from futu import OpenQuoteContext, OpenSecTradeContext

quote_ctx = OpenQuoteContext(host='127.0.0.1', port=11111)
trd_ctx = OpenSecTradeContext(host='127.0.0.1', port=11111)
```

---

# Rate Limits

## Subscription Quota
- Default: 100
- Higher tiers: 300, 1000, 2000 based on assets

## API Frequency Limits
- Varies per API (e.g., 60 requests per 30 seconds)

---

# Error Handling

## Common Error Codes
- `-1` - Invalid parameters
- `-100` - Timeout
- `-200` - Disconnected
- `-301` - Insufficient balance
- `-302` - Market closed

## Error Categories
- Connection errors
- Timeout errors  
- Trading errors
- Account errors