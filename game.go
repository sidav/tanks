package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sidav/golibrl/random/additive_random"
)

var gameTick = 0
var rnd additive_random.FibRandom

func runGame() {
	fmt.Printf("%d: START ", gameTick)
	if gameMap.playerTank.canMoveNow() {
		if rl.IsKeyDown(rl.KeyRight) {
			gameMap.playerTank.moveByVector(1, 0)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			gameMap.playerTank.moveByVector(-1, 0)
		}
		if rl.IsKeyDown(rl.KeyUp) {
			gameMap.playerTank.moveByVector(0, -1)
		}
		if rl.IsKeyDown(rl.KeyDown) {
			gameMap.playerTank.moveByVector(0, 1)
		}
	}
	if gameMap.playerTank.canShootNow() {
		if rl.IsKeyDown(rl.KeySpace) {
			gameMap.shootAsTank(gameMap.playerTank)
		}
	}

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
		gameMap.spawnTank(0, MAP_W-1, 0, MAP_H-1)
	}

	fmt.Printf("TURN FINISHED. \n")
	gameTick++
}
