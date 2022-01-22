package main

import (
	"fmt"
	"github.com/sidav/golibrl/random/additive_random"
)

var gameTick = 0
var gameWon = false
var gameOver = false
var rnd additive_random.FibRandom

func runGame() {
	fmt.Printf("%d: START ", gameTick)
	listenPlayerInput()

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
	if len(gameMap.tanks) < gameMap.maxTanksOnMap && rnd.OneChanceFrom(gameMap.chanceToSpawnEnemyEachTickOneFrom) && gameMap.totalTanksRemainingToSpawn > 0 {
		gameMap.spawnTank(0, MAP_W-1, 0, MAP_H-MAP_H/3)
	}


	gameOver = gameMap.playerTank == nil
	if gameOver {
		gameMap.maxTanksOnMap = MAP_W*MAP_H/2
		gameMap.totalTanksRemainingToSpawn = MAP_W*MAP_H/2
		gameMap.chanceToSpawnEnemyEachTickOneFrom = 5
	}
	gameWon = gameWon || gameMap.totalTanksRemainingToSpawn <= 0 && len(gameMap.tanks) == 1 && len(gameMap.effects) == 0
	if gameWon {
		for i := 0; i < 5; i++ {
			gameMap.spawnEffect("EXPLOSION", rnd.Rand(WINDOW_W), rnd.Rand(WINDOW_H), nil)
		}
	}

	fmt.Printf("TURN FINISHED. \n")
	gameTick++
}
