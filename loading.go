package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var (
	tankAtlaces = map[string]*horizSpriteAtlas{}
	tileAtlaces = map[string]*horizSpriteAtlas{}
	projectileAtlaces = map[string]*horizSpriteAtlas{}
	effectAtlaces = map[string]*horizSpriteAtlas{}
)

func loadImageResources() {
	var leftXForTank float32 = 128 // 0
	tankAtlaces["T1_TANK"] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*0, 16, 8, 2)
	tankAtlaces["T2_TANK"] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*1, 16, 8, 2)
	tankAtlaces["T3_TANK"] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*2, 16, 8, 2)
	tankAtlaces["T4_TANK"] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*3, 16, 8, 2)

	tileAtlaces["WALL"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*0, 16, 1, 1)
	tileAtlaces["ARMORED_WALL"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*1, 16, 1, 1)
	tileAtlaces["WOOD"] = CreateHorizAtlasFromFile("sprites.png", 16*17, 16*2, 16, 1, 1)
	tileAtlaces["WATER"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*3, 16, 1, 1)
	tileAtlaces["HQ"] = CreateHorizAtlasFromFile("sprites.png", 16*19, 16*2, 16, 1, 1)

	projectileAtlaces["BULLET"] = CreateHorizAtlasFromFile("sprites.png", 321, 100, 8, 4, 1)

	effectAtlaces["EXPLOSION"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*8, 16, 3, 7)
	effectAtlaces["SPAWN"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*6, 16, 4, 8)
	// effectAtlaces["SPAWN"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*9, 16, 2, 7)
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
	rl.ImageResizeNN(spritesImg, int32(float32(spritesImg.Width)*SPRITE_SCALE_FACTOR), int32(float32(spritesImg.Height)*SPRITE_SCALE_FACTOR))
	s.atlas = rl.LoadTextureFromImage(spritesImg)
	s.totalFrames = totalFrames
	s.totalSprites = int32(totalSprites)
	s.spriteSize = int(spriteSize* SPRITE_SCALE_FACTOR)
	return &s
}
