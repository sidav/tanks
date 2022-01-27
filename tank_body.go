package main

const (
	TBODY_TANK1 = iota
	TBODY_TANK2
	TBODY_TANK3
	TBODY_TANK4
	TBODY_TANK5
	TBODY_TANK6
	TBODY_TANK7
	TBODY_TANK8
)

type tankBodyStats struct {
	maxHp   int
	sprites *spriteAtlas
}

var tankBodyStatsList map[int]*tankBodyStats

func initTankBodyStatsList() {
	tankBodyStatsList = map[int]*tankBodyStats{
		/////// TANKS
		TBODY_TANK1: {
			sprites: tankAtlaces[TANK_T1],
			maxHp:   1,
		},
		TBODY_TANK2: {
			sprites: tankAtlaces[TANK_T2],
		},
	}
}
