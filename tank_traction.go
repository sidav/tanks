package main

const (
	TTRACTION_TRACKS = iota
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
	}
}
