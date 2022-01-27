package main

type tankAi struct {
	chanceToRotateAnywhere          int
	chanceToRotateAtTileMiddle      int

	chanceToShootOnTarget           int
	chanceToShootOnDestructibleTile int
	chanceToShootAnyway             int
}

func initSimpleTankAi() *tankAi {
	return &tankAi{
		chanceToRotateAnywhere:          25,
		chanceToRotateAtTileMiddle:      15,
		chanceToShootOnTarget:           15,
		chanceToShootOnDestructibleTile: 20,
		chanceToShootAnyway:             75,
	}
}

//func (b *battlefield) isTileInFrontOfTankImpassable(t *tank) bool {
//	tilex, tiley := t.centerX/TILE_PHYSICAL_SIZE+t.faceX, t.centerY/TILE_PHYSICAL_SIZE+t.faceY
//	if !areTileCoordsValid(tilex, tiley) {
//		return true
//	}
//	return b.tiles[tilex][tiley].isImpassable()
//}

func (b *battlefield) getVectorToRotateBy(t *tank) (int, int) {
	tileX, tileY := trueCoordsToTileCoords(t.centerX, t.centerY)
	vectorsPassable := make([][2] int, 0)
	vectorsDestructible := make([][2] int, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != j && (i == 0 || j == 0) {
				if areTileCoordsValid(tileX+i, tileY+j) {
					if !b.tiles[tileX+i][tileY+j].isImpassable() {
						vectorsPassable = append(vectorsPassable, [2]int{i, j})
					} else if b.tiles[tileX+i][tileY+j].isDestructible() {
						vectorsDestructible = append(vectorsDestructible, [2]int{i, j})
					}
				}
			}
		}
	}
	if len(vectorsDestructible) > 0 && rnd.OneChanceFrom(4) {
		ind := rnd.Rand(len(vectorsDestructible))
		return vectorsDestructible[ind][0], vectorsDestructible[ind][1]
	}
	if len(vectorsPassable) > 0 {
		ind := rnd.Rand(len(vectorsPassable))
		return vectorsPassable[ind][0], vectorsPassable[ind][1]
	}
	return randomUnitVector()
}

func (b *battlefield) isThereEnemyInFront(t *tank) bool {
	tileX, tileY := trueCoordsToTileCoords(t.centerX, t.centerY)

	for i := range b.tanks {
		if b.tanks[i].faction == t.faction {
			continue
		}
		ex, ey := b.tanks[i].getCenterCoords()
		ex, ey = trueCoordsToTileCoords(ex, ey)
		if ex != tileX && ey != tileY {
			continue
		}
		for areTileCoordsValid(tileX, tileY) {
			if ex == tileX && ey == tileY {
				return true
			}
			if b.tiles[tileX][tileY].stopsProjectiles() {
				// if it's HQ...
				if b.tiles[tileX][tileY].code == TILE_HQ {
					return true
				}
				// or if there's an HQ just behind of it...
				if areTileCoordsValid(tileX+t.faceX, tileY+t.faceY) && b.tiles[tileX+t.faceX][tileY+t.faceY].code == TILE_HQ {
					return true
				}

				return false
			}
			tileX += t.faceX
			tileY += t.faceY
		}
	}
	return false
}

func (b *battlefield) wantsToShoot(t *tank) bool {
	if rnd.OneChanceFrom(t.ai.chanceToShootAnyway) {
		return true
	}
	if rnd.OneChanceFrom(t.ai.chanceToShootOnTarget) && b.isThereEnemyInFront(t) {
		return true
	}
	tilex, tiley := trueCoordsToTileCoords(t.centerX, t.centerY)
	for areTileCoordsValid(tilex, tiley) {
		if b.tiles[tilex][tiley].stopsProjectiles() {
			if b.tiles[tilex][tiley].isDestructible() {
				return rnd.OneChanceFrom(t.ai.chanceToShootOnDestructibleTile) // destructible tile seen
			}
			return false
		}
		tilex += t.faceX
		tiley += t.faceY
	}
	return false
}

func (b *battlefield) actAiForTank(t *tank) {
	if t.faceX == 0 && t.faceY == 0 {
		t.faceX, t.faceY = b.getVectorToRotateBy(t)
	}
	debugWrite("check player; ")
	enemyInFront := b.wantsToShoot(t)
	debugWrite("rotate; ")
	wantsToRotate := rnd.OneChanceFrom(t.ai.chanceToRotateAnywhere)
	if t.isAtCenterOfTile() {
		wantsToRotate = wantsToRotate || rnd.OneChanceFrom(t.ai.chanceToRotateAtTileMiddle)
	}

	canMoveBy := b.howFarCanTankMoveByVectorInSingleTick(t, t.faceX, t.faceY)

	if t.canMoveNow() && !b.isThereEnemyInFront(t) && (wantsToRotate || canMoveBy == 0) {
		t.faceX, t.faceY = b.getVectorToRotateBy(t)
		t.nextTickToMove = gameTick + t.getStats().moveDelay
		return
	}
	debugWrite("shoot; ")
	if t.canShootNow() && enemyInFront {
		b.shootAsTank(t)
	}
	debugWrite("move; ")
	if t.canMoveNow() && canMoveBy > 0 {
		b.moveTankByVector(t, t.faceX, t.faceY)
	}
}
