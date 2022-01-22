package main

type tankStats struct {
	moveDelay  int
	shootDelay int
}

var tankStatsList = map[string]*tankStats {
	"EXPLOSION": {
		moveDelay: 3,
	},
	"SPAWN": {
		moveDelay: 4,
	},

	"PLAYER_TANK": {
		moveDelay:  2,
		shootDelay: 25,
	},
	"GRAY_T1_TANK": {
		moveDelay:  6,
		shootDelay: 40,
	},
	"GREEN_T1_TANK": {
		moveDelay:  8,
		shootDelay: 50,
	},
	"RED_T1_TANK": {
		moveDelay:  3,
		shootDelay: 60,
	},
}
