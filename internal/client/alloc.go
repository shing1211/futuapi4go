// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package futuapi

import (
	"sync"
)

var (
 marshalPool = sync.Pool{
  New: func() interface{} {
   return new(byteBuffer)
  },
 }

 packetPool = sync.Pool{
  New: func() interface{} {
   return new(packetBuffer)
  },
 }

 responsePool = sync.Pool{
  New: func() interface{} {
   return &responseBuffer{
    data: make([]byte, 0, 4096),
   }
  },
 }
)

type byteBuffer struct {
 data []byte
}

func getByteBuffer(size int) *byteBuffer {
 bb := marshalPool.Get().(*byteBuffer)
 if cap(bb.data) < size {
  bb.data = make([]byte, size)
 } else {
  bb.data = bb.data[:size]
 }
 return bb
}

func putByteBuffer(bb *byteBuffer) {
 bb.data = bb.data[:0]
 marshalPool.Put(bb)
}

type packetBuffer struct {
 buf []byte
}

func getPacketBuffer() *packetBuffer {
 pb := packetPool.Get().(*packetBuffer)
 pb.buf = pb.buf[:0]
 return pb
}

func putPacketBuffer(pb *packetBuffer) {
 pb.buf = pb.buf[:0]
 packetPool.Put(pb)
}

type responseBuffer struct {
 data []byte
 used  bool
}

func getResponseBuffer(capacity int) *responseBuffer {
 rb := responsePool.Get().(*responseBuffer)
 if cap(rb.data) < capacity {
  rb.data = make([]byte, 0, capacity)
 } else {
  rb.data = rb.data[:0]
 }
 rb.used = false
 return rb
}

func putResponseBuffer(rb *responseBuffer) {
 rb.data = rb.data[:0]
 rb.used = false
 responsePool.Put(rb)
}

type bufferPool struct {
 marshalBuf   sync.Pool
 responseBuf sync.Pool
 packetBuf  sync.Pool
}

func newBufferPool(marshalSize, responseSize int) *bufferPool {
 return &bufferPool{
  marshalBuf: sync.Pool{
   New: func() interface{} {
    return &byteBuffer{data: make([]byte, marshalSize)}
   },
  },
  responseBuf: sync.Pool{
   New: func() interface{} {
    return &responseBuffer{data: make([]byte, 0, responseSize)}
   },
  },
  packetBuf: sync.Pool{
   New: func() interface{} {
    return &packetBuffer{buf: make([]byte, 0, 8192)}
   },
  },
 }
}

func (bp *bufferPool) GetMarshalBuf() *byteBuffer {
 return bp.marshalBuf.Get().(*byteBuffer)
}

func (bp *bufferPool) PutMarshalBuf(bb *byteBuffer) {
 bb.data = bb.data[:0]
 bp.marshalBuf.Put(bb)
}

func (bp *bufferPool) GetResponseBuf() *responseBuffer {
 rb := bp.responseBuf.Get().(*responseBuffer)
 rb.used = false
 return rb
}

func (bp *bufferPool) PutResponseBuf(rb *responseBuffer) {
 rb.data = rb.data[:0]
 rb.used = false
 bp.responseBuf.Put(rb)
}

func (bp *bufferPool) GetPacketBuf() *packetBuffer {
 pb := bp.packetBuf.Get().(*packetBuffer)
 pb.buf = pb.buf[:0]
 return pb
}

func (bp *bufferPool) PutPacketBuf(pb *packetBuffer) {
 pb.buf = pb.buf[:0]
 bp.packetBuf.Put(pb)
}