package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	TILE_EMPTY = iota
	TILE_WALL
	TILE_ARMORED
	TILE_WOOD
	TILE_WATER
	TILE_ICE
	TILE_HQ
	TILE_FLAG
)

type tile struct {
	code        int
	damageTaken int
}

func (t *tile) isImpassable() bool {
	return tilesDictionary[t.code].impassable
}

func (t *tile) getMaxDamageTaken() int {
	return tilesDictionary[t.code].maxDamage
}

func (t *tile) stopsProjectiles() bool {
	return tilesDictionary[t.code].stopsBullets
}

func (t *tile) isSlowing() bool {
	return tilesDictionary[t.code].slows
}

func (t *tile) isNotArmored() bool {
	return !tilesDictionary[t.code].armored
}

func (t *tile) getSpritesAtlas() *spriteAtlas {
	return tilesDictionary[t.code].sprite
}

func (t *tile) getSprite() rl.Texture2D {
	totalFrames := t.getSpritesAtlas().totalFrames()
	frameNumber := (gameTick / 50) % totalFrames
	if t.getMaxDamageTaken() > 0 {
		frameNumber = t.damageTaken*totalFrames/t.getMaxDamageTaken()
	}
	return t.getSpritesAtlas().atlas[0][frameNumber]
}

type tileStats struct {
	sprite       *spriteAtlas
	maxDamage    int
	impassable   bool
	stopsBullets bool
	slows        bool
	armored      bool
}

var tilesDictionary map[int]tileStats

func initTileDictionary() {
	tilesDictionary = map[int]tileStats{
		TILE_EMPTY: {
			sprite:     nil,
			impassable: false,
			armored:    false,
		},
		TILE_WALL: {
			maxDamage:    4,
			sprite:       tileAtlaces["WALL"],
			impassable:   true,
			stopsBullets: true,
			armored:      false,
		},
		TILE_ARMORED: {
			maxDamage:    12,
			sprite:       tileAtlaces["ARMORED_WALL"],
			impassable:   true,
			stopsBullets: true,
			armored:      true,
		},
		TILE_WOOD: {
			sprite:     tileAtlaces["WOOD"],
			impassable: false,
			armored:    false,
		},
		TILE_ICE: {
			sprite:       tileAtlaces["ICE"],
			impassable:   false,
			stopsBullets: false,
			armored:      false,
			slows:        true,
		},
		TILE_WATER: {
			sprite:       tileAtlaces["WATER"],
			impassable:   true,
			stopsBullets: false,
			armored:      false,
		},
		TILE_HQ: {
			maxDamage:    2,
			sprite:       tileAtlaces["HQ"],
			impassable:   true,
			armored:      true,
			stopsBullets: true,
		},
		TILE_FLAG: {
			sprite:       tileAtlaces["FLAG"],
			impassable:   false,
			armored:      false,
			stopsBullets: false,
		},
	}
}
