// Package util provides shared utility functions for the Futu OpenD SDK,
// including stock code parsing, price conversion, JSON formatting,
// cryptographic operations, and date/time helpers.
//
// Key functions:
//   - ParseCode: Parse stock codes into market + symbol
//   - PriceToFloat/PriceToInt: Price conversion helpers
//   - ToJSON/FromJSON: JSON serialization for pb messages
//   - SHA1/SHA256: Cryptographic hash helpers (see sha1_test.go for quirks)
//   - DateRange: Generate date ranges for historical queries
//
// 简体中文:
// util 包提供富途 OpenD SDK 的共享工具函数，
// 包括股票代码解析、价格转换、JSON 格式化、加密操作和日期时间助手。
//
// 繁體中文:
// util 包提供富途 OpenD SDK 的共享工具函數，
// 包括股票代碼解析、價格轉換、JSON 格式化、加密操作和日期時間助手。
package util
