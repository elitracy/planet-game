package task

import "github.com/elitracy/planets/engine"

type Status int

const (
	Pending Status = iota
	Executing
	Complete
	Failed
)

type Task interface {
	GetID() TaskID
	SetID(TaskID)
	GetStartTick() engine.Tick
	GetDuration() engine.Tick
	GetStatus() Status
}

func (e Status) String() string {
	switch e {
	case Pending:
		return "Pending"
	case Executing:
		return "Executing"
	case Complete:
		return "Complete"
	case Failed:
		return "Failed"
	default:
		return ""
	}
}
