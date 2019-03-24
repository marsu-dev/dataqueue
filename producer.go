package dataqueue

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// DataProducer struct
type DataProducer struct {
	stream       chan DataType
	count        int
	producerFunc ProducerFunc
	wg           sync.WaitGroup
	stopingAll   uint64
	stopAll      context.CancelFunc
}

// NewProducer retrun new DataProducer
func NewProducer(count int, producerFunc ProducerFunc) (*DataProducer, error) {
	if producerFunc == nil {
		return nil, ErrInvalidFunc
	}
	return &DataProducer{
		stream:       make(chan DataType),
		count:        count,
		producerFunc: producerFunc,
	}, nil
}

// Stream retrun readonly data stream
func (p *DataProducer) Stream() <-chan DataType {
	return p.stream
}

// Run data producers
func (p *DataProducer) Run(ctx context.Context) {
	ctx, stopAll := context.WithCancel(ctx)
	p.stopAll = stopAll
	for i := 0; i < p.count; i++ {
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			p.produce(ctx, id)
		}(i)
	}
}

// Wait for data producers
func (p *DataProducer) Wait(ctx context.Context) {
	defer func() {
		if p.stopAll == nil {
			return
		}
		p.stopAll()
		p.stopAll = nil
	}()
	p.wg.Wait()
	close(p.stream)
}

func (p *DataProducer) produce(ctx context.Context, id int) {
	ch := make(chan DataType)
	quit := make(chan struct{})
	defer close(quit)

	go func() {
		for {
			select {
			case <-time.After(0 * time.Millisecond):
				if p.stopingAll > 0 {
					break
				}

				item, action := p.producerFunc(ctx)

				if action.IsStopAll() {
					atomic.StoreUint64(&p.stopingAll, 1)
				}
				if action.IsSend() {
					if item != nil {
						ch <- item
					}
				}
				if action.IsNoop() {
					break
				}
				if action.IsStop() {
					quit <- struct{}{}
					return
				}
				if action.IsStopAll() {
					quit <- struct{}{}
					if p.stopAll != nil {
						p.stopAll()
					}
					return
				}

			case <-ctx.Done():
				quit <- struct{}{}
				return
			}
		}
	}()

	for {
		select {
		case item, ok := <-ch:
			if !ok {
				return
			}
			p.stream <- item

		case <-quit:
			close(ch)
			return
		}
	}
}
