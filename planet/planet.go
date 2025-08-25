package planet

import (
	"fmt"
	"math/rand"

	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
)

var planetNames = []string{
	"Lyra",
	"Astra",
	"Seris",
	"Elara",
	"Veyra",
	"Calix",
	"Altair",
	"Liora",
	"Orren",
	"Selos",
	"Talmar",
	"Eryne",
	"Ivara",
	"Lumora",
	"Myra",
	"Caelum",
	"Corvus",
	"Althea",
	"Nyra",
	"Evros",
	"Arden",
	"Olenya",
	"Sylos",
	"Cyrene",
	"Velora",
}

type Planet struct {
	ID         uuid.UUID
	Name       string
	Population int64
}

// implement fmt.Stringer
func (p Planet) String() string {
	return fmt.Sprintf("🌍 %s (ID=%s, Population=%s)", p.Name, p.ID, humanize.Comma(p.Population))
}

func GeneratePlanet() Planet {
	randPlanetIdx := rand.Intn(len(planetNames))
	pop := rand.Int63n(10_000_000_000-10_000) + 10_000
	return Planet{uuid.New(), planetNames[randPlanetIdx], pop}
}
