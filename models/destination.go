package models

import (
	"github.com/elitracy/planets/core"
)

type Destination struct {
	Position core.Position
	Entity   Entity
}
