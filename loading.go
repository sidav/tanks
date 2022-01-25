package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"image"
	"image/png"
	"os"
	"time"
)

var (
	tankAtlaces       = map[int]*spriteAtlas{}
	tileAtlaces       = map[string]*spriteAtlas{}
	projectileAtlaces = map[int]*spriteAtlas{}
	effectAtlaces     = map[int]*spriteAtlas{}
)

func loadImageResources() {
	var leftXForTank float32 = 0
	tankAtlaces[TANK_T1] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*0, 16, true, 2)
	tankAtlaces[TANK_T2] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*1, 16, true, 2)
	tankAtlaces[TANK_T3] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*2, 16, true, 2)
	tankAtlaces[TANK_T4] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*3, 16, true, 2)
	tankAtlaces[TANK_T5] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*4, 16, true, 2)
	tankAtlaces[TANK_T6] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*5, 16, true, 2)
	tankAtlaces[TANK_T7] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*6, 16, true, 2)
	tankAtlaces[TANK_T8] = CreateAtlasFromFile("tanks.png", leftXForTank, 16*7, 16, true, 2)

	tileAtlaces["WALL"] = CreateAtlasFromFile("sprites.png", 16*0, 16*0, 16, false, 5)
	tileAtlaces["ARMORED_WALL"] = CreateAtlasFromFile("sprites.png", 16*0, 16*1, 16, false, 5)
	tileAtlaces["WATER"] = CreateAtlasFromFile("sprites.png", 16*0, 16*3, 16, false, 2)
	tileAtlaces["WOOD"] = CreateAtlasFromFile("sprites.png", 16*1, 16*2, 16, false, 1)
	tileAtlaces["ICE"] = CreateAtlasFromFile("sprites.png", 16*2, 16*2, 16, false, 1)
	tileAtlaces["HQ"] = CreateAtlasFromFile("sprites.png", 16*3, 16*2, 16, false, 1)

	projectileAtlaces[PROJ_BULLET] = CreateAtlasFromFile("projectiles.png", 0, 0, 8, true, 1)
	projectileAtlaces[PROJ_ROCKET] = CreateAtlasFromFile("projectiles.png", 0, 8, 8, true, 1)

	effectAtlaces[EFFECT_EXPLOSION] = CreateAtlasFromFile("sprites.png", 16*0, 16*6, 16, false, 1)
	effectAtlaces[EFFECT_BIG_EXPLOSION] = CreateAtlasFromFile("sprites.png", 16*3, 16*6, 16*2, false, 1)
	effectAtlaces[EFFECT_SPAWN] = CreateAtlasFromFile("sprites.png", 16*0, 16*4, 16, false, 1)
	// effectAtlaces["SPAWN"] = CreateAtlasFromFile("sprites.png", 16*16, 16*9, 16, 2, 7)
}

//func unloadResources() {
//	for k, v := range tankAtlaces {
//		debugWritef("Unload: ", k)
//		rl.UnloadTexture(v.atlas)
//	}
//}

func CreateAtlasFromFile(filename string, topleftx, toplefty, spriteSize float32, createAllDirections bool,
	totalFrames uint8) *spriteAtlas {

	file, _ := os.Open(filename)
	spritesImg, _ := png.Decode(file)
	defer file.Close()

	newAtlas := spriteAtlas{}
	newAtlas.atlas = make([][]rl.Texture2D, 0)

	sprites := 1
	if createAllDirections {
		sprites = 4
	}
	spriteSize = 19
	file.Seek(0, 0)
	spritesImg, _ = png.Decode(file)

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	//rl.DrawTexture(rl.LoadTextureFromImage(rl.NewImageFromImage(spritesImg)), 0, 0, DEFAULT_TINT)
	//
	//rl.DrawTexture(rl.LoadTextureFromImage(rl.NewImageFromImage(
	//	spritesImg.(*image.NRGBA).SubImage(
	//		image.Rect(0, 0, 10, 10),
	//	),
	//)), 32, 0, DEFAULT_TINT)
	//
	//rl.DrawTexture(rl.LoadTextureFromImage(rl.NewImageFromImage(
	//	spritesImg.(*image.NRGBA).SubImage(
	//		image.Rect(0, 0, 32, 16),
	//	),
	//)), 32, 16, DEFAULT_TINT)

	rl.DrawTexture(rl.LoadTextureFromImage(rl.NewImageFromImage(
		spritesImg.(*image.NRGBA).SubImage(
			image.Rect(16, 0, 32, 16),
		),
	)), 64, 32, DEFAULT_TINT)

	time.Sleep(1000 * time.Millisecond)
	rl.EndDrawing()

	for sprite := 0; sprite < sprites; sprite++ {
		newAtlas.atlas = append(newAtlas.atlas, make([]rl.Texture2D, 0))
		for frame := 0; frame < int(totalFrames); frame++ {
			subImg := spritesImg.(*image.NRGBA).SubImage(
				image.Rect(
					int(topleftx)+frame*int(spriteSize),
					int(toplefty),
					int(topleftx)+(frame+1)*int(spriteSize),
					int(toplefty)+int(spriteSize),
				),
			).(*image.NRGBA)
			rlImg := rl.NewImageFromImage(subImg)
			for r := 1; r <= sprite; r++ {
				rl.ImageRotateCCW(rlImg)
			}
			rl.ImageResizeNN(rlImg, int32(float32(spriteSize)*SPRITE_SCALE_FACTOR), int32(float32(spriteSize)*SPRITE_SCALE_FACTOR))
			newAtlas.atlas[sprite] = append(newAtlas.atlas[sprite], rl.LoadTextureFromImage(rlImg))
		}
	}

	newAtlas.totalFrames = totalFrames
	newAtlas.spriteSize = int(spriteSize * SPRITE_SCALE_FACTOR)
	return &newAtlas
}
