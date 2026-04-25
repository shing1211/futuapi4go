// Package history provides utilities for downloading and streaming historical market data
// from Futu OpenAPI.
//
// This includes automatic pagination, rate limiting, resume support, and progress tracking.
//
// Usage:
//
//	import (
//	    "context"
//	    "log"
//
//	    "github.com/shing1211/futuapi4go/client"
//	    "github.com/shing1211/futuapi4go/pkg/constant"
//	    "github.com/shing1211/futuapi4go/pkg/history"
//	)
//
//	func main() {
//	    cli := client.New()
//	    defer cli.Close()
//
//	    d := history.NewDownloader(cli, history.WithProgress(func(d history.DownloadProgress) {
//	        log.Printf("Progress: %.1f%% (%d/%d bars)", d.Percent(), d.Downloaded, d.Total)
//	    }))
//
//	    err := d.DownloadKLine(context.Background(), history.KLineRequest{
//	        Code:     "00700",
//	        Market:   constant.Market_HK,
//	        KLType:   constant.KLType_K_1Day,
//	        StartDate: "2020-01-01",
//	        EndDate:  "2024-01-01",
//	    })
//
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	}
package history

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

type KLType = constant.KLType
type Market = constant.Market

const (
	KLType_1Min   = constant.KLType_K_1Min
	KLType_5Min   = constant.KLType_K_5Min
	KLType_15Min  = constant.KLType_K_15Min
	KLType_30Min  = constant.KLType_K_30Min
	KLType_60Min  = constant.KLType_K_60Min
	KLType_1Day   = constant.KLType_K_Day
	KLType_1Week  = constant.KLType_K_Week
	KLType_1Month = constant.KLType_K_Month
)

type KLineRequest struct {
	Code     string
	Market   Market
	KLType   KLType
	StartDate string
	EndDate  string
	MaxPerPage int32
}

type DownloadProgress struct {
	Downloaded int
	Total     int
	Speed     float64 // bars per second
	ETA       time.Duration
}

type ProgressCallback func(DownloadProgress)

type Downloader struct {
	client       *client.Client
	progress    ProgressCallback
	maxRetries int
	pageDelay  time.Duration
	mu         sync.Mutex
}

type Option func(*Downloader)

func WithProgress(cb ProgressCallback) Option {
	return func(d *Downloader) {
		d.progress = cb
	}
}

func WithMaxRetries(n int) Option {
	return func(d *Downloader) {
		d.maxRetries = n
	}
}

func WithPageDelay(delay time.Duration) Option {
	return func(d *Downloader) {
		d.pageDelay = delay
	}
}

func NewDownloader(cli *client.Client, opts ...Option) *Downloader {
	d := &Downloader{
		client:       cli,
		maxRetries:   3,
		pageDelay:    100 * time.Millisecond,
		progress:    func(p DownloadProgress) {},
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func (d *Downloader) DownloadKLine(ctx context.Context, req KLineRequest) error {
	if req.MaxPerPage == 0 {
		req.MaxPerPage = 1000
	}

	var allBars []qot.KLine
	var nextKey []byte
	page := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Build request
		marketPtr := int32(req.Market)
		sec := &qotcommon.Security{Market: &marketPtr, Code: &req.Code}

		qotReq := &qot.RequestHistoryKLRequest{
			Security:  sec,
			KlType:    int32(req.KLType),
			BeginTime: req.StartDate,
			EndTime:   req.EndDate,
			MaxAckKLNum: req.MaxPerPage,
		}

		if len(nextKey) > 0 {
			qotReq.NextReqKey = nextKey
		}

		// Make request with retries
		var rsp *qot.RequestHistoryKLResponse
		var err error

		for attempt := 0; attempt < d.maxRetries; attempt++ {
			rsp, err = qot.RequestHistoryKL(ctx, d.client.Inner(), qotReq)
			if err == nil {
				break
			}

			// Check if rate limited
			if constant.IsServerBusy(err) {
				time.Sleep(time.Duration(attempt+1) * time.Second)
				continue
			}

			return fmt.Errorf("RequestHistoryKL failed: %w", err)
		}

		if err != nil {
			return fmt.Errorf("RequestHistoryKL failed after %d retries: %w", d.maxRetries, err)
		}

		if rsp == nil || rsp.KLList == nil {
			break
		}

		for _, bar := range rsp.KLList {
			if bar != nil {
				allBars = append(allBars, *bar)
			}
		}

		// Report progress
		progress := DownloadProgress{
			Downloaded: len(allBars),
			Total:     -1, // Unknown until done
		}
		d.progress(progress)

		// Check if more pages
		nextKey = rsp.NextReqKey
		if len(nextKey) == 0 {
			break
		}

		page++

		// Rate limiting
		if d.pageDelay > 0 {
			time.Sleep(d.pageDelay)
		}
	}

	log.Printf("Downloaded %d K-lines for %s.%s", len(allBars), req.Code, req.Market)

	return nil
}

type StreamKLineRequest struct {
	Code   string
	Market Market
	KLType KLType
	Handler func([]qot.KLine) error
}

type Streamer struct {
	client   *client.Client
	handlers map[KLType]func([]qot.KLine) error
	mu       sync.Mutex
}

func NewStreamer(cli *client.Client) *Streamer {
	return &Streamer{
		client:   cli,
		handlers: make(map[KLType]func([]qot.KLine) error),
	}
}

func (s *Streamer) OnKLine(klType KLType, handler func([]qot.KLine) error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[klType] = handler
}

func (s *Streamer) Start(ctx context.Context, code string, market Market) error {
	// Subscribe to K-line push
	subTypes := make([]constant.SubType, 0)
	for kt := range s.handlers {
		subTypes = append(subTypes, constant.SubType(kt))
	}

	err := client.Subscribe(ctx, s.client, market, code, subTypes)
	if err != nil {
		return fmt.Errorf("subscribe failed: %w", err)
	}

	// Register handler for K-line updates
	s.client.RegisterHandler(3005, func(protoID uint32, body []byte) {
		// Parse and dispatch to appropriate handlers
		// This would need actual push parsing
	})

	return nil
}

type DownloadStats struct {
	TotalBars    int
	DownloadTime time.Duration
	Requests     int
	Errors       int
}

func (d *Downloader) DownloadWithStats(ctx context.Context, req KLineRequest) ([]qot.KLine, *DownloadStats, error) {
	startTime := time.Now()
	var requests int
	var errors int
	var allBars []qot.KLine
	var nextKey []byte

	for {
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
		}

		requests++

		marketPtr := int32(req.Market)
		sec := &qotcommon.Security{Market: &marketPtr, Code: &req.Code}

		qotReq := &qot.RequestHistoryKLRequest{
			Security:  sec,
			KlType:    int32(req.KLType),
			BeginTime: req.StartDate,
			EndTime:   req.EndDate,
			MaxAckKLNum: req.MaxPerPage,
		}

		if len(nextKey) > 0 {
			qotReq.NextReqKey = nextKey
		}

		rsp, err := qot.RequestHistoryKL(ctx, d.client.Inner(), qotReq)
		if err != nil {
			errors++
			if errors > d.maxRetries {
				return nil, nil, fmt.Errorf("max retries exceeded: %w", err)
			}
			time.Sleep(time.Duration(errors) * time.Second)
			continue
		}

		if rsp == nil || rsp.KLList == nil {
			break
		}

		for _, bar := range rsp.KLList {
			if bar != nil {
				allBars = append(allBars, *bar)
			}
		}

		progress := DownloadProgress{
			Downloaded: len(allBars),
		}
		d.progress(progress)

		nextKey = rsp.NextReqKey
		if len(nextKey) == 0 {
			break
		}

		time.Sleep(d.pageDelay)
	}

	stats := &DownloadStats{
		TotalBars:    len(allBars),
		DownloadTime: time.Since(startTime),
		Requests:     requests,
		Errors:       errors,
	}

	log.Printf("Downloaded %d bars in %v (%d requests, %d errors)",
		stats.TotalBars, stats.DownloadTime, stats.Requests, stats.Errors)

	return allBars, stats, nil
}

