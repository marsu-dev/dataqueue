package dataqueue

// Actions bit flag enum
const (
	ActionSend    Action = 1 << 1
	ActionStop    Action = 1 << 2
	ActionStopAll Action = 1 << 3

	ActionSendAndStop    Action = ActionSend | ActionStop
	ActionSendAndStopAll Action = ActionSend | ActionStopAll
)

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
