// Package chanpkg provides channel-based asynchronous delivery for push
// notifications from Futu OpenD. Unlike the callback-based RegisterHandler
// approach, Chan delivers parsed push messages through Go channels.
//
// Usage:
//
//	ch := make(chan *push.UpdateBasicQot, 100)
//	stop := chanpkg.SubscribeQuote(cli, "HK.00700", ch)
//	defer stop()
//	for quote := range ch {
//	    fmt.Printf("price: %.3f", quote.Price)
//	}
//
// 简体中文:
// chanpkg 包提供基于 channel 的异步推送通知投递方式。
// 与基于回调的 RegisterHandler 不同，Chan 通过 Go channel 传递解析后的推送消息。
//
// 繁體中文:
// chanpkg 包提供基於 channel 的異步推送通知投遞方式。
// 與基於回調的 RegisterHandler 不同，Chan 通過 Go channel 傳遞解析後的推送消息。
package chanpkg
