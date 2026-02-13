package engine

import "time"

const (
	TICKS_PER_SECOND = 100 * 10
	TICK_SLEEP       = time.Second / TICKS_PER_SECOND
)
