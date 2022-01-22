package main

import (
	"fmt"
	"github.com/sidav/golibrl/random/additive_random"
)

var gameTick = 0
var rnd additive_random.FibRandom

func runGame() {
	fmt.Printf("%d: START ", gameTick)
	listenPlayerControls()

	fmt.Printf("PROJS ")
	gameMap.actForProjectiles()

	fmt.Printf("EFFECTS ")
	gameMap.actForEffects()

	fmt.Printf("AI {")
	for i := range gameMap.tanks {
		if gameMap.tanks[i] == gameMap.playerTank {
			continue
		}
		gameMap.actAiForTank(gameMap.tanks[i])
	}
	fmt.Printf("} ")

	fmt.Printf("RENDER ")
	renderBattlefield(gameMap)

	fmt.Printf("SPAWN ")
	if len(gameMap.tanks) < gameMap.desiredEnemiesCount && rnd.OneChanceFrom(gameMap.chanceToSpawnEnemyEachTickOneFrom) {
		gameMap.spawnTank(0, MAP_W-1, 0, MAP_H-MAP_H/3)
	}

	fmt.Printf("TURN FINISHED. \n")
	gameTick++
}
