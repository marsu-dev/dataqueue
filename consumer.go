package dataqueue

import (
	"context"
	"sync"
)

// DataConsumer struct
type DataConsumer struct {
	data         Data
	count        int
	consumerFunc ConsumerFunc
	wg           sync.WaitGroup
}

// NewConsumer retrun new DataConsumer
func NewConsumer(data Data, count int, consumerFunc ConsumerFunc) (*DataConsumer, error) {
	if consumerFunc == nil {
		return nil, ErrInvalidFunc
	}
	return &DataConsumer{
		data:         data,
		count:        count,
		consumerFunc: consumerFunc,
	}, nil
}

// Run data consumers
func (c *DataConsumer) Run(ctx context.Context) {
	for i := 0; i < c.count; i++ {
		c.wg.Add(1)

		go func(id int) {
			defer c.wg.Done()
			c.consume(ctx, id)
		}(i)
	}
}

// Wait for data consumers
func (c *DataConsumer) Wait(ctx context.Context) {
	c.wg.Wait()
}

func (c *DataConsumer) consume(ctx context.Context, id int) {
	var skip bool
	for {
		select {
		case item, ok := <-c.data.Stream():
			if !ok {
				return
			}
			// read all data left in stream to avoid deadlock
			if skip {
				break
			}
			if item == nil {
				break
			}
			c.consumerFunc(ctx, item)

		case <-ctx.Done():
			// wait for stream stop
			skip = true
			break
		}
	}
}
