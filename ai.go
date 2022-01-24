package main

type tankAi struct {
	chanceToRotateAnywhere          int
	chanceToRotateAtTileMiddle      int
	chanceToShootOnTarget           int
	chanceToShootOnDestructibleTile int
}

func initSimpleTankAi() *tankAi {
	return &tankAi{
		chanceToRotateAnywhere:          100,
		chanceToRotateAtTileMiddle:      35,
		chanceToShootOnTarget:           15,
		chanceToShootOnDestructibleTile: 30,
	}
}

func (b *battlefield) isTileInFrontOfTankImpassable(t *tank) bool {
	tilex, tiley := t.centerX/TILE_PHYSICAL_SIZE+t.faceX, t.centerY/TILE_PHYSICAL_SIZE+t.faceY
	if !areTileCoordsValid(tilex, tiley) {
		return true
	}
	return b.tiles[tilex][tiley].isImpassable()
}

func (b *battlefield) wantsToShoot(t *tank) bool {
	tilex, tiley := trueCoordsToTileCoords(t.centerX, t.centerY)

	for i := range b.tanks {
		if b.tanks[i].faction == t.faction {
			continue
		}
		ex, ey := b.tanks[i].getCenterCoords()
		ex, ey = trueCoordsToTileCoords(ex, ey)
		if ex != tilex && ey != tiley {
			continue
		}
		for areTileCoordsValid(tilex, tiley) {
			if ex == tilex && ey == tiley {
				return rnd.OneChanceFrom(t.ai.chanceToShootOnTarget) // enemy seen
			}

			if b.tiles[tilex][tiley].stopsProjectiles() {
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
		t.faceX, t.faceY = randomUnitVector()
	}
	debugWrite("check player; ")
	enemyInFront := b.wantsToShoot(t)
	debugWrite("rotate; ")
	wantsToRotate := rnd.OneChanceFrom(t.ai.chanceToRotateAnywhere)
	if (t.centerX %TILE_PHYSICAL_SIZE == TILE_PHYSICAL_SIZE/2+1) || (t.centerY %TILE_PHYSICAL_SIZE == TILE_PHYSICAL_SIZE/ 2+1) {
		wantsToRotate = wantsToRotate || rnd.OneChanceFrom(t.ai.chanceToRotateAtTileMiddle)
	}
	if t.canMoveNow() && !enemyInFront && (wantsToRotate || b.isTileInFrontOfTankImpassable(t)) {
		t.faceX, t.faceY = randomUnitVector()
		t.nextTickToMove = gameTick + t.getStats().moveDelay
		return
	}
	debugWrite("shoot; ")
	if t.canShootNow() && enemyInFront {
		b.shootAsTank(t)
	}
	debugWrite("move; ")
	if t.canMoveNow() && b.howFarCanTankMoveByVectorInSingleTick(t, t.faceX, t.faceY) > 0 {
		b.moveTankByVector(t, t.faceX, t.faceY)
	}
}
