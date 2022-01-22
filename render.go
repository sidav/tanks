package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var DEFAULT_TINT = rl.RayWhite

var gameOverLineH int32 = -TILE_SIZE_IN_PIXELS
var gameOverRgb color.RGBA

func renderBattlefield(b *battlefield) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	renderTiles(b)
	for i := range b.tanks {
		renderTank(b.tanks[i], true)
	}
	renderProjectiles(b)
	renderEffects(b)
	renderWood(b)

	if gameOver {
		rl.DrawText("GAME OVER.", int32(WINDOW_W)/3, gameOverLineH, TILE_SIZE_IN_PIXELS+4, gameOverRgb)
		rl.DrawText("Press ESC for menu", int32(WINDOW_W)/4, gameOverLineH+TILE_SIZE_IN_PIXELS+4, TILE_SIZE_IN_PIXELS+4, gameOverRgb)
		gameOverLineH++
		gameOverRgb.A = 255
		gameOverRgb.R += uint8(rnd.Rand(2))
		gameOverRgb.G += uint8(rnd.Rand(2))
		gameOverRgb.B += uint8(rnd.Rand(2))
		if gameOverLineH > WINDOW_H {
			gameOverLineH = -TILE_SIZE_IN_PIXELS
		}
	}

	if gameWon {
		rl.DrawText("YOU WON!", int32(WINDOW_W)/3, gameOverLineH, TILE_SIZE_IN_PIXELS+4, gameOverRgb)
		gameOverLineH++
		gameOverRgb.A = 255
		gameOverRgb.R += uint8(rnd.Rand(2))
		gameOverRgb.G += uint8(rnd.Rand(2))
		gameOverRgb.B += uint8(rnd.Rand(2))
		if gameOverLineH > WINDOW_H {
			gameOverLineH = -TILE_SIZE_IN_PIXELS
		}
	}

	rl.DrawText(fmt.Sprintf("Remaining tanks %d", b.totalTanksRemainingToSpawn), 0, TILE_SIZE_IN_PIXELS*MAP_H, TEXT_SIZE, color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})

	rl.EndDrawing()
}

func renderTank(t *tank, useFactionTint bool) {
	cx, cy := ingameCoordsToOnScreenCoords(t.centerX, t.centerY)
	x, y := float32(cx - t.getSpritesAtlas().spriteSize/2), float32(cy - t.getSpritesAtlas().spriteSize/2)
	if useFactionTint {
		rl.DrawTextureRec(
			t.getSpritesAtlas().atlas,
			t.getCurrentSpriteRect(),
			rl.Vector2{
				X: x,
				Y: y,
			},
			factionTints[t.faction],
		)
	} else {
		rl.DrawTextureRec(
			t.getSpritesAtlas().atlas,
			t.getCurrentSpriteRect(),
			rl.Vector2{
				X: x,
				Y: y,
			},
			DEFAULT_TINT,
		)
	}
}

func renderProjectiles(b *battlefield) {
	for _, p := range b.projectiles {
		renderTank(p, false)
	}
}

func renderEffects(b *battlefield) {
	for _, p := range b.effects {
		renderTank(p, false)
	}
}

func renderTiles(b *battlefield) {
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			spr := t.getSpritesAtlas()
			if spr != nil {
				rl.DrawTextureRec(
					t.getSpritesAtlas().atlas,
					t.getSpriteRect(),
					rl.Vector2{
						X: float32(x*spr.spriteSize),
						Y: float32(y*spr.spriteSize),
					},
					DEFAULT_TINT,
				)
			}
		}
	}
}

func renderWood(b *battlefield) {
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			spr := t.getSpritesAtlas()
			if t.code == TILE_WOOD {
				rl.DrawTexture(spr.atlas, int32(x*spr.spriteSize), int32(y*spr.spriteSize), DEFAULT_TINT)
			}
		}
	}
}

func ingameCoordsToOnScreenCoords(igx, igy int) (int, int) {
	return int(float32(igx)* PIXEL_TO_PHYSICAL_RATIO), int(float32(igy)* PIXEL_TO_PHYSICAL_RATIO)
}
