package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var gameMap *battlefield
var gameIsRunning, playerPressedExit bool

func main() {
	rl.InitWindow(WINDOW_W, WINDOW_H, "TANKS!")
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyF12)

	rnd.InitDefault()

	loadImageResources()

	initEventsStatsList()
	initProjectileStatsList()
	initTankStatsList()
	initTankBodyStatsList()
	initTankWeaponStatsList()
	initTankTractionStatsList()
	initTileDictionary()

	// generateSpriteSheetFromParts()
	// defer unloadResources()

	gameIsRunning = true

	for !rl.WindowShouldClose() {
		showGameMenu()
		if playerPressedExit {
			break
		}
		for gameIsRunning {
			runGame()
		}
	}

	rl.CloseWindow()
}
