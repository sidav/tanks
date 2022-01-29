package main

const (
	TBODY_PLAYER1 = iota
	TBODY_PLAYER2
	TBODY_TANK1
	TBODY_TANK2
	TBODY_TANK3
	TBODY_TANK4
	TBODY_TANK5
	TBODY_TANK6
	TBODY_TANK7
	TBODY_TANK8
	TBODY_PART1
)

type tankBodyStats struct {
	maxHp   int
	sprites *spriteAtlas
}

var tankBodyStatsList map[int]*tankBodyStats

func initTankBodyStatsList() {
	tankBodyStatsList = map[int]*tankBodyStats{
		/////// TANKS
		TBODY_PLAYER1: {
			sprites: tankAtlaces[TANK_T1],
			maxHp:   3,
		},
		TBODY_PLAYER2: {
			sprites: tankAtlaces[TANK_T5],
			maxHp:   3,
		},
		TBODY_TANK1: {
			sprites: tankAtlaces[TANK_T1],
			maxHp:   1,
		},
		TBODY_TANK2: {
			sprites: tankAtlaces[TANK_T2],
			maxHp:   2,
		},
		TBODY_TANK3: {
			sprites: tankAtlaces[TANK_T3],
			maxHp:   3,
		},
		TBODY_TANK4: {
			sprites: tankAtlaces[TANK_T4],
			maxHp:   5,
		},
		TBODY_TANK5: {
			sprites: tankAtlaces[TANK_T5],
			maxHp:   1,
		},
		TBODY_TANK6: {
			sprites: tankAtlaces[TANK_T6],
			maxHp:   1,
		},
		TBODY_TANK7: {
			sprites: tankAtlaces[TANK_T7],
			maxHp:   2,
		},
		TBODY_TANK8: {
			sprites: tankAtlaces[TANK_T8],
			maxHp:   4,
		},

		TBODY_PART1: {
			sprites: bodiesAtlaces[rnd.Rand(len(bodiesAtlaces))],
			maxHp:   1,
		},
	}
}
