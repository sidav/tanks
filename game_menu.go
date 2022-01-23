package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"time"
)

var (
	menuPressDelay time.Duration = 100

	menuValuesData = [][]int{
		// {(default) value, min, max, step}
		{0, 0, 25, 1},
		{10, 2, 1000, 1},
		{30, 5, 10000, 10},
		{120, 5, 1000, 5},
		{4, 2, len(factionTints), 1},

		{25, 11, 1001, 2},
		{14, 11, 1000, 2},
		{MAP_H*MAP_W/3, 0, 100000, 4},
		{MAP_H*MAP_W/8, 0, 10000, 4},
		{MAP_H*MAP_W/12, 0, 10000, 4},
		{MAP_H*MAP_W/12, 0, 10000, 4},
		{MAP_H*MAP_W/20, 0, 10000, 4},
	}

	menuEntries = []string{
		"INITIAL ENEMIES       <%d>",
		"MAX TANKS ON MAP      <%d>",
		"TOTAL TANKS PER GAME  <%d>",
		"HOW RARE ARE SPAWNS   <%d>",
		"FACTIONS              <%d>",

		"MAP WIDTH             <%d>",
		"MAP HEIGHT            <%d>",
		"DESIRED WALLS         <%d>",
		"DESIRED ARMORED_WALLS <%d>",
		"DESIRED WOODS         <%d>",
		"DESIRED WATER         <%d>",
		"DESIRED ICE           <%d>",
	}
	colorText = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	colorCursor = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
)

func showGameMenu() {
	cursor := 0

	for {
		MAP_W = menuValuesData[5][0]
		MAP_H = menuValuesData[6][0]
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{0, 0, 0, 255})

		for i := range menuEntries {
			if i == cursor {
				rl.DrawText(fmt.Sprintf(menuEntries[i], menuValuesData[i][0]), 0, int32(i*(TEXT_SIZE+TEXT_MARGIN)), TEXT_SIZE, colorCursor)
			} else {
				rl.DrawText(fmt.Sprintf(menuEntries[i], menuValuesData[i][0]), 0, int32(i*(TEXT_SIZE+TEXT_MARGIN)), TEXT_SIZE, colorText)
			}
		}
		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeyUp) {
			cursor--
			if cursor < 0 {
				cursor = len(menuValuesData) - 1
			}

		}
		if rl.IsKeyDown(rl.KeySpace) {
			value := menuValuesData[cursor][0]
			if value == menuValuesData[cursor][1] {
				menuValuesData[cursor][0] = menuValuesData[cursor][2]
			} else if value == menuValuesData[cursor][2] {
				menuValuesData[cursor][0] = menuValuesData[cursor][1]
			} else {
				menuValuesData[cursor][0] = menuValuesData[cursor][2]
			}
			time.Sleep(menuPressDelay * time.Millisecond)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			menuValuesData[cursor][0] -= menuValuesData[cursor][3]
			if menuValuesData[cursor][0] < menuValuesData[cursor][1] {
				menuValuesData[cursor][0] = menuValuesData[cursor][1]
			}
			time.Sleep(menuPressDelay * time.Millisecond)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			menuValuesData[cursor][0] += menuValuesData[cursor][3]
			if menuValuesData[cursor][0] > menuValuesData[cursor][2] {
				menuValuesData[cursor][0] = menuValuesData[cursor][2]
			}
			time.Sleep(menuPressDelay * time.Millisecond)
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			cursor++
			if cursor == len(menuValuesData) {
				cursor = 0
			}
		}
		if rl.IsKeyPressed(rl.KeyEnter) {
			gameIsRunning = true
			gameWon = false
			gameOver = false
			break
		}
		if rl.IsKeyPressed(rl.KeyEscape) {
			gameIsRunning = false
			playerPressedExit = true
			break
		}

	}
	gameMap = &battlefield{
		initialEnemiesCount:               menuValuesData[0][0],
		maxTanksOnMap:                     menuValuesData[1][0],
		totalTanksRemainingToSpawn:        menuValuesData[2][0],
		chanceToSpawnEnemyEachTickOneFrom: menuValuesData[3][0],
		numFactions:                       menuValuesData[4][0],
	}
	gameMap.init(menuValuesData[7][0], menuValuesData[8][0], menuValuesData[9][0], menuValuesData[10][0], menuValuesData[11][0])
}
