package dataqueue

import (
	"testing"
)

// ActionNoop    Action = 0
// ActionSend    Action = 1 << 1
// ActionStop    Action = 1 << 2
// ActionStopAll Action = 1 << 3

// ActionSendAndStop    Action = ActionSend | ActionStop
// ActionSendAndStopAll Action = ActionSend | ActionStopAll

func TestAction_IsNoop(t *testing.T) {
	tests := []struct {
		name   string
		action Action
		want   bool
	}{
		{"ActionNoop", ActionNoop, true},

		{"ActionSend", ActionSend, false},
		{"ActionStop", ActionStop, false},
		{"ActionStopAll", ActionStopAll, false},
		{"ActionSendAndStop", ActionSendAndStop, false},
		{"ActionSendAndStopAll", ActionSendAndStopAll, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.IsNoop(); got != tt.want {
				t.Errorf("Action.IsNoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAction_IsSend(t *testing.T) {
	tests := []struct {
		name   string
		action Action
		want   bool
	}{
		{"ActionSend", ActionSend, true},
		{"ActionSendAndStop", ActionSendAndStop, true},

		{"ActionNoop", ActionNoop, false},
		{"ActionStop", ActionStop, false},
		{"ActionStopAll", ActionStopAll, false},
		{"ActionSendAndStopAll", ActionSendAndStopAll, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.IsSend(); got != tt.want {
				t.Errorf("Action.IsSend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAction_IsStop(t *testing.T) {
	tests := []struct {
		name   string
		action Action
		want   bool
	}{
		{"ActionStop", ActionStop, true},
		{"ActionSendAndStop", ActionSendAndStop, true},

		{"ActionNoop", ActionNoop, false},
		{"ActionSend", ActionSend, false},
		{"ActionStopAll", ActionStopAll, false},
		{"ActionSendAndStopAll", ActionSendAndStopAll, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.IsStop(); got != tt.want {
				t.Errorf("Action.IsStop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAction_IsStopAll(t *testing.T) {
	tests := []struct {
		name   string
		action Action
		want   bool
	}{
		{"ActionStopAll", ActionStopAll, true},
		{"ActionSendAndStopAll", ActionSendAndStopAll, true},

		{"ActionNoop", ActionNoop, false},
		{"ActionSend", ActionSend, false},
		{"ActionStop", ActionStop, false},
		{"ActionSendAndStop", ActionSendAndStop, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.IsStopAll(); got != tt.want {
				t.Errorf("Action.IsStopAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
