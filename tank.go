package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type tank struct {
	centerX, centerY   int
	faceX, faceY       int
	sprites            *horizSpriteAtlas
	currentFrameNumber uint8
	stats              *tankStats
	faction            int

	owner          *tank
	markedToRemove bool
	code           string

	nextTickToMove, nextTickToShoot int

	ai *tankAi
}

func (t *tank) getCenterCoords() (int, int) {
	return t.centerX, t.centerY
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

func (t *tank) moveByVector(x, y int) {
	t.nextTickToMove = gameTick + t.stats.moveDelay
	tx, ty := trueCoordsToTileCoords(t.centerX, t.centerY)
	if gameMap.tiles[tx][ty].code == TILE_WATER {
		t.nextTickToMove += t.stats.moveDelay
	}
	if gameMap.canTankMoveByVector(t, x, y) {
		t.centerX += x
		t.centerY += y
	}
	t.faceX = x
	t.faceY = y
	t.currentFrameNumber = (t.currentFrameNumber + 1) % t.sprites.totalFrames
}

func (t *tank) getCurrentSpriteRect() rl.Rectangle {
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
	spriteNumber := int(spriteGroup*t.sprites.totalFrames + (t.currentFrameNumber % t.sprites.totalFrames))
	return t.sprites.getRectForSpriteFromAtlas(spriteNumber)
}
