package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "TANKS!")
	rl.SetTargetFPS(60)
	loadImageResources()
	// defer unloadResources()

	t := tank{
		centerX:            0,
		centerY:            0,
		radius:             8,
		sprites:            tankAtlaces["YELLOW_DEFAULT_TANK"],
		currentFrameNumber: 0,
	}

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyRight) {
			t.moveByVector(1, 0)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			t.moveByVector(-1, 0)

		}
		if rl.IsKeyDown(rl.KeyUp) {
			t.moveByVector(0, -1)

		}
		if rl.IsKeyDown(rl.KeyDown) {
			t.moveByVector(0, 1)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawTexture(tileAtlaces["GRASS"].atlas, 0, 0, rl.White)
		renderTank(&t)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
