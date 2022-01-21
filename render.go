package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var TINT = rl.RayWhite

func renderBattlefield(b *battlefield) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	renderTiles(b)

	renderTank(b.playerTank)
	// rl.DrawText(fmt.Sprintf("%d, %d", b.playerTank.centerX, b.playerTank.centerY),0,0, 32, color.RGBA{255, 255, 255, 255})
	for i := range b.enemies {
		renderTank(b.enemies[i])
	}

	rl.EndDrawing()
}

func renderTank(t *tank) {
	x, y := t.getTopLeftCoordForDraw()
	rl.DrawTextureRec(
		t.sprites.atlas,
		t.getCurrentSpriteRect(),
		rl.Vector2{
			X: x,
			Y: y,
		},
		TINT,
	)
}

func renderTiles(b *battlefield) {
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			if t.sprite != nil {
				rl.DrawTexture(t.sprite.atlas, int32(x*t.sprite.spriteSize), int32(y*t.sprite.spriteSize), TINT)
			}
		}
	}
}
