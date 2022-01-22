package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var gameMap *battlefield

func main() {
	rl.InitWindow(WINDOW_W, WINDOW_H, "TANKS!")
	rl.SetTargetFPS(60)
	loadImageResources()
	initTileDictionary()
	// defer unloadResources()

	rnd.InitDefault()

	gameMap = &battlefield{
		initialEnemiesCount: 0,
		desiredEnemiesCount: 10,
		chanceToSpawnEnemyEachTickOneFrom: 150,
	}
	gameMap.init()

	for !rl.WindowShouldClose() {
		runGame()
	}

	rl.CloseWindow()
}
