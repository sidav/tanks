package main

import (
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
		centerTank := b.playerTanks[playerNumber]
		if !(playerNumber > 0 && r.doesLevelFitInScreenHorizontally()) {

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
		}
		r.renderUI(playerNumber, b, centerTank)
		rl.EndScissorMode()
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
		if r.AreOnScreenCoordsInViewport(osx, osy) {
			rl.DrawTexture(
				t.getSprite(),
				int32(osx),
				int32(osy),
				DEFAULT_TINT,
			)
		}
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

func (r *renderer) AreOnScreenCoordsInViewport(osx, osy int) bool {
	return osx >= 0 && osx < r.viewportW && osy >= 0 && osy < WINDOW_H
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
