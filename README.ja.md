# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.19.2-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.10.7008-blue?style=flat-square" alt="Futu Proto Version">
  <a href="https://shing1211.github.io/futuapi4go/"><img src="https://img.shields.io/badge/Docs-GitHub%20Pages-97CAFF?style=flat-square&logo=github" alt="Docs"></a>
</p>

> **⚠️ 活発に開発中**  
> この SDK は活発に開発中です。実際の Futu OpenD インスタンスに対しては動作しますが、
> proto レスポンスの一部フィールドはまだマッピングされていない可能性があります。API と型はマイナー
> バージョン間で変更されることがあります。依存する前に、特定のユースケースについて `client/types.go` を
> [Futu Proto Reference](https://openapi.futunn.com/mds/Futu-API-Doc-zh-Proto.md) と照合して監査してください。

> **Go ネイティブ。型安全。本番運用対応。** [Futu OpenAPI](https://www.futunn.com/en/overview) 向けの最も完全で使いやすい Go SDK — 市場データ、取引、リアルタイムプッシュ。すべての通信は TCP 上の Protocol Buffers 経由です。

[English](./README.md) · [简体中文](./README.zh-Hans.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Español](./README.es.md)

> 本ファイルは英語版 [README](./README.md) のコミュニティ翻訳です。**英語版が正となります。**
> 同期 / Last synced: d5bc023

- すべての Futu OpenAPI サービスをカバーする 184 個の protobuf 型
- 環境変数設定によるワンライナー接続 (`NewClientFromEnv`)
- チャネルまたは型付きコールバックによるリアルタイムプッシュ
- 流暢な API: `cli.Quote().GetBasicQot()`、`cli.Trade().PlaceOrder()`
- OpenTelemetry による分散トレーシング + メトリクス (`pkg/tracing/otel` でオプトイン)
- 接続状態マシン、グレースフルシャットダウン、自動再接続
- すべての API 呼び出しに組み込まれたレートリミッター、サーキットブレーカー、リトライ
- K ラインデータキャッシュ (LRU + TTL)、注文の事前検証、監査ログ
- タグ push で GitHub リリースを自動作成（CHANGELOG からノート生成）

## 目次

- [インストール](#インストール)
- [クイックスタート](#クイックスタート)
- [主な機能](#主な機能)
- [サンプル](#サンプル)
- [パッケージ構成](#パッケージ構成)
- [よく使う API](#よく使う-api)
- [ビルドとテスト](#ビルドとテスト)
- [アーキテクチャ](#アーキテクチャ)
- [トラブルシューティング](#トラブルシューティング)
- [コントリビューション](#コントリビューション)
- [ライセンス](#ライセンス)

## インストール

```bash
go get github.com/shing1211/futuapi4go@v0.19.2
```

Go 1.26+ と、実行中の [Futu OpenD](https://www.futunn.com/en/overview) インスタンスが必要です。

## クイックスタート

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
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
	quote, err := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: price=%.2f high=%.2f low=%.2f vol=%d\n",
		quote.Symbol, quote.Price, quote.High, quote.Low, quote.Volume)
}
```

> **注:** 米国株は `GetQuote` が機能する前にサブスクライブが必要です。香港株は不要です。

## 主な機能

### リアルタイムプッシュ

ポーリングはやめましょう — データが届き次第受け取れます。2 つの配信モデルがあります:

```go
// Option 1: Channels (streaming)
ch := make(chan *push.UpdateBasicQot, 100)
stop, _ := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
defer stop()
for q := range ch {
	fmt.Printf("[%s] price=%.2f\n", q.Security.GetCode(), q.CurPrice)
}

// Option 2: Typed callbacks (chainable on client)
cli.OnQuote(func(q *client.PushQuote) error {
	fmt.Printf("[%s] price=%.2f\n", q.Code, q.CurPrice)
	return nil
}).OnOrder(func(o *client.PushOrderUpdate) error {
	fmt.Printf("Order %s: status=%d\n", o.OrderIDEx, o.OrderStatus)
	return nil
})
```

### 市場データ

```go
// One-shot
quote, _ := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
snapshots, _ := client.GetSecuritySnapshot(ctx, cli, securities)

// Auto-paginated historical K-lines
klines, _ := client.RequestHistoryKL(ctx, cli, constant.Market_HK, "00700",
	constant.KLType_K_Day, "2024-01-01", "2025-01-01")
```

### 取引

```go
accounts, _ := client.GetAccountList(ctx, cli)
accID := accounts[0].AccID

client.UnlockTrading(ctx, cli, "md5_password")
result, _ := client.PlaceOrder(ctx, cli, accID,
	constant.TrdMarket_HK, "00700",
	constant.TrdSide_Buy, constant.OrderType_Normal, 350.0, 100,
	constant.TrdSecMarket_HK)

// Fluent order builder (Build returns the request and an error)
req, err := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
	Buy("00700", 100).At(350.0).Build()
```

### ユーティリティ

```go
// Circuit breaker
cb := breaker.New(breaker.WithThreshold(5), breaker.WithCooldown(30*time.Second))
res, err := cb.Do(func() (interface{}, error) {
	return client.GetQuote(ctx, cli, constant.Market_HK, "00700")
})

// Structured logging
l := logger.New(logger.WithLevel(logger.LevelDebug))
l.Info("connected", "addr", "127.0.0.1:11111")

// Code helpers
mkt, code := util.ParseCode("HK.00700")  // market=1, code="00700"
s := util.FormatCode(mkt, code)          // "HK.00700"
```

## サンプル

すべての API サーフェスを網羅する完全で実行可能なサンプル — リアルタイムプッシュ、取引ワークフロー、履歴データ、戦略パターンを含みます:

**[futuapi4go-demo →](https://github.com/shing1211/futuapi4go-demo)**

## パッケージ構成

| パッケージ | 用途 |
|---------|---------|
| `client` | 高レベルラッパー — 推奨エントリポイント |
| `pkg/qot` | 市場データ: 気配値、K ライン、板情報、ティックデータ... |
| `pkg/trd` | 取引: 注文、ポジション、資金、履歴... |
| `pkg/sys` | システム: グローバル状態、ユーザー情報 |
| `pkg/push` | プッシュ通知パーサー |
| `pkg/push/chan` | チャネルベースのリアルタイムプッシュ配信 |
| `pkg/breaker` | サーキットブレーカーパターン |
| `pkg/cache` | K ラインデータキャッシュ (LRU + TTL) |
| `pkg/logger` | 構造化レベル付きログ |
| `pkg/util` | コード解析 (`ParseCode`、`FormatCode`)、市場ヘルパー |
| `pkg/constant` | `String()` メソッド付きの型付き定数 |
| `pkg/degradation` | 接続断時のグレースフルデグラデーション |
| `pkg/futuapi` | 便利な再エクスポート — `NewClient()`、`NewClientFromEnv()` |
| `pkg/health` | OpenD のライブネス/レディネスプローブ用ヘルスチェック |
| `pkg/history` | 自動ページングされた履歴 K ラインのダウンロード |
| `pkg/market` | 市場時間、取引カレンダー、セッション検出 |
| `pkg/metrics` | クライアント側のパフォーマンスメトリクス収集 |
| `pkg/option` | オプションチェーン照会、コード解析、ギリシャスヘルパー |
| `pkg/pb/*` | 184 個の protobuf 型 (v10.10.7008) |
| `pkg/ratelimit` | API レート制限 (protoID ごとのトークンバケット) |
| `pkg/retry` | 指数バックオフ付きの設定可能なリトライ |
| `pkg/trd/audit.go` | 取引監査ログ |
| `pkg/trd/validation.go` | 注文の事前検証 |
| `pkg/tracing` | コアトレーシングインターフェース (Tracer、Span、デフォルト no-op) |
| `pkg/tracing/otel` | OpenTelemetry ベースのトレーシングアダプター (オプトイン) |

## よく使う API

### 接続

```go
// Manual config
cli := client.New(
	client.WithDialTimeout(10*time.Second),
	client.WithAPISetTimeout(30*time.Second),
).WithTradeEnv(constant.TrdEnv_Simulate)

// From env vars: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, FUTU_ENCRYPT, FUTU_LOG_LEVEL
cli, _ := futuapi.NewClientFromEnv()

cli.Connect("127.0.0.1:11111")
// cli.GetConnID(), cli.GetServerVer(), cli.IsEncrypt(), cli.GetLoginUserID()
// cli.CanSendProto(protoID)
```

### 市場データ

| 関数 | 説明 |
|---|---|
| `GetQuote(ctx, c, market, code)` | リアルタイム気配値 |
| `GetKLines(ctx, c, market, code, klType, num)` | 最新の K ライン足 |
| `GetOrderBook(ctx, c, market, code, num)` | 気配値の深さ (板) |
| `GetTicker(ctx, c, market, code, num)` | ティックごとの約定 |
| `GetStaticInfo(ctx, c, market, code)` | 銘柄名、種類、売買単位 |
| `GetSecuritySnapshot(ctx, c, securities)` | 複数銘柄の完全なスナップショット |
| `GetCapitalFlow(ctx, c, market, code)` | 資金フロー |
| `RequestHistoryKL(ctx, c, market, code, klType, start, end)` | 履歴 K ライン (自動ページング) |
| `GetHistoryKLQuota(ctx, c)` | API クォータ使用量 |

### 取引

| 関数 | 説明 |
|---|---|
| `GetAccountList(ctx, c)` | すべての取引口座 |
| `UnlockTrading(ctx, c, pwdMD5)` | 取引ロック解除 |
| `GetFunds(ctx, c, accID)` | 口座資金と購買力 |
| `PlaceOrder(ctx, c, accID, market, code, side, orderType, price, qty)` | 注文発注 |
| `ModifyOrder(ctx, c, accID, market, orderID, op, price, qty)` | 注文の変更または取消 |
| `GetOrderList(ctx, c, accID)` | 有効な注文 |
| `GetPositionList(ctx, c, accID)` | 損益付きの現在のポジション |
| `GetHistoryOrderList(ctx, c, accID, market, start, end)` | 履歴注文 |
| `GetOrderFillList(ctx, c, accID)` | 注文の約定 |

### サブスクリプション

| 関数 | 説明 |
|---|---|
| `Subscribe(ctx, c, market, code, []constant.SubType)` | プッシュタイプを購読 |
| `Unsubscribe(ctx, c, market, code, []constant.SubType)` | 購読解除 |
| `chanpkg.SubscribeQuote(ctx, cli, market, code, ch)` | チャネル経由の気配値プッシュ |
| `chanpkg.SubscribeKLine(ctx, cli, market, code, klType, ch)` | チャネル経由の単一 K ラインプッシュ |
| `chanpkg.SubscribeKLines(ctx, cli, market, code, []klTypes, ch)` | フィルター付きの複数 K ラインプッシュ |
| `chanpkg.SubscribeTicker(ctx, cli, market, code, ch)` | チャネル経由のティックプッシュ |
| `chanpkg.SubscribeOrderBook(ctx, cli, market, code, ch)` | チャネル経由の板情報プッシュ |

## ビルドとテスト

```bash
go build ./...      # Compile everything
go vet ./...        # Lint
go test -race ./... # Full suite with race detector
```

## アーキテクチャ

```
Application
  └── client/Client         (public wrappers)
       └── pkg/*            (qot, trd, sys — business logic)
            └── internal/client/Client   (connection, reconnect)
                 └── internal/client/Conn  (TCP I/O, packet framing)
                      └── Futu OpenD (TCP socket)
```

すべての通信は TCP 上の Protocol Buffers 経由です。全体のアーキテクチャ決定については [DESIGN.md](DESIGN.md) を、テストで使用するモック OpenD サーバーについては [internal/testutil/mock](internal/testutil/mock/) を参照してください。

## トラブルシューティング

| エラー | 考えられる原因 |
|-------|-------------|
| `connection refused` | OpenD が起動していません。`FUTU_OPEND_ADDR` を確認してください。 |
| no data from `GetQuote` (US stocks) | 米国市場では先に `Subscribe` を呼ぶ必要があります。香港は不要です。 |
| `The packet body SHA1 signature is incorrect` (very old OpenD) | OpenD を v10.5+ にアップグレードしてください。SDK は SHA1(ciphertext) を使用しており、OpenD はこれを受け入れます。 |
| `解析protobuf协议失败` | リクエスト本文に必須の C2S フィールドがありません。 |
| `模拟交易不支持` | シミュレートモードでは利用できない機能です。`WithTradeEnv(TrdEnv_Real)` を使用してください。 |

## コントリビューション

1. リポジトリをフォークします。
2. フィーチャーブランチを作成します (`git checkout -b feat/my-change`)。
3. 既存のテストがすべて通ることを確認します: `go test -race ./...`
4. 新機能にはテストを追加します。
5. `go vet ./...` を実行し、警告を修正します。
6. プルリクエストを開きます。

バージョン履歴については [CHANGELOG.md](CHANGELOG.md) を、ロードマップについては [ENHANCEMENT_PLAN.md](docs/IMPLEMENTATION_COMPLETE.md) を参照してください。

## 関連情報

- [CHANGELOG](CHANGELOG.md) — バージョン履歴とリリースノート
- [Version Map](docs/VERSION_MAP.md) — 各 SDK リリースが携える Futu OpenD プロトコル / proto 数 / `clientVer`
- [USAGE Guide](docs/USAGE.md) — 詳細なセットアップ、環境、高度なパターン
- [DESIGN](DESIGN.md) — アーキテクチャ、設計判断、API パターン
- [ENHANCEMENT_PLAN](docs/IMPLEMENTATION_COMPLETE.md) — 今後の機能とロードマップ
- [futuapi4go-demo](https://github.com/shing1211/futuapi4go-demo) — すべての機能の実行可能なサンプル

## ライセンス

Apache License 2.0 — [LICENSE](LICENSE) を参照してください。

> **取引に関する免責事項**: 金融商品の取引には重大なリスクが伴います。実際の資金を使用する前に、必ずシミュレートモードで十分にテストしてください。
