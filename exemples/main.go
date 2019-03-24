package main

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"time"

	"dataqueue"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ChalengeFinder struct
type ChalengeFinder struct {
	difficulty uint64
	chalenge   uint64
}

func (p *ChalengeFinder) chalengeFinder(cxt context.Context) (dataqueue.DataType, dataqueue.Action) {
	r := rand.Uint64() % p.difficulty

	if r > p.chalenge {
		return r, dataqueue.ActionSend
	}

	log.Printf("Chalenge Found")
	return r, dataqueue.ActionSendAndStopAll
}

func (p *ChalengeFinder) chalengeFound(ctx context.Context, item dataqueue.DataType) {
	switch item.(type) {
	case uint64:
		if item.(uint64) <= p.chalenge {
			log.Printf("Chalenge Recieved: %v", item)
		}
	default:
		return
	}
}

// Callbacks Runner interface
func (p *ChalengeFinder) Callbacks() (dataqueue.ProducerFunc, dataqueue.ConsumerFunc) {
	return p.chalengeFinder, p.chalengeFound
}

func main() {
	ctx := context.Background()
	ctx, ctxCancel := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer ctxCancel()

	startTime := time.Now()

	dataqueue.StartSimpleConsumer(ctx,
		&ChalengeFinder{
			difficulty: 1 << 30,
			chalenge:   1000,
		},
		runtime.NumCPU()*4,
	)

	log.Printf("Execution time: %s", time.Now().Sub(startTime).Round(time.Millisecond))
}
