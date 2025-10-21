package consts

//go:generate stringer -type=EventStatus
type EventStatus int

const (
	Pending EventStatus = iota
	Executing
	Complete
	Failed
)
