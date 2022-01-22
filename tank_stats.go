package main

type tankStats struct {
	radius int

	moveDelay  int
	shootDelay int
}

var tankStatsList = map[string]*tankStats{
	"BULLET": {
		radius: 2,
	},

	"EXPLOSION": {
		radius: TILE_PHYSICAL_SIZE/2-1,
		moveDelay: 3,
	},
	"SPAWN": {
		radius: TILE_PHYSICAL_SIZE/2-1,
		moveDelay: 7,
	},

	"PLAYER_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  2,
		shootDelay: 25,
	},
	"T1_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  6,
		shootDelay: 40,
	},
	"T2_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  8,
		shootDelay: 50,
	},
	"T3_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  3,
		shootDelay: 60,
	},
	"T4_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  3,
		shootDelay: 60,
	},
}
