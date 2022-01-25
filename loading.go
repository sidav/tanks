package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image"
	"image/png"
	"os"
)

var (
	tankAtlaces       = map[int]*spriteAtlas{}
	tileAtlaces       = map[string]*spriteAtlas{}
	projectileAtlaces = map[int]*spriteAtlas{}
	effectAtlaces     = map[int]*spriteAtlas{}
)

func loadImageResources() {
	var leftXForTank = 0
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

	effectAtlaces[EFFECT_EXPLOSION] = CreateAtlasFromFile("sprites.png", 16*0, 16*6, 16, false, 3)
	effectAtlaces[EFFECT_BIG_EXPLOSION] = CreateAtlasFromFile("sprites.png", 16*3, 16*6, 16*2, false, 2)
	effectAtlaces[EFFECT_SPAWN] = CreateAtlasFromFile("sprites.png", 16*0, 16*4, 16, false, 4)
	// effectAtlaces["SPAWN"] = CreateAtlasFromFile("sprites.png", 16*16, 16*9, 16, 2, 7)
}

//func unloadResources() {
//	for k, v := range tankAtlaces {
//		debugWritef("Unload: ", k)
//		rl.UnloadTexture(v.atlas)
//	}
//}

func extractSubimageFromImage(img image.Image, fromx, fromy, w, h int) image.Image {
	minx, miny := img.Bounds().Min.X, img.Bounds().Min.Y
	//maxx, maxy := img.Bounds().Min.X, img.Bounds().Max.Y
	subImg := img.(*image.NRGBA).SubImage(
		image.Rect(minx+fromx, miny+fromy, minx+fromx+w, miny+fromy+h),
	)
	// reset img bounds, because RayLib goes nuts about it otherwise
	subImg.(*image.NRGBA).Rect = image.Rect(0, 0, w, h)
	return subImg
}

func CreateAtlasFromFile(filename string, topleftx, toplefty, spriteSize int, createAllDirections bool,
	totalFrames int) *spriteAtlas {

	file, _ := os.Open(filename)
	img, _ := png.Decode(file)
	file.Close()

	newAtlas := spriteAtlas{
		spriteSize:   spriteSize*int(SPRITE_SCALE_FACTOR),
	}
	if createAllDirections {
		newAtlas.atlas = make([][]rl.Texture2D, 4)
	} else {
		newAtlas.atlas = make([][]rl.Texture2D, 1)
	}
	// newAtlas.atlas
	for currFrame := 0; currFrame < totalFrames; currFrame++ {
		currPic := extractSubimageFromImage(img, topleftx+currFrame*spriteSize, toplefty, spriteSize, spriteSize)
		rlImg := rl.NewImageFromImage(currPic)
		rl.ImageResizeNN(rlImg, int32(spriteSize)*int32(SPRITE_SCALE_FACTOR), int32(spriteSize)*int32(SPRITE_SCALE_FACTOR))
		newAtlas.atlas[0] = append(newAtlas.atlas[0], rl.LoadTextureFromImage(rlImg))
		if createAllDirections {
			for i := 1; i < 4; i++ {
				rl.ImageRotateCCW(rlImg)
				newAtlas.atlas[i] = append(newAtlas.atlas[i], rl.LoadTextureFromImage(rlImg))
			}
		}
	}

	return &newAtlas
}
