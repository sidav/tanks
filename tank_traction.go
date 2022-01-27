package main

const (
	TTRACTION_TRACKS = iota
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
		TTRACTION_TRACKS: {
			sprites:   nil,
			speed:     1,
			moveDelay: 2,
		},

		TTRACTION_PART1: {
			sprites:   trackAtlaces[rnd.Rand(len(trackAtlaces))],
			speed:     1,
			moveDelay: 2,
		},
	}
}
