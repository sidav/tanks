package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var (
	tankAtlaces = map[string]*horizSpriteAtlas{}
	tileAtlaces = map[string]*horizSpriteAtlas{}
)

const resizeCoeff = 2.0

func loadImageResources() {
	tankAtlaces["YELLOW_DEFAULT_TANK"] = CreateHorizAtlasFromFile("sprites.png", 0, 0, 16, 8, 2)

	tileAtlaces["GRASS"] = CreateHorizAtlasFromFile("sprites.png", 17*16, 32, 16, 1, 1)
}

func unloadResources() {
	for k, v := range tankAtlaces {
		fmt.Println("Unload: " + k)
		rl.UnloadTexture(v.atlas)
	}
}

func CreateHorizAtlasFromFile(filename string, topleftx, toplefty, spriteSize, totalSprites float32,
	totalFrames uint8) *horizSpriteAtlas {

	s := horizSpriteAtlas{}
	spritesImg := rl.LoadImage(filename)
	rl.ImageColorReplace(spritesImg,
		color.RGBA{
			R: 0,
			G: 0,
			B: 1,
			A: 255,
		},
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		})
	rl.ImageCrop(spritesImg, rl.Rectangle{
		X:      topleftx,
		Y:      toplefty,
		Width:  spriteSize * totalSprites,
		Height: spriteSize,
	})
	rl.ImageResizeNN(spritesImg, int32(float32(spritesImg.Width)*resizeCoeff), int32(float32(spritesImg.Height)*resizeCoeff))
	s.atlas = rl.LoadTextureFromImage(spritesImg)
	s.totalFrames = totalFrames
	s.totalSprites = int32(totalSprites)
	return &s
}
