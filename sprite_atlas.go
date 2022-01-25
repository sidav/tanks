package main

import rl "github.com/gen2brain/raylib-go/raylib"

type spriteAtlas struct {
	// first index is sprite number, second is frame number
	atlas        [][]rl.Texture2D
	spriteSize   int   // width of square sprite
}

func (sa *spriteAtlas) totalFrames() uint8 {
	if len(sa.atlas) > 0 {
		return uint8(len(sa.atlas[0]))
	}
	return 0
}
