package consts

//go:generate stringer -type=EventStatus
type EventStatus int

const (
	EventPending EventStatus = iota
	EventExecuting
	EventComplete
	Failed
)
