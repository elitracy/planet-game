package models

type Player struct {
	*CoreEntity
}

func NewPlayer(location Location) *Player {
	return &Player{
		CoreEntity: &CoreEntity{
			Name:     "Player0",
			Location: location,
		},
	}
}
