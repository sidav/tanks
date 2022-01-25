package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type tank struct {
	playerControlled bool

	centerX, centerY   int
	faceX, faceY       int
	currentFrameNumber uint8
	faction            int

	owner          *tank
	markedToRemove bool
	code           int

	nextTickToMove, nextTickToShoot int

	ai *tankAi
}

func (t *tank) getCenterCoords() (int, int) {
	return t.centerX, t.centerY
}

func (t *tank) isAtCenterOfTile() bool {
	xInTile := t.centerX % TILE_PHYSICAL_SIZE
	yInTile := t.centerY % TILE_PHYSICAL_SIZE
	precision := TILE_PHYSICAL_SIZE / 10
	if abs(halfPhysicalTileSize()-xInTile) <= precision && abs(halfPhysicalTileSize()-yInTile) <= precision {
		return true
	}
	return false
}

func (t *tank) getRadius() int {
	return tankStatsList[t.code].radius
}

func (t *tank) canMoveNow() bool {
	return gameTick >= t.nextTickToMove
}

func (t *tank) canShootNow() bool {
	return gameTick >= t.nextTickToShoot
}

func (t *tank) getStats() *tankStats {
	return tankStatsList[t.code]
}

func (t *tank) getSpritesAtlas() *spriteAtlas {
	//debugWritef("ATLAS{%v}", t.code)
	return t.getStats().sprites
}

func (t *tank) getCurrentSprite() rl.Texture2D {
	var spriteGroup uint8 = 0
	if t.faceX == 1 {
		spriteGroup = 3
	}
	if t.faceX == -1 {
		spriteGroup = 1
	}
	if t.faceY == 1 {
		spriteGroup = 2
	}
	// spriteNumber := int(spriteGroup*t.getSpritesAtlas().totalFrames + (t.currentFrameNumber % t.getSpritesAtlas().totalFrames))
	return t.getSpritesAtlas().atlas[spriteGroup][t.currentFrameNumber]
}
