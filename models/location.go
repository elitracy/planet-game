package models

import (
	"fmt"

	"github.com/elitracy/planets/core"
)

type Location struct {
	Position core.Position
	Entity   Entity
}

func (l Location) String() string {
	output := l.Position.String()
	if l.Entity != nil {
		output = fmt.Sprintf("%v %v", l.Entity.GetName(), output)
	}
	return output
}
