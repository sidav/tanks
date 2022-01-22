package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"time"
)

func showGameMenu() {
	cursor := 0
	valuesData := [][]int{
		// {value, min, max, step}
		{0, 0, 25, 1},
		{10, 2, 100, 1},
		{30, 5, 1000, 5},
		{150, 5, 1000, 5},
		{4, 2, 4, 1},
	}
	entries := []string{
		"INITIAL ENEMIES      <%d>",
		"MAX TANKS ON MAP     <%d>",
		"TOTAL TANKS PER GAME <%d>",
		"HOW RARE ARE SPAWNS  <%d>",
		"FACTIONS             <%d>",
	}
	colorText := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	colorCursor := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	for {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{0, 0, 0, 255})

		for i := range entries {
			if i == cursor {
				rl.DrawText(fmt.Sprintf(entries[i], valuesData[i][0]), 0, int32(i*(TEXT_SIZE+TEXT_MARGIN)), TEXT_SIZE, colorCursor)
			} else {
				rl.DrawText(fmt.Sprintf(entries[i], valuesData[i][0]), 0, int32(i*(TEXT_SIZE+TEXT_MARGIN)), TEXT_SIZE, colorText)
			}
		}
		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeyUp) {
			cursor--
			if cursor < 0 {
				cursor = len(valuesData) - 1
			}

		}
		if rl.IsKeyDown(rl.KeyLeft) {
			valuesData[cursor][0] -= valuesData[cursor][3]
			if valuesData[cursor][0] < valuesData[cursor][1] {
				valuesData[cursor][0] = valuesData[cursor][1]
			}
			time.Sleep(150 * time.Millisecond)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			valuesData[cursor][0] += valuesData[cursor][3]
			if valuesData[cursor][0] > valuesData[cursor][2] {
				valuesData[cursor][0] = valuesData[cursor][2]
			}
			time.Sleep(150 * time.Millisecond)
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			cursor++
			if cursor == len(valuesData) {
				cursor = 0
			}
		}
		if rl.IsKeyPressed(rl.KeyEnter) {
			gameIsRunning = true
			break
		}
		if rl.IsKeyPressed(rl.KeyEscape) {
			gameIsRunning = false
			playerPressedExit = true
			break
		}

	}
	gameMap = &battlefield{
		initialEnemiesCount:               valuesData[0][0],
		MaxTanksOnMap:                     valuesData[1][0],
		totalTanksRemainingToSpawn:        valuesData[2][0],
		chanceToSpawnEnemyEachTickOneFrom: valuesData[3][0],
		numFactions:                       valuesData[4][0],
	}
	gameMap.init()
}
