package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(MAP_SIZE*TILE_SIZE_IN_PIXELS, MAP_SIZE*TILE_SIZE_IN_PIXELS, "TANKS!")
	rl.SetTargetFPS(60)
	loadImageResources()
	// defer unloadResources()

	gameMap := &battlefield{}
	gameMap.init()
	for !rl.WindowShouldClose() {
		runGame(gameMap)
	}

	rl.CloseWindow()
}
