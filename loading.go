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
	bonusAtlaces      = map[int]*spriteAtlas{}

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

	bonusAtlaces[BONUS_HELM] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*5, 16, 16, 1, false)
	bonusAtlaces[BONUS_CLOCK] = CreateAtlasFromFile("assets/sprites.png", 16*1, 16*5, 16, 16, 1, false)
	bonusAtlaces[BONUS_SHOVEL] = CreateAtlasFromFile("assets/sprites.png", 16*2, 16*5, 16, 16, 1, false)
	bonusAtlaces[BONUS_STAR] = CreateAtlasFromFile("assets/sprites.png", 16*3, 16*5, 16, 16, 1, false)
	bonusAtlaces[BONUS_GRENADE]= CreateAtlasFromFile("assets/sprites.png", 16*4, 16*5, 16, 16, 1, false)
	bonusAtlaces[BONUS_TANK]= CreateAtlasFromFile("assets/sprites.png", 16*5, 16*5, 16, 16, 1, false)
	bonusAtlaces[BONUS_GUN]= CreateAtlasFromFile("assets/sprites.png", 16*6, 16*5, 16, 16, 1, false)

	projectileAtlaces[PROJ_BULLET] = CreateAtlasFromFile("assets/projectiles.png", 0, 0, 8, 8, 2, true)
	projectileAtlaces[PROJ_ROCKET] = CreateAtlasFromFile("assets/projectiles.png", 0, 8, 8, 8, 2, true)
	projectileAtlaces[PROJ_LIGHTNING] = CreateAtlasFromFile("assets/projectiles.png", 0, 16, 8, 8, 2, true)
	projectileAtlaces[PROJ_BIG] = CreateAtlasFromFile("assets/projectiles.png", 0, 24, 8, 8, 2, true)
	projectileAtlaces[PROJ_ANNIHILATOR] = CreateAtlasFromFile("assets/projectiles.png", 0, 32, 8, 8, 2, true)

	effectAtlaces[EFFECT_EXPLOSION] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*6, 16, 16, 3, false)
	effectAtlaces[EFFECT_BIG_EXPLOSION] = CreateAtlasFromFile("assets/sprites.png", 16*3, 16*6, 32, 32, 2, false)
	effectAtlaces[EFFECT_SPAWN] = CreateAtlasFromFile("assets/sprites.png", 16*0, 16*4, 16, 16, 4, false)

	// parts
	for i := 0; i < 27; i++ {
		bodiesAtlaces = append(bodiesAtlaces, CreateAtlasFromFile("assets/parts/TankBase24x24.png", 0, i*24, 24, ORIGINAL_TILE_SIZE_IN_PIXELS, 1, true))
	}
	for i := 0; i < 16; i++ {
		weaponAtlaces = append(weaponAtlaces, CreateAtlasFromFile("assets/parts/TankWeapon24x24.png", 0, i*24, 24, ORIGINAL_TILE_SIZE_IN_PIXELS, 1, true))
	}
	for i := 0; i < 6; i++ {
		trackAtlaces = append(trackAtlaces, CreateAtlasFromFile("assets/parts/TankTracksAnimated(2)24x24.png", 0, i*24, 24, ORIGINAL_TILE_SIZE_IN_PIXELS, 2, true))
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

func generateSpriteSheetFromParts() {
	const partSize = TILE_SIZE_IN_PIXELS + 5

	finishedPic := image.NewNRGBA(image.Rect(0, 0, len(trackAtlaces)*len(weaponAtlaces)*partSize, len(bodiesAtlaces)*partSize))
	currLine := 0
	for body := range bodiesAtlaces {
		currColumn := 0
		for legs := range trackAtlaces {
			for weapon := range weaponAtlaces {
				currNewFrame := image.NewRGBA(image.Rect(0, 0, partSize, partSize))
				mergeImages(currNewFrame,
					textureToGolangImage(trackAtlaces[legs].atlas[0][0]),
					textureToGolangImage(bodiesAtlaces[body].atlas[0][0]),
					textureToGolangImage(weaponAtlaces[weapon].atlas[0][0]),
					partSize,
				)
				draw.Draw(finishedPic, image.Rect(currColumn*partSize, currLine*partSize, (currColumn+1)*partSize, (currLine+1)*partSize), currNewFrame, image.Point{0, 0}, draw.Over)
				currColumn++
			}
		}
		currLine++
	}

	file, _ := os.Create(fmt.Sprintf("build/generated.png"))
	png.Encode(file, finishedPic)
	file.Close()
}

func textureToGolangImage(t rl.Texture2D) *image.RGBA {
	img := rl.LoadImageFromTexture(t)
	nrgba := img.ToImage().(*image.RGBA)
	return nrgba
}

func mergeImages(newImg, legs, bodies, guns image.Image, partSize int) {
	newImg.(*image.RGBA).Rect = image.Rect(0, 0, partSize, partSize)
	draw.Draw(newImg.(*image.RGBA), image.Rect(0, 0, partSize, partSize), legs, image.Point{0, 0}, draw.Over)
	newImg.(*image.RGBA).Rect = image.Rect(0, 0, partSize, partSize)
	draw.Draw(newImg.(*image.RGBA), image.Rect(0, 0, partSize, partSize), bodies, image.Point{0, 0}, draw.Over)
	newImg.(*image.RGBA).Rect = image.Rect(0, 0, partSize, partSize)
	draw.Draw(newImg.(*image.RGBA), image.Rect(0, 0, partSize, partSize), guns, image.Point{0, 0}, draw.Over)
	newImg.(*image.RGBA).Rect = image.Rect(0, 0, partSize, partSize)
}
