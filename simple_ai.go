package main

import "fmt"

type tankAi struct {
	chanceToRotateOneFrom int
	chanceToShootOneFrom  int
}

func initSimpleTankAi() *tankAi {
	return &tankAi{
		chanceToRotateOneFrom: 35,
		chanceToShootOneFrom:  1,
	}
}

func (b *battlefield) isTileInFrontOfTankImpassable(t *tank) bool {
	tilex, tiley := t.centerX/TILE_SIZE_TRUE+t.faceX, t.centerY/TILE_SIZE_TRUE+t.faceY
	if !b.areTileCoordsValid(tilex, tiley) {
		return true
	}
	return b.tiles[tilex][tiley].isImpassable()
}

func (b *battlefield) isThereEnemyInFront(t *tank) bool {
	tilex, tiley := b.trueCoordsToTileCoords(t.centerX, t.centerY)

	for i := range b.tanks {
		if b.tanks[i].faction == t.faction {
			continue
		}
		ex, ey := b.tanks[i].getCenterCoords()
		ex, ey = b.trueCoordsToTileCoords(ex, ey)
		if ex != tilex && ey != tiley {
			continue
		}
		for b.areTileCoordsValid(tilex, tiley) {
			if ex == tilex && ey == tiley {
				return true
			}
			if b.tiles[tilex][tiley].isImpassable() {
				return false
			}
			tilex += t.faceX
			tiley += t.faceY
		}
	}
	return false
}

func (b *battlefield) actAiForTank(t *tank) {
	if t.faceX == 0 && t.faceY == 0 {
		t.faceX = -1
	}
	fmt.Printf("check player; ")
	enemyInFront := b.isThereEnemyInFront(t)
	fmt.Printf("rotate; ")
	if t.canMoveNow() && !enemyInFront && rnd.OneChanceFrom(t.ai.chanceToRotateOneFrom) || b.isTileInFrontOfTankImpassable(t){
		for {
			t.faceX, t.faceY = rnd.RandomUnitVectorInt()
			if t.faceX == 0 || t.faceY == 0 {
				break
			}
		}
		return
	}
	fmt.Printf("shoot; ")
	if t.canShootNow() && enemyInFront && rnd.OneChanceFrom(t.ai.chanceToShootOneFrom) {
		b.shootAsTank(t)
	}
	fmt.Printf("move; ")
	if t.canMoveNow() && b.canTankMoveByVector(t, t.faceX, t.faceY) {
		t.moveByVector(t.faceX, t.faceY)
	}
}
