package main

const (
	TTRACTION_DEFAULT_SLOWEST = iota
	TTRACTION_DEFAULT_SLOW
	TTRACTION_DEFAULT_MEDIUM
	TTRACTION_DEFAULT_FAST
	TTRACTION_DEFAULT_FASTEST

	TTRACTION_PART1
)

type tankTraction struct {
	code int
}

type tankTractionStats struct {
	sprites          *spriteAtlas
	moveDelay, speed int
}

var tankTractionStatsList map[int]*tankTractionStats

func initTankTractionStatsList() {
	tankTractionStatsList = map[int]*tankTractionStats{
		/////// TANKS
		TTRACTION_DEFAULT_SLOWEST: {
			sprites:   nil,
			speed:     1,
			moveDelay: 10,
		},
		TTRACTION_DEFAULT_SLOW: {
			sprites:   nil,
			speed:     1,
			moveDelay: 3,
		},
		TTRACTION_DEFAULT_MEDIUM: {
			sprites:   nil,
			speed:     2,
			moveDelay: 2,
		},
		TTRACTION_DEFAULT_FAST: {
			sprites:   nil,
			speed:     4,
			moveDelay: 3,
		},
		TTRACTION_DEFAULT_FASTEST: {
			sprites:   nil,
			speed:     3,
			moveDelay: 2,
		},

		TTRACTION_PART1: {
			sprites:   trackAtlaces[rnd.Rand(len(trackAtlaces))],
			speed:     1,
			moveDelay: 2,
		},
	}
}
