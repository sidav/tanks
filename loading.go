package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var (
	tankAtlaces = map[int]*horizSpriteAtlas{}
	tileAtlaces = map[string]*horizSpriteAtlas{}
	projectileAtlaces = map[int]*horizSpriteAtlas{}
	effectAtlaces = map[int]*horizSpriteAtlas{}
)

func loadImageResources() {
	var leftXForTank float32 = 128 // 0
	tankAtlaces[TANK_T1] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*0, 16, 8, 2)
	tankAtlaces[TANK_T2] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*1, 16, 8, 2)
	tankAtlaces[TANK_T3] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*2, 16, 8, 2)
	tankAtlaces[TANK_T4] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*3, 16, 8, 2)
	tankAtlaces[TANK_T5] = CreateHorizAtlasFromFile("sprites.png", leftXForTank, 16*4, 16, 8, 2)

	tileAtlaces["WALL"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*0, 16, 1, 1)
	tileAtlaces["ARMORED_WALL"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*1, 16, 1, 1)
	tileAtlaces["WOOD"] = CreateHorizAtlasFromFile("sprites.png", 16*17, 16*2, 16, 1, 1)
	tileAtlaces["WATER"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*3, 16, 1, 1)
	tileAtlaces["HQ"] = CreateHorizAtlasFromFile("sprites.png", 16*19, 16*2, 16, 1, 1)

	projectileAtlaces[PROJ_BULLET] = CreateHorizAtlasFromFile("sprites.png", 321, 100, 8, 4, 1)

	effectAtlaces[EFFECT_EXPLOSION] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*8, 16, 3, 7)
	effectAtlaces[EFFECT_SPAWN] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*6, 16, 4, 8)
	// effectAtlaces["SPAWN"] = CreateHorizAtlasFromFile("sprites.png", 16*16, 16*9, 16, 2, 7)
}

func unloadResources() {
	for k, v := range tankAtlaces {
		debugWritef("Unload: ", k)
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
