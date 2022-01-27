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

	weaponAtlaces = []*spriteAtlas{}
	trackAtlaces  = []*spriteAtlas{}
	bodiesAtlaces = []*spriteAtlas{}
)

func loadImageResources() {
	var leftXForTank = 0
	tankAtlaces[TANK_T1] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*0, 16, 16, 2, true)
	tankAtlaces[TANK_T2] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*1, 16, 16, 2, true)
	tankAtlaces[TANK_T3] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*2, 16, 16, 2, true)
	tankAtlaces[TANK_T4] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*3, 16, 16, 2, true)
	tankAtlaces[TANK_T5] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*4, 16, 16, 2, true)
	tankAtlaces[TANK_T6] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*5, 16, 16, 2, true)
	tankAtlaces[TANK_T7] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*6, 16, 16, 2, true)
	tankAtlaces[TANK_T8] = CreateAtlasFromFile("assets/tanks.png", leftXForTank, 16*7, 16, 16, 2, true)

	tileAtlaces["WALL"] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*0, 16, 16, 5, false)
	tileAtlaces["ARMORED_WALL"] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*1, 16, 16, 5, false)
	tileAtlaces["WATER"] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*3, 16, 16, 2, false)
	tileAtlaces["WOOD"] = CreateAtlasFromFile("assets/sprites.png", 16*1, 16*2, 16, 16, 1, false)
	tileAtlaces["ICE"] = CreateAtlasFromFile("assets/sprites.png", 16*2, 16*2, 16, 16, 1, false)
	tileAtlaces["HQ"] = CreateAtlasFromFile("assets/sprites.png", 16*3, 16*2, 16, 16, 1, false)
	tileAtlaces["FLAG"] = CreateAtlasFromFile("assets/sprites.png", 16*4, 16*2, 16, 16, 1, false)

	projectileAtlaces[PROJ_BULLET] = CreateAtlasFromFile("assets/projectiles.png", 0, 0, 8, 8, 1, true)
	projectileAtlaces[PROJ_ROCKET] = CreateAtlasFromFile("assets/projectiles.png", 0, 8, 8, 8, 1, true)

	effectAtlaces[EFFECT_EXPLOSION] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*6, 16, 16, 3, false)
	effectAtlaces[EFFECT_BIG_EXPLOSION] = CreateAtlasFromFile("assets/sprites.png", 16*3, 16*6, 32, 32, 2, false)
	effectAtlaces[EFFECT_SPAWN] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*4, 16, 16, 4, false)

	// parts
	for i := 0; i < 27; i++ {
		bodiesAtlaces = append(bodiesAtlaces, CreateAtlasFromFile("assets/parts/TankBase24x24.png", 0, i*24, 24, 16, 1, true))
	}
	for i := 0; i < 16; i++ {
		weaponAtlaces = append(weaponAtlaces, CreateAtlasFromFile("assets/parts/TankWeapon24x24.png", 0, i*24, 24, 16, 1, true))
	}
	for i := 0; i < 6; i++ {
		trackAtlaces = append(trackAtlaces, CreateAtlasFromFile("assets/parts/TankTracksAnimated(2)24x24.png", 0, i*24, 24, 16, 2, true))
	}
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

func CreateAtlasFromFile(filename string, topleftx, toplefty, originalSpriteSize, desiredSpriteSize, totalFrames int, createAllDirections bool) *spriteAtlas {

	file, _ := os.Open(filename)
	img, _ := png.Decode(file)
	file.Close()

	newAtlas := spriteAtlas{
		spriteSize: desiredSpriteSize * int(SPRITE_SCALE_FACTOR),
	}
	if createAllDirections {
		newAtlas.atlas = make([][]rl.Texture2D, 4)
	} else {
		newAtlas.atlas = make([][]rl.Texture2D, 1)
	}
	// newAtlas.atlas
	for currFrame := 0; currFrame < totalFrames; currFrame++ {
		currPic := extractSubimageFromImage(img, topleftx+currFrame*originalSpriteSize, toplefty, originalSpriteSize, originalSpriteSize)
		rlImg := rl.NewImageFromImage(currPic)
		rl.ImageResizeNN(rlImg, int32(desiredSpriteSize)*int32(SPRITE_SCALE_FACTOR), int32(desiredSpriteSize)*int32(SPRITE_SCALE_FACTOR))
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

//func generateSpriteSheetFromParts() {
//	const partSize = 24
//	const types = 7
//	file, _ := os.Open("parts.png")
//	img, _ := png.Decode(file)
//	file.Close()
//
//	bodies := make([]image.Image, 0)
//	guns := make([]image.Image, 0)
//	legs := make([][]image.Image, 0) // legs have frames, so 2-dimensional
//
//	legFrames := []int{2, 2, 3, 1, 2, 2, 4}
//	// legFrames := []int{1, 1, 1, 1, 1, 1, 1}
//	for currLine := 0; currLine < types; currLine++ {
//		bodies = append(bodies, extractSubimageFromImage(img, 0, currLine*partSize, partSize, partSize))
//		guns = append(guns, extractSubimageFromImage(img, partSize, currLine*partSize, partSize, partSize))
//		legFramesCurrType := make([]image.Image, 0)
//		for j := 0; j < legFrames[currLine]; j++ {
//			legFramesCurrType = append(legFramesCurrType, extractSubimageFromImage(img, partSize*2+(j*partSize), currLine*partSize, partSize, partSize))
//		}
//		legs = append(legs, legFramesCurrType)
//	}
//	finishedPic := image.NewNRGBA(image.Rect(0, 0, 4*partSize, types*types*types*partSize))
//	currLine := 0
//	for bnum := 0; bnum < types; bnum++ {
//		for gnum := 0; gnum < types; gnum++ {
//			for lnum := 0; lnum < types; lnum++ {
//				for lframe := 0; lframe < 4; lframe++ {
//					frameNum := lframe % legFrames[lnum]
//					currNewFrame := image.NewNRGBA(image.Rect(0, 0, partSize, partSize))
//					mergeImages(currNewFrame, legs[lnum][frameNum], bodies[bnum], guns[gnum], partSize)
//					draw.Draw(finishedPic, image.Rect(lframe*partSize, currLine*partSize, (lframe+1)*partSize, (currLine+1)*partSize), currNewFrame, image.Point{0, 0}, draw.Over)
//					finishedPic.Rect = image.Rect(0, 0, 4*partSize, types*types*types*partSize)
//				}
//				currLine++
//			}
//		}
//	}
//	file, _ = os.Create(fmt.Sprintf("generated.png"))
//	png.Encode(file, finishedPic)
//	file.Close()
//}
//
//func mergeImages(newImg, legs, bodies, guns image.Image, partSize int) {
//	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
//	draw.Draw(newImg.(*image.NRGBA), image.Rect(0, 0, partSize, partSize), legs, image.Point{0, 0}, draw.Over)
//	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
//	draw.Draw(newImg.(*image.NRGBA), image.Rect(0, 0, partSize, partSize), bodies, image.Point{0, 0}, draw.Over)
//	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
//	draw.Draw(newImg.(*image.NRGBA), image.Rect(0, 0, partSize, partSize), guns, image.Point{0, 0}, draw.Over)
//	newImg.(*image.NRGBA).Rect = image.Rect(0, 0, partSize, partSize)
//}
//
//func createAtlasFromRandomGenerated() *spriteAtlas {
//	return CreateAtlasFromFile(
//		"generated.png",
//		0,
//		rnd.Rand(100)*24,
//		24,
//		16,
//		4,
//		true,
//		)
//}
