// Package push provides handlers for parsing push notification payloads
// from Futu OpenD. Use RegisterHandler on the client to receive real-time
// market data and order updates.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/push"
//
//	cli.RegisterHandler(push.ProtoID_Qot_UpdateBasicQot, func(protoID uint32, body []byte) {
//	    data, err := push.ParseUpdateBasicQot(body)
//	    // ...
//	})
//
// For channel-based async delivery (Go-native alternative to callbacks),
// use the Chan subpackage:
//
//	import "github.com/shing1211/futuapi4go/pkg/push/chanpkg"
//
//	ch := make(chan *push.UpdateBasicQot, 100)
//	stop := chanpkg.SubscribeQuote(cli, "HK.00700", ch)
//	defer stop()
//
// # Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package push
