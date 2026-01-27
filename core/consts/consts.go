package consts

import "time"

const (
	FOOD_PER_PERSON             = 3
	POPULATION_GROWTH_PER_PULSE = 1000
	NUM_DAYS_FED                = 10

	//time
	TICKS_PER_SECOND_UI = 2
	TICK_SLEEP_UI       = time.Second / TICKS_PER_SECOND_UI
	TICKS_PER_SECOND    = 100
	TICK_SLEEP          = time.Second / TICKS_PER_SECOND
	TICKS_PER_PULSE     = 1_440
	TICKS_PER_CYCLE     = TICKS_PER_PULSE * PULSES_PER_CYCLE
	PULSES_PER_CYCLE    = 365

	//ships
	SCOUT_VELOCITY = 500
)
