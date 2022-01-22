package main

import "fmt"

type tankAi struct {
	chanceToRotateOneFrom int
	chanceToShootOneFrom  int
}

func initSimpleTankAi() *tankAi {
	return &tankAi{
		chanceToRotateOneFrom: 35,
		chanceToShootOneFrom:  5,
	}
}

func (b *battlefield) isTileInFrontOfTankImpassable(t *tank) bool {
	tilex, tiley := t.centerX/TILE_SIZE_TRUE+t.faceX, t.centerY/TILE_SIZE_TRUE+t.faceY
	if !b.areTileCoordsValid(tilex, tiley) {
		return true
	}
	return b.tiles[tilex][tiley].impassable
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
		if b.tiles[tilex][tiley].impassable {
			return false
		}
		tilex += t.faceX
		tiley += t.faceY
	}
	return false
}

func (b *battlefield) actAiForTank(t *tank) {
	fmt.Printf("%d: AI\n", gameTick)
	playerInFront := b.isTherePlayerInFrontOfTank(t)
	if t.canMoveNow() && !playerInFront && rnd.OneChanceFrom(t.ai.chanceToRotateOneFrom) || b.isTileInFrontOfTankImpassable(t){
		for {
			t.faceX, t.faceY = rnd.RandomUnitVectorInt()
			if t.faceX == 0 || t.faceY == 0 {
				break
			}
		}
		return
	}
	if t.canShootNow() && playerInFront && rnd.OneChanceFrom(t.ai.chanceToShootOneFrom) {
		b.shootAsTank(t)
	}
	if t.canMoveNow() && b.canTankMoveByVector(t, t.faceX, t.faceY) {
		t.moveByVector(t.faceX, t.faceY)
	}
}
