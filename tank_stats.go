package main

type tankStats struct {
	moveDelay  int
	shootDelay int
}

var tankStatsList = map[string]*tankStats {
	"PLAYER_TANK": {
		moveDelay:  2,
		shootDelay: 15,
	},
	"ENEMY_TANK": {
		moveDelay:  5,
		shootDelay: 25,
	},
}
