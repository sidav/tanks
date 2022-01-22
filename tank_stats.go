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
	"GRAY_T1_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  6,
		shootDelay: 40,
	},
	"GREEN_T1_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  8,
		shootDelay: 50,
	},
	"RED_T1_TANK": {
		radius: TILE_PHYSICAL_SIZE/2-2,
		moveDelay:  3,
		shootDelay: 60,
	},
}
