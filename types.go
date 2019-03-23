package dataqueue

import (
	"context"
	"errors"
)

// DataType for data queue
type DataType interface{}

// Action type
type Action int

// Data queue interface
type Data interface {
	Stream() <-chan DataType
}

// Producer queue interface
type Producer interface {
	Run(ctx context.Context)
	Wait(ctx context.Context)
}

// Consumer queue interface
type Consumer interface {
	Run(ctx context.Context)
	Wait(ctx context.Context)
}

// ProducerFunc callback
type ProducerFunc func(ctx context.Context) (DataType, Action)

// ConsumerFunc callback
type ConsumerFunc func(ctx context.Context, item DataType)

// Runner queue interface
type Runner interface {
	Callbacks() (ProducerFunc, ConsumerFunc)
}

// Errors
var (
	ErrInvalidFunc     = errors.New("Invalid Func")
	ErrInvalidProducer = errors.New("Invalid Produce")
	ErrInvalidConsumer = errors.New("Invalid Consumer")
)
