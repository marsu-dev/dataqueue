package dataqueue

// Actions bit flag enum
const (
	ActionNoop    Action = 0
	ActionSend    Action = 1 << 1
	ActionStop    Action = 1 << 2
	ActionStopAll Action = 1 << 3

	ActionSendAndStop    Action = ActionSend | ActionStop
	ActionSendAndStopAll Action = ActionSend | ActionStopAll
)

// IsNoop return if no action is required
func (action Action) IsNoop() bool {
	return action == ActionNoop
}

// IsSend return if action is to send
func (action Action) IsSend() bool {
	return action&ActionSend == ActionSend
}

// IsStop return if action is to stop
func (action Action) IsStop() bool {
	return action&ActionStop == ActionStop
}

// IsStopAll return if action is to stop all
func (action Action) IsStopAll() bool {
	return action&ActionStopAll == ActionStopAll
}
