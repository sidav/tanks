package main

import rl "github.com/gen2brain/raylib-go/raylib"

type spriteAtlas struct {
	// first index is sprite number, second is frame number
	atlas        [][]rl.Texture2D
	totalSprites int32
	spriteSize   int   // width of square sprite
	totalFrames  uint8 // frames per sprite group
}
