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
			r.cameraCenterX = int(float64(r.cameraCenterX) * PIXEL_TO_PHYSICAL_RATIO)
			r.cameraCenterY = int(float64(r.cameraCenterY) * PIXEL_TO_PHYSICAL_RATIO)
		}

		r.renderTiles(b)
		for i := range b.tanks {
			r.renderTank(b.tanks[i], true)
		}
		r.renderProjectiles(b)
		r.renderEvents(b)
		r.renderWood(b)
		r.renderLevelOutline()
		rl.EndScissorMode()
	}

	if b.numPlayers > 1 && !r.doesLevelFitInScreenHorizontally() {
		separatorWidth := int32(TILE_PHYSICAL_SIZE / 2)
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
	spriteSize := tankBodyStatsList[t.getStats().bodyCode].sprites.spriteSize
	x, y := float32(cx-spriteSize/2), float32(cy-spriteSize/2)
	if tankTractionStatsList[t.getStats().tractionCode].sprites != nil {
		rl.DrawTexture(
			tankTractionStatsList[t.getStats().tractionCode].sprites.getSpriteByDirectionAndFrameNumber(t.faceX, t.faceY, t.currentFrameNumber),
			int32(x),
			int32(y),
			factionTints[t.faction],
		)
	}
	rl.DrawTexture(
		tankBodyStatsList[t.getStats().bodyCode].sprites.getSpriteByDirectionAndFrameNumber(t.faceX, t.faceY, t.currentFrameNumber),
		int32(x),
		int32(y),
		factionTints[t.faction],
	)
	if tankWeaponStatsList[t.weapons[0].code].sprites != nil {
		rl.DrawTexture(
			tankWeaponStatsList[t.weapons[t.currentWeaponNumber].code].sprites.getSpriteByDirectionAndFrameNumber(t.faceX, t.faceY, t.currentFrameNumber),
			int32(x),
			int32(y),
			factionTints[t.faction],
		)
	}
}

func (r *renderer) renderProjectiles(b *battlefield) {
	for _, p := range b.projectiles {
		sprites := p.getStats().sprites
		if sprites != nil {
			cx, cy := r.physicalToOnScreenCoords(p.centerX, p.centerY)
			x, y := float32(cx-sprites.spriteSize/2), float32(cy-sprites.spriteSize/2)
			rl.DrawTexture(
				sprites.getSpriteByDirectionAndFrameNumber(p.faceX, p.faceY, p.currentFrameNumber),
				int32(x),
				int32(y),
				DEFAULT_TINT,
			)
		}
	}
}

func (r *renderer) renderEvents(b *battlefield) {
	for _, p := range b.bonuses {
		sprites := p.getStats().sprites
		if sprites != nil {
			cx, cy := r.physicalToOnScreenCoords(p.centerX, p.centerY)
			x, y := float32(cx-sprites.spriteSize/2), float32(cy-sprites.spriteSize/2)
			rl.DrawTexture(
				sprites.getSpriteByDirectionAndFrameNumber(0, 0, p.currentFrameNumber),
				int32(x),
				int32(y),
				DEFAULT_TINT,
			)
		}
	}
	for _, p := range b.effects {
		sprites := p.getStats().sprites
		if sprites != nil {
			cx, cy := r.physicalToOnScreenCoords(p.centerX, p.centerY)
			x, y := float32(cx-sprites.spriteSize/2), float32(cy-sprites.spriteSize/2)
			rl.DrawTexture(
				sprites.getSpriteByDirectionAndFrameNumber(0, 0, p.currentFrameNumber),
				int32(x),
				int32(y),
				DEFAULT_TINT,
			)
		}
	}
}

func (r *renderer) renderTile(b *battlefield, x, y int) {
	t := b.tiles[x][y]
	spr := t.getSpritesAtlas()
	if spr != nil {
		osx, osy := r.physicalToOnScreenCoords(x*TILE_PHYSICAL_SIZE, y*TILE_PHYSICAL_SIZE)
		rl.DrawTexture(
			t.getSprite(),
			int32(osx),
			int32(osy),
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

func (r *renderer) renderLevelOutline() {
	if r.doesLevelFitInScreenHorizontally() && r.doesLevelFitInScreenVertically() {
		return
	}
	const thickness = 5
	x, y := r.physicalToOnScreenCoords(0, 0)
	rl.DrawRectangleLinesEx(
		rl.Rectangle{
			X:      float32(x),
			Y:      float32(y),
			Width:  float32(TILE_SIZE_IN_PIXELS * MAP_W),
			Height: float32(TILE_SIZE_IN_PIXELS * MAP_H),
		},
		thickness,
		color.RGBA{
			R: 255,
			G: 96,
			B: 96,
			A: 255,
		},
	)
}

func (r *renderer) physicalToOnScreenCoords(physX, physY int) (int, int) {
	pixx, pixy := r.physicalToPixelCoords(physX, physY)
	if !r.doesLevelFitInScreenHorizontally() {
		if r.cameraCenterX > MAP_W*TILE_SIZE_IN_PIXELS-r.viewportW/2 {
			pixx = pixx - MAP_W*TILE_SIZE_IN_PIXELS + r.viewportW
		} else if r.cameraCenterX > r.viewportW/2 {
			pixx = pixx - r.cameraCenterX + r.viewportW/2
		}
	}
	if !r.doesLevelFitInScreenVertically() {
		if r.cameraCenterY > MAP_H*TILE_SIZE_IN_PIXELS-WINDOW_H/2 {
			pixy = pixy - MAP_H*TILE_SIZE_IN_PIXELS + WINDOW_H
		} else if r.cameraCenterY > WINDOW_H/2 {
			pixy = pixy - r.cameraCenterY + WINDOW_H/2
		}
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
