package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "TANKS!")
	rl.SetTargetFPS(60)
	loadImageResources()
	// defer unloadResources()

	gameMap := &battlefield{}
	gameMap.init()
	for !rl.WindowShouldClose() {
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

		renderBattlefield(gameMap)
	}

	rl.CloseWindow()
}
