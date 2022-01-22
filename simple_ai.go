package main

import "fmt"

type tankAi struct {
	chanceToRotateOneFrom int
	chanceToShootOneFrom  int
}

func initSimpleTankAi() *tankAi {
	return &tankAi{
		chanceToRotateOneFrom: 35,
		chanceToShootOneFrom:  15,
	}
}

func (b *battlefield) isTileInFrontOfTankImpassable(t *tank) bool {
	tilex, tiley := t.centerX/TILE_SIZE_TRUE+t.faceX, t.centerY/TILE_SIZE_TRUE+t.faceY
	if !b.areTileCoordsValid(tilex, tiley) {
		return true
	}
	return b.tiles[tilex][tiley].isImpassable()
}

func (b *battlefield) isTherePlayerInFrontOfTank(t *tank) bool {
	tilex, tiley := b.trueCoordsToTileCoords(t.centerX, t.centerY)
	px, py := b.playerTank.getCenterCoords()
	px, py = b.trueCoordsToTileCoords(px, py)
	if px != tilex && py != tiley {
		return false
	}
	for b.areTileCoordsValid(tilex, tiley) {
		if px == tilex && py == tiley {
			return true
		}
		if b.tiles[tilex][tiley].isImpassable() {
			return false
		}
		tilex += t.faceX
		tiley += t.faceY
	}
	return false
}

func (b *battlefield) actAiForTank(t *tank) {
	if t.faceX == 0 && t.faceY == 0 {
		t.faceX = -1
	}
	fmt.Printf("check player; ")
	playerInFront := b.isTherePlayerInFrontOfTank(t)
	fmt.Printf("rotate; ")
	if t.canMoveNow() && !playerInFront && rnd.OneChanceFrom(t.ai.chanceToRotateOneFrom) || b.isTileInFrontOfTankImpassable(t){
		for {
			t.faceX, t.faceY = rnd.RandomUnitVectorInt()
			if t.faceX == 0 || t.faceY == 0 {
				break
			}
		}
		return
	}
	fmt.Printf("shoot; ")
	if t.canShootNow() && playerInFront && rnd.OneChanceFrom(t.ai.chanceToShootOneFrom) {
		b.shootAsTank(t)
	}
	fmt.Printf("move; ")
	if t.canMoveNow() && b.canTankMoveByVector(t, t.faceX, t.faceY) {
		t.moveByVector(t.faceX, t.faceY)
	}
}
