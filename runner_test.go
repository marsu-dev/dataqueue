package dataqueue

import (
	"context"
	"sync/atomic"
	"testing"
)

type NilRunner struct {
}

func (p *NilRunner) Callbacks() (ProducerFunc, ConsumerFunc) {
	return nil, nil
}

type CountRunner struct {
	action Action
	pCount uint64
	rCount uint64
}

func (p *CountRunner) producer(ctx context.Context) (DataType, Action) {
	atomic.AddUint64(&p.pCount, 1)
	return 42, p.action
}

func (p *CountRunner) consumer(ctx context.Context, item DataType) {
	atomic.AddUint64(&p.rCount, 1)
}

func (p *CountRunner) Callbacks() (ProducerFunc, ConsumerFunc) {
	return p.producer, p.consumer
}

func TestStartSimpleConsumer(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		runner Runner
		nbP    int
	}
	type want struct {
		err    bool
		pCount uint64
		rCount uint64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"NilRunner", args{ctx, &NilRunner{}, 1}, want{true, 0, 0}},
		{"CountRunner", args{ctx, &CountRunner{action: ActionSendAndStop}, 4}, want{false, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartSimpleConsumer(ctx, tt.args.runner, tt.args.nbP); (err != nil) != tt.want.err {
				t.Errorf("StartSimpleConsumer() error = %v, want.err %v", err, tt.want.err)
			}
			switch tt.args.runner.(type) {
			case *CountRunner:
				r := tt.args.runner.(*CountRunner)
				if r.pCount != tt.want.pCount {
					t.Errorf("StartSimpleConsumer() pCount = %v, want.pCount %v", r.pCount, tt.want.pCount)
				}
				if r.rCount != tt.want.rCount {
					t.Errorf("StartSimpleConsumer() rCount = %v, want.rCount %v", r.rCount, tt.want.rCount)
				}
			}
		})
	}
}

func TestStartSimpleProducer(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		runner Runner
		nbC    int
	}
	type want struct {
		err    bool
		pCount uint64
		rCount uint64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"NilRunner", args{ctx, &NilRunner{}, 1}, want{true, 0, 0}},
		{"CountRunner", args{ctx, &CountRunner{action: ActionSendAndStop}, 4}, want{false, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartSimpleProducer(ctx, tt.args.runner, tt.args.nbC); (err != nil) != tt.want.err {
				t.Errorf("StartSimpleProducer() error = %v, want.err %v", err, tt.want.err)
			}
			switch tt.args.runner.(type) {
			case *CountRunner:
				r := tt.args.runner.(*CountRunner)
				if r.pCount != tt.want.pCount {
					t.Errorf("StartSimpleProducer() pCount = %v, want.pCount %v", r.pCount, tt.want.pCount)
				}
				if r.rCount != tt.want.rCount {
					t.Errorf("StartSimpleProducer() rCount = %v, want.rCount %v", r.rCount, tt.want.rCount)
				}
			}
		})
	}
}

func TestStartBoth(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		runner Runner
		nbP    int
		nbC    int
	}
	type want struct {
		err    bool
		pCount uint64
		rCount uint64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"NilRunner", args{ctx, &NilRunner{}, 1, 1}, want{true, 0, 0}},
		{"CountRunner", args{ctx, &CountRunner{action: ActionSendAndStop}, 4, 4}, want{false, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Start(ctx, tt.args.runner, tt.args.nbP, tt.args.nbC); (err != nil) != tt.want.err {
				t.Errorf("Start() error = %v, want.err %v", err, tt.want.err)
			}
			switch tt.args.runner.(type) {
			case *CountRunner:
				r := tt.args.runner.(*CountRunner)
				if r.pCount != tt.want.pCount {
					t.Errorf("Start() pCount = %v, want.pCount %v", r.pCount, tt.want.pCount)
				}
				if r.rCount != tt.want.rCount {
					t.Errorf("Start() rCount = %v, want.rCount %v", r.rCount, tt.want.rCount)
				}
			}
		})
	}
}

