package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var TINT = rl.RayWhite

var gameOverLineH int32 = -TILE_SIZE_IN_PIXELS
var gameOverRgb color.RGBA

func renderBattlefield(b *battlefield) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// rl.DrawText(fmt.Sprintf("%d, %d", b.playerTank.centerX, b.playerTank.centerY),0,0, 32, color.RGBA{255, 255, 255, 255})
	for i := range b.tanks {
		renderTank(b.tanks[i])
	}
	renderProjectiles(b)
	renderTiles(b)
	renderEffects(b)

	if b.playerTank == nil {
		rl.DrawText("GAME OVER.", WINDOW_W/3, gameOverLineH, TILE_SIZE_IN_PIXELS+4, gameOverRgb)
		gameOverLineH++
		gameOverRgb.A = 255
		gameOverRgb.R += uint8(rnd.Rand(2))
		gameOverRgb.G += uint8(rnd.Rand(2))
		gameOverRgb.B += uint8(rnd.Rand(2))
		if gameOverLineH > WINDOW_H {
			gameOverLineH = -TILE_SIZE_IN_PIXELS
		}
	}

	rl.EndDrawing()
}

func renderTank(t *tank) {
	cx, cy := ingameCoordsToOnScreenCoords(t.centerX, t.centerY)
	x, y := float32(cx - t.sprites.spriteSize/2), float32(cy - t.sprites.spriteSize/2)
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

func renderProjectiles(b *battlefield) {
	for _, p := range b.projectiles {
		renderTank(p)
	}
}

func renderEffects(b *battlefield) {
	for _, p := range b.effects {
		renderTank(p)
	}
}

func renderTiles(b *battlefield) {
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			spr := t.getSpritesAtlas()
			if spr != nil {
				rl.DrawTexture(spr.atlas, int32(x*spr.spriteSize), int32(y*spr.spriteSize), TINT)
			}
		}
	}
}

func ingameCoordsToOnScreenCoords(igx, igy int) (int, int) {
	return int(float32(igx)*PIXEL_TO_GAMECOORD_RATIO), int(float32(igy)*PIXEL_TO_GAMECOORD_RATIO)
}
