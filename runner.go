package dataqueue

import (
	"context"
	"reflect"
)

// StartSimpleConsumer runner with nbP producers
func StartSimpleConsumer(ctx context.Context, runner Runner, nbP int) error {
	return Start(ctx, runner, nbP, 1)
}

// StartSimpleProducer runner with nbC consumers
func StartSimpleProducer(ctx context.Context, runner Runner, nbC int) error {
	return Start(ctx, runner, 1, nbC)
}

// Start runner with nbP producers and nbC consumers
func Start(ctx context.Context, runner Runner, nbP, nbC int) error {
	pFun, cFun := runner.Callbacks()
	return RunFunc(ctx, nbP, pFun, nbC, cFun)
}

// RunFunc create and producers and consumers with callbacks
func RunFunc(ctx context.Context, nbP int, pFunc ProducerFunc, nbC int, cFunc ConsumerFunc) error {
	p, err := NewProducer(nbP, pFunc)
	if err != nil {
		return ErrInvalidProducer
	}
	c, err := NewConsumer(p, nbC, cFunc)
	if err != nil {
		return ErrInvalidConsumer
	}

	return Run(ctx, p, c)
}

func isNilValue(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
	// Faster implementation but avoiding use of unsafe
	// return (*[2]uintptr)(unsafe.Pointer(&i))[1] == 0
}

// Run producers and consumers
func Run(ctx context.Context, p Producer, c Consumer) error {
	if isNilValue(p) {
		return ErrInvalidProducer
	}
	if isNilValue(c) {
		return ErrInvalidConsumer
	}
	p.Run(ctx)
	c.Run(ctx)

	p.Wait(ctx)
	c.Wait(ctx)

	return nil
}
