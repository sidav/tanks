package main

import rl "github.com/gen2brain/raylib-go/raylib"

type tank struct {
	centerX, centerY   int
	faceX, faceY       int
	radius             int
	sprites            *horizSpriteAtlas
	currentFrameNumber uint8
}

func (t *tank) getCenterCoords() (int, int) {
	return t.centerX, t.centerY
}

func (t *tank) moveByVector(x, y int) {
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
	spriteNumber := int(spriteGroup*2 + (t.currentFrameNumber % 2))
	return t.sprites.getRectForSpriteFromAtlas(spriteNumber)
}