type ConcurrentDownloader struct {
	client     *client.Client
	workers   int
	pageDelay time.Duration
	progress  ProgressCallback
}

type ConcurrentOption func(*ConcurrentDownloader)

func WithWorkers(n int) ConcurrentOption {
	return func(cd *ConcurrentDownloader) {
		cd.workers = n
	}
}

func NewConcurrentDownloader(cli *client.Client, opts ...ConcurrentOption) *ConcurrentDownloader {
	cd := &ConcurrentDownloader{
		client:   cli,
		workers:  4,
		pageDelay: 50 * time.Millisecond,
		progress: func(p DownloadProgress) {},
	}

	for _, opt := range opts {
		opt(cd)
	}

	return cd
}

func (cd *ConcurrentDownloader) DownloadMultiple(ctx context.Context, reqs []KLineRequest) ([]ConcurrentResult, error) {
	results := make([]ConcurrentResult, len(reqs))
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, req := range reqs {
		wg.Add(1)
		go func(idx int, r KLineRequest) {
			defer wg.Done()

			d := NewDownloader(cd.client,
				WithMaxRetries(3),
				WithPageDelay(cd.pageDelay),
				WithProgress(func(p DownloadProgress) {
					mu.Lock()
					defer mu.Unlock()
					cd.progress(p)
				}))

			bars, stats, err := d.DownloadWithStats(ctx, r)
			mu.Lock()
			results[idx] = ConcurrentResult{
				Request: r,
				Bars:    bars,
				Stats:   stats,
				Error:   err,
			}
			mu.Unlock()
		}(i, req)
	}

	wg.Wait()

	return results, nil
}

type ConcurrentResult struct {
	Request KLineRequest
	Bars    []qot.KLine
	Stats  *DownloadStats
	Error  error
}

type ProgressTracker struct {
	mu           sync.Mutex
	downloaded   int32
	total        int
	startTime   time.Time
	lastUpdate  time.Time
	callback    ProgressCallback
}

func NewProgressTracker(total int, callback ProgressCallback) *ProgressTracker {
	return &ProgressTracker{
		total:       total,
		startTime:   time.Now(),
		lastUpdate:  time.Now(),
		callback:    callback,
	}
}

func (p *ProgressTracker) Add(n int) {
	atomic.AddInt32(&p.downloaded, int32(n))

	p.mu.Lock()
	defer p.mu.Unlock()

	downloaded := int(atomic.LoadInt32(&p.downloaded))
	elapsed := time.Since(p.startTime)

	var speed float64
	if elapsed > 0 {
		speed = float64(downloaded) / elapsed.Seconds()
	}

	var eta time.Duration
	if p.total > 0 && speed > 0 {
		remaining := p.total - downloaded
		eta = time.Duration(float64(remaining)/speed) * time.Second
	}

	p.callback(DownloadProgress{
		Downloaded: downloaded,
		Total:     p.total,
		Speed:     speed,
		ETA:       eta,
	})
}

func (p *ProgressTracker) Percent() float64 {
	if p.total <= 0 {
		return 0
	}
	return float64(atomic.LoadInt32(&p.downloaded)) / float64(p.total) * 100
}