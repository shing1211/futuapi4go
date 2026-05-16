package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

type entry struct {
	key       string
	value     []*qot.KLine
	expiresAt time.Time
	prev      *entry
	next      *entry
}

type KLCache struct {
	mu         sync.RWMutex
	entries    map[string]*entry
	head       *entry
	tail       *entry
	maxEntries int
	ttl        time.Duration
}

type KLCacheOption func(*KLCache)

func WithMaxEntries(n int) KLCacheOption {
	return func(c *KLCache) {
		if n > 0 {
			c.maxEntries = n
		}
	}
}

func WithTTL(d time.Duration) KLCacheOption {
	return func(c *KLCache) {
		if d > 0 {
			c.ttl = d
		}
	}
}

func NewKLCache(opts ...KLCacheOption) *KLCache {
	c := &KLCache{
		entries:    make(map[string]*entry),
		maxEntries: 1000,
		ttl:        5 * time.Minute,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func cacheKey(security *qotcommon.Security, klType, rehabType int32) string {
	return fmt.Sprintf("%d|%s|%d|%d", security.GetMarket(), security.GetCode(), klType, rehabType)
}

func (c *KLCache) remove(e *entry) {
	if e.prev != nil {
		e.prev.next = e.next
	} else {
		c.head = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	} else {
		c.tail = e.prev
	}
	delete(c.entries, e.key)
}

func (c *KLCache) moveToFront(e *entry) {
	if c.head == e {
		return
	}
	if e.prev != nil {
		e.prev.next = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	}
	if c.tail == e {
		c.tail = e.prev
	}
	e.prev = nil
	e.next = c.head
	if c.head != nil {
		c.head.prev = e
	}
	c.head = e
	if c.tail == nil {
		c.tail = e
	}
}

func (c *KLCache) evictLRU() {
	if c.tail == nil {
		return
	}
	c.remove(c.tail)
}

func (c *KLCache) evictExpired() {
	now := time.Now()
	e := c.tail
	for e != nil && now.After(e.expiresAt) {
		prev := e.prev
		c.remove(e)
		e = prev
	}
}

func (c *KLCache) Get(security *qotcommon.Security, klType, rehabType int32) ([]*qot.KLine, bool) {
	key := cacheKey(security, klType, rehabType)
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(e.expiresAt) {
		c.remove(e)
		return nil, false
	}

	c.moveToFront(e)
	result := make([]*qot.KLine, len(e.value))
	copy(result, e.value)
	return result, true
}

func (c *KLCache) Set(security *qotcommon.Security, klType, rehabType int32, klines []*qot.KLine) {
	key := cacheKey(security, klType, rehabType)
	c.mu.Lock()
	defer c.mu.Unlock()

	if existing, ok := c.entries[key]; ok {
		c.remove(existing)
	}

	for len(c.entries) >= c.maxEntries {
		c.evictLRU()
	}

	e := &entry{
		key:       key,
		value:     klines,
		expiresAt: time.Now().Add(c.ttl),
	}
	c.entries[key] = e
	c.moveToFront(e)
}

func (c *KLCache) Invalidate(security *qotcommon.Security, klType, rehabType int32) {
	key := cacheKey(security, klType, rehabType)
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.entries[key]; ok {
		c.remove(e)
	}
}

func (c *KLCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[string]*entry)
	c.head = nil
	c.tail = nil
}

func (c *KLCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.entries)
}

func (c *KLCache) cleanupExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.evictExpired()
}

func (c *KLCache) StartCleanup(interval time.Duration, stopCh <-chan struct{}) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.cleanupExpired()
			case <-stopCh:
				return
			}
		}
	}()
}

type KLCachedClient struct {
	client *futuapi.Client
	cache  *KLCache
}

func NewKLCachedClient(client *futuapi.Client, cache *KLCache) *KLCachedClient {
	return &KLCachedClient{
		client: client,
		cache:  cache,
	}
}

func (kc *KLCachedClient) GetKL(ctx context.Context, rehabType, klType int32, security *qotcommon.Security) ([]*qot.KLine, error) {
	if cached, ok := kc.cache.Get(security, klType, rehabType); ok {
		return cached, nil
	}

	klines, err := qot.RequestHistoryKL(ctx, kc.client, &qot.RequestHistoryKLRequest{
		RehabType: rehabType,
		KlType:    klType,
		Security:  security,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*qot.KLine, len(klines.KLList))
	for i, kl := range klines.KLList {
		result[i] = kl
	}

	kc.cache.Set(security, klType, rehabType, result)
	return result, nil
}
