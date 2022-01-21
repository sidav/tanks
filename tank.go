package main

import rl "github.com/gen2brain/raylib-go/raylib"

type tank struct {
	centerX, centerY   int
	faceX, faceY       int
	radius             int
	sprites            *horizSpriteAtlas
	currentFrameNumber uint8
}

func (t *tank) moveByVector(x, y int) {
	t.centerX += x
	t.centerY += y
	t.faceX = x
	t.faceY = y
	t.currentFrameNumber = (t.currentFrameNumber + 1) % t.sprites.totalFrames
}

func (t *tank) getTopLeftCoordForDraw() (float32, float32) {
	return float32(t.centerX - t.sprites.spriteSize/2), float32(t.centerY - t.sprites.spriteSize/2)
}

func (a *tank) getCurrentSpriteRect() rl.Rectangle {
	var spriteGroup uint8 = 0
	if a.faceX == 1 {
		spriteGroup = 3
	}
	if a.faceX == -1 {
		spriteGroup = 1
	}
	if a.faceY == 1 {
		spriteGroup = 2
	}
	spriteNumber := int(spriteGroup*2 + (a.currentFrameNumber % 2))
	return a.sprites.getRectForSpriteFromAtlas(spriteNumber)
}
