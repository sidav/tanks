package main

import "fmt"

type tankAi struct {
	chanceToRotateAnywhere          int
	chanceToRotateAtTilePerfectSpot int
	chanceToShootOnTarget           int
	chanceToShootOnDestructibleTile int
}

func initSimpleTankAi() *tankAi {
	return &tankAi{
		chanceToRotateAnywhere: 100,
		chanceToRotateAtTilePerfectSpot: 35,
		chanceToShootOnTarget: 15,
		chanceToShootOnDestructibleTile: 50,
	}
}

func (b *battlefield) isTileInFrontOfTankImpassable(t *tank) bool {
	tilex, tiley := t.centerX/TILE_SIZE_TRUE+t.faceX, t.centerY/TILE_SIZE_TRUE+t.faceY
	if !b.areTileCoordsValid(tilex, tiley) {
		return true
	}
	return b.tiles[tilex][tiley].isImpassable()
}

func (b *battlefield) wantsToShoot(t *tank) bool {
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
				return rnd.OneChanceFrom(t.ai.chanceToShootOnTarget) // enemy seen
			}

			if b.tiles[tilex][tiley].isImpassable() {
				if b.tiles[tilex][tiley].isDestructible() {
					return rnd.OneChanceFrom(t.ai.chanceToShootOnDestructibleTile) // destructible tile seen
				}
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
	enemyInFront := b.wantsToShoot(t)
	fmt.Printf("rotate; ")
	wantsToRotate := rnd.OneChanceFrom(t.ai.chanceToRotateAnywhere)
	if (t.centerX % TILE_SIZE_TRUE == TILE_SIZE_TRUE/2+1) || (t.centerY % TILE_SIZE_TRUE == TILE_SIZE_TRUE / 2+1) {
		wantsToRotate = wantsToRotate || rnd.OneChanceFrom(t.ai.chanceToRotateAtTilePerfectSpot)
	}
	if t.canMoveNow() && !enemyInFront && (wantsToRotate || b.isTileInFrontOfTankImpassable(t)) {
		for {
			t.faceX, t.faceY = rnd.RandomUnitVectorInt()
			if t.faceX == 0 || t.faceY == 0 {
				break
			}
		}
		return
	}
	fmt.Printf("shoot; ")
	if t.canShootNow() && enemyInFront {
		b.shootAsTank(t)
	}
	fmt.Printf("move; ")
	if t.canMoveNow() && b.canTankMoveByVector(t, t.faceX, t.faceY) {
		t.moveByVector(t.faceX, t.faceY)
	}
}
