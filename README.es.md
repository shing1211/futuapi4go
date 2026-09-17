# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.18.1-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.10.7008-blue?style=flat-square" alt="Futu Proto Version">
  <a href="https://shing1211.github.io/futuapi4go/"><img src="https://img.shields.io/badge/Docs-GitHub%20Pages-97CAFF?style=flat-square&logo=github" alt="Docs"></a>
</p>

> **⚠️ En desarrollo activo**  
> Este SDK está en desarrollo activo. Aunque funciona contra instancias reales de Futu OpenD,
> algunos campos de respuesta proto aún pueden estar sin mapear. Las APIs y los tipos pueden cambiar entre
> versiones menores. Audita `client/types.go` contra la [Referencia de Futu Proto](https://openapi.futunn.com/mds/Futu-API-Doc-zh-Proto.md)
> para tu caso de uso específico antes de depender de cualquier campo.

> **Nativo de Go. Seguro en tipos. Listo para producción.** El SDK de Go más completo y ergonómico para [Futu OpenAPI](https://www.futunn.com/en/overview): datos de mercado, trading y push en tiempo real. Toda la comunicación mediante Protocol Buffers sobre TCP.

[English](./README.md) · [简体中文](./README.zh-Hans.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Español](./README.es.md)

> Este archivo es una traducción comunitaria del [README](./README.md) en inglés. **La versión en inglés es la autoritativa.**
> Sincronizado / Last synced: c417534

- 184 tipos protobuf que cubren todos los servicios de Futu OpenAPI
- Conexión en una sola línea con configuración automática desde variables de entorno (`NewClientFromEnv`)
- Push en tiempo real mediante canales o callbacks tipados
- API fluida: `cli.Quote().GetBasicQot()`, `cli.Trade().PlaceOrder()`
- Trazado distribuido + métricas con OpenTelemetry (opcional mediante `pkg/tracing/otel`)
- Máquina de estados de conexión, cierre ordenado y reconexión automática
- Limitador de tasa, cortacircuitos y reintentos integrados en cada llamada a la API
- Caché de datos K-Line (LRU + TTL), validación previa de órdenes y registro de auditoría
- Automatización de versiones con goreleaser

## Tabla de contenidos

- [Instalación](#instalación)
- [Inicio rápido](#inicio-rápido)
- [Características principales](#características-principales)
- [Ejemplos](#ejemplos)
- [Mapa de paquetes](#mapa-de-paquetes)
- [APIs comunes](#apis-comunes)
- [Compilación y pruebas](#compilación-y-pruebas)
- [Arquitectura](#arquitectura)
- [Solución de problemas](#solución-de-problemas)
- [Contribuir](#contribuir)
- [Licencia](#licencia)

## Instalación

```bash
go get github.com/shing1211/futuapi4go@v0.18.1
```

Requiere Go 1.26+ y una instancia de [Futu OpenD](https://www.futunn.com/en/overview) en ejecución.

## Inicio rápido

```go
package main

import (
	"context"
	"fmt"
	"log"

	futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
)

func main() {
	// Conexión en una sola llamada (lee variables de entorno: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, ...)
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

> **Nota:** Las acciones de EE. UU. requieren suscribirse antes de que `GetQuote` funcione. Las acciones de Hong Kong no.

## Características principales

### Push en tiempo real

Deja de sondear: recibe los datos a medida que llegan. Dos modelos de entrega:

```go
// Opción 1: Canales (streaming)
ch := make(chan *push.UpdateBasicQot, 100)
stop, _ := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
defer stop()
for q := range ch {
	fmt.Printf("[%s] price=%.2f\n", q.Security.GetCode(), q.CurPrice)
}

// Opción 2: Callbacks tipados (encadenables en el cliente)
cli.OnQuote(func(q *push.UpdateBasicQot) {
	fmt.Printf("[%s] price=%.2f\n", q.Security.GetCode(), q.CurPrice)
}).OnOrder(func(o *push.TrdUpdateOrder) {
	fmt.Printf("Order %s: status=%d\n", o.GetOrderIDEx(), o.GetOrderStatus())
})
```

### Datos de mercado

```go
// Consulta única
quote, _ := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
snapshots, _ := client.GetSecuritySnapshot(ctx, cli, securities)

// K-lines históricas con paginación automática
klines, _ := client.RequestHistoryKL(ctx, cli, constant.Market_HK, "00700",
	constant.KLType_K_Day, "2024-01-01", "2025-01-01")
```

### Trading

```go
accounts, _ := client.GetAccountList(ctx, cli)
accID := accounts[0].AccID

client.UnlockTrading(ctx, cli, "md5_password")
result, _ := client.PlaceOrder(ctx, cli, accID,
	constant.TrdMarket_HK, "00700",
	constant.TrdSide_Buy, constant.OrderType_Normal, 350.0, 100)

// Constructor de órdenes fluido
order := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
	Buy("00700", 100).At(350.0).Build()
```

### Utilidades

```go
// Cortacircuitos
cb := breaker.New(breaker.WithThreshold(5), breaker.WithCooldown(30*time.Second))
result, _ := cb.Do(func() (interface{}, error) {
	return client.PlaceOrder(ctx, cli, accID, ...)
})

// Registro estructurado
l := futulogger.New(futulogger.WithLevel(futulogger.LevelDebug))
l.Info("connected", "addr", "127.0.0.1:11111")

// Utilidades de códigos
mkt, code := util.ParseCode("HK.00700")  // market=1, code="00700"
s := util.FormatCode(mkt, code)          // "HK.00700"
```

## Ejemplos

Para ver ejemplos completos y ejecutables que cubren toda la superficie de la API —incluidos push en tiempo real, flujos de trading, datos históricos y patrones de estrategia—:

**[futuapi4go-demo →](https://github.com/shing1211/futuapi4go-demo)**

## Mapa de paquetes

| Paquete | Propósito |
|---------|---------|
| `client` | Envoltorios de alto nivel: punto de entrada recomendado |
| `pkg/qot` | Datos de mercado: cotizaciones, K-lines, libro de órdenes, datos de tick... |
| `pkg/trd` | Trading: órdenes, posiciones, fondos, historial... |
| `pkg/sys` | Sistema: estado global, información del usuario |
| `pkg/push` | Analizadores de notificaciones push |
| `pkg/push/chan` | Entrega de push en tiempo real basada en canales |
| `pkg/breaker` | Patrón de cortacircuitos |
| `pkg/cache` | Caché de datos K-Line (LRU + TTL) |
| `pkg/logger` | Registro estructurado por niveles |
| `pkg/util` | Análisis de códigos (`ParseCode`, `FormatCode`), utilidades de mercado |
| `pkg/constant` | Constantes tipadas con métodos `String()` |
| `pkg/degradation` | Degradación elegante ante pérdida de conexión |
| `pkg/futuapi` | Reexportación de conveniencia: `NewClient()`, `NewClientFromEnv()` |
| `pkg/health` | Comprobaciones de estado para sondas de vida/disponibilidad de OpenD |
| `pkg/history` | Descargas de K-lines históricas con paginación automática |
| `pkg/market` | Horarios de mercado, calendario de trading, detección de sesiones |
| `pkg/metrics` | Recopilación de métricas de rendimiento del lado del cliente |
| `pkg/option` | Consulta de cadenas de opciones, análisis de códigos, utilidades de Greeks |
| `pkg/pb/*` | 184 tipos protobuf (v10.10.7008) |
| `pkg/ratelimit` | Limitación de tasa de la API (token bucket por protoID) |
| `pkg/retry` | Reintentos configurables con retroceso exponencial |
| `pkg/trd/audit.go` | Registro de auditoría de operaciones |
| `pkg/trd/validation.go` | Validación previa de órdenes |
| `pkg/tracing` | Interfaces centrales de trazado (Tracer, Span, no-op por defecto) |
| `pkg/tracing/otel` | Adaptador de trazado basado en OpenTelemetry (opcional) |

## APIs comunes

### Conexión

```go
// Configuración manual
cli := client.New(
	client.WithDialTimeout(10*time.Second),
	client.WithAPISetTimeout(30*time.Second),
).WithTradeEnv(constant.TrdEnv_Simulate)

// Desde variables de entorno: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, FUTU_ENCRYPT, FUTU_LOG_LEVEL
cli, _ := client.NewClientFromEnv()

cli.Connect("127.0.0.1:11111")
// cli.GetConnID(), cli.GetServerVer(), cli.IsEncrypt(), cli.GetLoginUserID()
// cli.CanSendProto(protoID)
```

### Datos de mercado

| Función | Descripción |
|---|---|
| `GetQuote(ctx, c, market, code)` | Cotización en tiempo real |
| `GetKLines(ctx, c, market, code, klType, num)` | Últimas barras K-line |
| `GetOrderBook(ctx, c, market, code, num)` | Profundidad de compra/venta |
| `GetTicker(ctx, c, market, code, num)` | Operaciones tick a tick |
| `GetStaticInfo(ctx, c, market, code)` | Nombre del valor, tipo, tamaño del lote |
| `GetSecuritySnapshot(ctx, c, securities)` | Instantánea completa de múltiples valores |
| `GetCapitalFlow(ctx, c, market, code)` | Flujo de capital |
| `RequestHistoryKL(ctx, c, market, code, klType, start, end)` | K-lines históricas (paginación automática) |
| `RequestHistoryKLQuota(ctx, c)` | Uso de cuota de la API |

### Trading

| Función | Descripción |
|---|---|
| `GetAccountList(ctx, c)` | Todas las cuentas de trading |
| `UnlockTrading(ctx, c, pwdMD5)` | Desbloquear el trading |
| `GetFunds(ctx, c, accID)` | Fondos y poder de la cuenta |
| `PlaceOrder(ctx, c, accID, market, code, side, orderType, price, qty)` | Realizar orden |
| `ModifyOrder(ctx, c, accID, market, orderID, op, price, qty)` | Modificar o cancelar orden |
| `GetOrderList(ctx, c, accID)` | Órdenes activas |
| `GetPositionList(ctx, c, accID)` | Posiciones actuales con P&L |
| `GetHistoryOrderList(ctx, c, accID, market, start, end)` | Órdenes históricas |
| `GetOrderFillList(ctx, c, accID)` | Ejecuciones de órdenes |

### Suscripciones

| Función | Descripción |
|---|---|
| `Subscribe(ctx, c, market, code, []SubType)` | Suscribirse a tipos de push |
| `Unsubscribe(ctx, c, market, code, []SubType)` | Cancelar suscripción |
| `chanpkg.SubscribeQuote(ctx, cli, market, code, ch)` | Push de cotizaciones por canal |
| `chanpkg.SubscribeKLine(ctx, cli, market, code, klType, ch)` | Push de una sola K-line por canal |
| `chanpkg.SubscribeKLines(ctx, cli, market, code, []klTypes, ch)` | Push de múltiples K-lines con filtro |
| `chanpkg.SubscribeTicker(ctx, cli, market, code, ch)` | Push de ticker por canal |
| `chanpkg.SubscribeOrderBook(ctx, cli, market, code, ch)` | Push del libro de órdenes por canal |

## Compilación y pruebas

```bash
go build ./...      # Compila todo
go vet ./...        # Análisis estático
go test -race ./... # Suite completa con detector de carreras
```

## Arquitectura

```
Application
  └── client/Client         (public wrappers)
       └── pkg/*            (qot, trd, sys — business logic)
            └── internal/client/Client   (connection, reconnect)
                 └── internal/client/Conn  (TCP I/O, packet framing)
                      └── Futu OpenD (TCP socket)
```

Toda la comunicación es mediante Protocol Buffers sobre TCP. Consulta [DESIGN.md](DESIGN.md) para ver las decisiones completas de arquitectura y [internal/testutil/mock](internal/testutil/mock/) para el servidor OpenD simulado que se usa en las pruebas.

## Solución de problemas

| Error | Causa probable |
|-------|-------------|
| `connection refused` | OpenD no está en ejecución. Comprueba `FUTU_OPEND_ADDR`. |
| sin datos de `GetQuote` (acciones de EE. UU.) | Primero debe llamarse a `Subscribe` para el mercado de EE. UU. Las acciones de Hong Kong no lo necesitan. |
| `The packet body SHA1 signature is incorrect` (OpenD muy antiguo) | Actualiza OpenD a v10.5+. El SDK usa SHA1(texto cifrado), que OpenD acepta. |
| `解析protobuf协议失败` | Faltan campos C2S obligatorios en el cuerpo de la solicitud. |
| `模拟交易不支持` | Función no disponible en modo simulado; usa `WithTradeEnv(TrdEnv_Real)`. |

## Contribuir

1. Haz un fork del repositorio.
2. Crea una rama de funcionalidad (`git checkout -b feat/my-change`).
3. Asegúrate de que todas las pruebas existentes pasen: `go test -race ./...`
4. Añade pruebas para cualquier funcionalidad nueva.
5. Ejecuta `go vet ./...` y corrige cualquier advertencia.
6. Abre una pull request.

Consulta [CHANGELOG.md](CHANGELOG.md) para el historial de versiones y [ENHANCEMENT_PLAN.md](docs/IMPLEMENTATION_COMPLETE.md) para la hoja de ruta.

## Véase también

- [CHANGELOG](CHANGELOG.md) — historial de versiones y notas de la versión
- [Version Map](docs/VERSION_MAP.md) — qué protocolo de Futu OpenD / recuento de proto / `clientVer` incluye cada versión del SDK
- [USAGE Guide](docs/USAGE.md) — configuración detallada, entorno y patrones avanzados
- [DESIGN](DESIGN.md) — arquitectura, decisiones de diseño, patrones de API
- [ENHANCEMENT_PLAN](docs/IMPLEMENTATION_COMPLETE.md) — próximas funcionalidades y hoja de ruta
- [futuapi4go-demo](https://github.com/shing1211/futuapi4go-demo) — ejemplos ejecutables para cada funcionalidad

## Licencia

Apache License 2.0 — see [LICENSE](LICENSE).

> **Trading Disclaimer**: Trading financial instruments carries significant risk. Always test thoroughly in simulate mode before using real funds.
