package models

type Action interface {
	Execute(state *GameState, order *Order) error
	Describe() string
}