func TestStartFunc(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		runner Runner
		nbP    int
		nbC    int
	}
	type want struct {
		err    bool
		pCount uint64
		rCount uint64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"NilRunner", args{ctx, &NilRunner{}, 1, 1}, want{true, 0, 0}},
		{"CountRunner", args{ctx, &CountRunner{action: ActionSendAndStop}, 1, 1}, want{false, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pFunc, cFunc := tt.args.runner.Callbacks()
			if err := RunFunc(ctx, tt.args.nbP, pFunc, tt.args.nbC, cFunc); (err != nil) != tt.want.err {
				t.Errorf("RunFunc() error = %v, want.err %v", err, tt.want.err)
			}
			switch tt.args.runner.(type) {
			case *CountRunner:
				r := tt.args.runner.(*CountRunner)
				if r.pCount != tt.want.pCount {
					t.Errorf("RunFunc() pCount = %v, want.pCount %v", r.pCount, tt.want.pCount)
				}
				if r.rCount != tt.want.rCount {
					t.Errorf("RunFunc() rCount = %v, want.rCount %v", r.rCount, tt.want.rCount)
				}
			}
		})
	}
}

func TestRunInterface(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		runner Runner
		nbP    int
		nbC    int
	}
	type want struct {
		err    bool
		pCount uint64
		rCount uint64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"NilRunner", args{ctx, &NilRunner{}, 1, 1}, want{true, 0, 0}},

		{"CountRunnerStop", args{ctx, &CountRunner{action: ActionStop}, 4, 4}, want{false, 4, 0}},
		{"CountRunnerStopAll", args{ctx, &CountRunner{action: ActionStopAll}, 4, 4}, want{false, 1, 0}},
		{"CountRunnerSendAndStop", args{ctx, &CountRunner{action: ActionSendAndStop}, 4, 4}, want{false, 4, 4}},
		{"CountRunnerSendAndStopAll", args{ctx, &CountRunner{action: ActionSendAndStopAll}, 4, 4}, want{false, 1, 1}},

		{"CountRunnerStop", args{ctx, &CountRunner{action: ActionStop}, 8, 4}, want{false, 8, 0}},
		{"CountRunnerStopAll", args{ctx, &CountRunner{action: ActionStopAll}, 8, 4}, want{false, 1, 0}},
		{"CountRunnerSendAndStop", args{ctx, &CountRunner{action: ActionSendAndStop}, 8, 4}, want{false, 8, 8}},
		{"CountRunnerSendAndStopAll", args{ctx, &CountRunner{action: ActionSendAndStopAll}, 8, 4}, want{false, 1, 1}},

		{"CountRunnerStop", args{ctx, &CountRunner{action: ActionStop}, 4, 8}, want{false, 4, 0}},
		{"CountRunnerStopAll", args{ctx, &CountRunner{action: ActionStopAll}, 4, 8}, want{false, 1, 0}},
		{"CountRunnerSendAndStop", args{ctx, &CountRunner{action: ActionSendAndStop}, 4, 8}, want{false, 4, 4}},
		{"CountRunnerSendAndStopAll", args{ctx, &CountRunner{action: ActionSendAndStopAll}, 4, 8}, want{false, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pFunc, cFunc := tt.args.runner.Callbacks()
			p, _ := NewProducer(tt.args.nbP, pFunc)
			c, _ := NewConsumer(p, tt.args.nbC, cFunc)

			if err := Run(ctx, p, c); (err != nil) != tt.want.err {
				t.Errorf("RunFunc() error = %v, want.err %v", err, tt.want.err)
			}
			switch tt.args.runner.(type) {
			case *CountRunner:
				r := tt.args.runner.(*CountRunner)
				if r.pCount != tt.want.pCount {
					t.Errorf("RunFunc() pCount = %v, want.pCount %v", r.pCount, tt.want.pCount)
				}
				if r.rCount != tt.want.rCount {
					t.Errorf("RunFunc() rCount = %v, want.rCount %v", r.rCount, tt.want.rCount)
				}
			}
		})
	}
}
