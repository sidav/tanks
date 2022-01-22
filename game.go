package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sidav/golibrl/random/additive_random"
)

var gameTick = 0
var rnd additive_random.FibRandom

func runGame() {
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

	gameMap.actForProjectiles()
	for i := range gameMap.enemies {
		gameMap.actAiForTank(gameMap.enemies[i])
	}
	renderBattlefield(gameMap)
	gameTick++
}
