package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var gameMap *battlefield

func main() {
	rl.InitWindow(MAP_W*TILE_SIZE_IN_PIXELS, MAP_H*TILE_SIZE_IN_PIXELS, "TANKS!")
	rl.SetTargetFPS(60)
	loadImageResources()
	// defer unloadResources()

	rnd.InitDefault()

	gameMap = &battlefield{
		initialEnemiesCount: 1,
		desiredEnemiesCount: 5,
		chanceToSpawnEnemyEachTickOneFrom: 250,
	}
	gameMap.init()

	for !rl.WindowShouldClose() {
		runGame()
	}

	rl.CloseWindow()
}
