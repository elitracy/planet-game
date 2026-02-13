package models

import (
	"fmt"

	"github.com/elitracy/planets/engine"
)

type Location struct {
	Position engine.Position
	Entity   Entity
}

func (l Location) String() string {
	output := l.Position.String()
	if l.Entity != nil {
		output = fmt.Sprintf("%v %v", l.Entity.GetName(), output)
	}
	return output
}
