package main

import rl "github.com/gen2brain/raylib-go/raylib"

type horizSpriteAtlas struct {
	atlas        rl.Texture2D
	totalSprites int32
	spriteSize   int   // width of square sprite
	totalFrames  uint8 // frames per sprite group
}

func (s *horizSpriteAtlas) getRectForSpriteFromAtlas(spriteNumber int) rl.Rectangle {
	spriteWidth := float32(s.atlas.Width / s.totalSprites)
	return rl.Rectangle{
		X:      spriteWidth * float32(spriteNumber),
		Y:      0,
		Width:  spriteWidth,
		Height: float32(s.atlas.Height),
	}
}
