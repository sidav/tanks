package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image"
	"image/draw"
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
	tankAtlaces[TANK_BOSS] = createAtlasFromRandomGenerated()

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

func generateSpriteSheetFromParts() {
	const partSize = 24
	const types = 7
	file, _ := os.Open("parts.png")
	img, _ := png.Decode(file)
	file.Close()

	bodies := make([]image.Image, 0)
	guns := make([]image.Image, 0)
	legs := make([][]image.Image, 0) // legs have frames, so 2-dimensional

	legFrames := []int{2, 2, 3, 1, 2, 2, 4}
	// legFrames := []int{1, 1, 1, 1, 1, 1, 1}
	for currLine := 0; currLine < types; currLine++ {
		bodies = append(bodies, extractSubimageFromImage(img, 0, currLine*partSize, partSize, partSize))
		guns = append(guns, extractSubimageFromImage(img, partSize, currLine*partSize, partSize, partSize))
		legFramesCurrType := make([]image.Image, 0)
		for j := 0; j < legFrames[currLine]; j++ {
			legFramesCurrType = append(legFramesCurrType, extractSubimageFromImage(img, partSize*2+(j*partSize), currLine*partSize, partSize, partSize))
		}
		legs = append(legs, legFramesCurrType)
	}
	finishedPic := image.NewNRGBA(image.Rect(0, 0, 4*partSize, types*types*types*partSize))
	currLine := 0
	for bnum := 0; bnum < types; bnum++ {
		for gnum := 0; gnum < types; gnum++ {
			for lnum := 0; lnum < types; lnum++ {
				for lframe := 0; lframe < 4; lframe++ {
					frameNum := lframe % legFrames[lnum]
					currNewFrame := image.NewNRGBA(image.Rect(0, 0, partSize, partSize))
					mergeImages(currNewFrame, legs[lnum][frameNum], bodies[bnum], guns[gnum], partSize)
					draw.Draw(finishedPic, image.Rect(lframe*partSize, currLine*partSize, (lframe+1)*partSize, (currLine+1)*partSize), currNewFrame, image.Point{0, 0}, draw.Over)
					finishedPic.Rect = image.Rect(0, 0, 4*partSize, types*types*types*partSize)
				}
				currLine++
			}
		}
	}
	file, _ = os.Create(fmt.Sprintf("generated.png"))
	png.Encode(file, finishedPic)
	file.Close()
}

func mergeImages(newImg, legs, bodies, guns image.Image, partSize int) {
	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
	draw.Draw(newImg.(*image.NRGBA), image.Rect(0, 0, partSize, partSize), legs, image.Point{0, 0}, draw.Over)
	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
	draw.Draw(newImg.(*image.NRGBA), image.Rect(0, 0, partSize, partSize), bodies, image.Point{0, 0}, draw.Over)
	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
	draw.Draw(newImg.(*image.NRGBA), image.Rect(0, 0, partSize, partSize), guns, image.Point{0, 0}, draw.Over)
	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
}

func createAtlasFromRandomGenerated() *spriteAtlas {
	return CreateAtlasFromFile(
		"generated.png",
		0,
		rnd.Rand(100)*24,
		24,
		true,
		4,
		)
}
