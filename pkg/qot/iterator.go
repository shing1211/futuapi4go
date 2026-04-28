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

package qot

import (
	"context"
	"sync"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

type HistoryKLineIterator struct {
	ctx         context.Context
	client      *futuapi.Client
	req         *RequestHistoryKLRequest
	mu          sync.Mutex
	err         error
	atEnd       bool
	totalFetched int
	pageCount   int
}

func NewHistoryKLineIterator(ctx context.Context, c *futuapi.Client, req *RequestHistoryKLRequest) *HistoryKLineIterator {
	return &HistoryKLineIterator{
		ctx:    ctx,
		client: c,
		req:    req,
		atEnd:  false,
	}
}

func (it *HistoryKLineIterator) HasNext() bool {
	it.mu.Lock()
	defer it.mu.Unlock()
	return !it.atEnd && it.err == nil
}

func (it *HistoryKLineIterator) Next() ([]*KLine, error) {
	it.mu.Lock()
	defer it.mu.Unlock()

	if it.err != nil {
		return nil, it.err
	}
	if it.atEnd {
		return nil, nil
	}

	select {
	case <-it.ctx.Done():
		it.err = it.ctx.Err()
		return nil, it.err
	default:
	}

	rsp, err := RequestHistoryKL(it.ctx, it.client, it.req)
	it.pageCount++
	if err != nil {
		it.err = err
		return nil, err
	}

	it.totalFetched += len(rsp.KLList)

	if len(rsp.NextReqKey) == 0 {
		it.atEnd = true
	} else {
		it.req.NextReqKey = rsp.NextReqKey
	}

	return rsp.KLList, nil
}

func (it *HistoryKLineIterator) Err() error {
	it.mu.Lock()
	defer it.mu.Unlock()
	return it.err
}

func (it *HistoryKLineIterator) TotalFetched() int {
	it.mu.Lock()
	defer it.mu.Unlock()
	return it.totalFetched
}

func (it *HistoryKLineIterator) PageCount() int {
	it.mu.Lock()
	defer it.mu.Unlock()
	return it.pageCount
}
