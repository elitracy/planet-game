package models

type GameState struct {
	CurrentTick int
	StarSystems []*StarSystem
}

func (gs *GameState) CreatePlayer(location Location) Player {
	player := Player{Location: location}

	for _, s := range gs.StarSystems {
		for _, p := range s.Planets {
			if p.Location.Coordinates == location.Coordinates {

				p.Players = append(p.Players, &player)
			}
		}
	}

	return player
}
