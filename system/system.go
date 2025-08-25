package system

import (
	"fmt"
	"math/rand"

	"github.com/elitracy/planets/planet"
	"github.com/google/uuid"
)

var systemNames = []string{
	"Andara",
	"Velorum",
	"Caelora",
	"Nytheris",
	"Altairis",
	"Eryndor",
	"Zorath",
	"Luminar",
	"Orvion",
	"Thalora",
	"Corvane",
	"Syralis",
	"Aequinox",
	"Veythra",
	"Solara",
	"Drakoris",
	"Lythera",
	"Astralis",
	"Cygnara",
	"Voralis",
	"Seraphel",
	"Arionis",
	"Talmaris",
	"Kalthera",
	"Orryxis",
}

type System struct {
	ID              uuid.UUID
	Name            string
	Planets         []planet.Planet
	SystemDistances map[uuid.UUID]int
}

// implement fmt.Stringer
func (s System) String() string {
	var output string
	output += fmt.Sprintf("> 🌌 %s (ID=%s)\n", s.Name, s.ID)

	for _, p := range s.Planets {
		output += "  | " + p.String() + "\n"
	}

	return output
}

func GenerateSystem() System {
	randSystemIdx := rand.Intn(len(systemNames))
	return System{uuid.New(), systemNames[randSystemIdx], []planet.Planet{}, make(map[uuid.UUID]int)}
}
