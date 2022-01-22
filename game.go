package main

import rl "github.com/gen2brain/raylib-go/raylib"

var gameTick = 0

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
	renderBattlefield(gameMap)
	gameTick++
}
