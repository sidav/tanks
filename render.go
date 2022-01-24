package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var DEFAULT_TINT = rl.RayWhite

var gameOverLineH int32 = -TILE_SIZE_IN_PIXELS
var gameOverRgb color.RGBA

type renderer struct {
	cameraCenterX, cameraCenterY                int
	viewportW                                   int
	verticalViewportOffset, horizViewportOffset int // for split view
}

func (r *renderer) renderBattlefield(b *battlefield) {
	r.viewportW = WINDOW_W / b.numPlayers
	if r.doesLevelFitInScreenHorizontally() {
		r.viewportW = WINDOW_W
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for playerNumber := 0; playerNumber < b.numPlayers; playerNumber++ {
		if playerNumber > 0 && r.doesLevelFitInScreenHorizontally() {
			break
		}
		centerTank := b.playerTanks[playerNumber]

		rl.BeginScissorMode(int32(playerNumber*r.viewportW), 0, int32(r.viewportW), WINDOW_H)
		r.horizViewportOffset = r.viewportW * playerNumber

		if centerTank != nil {
			r.cameraCenterX, r.cameraCenterY = centerTank.getCenterCoords()
			r.cameraCenterX *= PIXEL_TO_PHYSICAL_RATIO
			r.cameraCenterY *= PIXEL_TO_PHYSICAL_RATIO
		}

		r.renderTiles(b)
		for i := range b.tanks {
			r.renderTank(b.tanks[i], true)
		}
		r.renderProjectiles(b)
		r.renderEffects(b)
		r.renderWood(b)
		rl.EndScissorMode()
	}

	if b.numPlayers > 1 && !r.doesLevelFitInScreenHorizontally() {
		separatorWidth := int32(TILE_PHYSICAL_SIZE/2)
		rl.DrawRectangleGradientV(WINDOW_W/2-separatorWidth/2, 0, separatorWidth, WINDOW_H-(TEXT_MARGIN*2+TEXT_SIZE),
			color.RGBA{
				R: 32,
				G: 32,
				B: 32,
				A: 255,
			},
			color.RGBA{
				R: 64,
				G: 64,
				B: 64,
				A: 255,
			},
		)
	}

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

	rl.DrawRectangleGradientV(0, WINDOW_H-2*TEXT_MARGIN-TEXT_SIZE, WINDOW_W, TEXT_MARGIN*2+TEXT_SIZE,
		color.RGBA{
			R: 64,
			G: 64,
			B: 64,
			A: 255,
		},
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 64,
		},
	)
	if !gameOver {
		rl.DrawText(fmt.Sprintf("Remaining tanks %d", b.totalTanksRemainingToSpawn), 0,
			WINDOW_H-TEXT_MARGIN-TEXT_SIZE, TEXT_SIZE,
			color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
	}

	rl.EndDrawing()
}

func (r *renderer) renderTank(t *tank, useFactionTint bool) {
	cx, cy := r.physicalToOnScreenCoords(t.centerX, t.centerY)
	x, y := float32(cx-t.getSpritesAtlas().spriteSize/2), float32(cy-t.getSpritesAtlas().spriteSize/2)
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

func (r *renderer) renderProjectiles(b *battlefield) {
	for _, p := range b.projectiles {
		r.renderTank(p, false)
	}
}

func (r *renderer) renderEffects(b *battlefield) {
	for _, p := range b.effects {
		r.renderTank(p, false)
	}
}

func (r *renderer) renderTile(b *battlefield, x, y int) {
	t := b.tiles[x][y]
	spr := t.getSpritesAtlas()
	if spr != nil {
		osx, osy := r.physicalToOnScreenCoords(x*TILE_PHYSICAL_SIZE, y*TILE_PHYSICAL_SIZE)
		rl.DrawTextureRec(
			t.getSpritesAtlas().atlas,
			t.getSpriteRect(),
			rl.Vector2{
				X: float32(osx),
				Y: float32(osy),
			},
			DEFAULT_TINT,
		)
	}
}

func (r *renderer) renderTiles(b *battlefield) {
	for x := range b.tiles {
		for y := range b.tiles[x] {
			r.renderTile(b, x, y)
		}
	}
}

func (r *renderer) renderWood(b *battlefield) {
	for x := range b.tiles {
		for y, t := range b.tiles[x] {
			if t.code == TILE_WOOD {
				r.renderTile(b, x, y)
			}
		}
	}
}

func (r *renderer) physicalToOnScreenCoords(physX, physY int) (int, int) {
	pixx, pixy := r.physicalToPixelCoords(physX, physY)
	if !r.doesLevelFitInScreenHorizontally() {
		pixx = pixx - r.cameraCenterX + r.viewportW/2
	}
	if !r.doesLevelFitInScreenVertically() {
		pixy = pixy - r.cameraCenterY + WINDOW_H/2
	}
	return pixx + r.horizViewportOffset, pixy
}

func (r *renderer) physicalToPixelCoords(px, py int) (int, int) {
	return int(float32(px) * PIXEL_TO_PHYSICAL_RATIO), int(float32(py) * PIXEL_TO_PHYSICAL_RATIO)
}

func (r *renderer) doesLevelFitInScreenHorizontally() bool {
	return MAP_W*TILE_SIZE_IN_PIXELS <= WINDOW_W
}

func (r *renderer) doesLevelFitInScreenVertically() bool {
	return MAP_H*TILE_SIZE_IN_PIXELS <= WINDOW_H
}
