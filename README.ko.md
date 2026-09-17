# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.18.1-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.10.7008-blue?style=flat-square" alt="Futu Proto Version">
  <a href="https://shing1211.github.io/futuapi4go/"><img src="https://img.shields.io/badge/Docs-GitHub%20Pages-97CAFF?style=flat-square&logo=github" alt="Docs"></a>
</p>

> **⚠️ 개발 진행 중**  
> 이 SDK는 활발히 개발 중입니다. 실제 Futu OpenD 인스턴스에 대해 동작하지만,
> 일부 proto 응답 필드는 아직 매핑되지 않았을 수 있습니다. API와 타입은 마이너
> 버전 간에 변경될 수 있습니다. 특정 필드에 의존하기 전에 사용 사례에 맞게 `client/types.go`를
> [Futu Proto Reference](https://openapi.futunn.com/mds/Futu-API-Doc-zh-Proto.md)와 대조하여 검토하십시오.

> **Go 네이티브. 타입 안전. 프로덕션 지원.** [Futu OpenAPI](https://www.futunn.com/en/overview)를 위한 가장 완전하고 사용하기 편한 Go SDK — 시세 데이터, 거래, 실시간 푸시. 모든 통신은 TCP 상의 Protocol Buffers를 통해 이루어집니다.

[English](./README.md) · [简体中文](./README.zh-Hans.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Español](./README.es.md)

> 이 문서는 영어 [README](./README.md)의 커뮤니티 번역입니다. **영어 버전이 기준입니다.**
> 동기화 / Last synced: c417534

- 모든 Futu OpenAPI 서비스를 포괄하는 184개의 protobuf 타입
- 자동 환경 설정으로 한 줄 연결 (`NewClientFromEnv`)
- 채널 또는 타입 지정 콜백을 통한 실시간 푸시
- 플루언트 API: `cli.Quote().GetBasicQot()`, `cli.Trade().PlaceOrder()`
- OpenTelemetry를 통한 분산 추적 + 메트릭 (`pkg/tracing/otel`로 선택적 활성화)
- 연결 상태 머신, 정상 종료(graceful shutdown), 자동 재연결
- 모든 API 호출에 내장된 속도 제한기, 서킷 브레이커, 재시도
- K-Line 데이터 캐시(LRU + TTL), 주문 사전 검증, 감사 로깅
- goreleaser를 통한 릴리스 자동화

## 목차

- [설치](#설치)
- [빠른 시작](#빠른-시작)
- [주요 기능](#주요-기능)
- [예제](#예제)
- [패키지 맵](#패키지-맵)
- [공통 API](#공통-api)
- [빌드 및 테스트](#빌드-및-테스트)
- [아키텍처](#아키텍처)
- [문제 해결](#문제-해결)
- [기여하기](#기여하기)
- [라이선스](#라이선스)

## 설치

```bash
go get github.com/shing1211/futuapi4go@v0.18.1
```

Go 1.26+ 및 실행 중인 [Futu OpenD](https://www.futunn.com/en/overview) 인스턴스가 필요합니다.

## 빠른 시작

```go
package main

import (
	"context"
	"fmt"
	"log"

	futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
)

func main() {
	// One-call connect (reads env: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, ...)
	cli, err := futuapi.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx := context.Background()
	quote, err := cli.GetQuote(ctx, "HK.00700")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: price=%.2f high=%.2f low=%.2f vol=%d\n",
		quote.Code, quote.CurPrice, quote.HighPrice, quote.LowPrice, quote.Volume)
}
```

> **참고:** 미국 주식은 `GetQuote`가 동작하기 전에 구독이 필요합니다. 홍콩 주식은 필요하지 않습니다.

## 주요 기능

### 실시간 푸시

폴링을 중단하세요 — 데이터가 도착하는 즉시 수신합니다. 두 가지 전달 모델:

```go
// Option 1: Channels (streaming)
ch := make(chan *push.UpdateBasicQot, 100)
stop, _ := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
defer stop()
for q := range ch {
	fmt.Printf("[%s] price=%.2f\n", q.Security.GetCode(), q.CurPrice)
}

// Option 2: Typed callbacks (chainable on client)
cli.OnQuote(func(q *push.UpdateBasicQot) {
	fmt.Printf("[%s] price=%.2f\n", q.Security.GetCode(), q.CurPrice)
}).OnOrder(func(o *push.TrdUpdateOrder) {
	fmt.Printf("Order %s: status=%d\n", o.GetOrderIDEx(), o.GetOrderStatus())
})
```

### 시세 데이터

```go
// One-shot
quote, _ := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
snapshots, _ := client.GetSecuritySnapshot(ctx, cli, securities)

// Auto-paginated historical K-lines
klines, _ := client.RequestHistoryKL(ctx, cli, constant.Market_HK, "00700",
	constant.KLType_K_Day, "2024-01-01", "2025-01-01")
```

### 거래

```go
accounts, _ := client.GetAccountList(ctx, cli)
accID := accounts[0].AccID

client.UnlockTrading(ctx, cli, "md5_password")
result, _ := client.PlaceOrder(ctx, cli, accID,
	constant.TrdMarket_HK, "00700",
	constant.TrdSide_Buy, constant.OrderType_Normal, 350.0, 100)

// Fluent order builder
order := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
	Buy("00700", 100).At(350.0).Build()
```

### 유틸리티

```go
// Circuit breaker
cb := breaker.New(breaker.WithThreshold(5), breaker.WithCooldown(30*time.Second))
result, _ := cb.Do(func() (interface{}, error) {
	return client.PlaceOrder(ctx, cli, accID, ...)
})

// Structured logging
l := futulogger.New(futulogger.WithLevel(futulogger.LevelDebug))
l.Info("connected", "addr", "127.0.0.1:11111")

// Code helpers
mkt, code := util.ParseCode("HK.00700")  // market=1, code="00700"
s := util.FormatCode(mkt, code)          // "HK.00700"
```

## 예제

실시간 푸시, 거래 워크플로, 과거 데이터, 전략 패턴을 포함한 모든 API 표면을 다루는 완전하고 실행 가능한 예제:

**[futuapi4go-demo →](https://github.com/shing1211/futuapi4go-demo)**

## 패키지 맵

| 패키지 | 목적 |
|---------|---------|
| `client` | 고수준 래퍼 — 권장 진입점 |
| `pkg/qot` | 시세 데이터: 호가, K-Line, 호가창, 틱 데이터... |
| `pkg/trd` | 거래: 주문, 포지션, 자금, 내역... |
| `pkg/sys` | 시스템: 전역 상태, 사용자 정보 |
| `pkg/push` | 푸시 알림 파서 |
| `pkg/push/chan` | 채널 기반 실시간 푸시 전달 |
| `pkg/breaker` | 서킷 브레이커 패턴 |
| `pkg/cache` | K-Line 데이터 캐시(LRU + TTL) |
| `pkg/logger` | 구조화된 레벨별 로깅 |
| `pkg/util` | 코드 파싱(`ParseCode`, `FormatCode`), 시장 헬퍼 |
| `pkg/constant` | `String()` 메서드를 갖춘 타입 지정 상수 |
| `pkg/degradation` | 연결 손실 시 우아한 성능 저하 |
| `pkg/futuapi` | 편의 재내보내기 — `NewClient()`, `NewClientFromEnv()` |
| `pkg/health` | OpenD 라이브니스/레디니스 프로브용 상태 확인 |
| `pkg/history` | 자동 페이지네이션 과거 K-line 다운로드 |
| `pkg/market` | 시장 시간, 거래 달력, 세션 감지 |
| `pkg/metrics` | 클라이언트 측 성능 메트릭 수집 |
| `pkg/option` | 옵션 체인 조회, 코드 파싱, 그릭스 헬퍼 |
| `pkg/pb/*` | 184개의 protobuf 타입 (v10.10.7008) |
| `pkg/ratelimit` | API 속도 제한 (protoID별 토큰 버킷) |
| `pkg/retry` | 지수 백오프를 지원하는 구성 가능한 재시도 |
| `pkg/trd/audit.go` | 거래 감사 로깅 |
| `pkg/trd/validation.go` | 주문 사전 검증 |
| `pkg/tracing` | 핵심 추적 인터페이스 (Tracer, Span, 기본 no-op) |
| `pkg/tracing/otel` | OpenTelemetry 기반 추적 어댑터 (선택적) |

## 공통 API

### 연결

```go
// Manual config
cli := client.New(
	client.WithDialTimeout(10*time.Second),
	client.WithAPISetTimeout(30*time.Second),
).WithTradeEnv(constant.TrdEnv_Simulate)

// From env vars: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, FUTU_ENCRYPT, FUTU_LOG_LEVEL
cli, _ := client.NewClientFromEnv()

cli.Connect("127.0.0.1:11111")
// cli.GetConnID(), cli.GetServerVer(), cli.IsEncrypt(), cli.GetLoginUserID()
// cli.CanSendProto(protoID)
```

### 시세 데이터

| 함수 | 설명 |
|---|---|
| `GetQuote(ctx, c, market, code)` | 실시간 호가 |
| `GetKLines(ctx, c, market, code, klType, num)` | 최신 K-line 봉 |
| `GetOrderBook(ctx, c, market, code, num)` | 매수/매도 호가 깊이 |
| `GetTicker(ctx, c, market, code, num)` | 틱별 체결 |
| `GetStaticInfo(ctx, c, market, code)` | 종목 이름, 유형, 최소 거래 단위 |
| `GetSecuritySnapshot(ctx, c, securities)` | 여러 종목의 전체 스냅샷 |
| `GetCapitalFlow(ctx, c, market, code)` | 자금 흐름 |
| `RequestHistoryKL(ctx, c, market, code, klType, start, end)` | 과거 K-line (자동 페이지네이션) |
| `RequestHistoryKLQuota(ctx, c)` | API 할당량 사용량 |

### 거래

| 함수 | 설명 |
|---|---|
| `GetAccountList(ctx, c)` | 모든 거래 계좌 |
| `UnlockTrading(ctx, c, pwdMD5)` | 거래 잠금 해제 |
| `GetFunds(ctx, c, accID)` | 계좌 자금 및 매수 여력 |
| `PlaceOrder(ctx, c, accID, market, code, side, orderType, price, qty)` | 주문 제출 |
| `ModifyOrder(ctx, c, accID, market, orderID, op, price, qty)` | 주문 수정 또는 취소 |
| `GetOrderList(ctx, c, accID)` | 활성 주문 |
| `GetPositionList(ctx, c, accID)` | 손익 포함 현재 포지션 |
| `GetHistoryOrderList(ctx, c, accID, market, start, end)` | 과거 주문 |
| `GetOrderFillList(ctx, c, accID)` | 주문 체결 |

### 구독

| 함수 | 설명 |
|---|---|
| `Subscribe(ctx, c, market, code, []SubType)` | 푸시 유형 구독 |
| `Unsubscribe(ctx, c, market, code, []SubType)` | 구독 해제 |
| `chanpkg.SubscribeQuote(ctx, cli, market, code, ch)` | 채널을 통한 호가 푸시 |
| `chanpkg.SubscribeKLine(ctx, cli, market, code, klType, ch)` | 채널을 통한 단일 K-line 푸시 |
| `chanpkg.SubscribeKLines(ctx, cli, market, code, []klTypes, ch)` | 필터를 지원하는 다중 K-line 푸시 |
| `chanpkg.SubscribeTicker(ctx, cli, market, code, ch)` | 채널을 통한 틱 푸시 |
| `chanpkg.SubscribeOrderBook(ctx, cli, market, code, ch)` | 채널을 통한 호가창 푸시 |

## 빌드 및 테스트

```bash
go build ./...      # Compile everything
go vet ./...        # Lint
go test -race ./... # Full suite with race detector
```

## 아키텍처

```
Application
  └── client/Client         (public wrappers)
       └── pkg/*            (qot, trd, sys — business logic)
            └── internal/client/Client   (connection, reconnect)
                 └── internal/client/Conn  (TCP I/O, packet framing)
                      └── Futu OpenD (TCP socket)
```

모든 통신은 TCP 상의 Protocol Buffers를 통해 이루어집니다. 전체 아키텍처 결정은 [DESIGN.md](DESIGN.md)를, 테스트에 사용되는 모의 OpenD 서버는 [internal/testutil/mock](internal/testutil/mock/)을 참조하십시오.

## 문제 해결

| 오류 | 가능한 원인 |
|-------|-------------|
| `connection refused` | OpenD가 실행 중이 아닙니다. `FUTU_OPEND_ADDR`를 확인하십시오. |
| no data from `GetQuote` (US stocks) | 미국 시장에서는 먼저 `Subscribe`를 호출해야 합니다. 홍콩은 필요하지 않습니다. |
| `The packet body SHA1 signature is incorrect` (very old OpenD) | OpenD를 v10.5+로 업그레이드하십시오. SDK는 OpenD가 허용하는 SHA1(ciphertext)을 사용합니다. |
| `解析protobuf协议失败` | 요청 본문에 필수 C2S 필드가 누락되었습니다. |
| `模拟交易不支持` | 모의 모드에서는 사용할 수 없는 기능입니다. `WithTradeEnv(TrdEnv_Real)`을 사용하십시오. |

## 기여하기

1. 저장소를 포크합니다.
2. 기능 브랜치를 만듭니다 (`git checkout -b feat/my-change`).
3. 기존 테스트가 모두 통과하는지 확인합니다: `go test -race ./...`
4. 새로운 기능에 대한 테스트를 추가합니다.
5. `go vet ./...`를 실행하고 경고를 수정합니다.
6. 풀 리퀘스트를 엽니다.

버전 내역은 [CHANGELOG.md](CHANGELOG.md)를, 로드맵은 [ENHANCEMENT_PLAN.md](docs/IMPLEMENTATION_COMPLETE.md)를 참조하십시오.

## 참고 항목

- [CHANGELOG](CHANGELOG.md) — 버전 내역 및 릴리스 노트
- [Version Map](docs/VERSION_MAP.md) — 각 SDK 릴리스가 포함하는 Futu OpenD 프로토콜 / proto 수 / `clientVer`
- [USAGE Guide](docs/USAGE.md) — 자세한 설정, 환경, 고급 패턴
- [DESIGN](DESIGN.md) — 아키텍처, 설계 결정, API 패턴
- [ENHANCEMENT_PLAN](docs/IMPLEMENTATION_COMPLETE.md) — 예정된 기능 및 로드맵
- [futuapi4go-demo](https://github.com/shing1211/futuapi4go-demo) — 모든 기능에 대한 실행 가능한 예제

## 라이선스

Apache License 2.0 — see [LICENSE](LICENSE).

> **Trading Disclaimer**: Trading financial instruments carries significant risk. Always test thoroughly in simulate mode before using real funds.
