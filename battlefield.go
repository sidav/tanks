package main

type battlefield struct {
	tiles [][]tile

	playerTank *tank

	desiredEnemiesCount               int
	initialEnemiesCount               int
	chanceToSpawnEnemyEachTickOneFrom int
	enemies                           []*tank

	projectiles []*tank // haha, projectiles are tanks. TODO: refactor
	effects []*tank // haha, effecrs are too. TODO: refactor
}

func (b *battlefield) areTileCoordsValid(tx, ty int) bool {
	return tx >= 0 && tx < MAP_W && ty >= 0 && ty < MAP_H
}

func (b *battlefield) trueCoordsToTileCoords(tx, ty int) (int, int) {
	return tx / TILE_SIZE_TRUE, ty / TILE_SIZE_TRUE
}

func (b *battlefield) spawnEffect(code string, cx, cy int) {
	b.effects = append(b.effects, &tank{
		centerX:            cx,
		centerY:            cy,
		sprites:            effectAtlaces[code],
		nextTickToMove:     gameTick+tankStatsList[code].moveDelay,
	})
}

func (b *battlefield) spawnEnemyTank() {
	x, y := rnd.RandInRange(3, 12), rnd.RandInRange(0, 12)
	for b.tiles[x][y].isImpassable() || b.getAnotherTankPresentAtTrueCoords(nil, x*TILE_SIZE_TRUE, y*TILE_SIZE_TRUE) != nil {
		x, y = rnd.RandInRange(3, 12), rnd.RandInRange(0, 12)
	}
	b.enemies = append(b.enemies, &tank{
		centerX:            x*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2,
		centerY:            y*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2,
		radius:             TILE_SIZE_TRUE / 2,
		sprites:            tankAtlaces["RED_T1_TANK"],
		stats:              tankStatsList["ENEMY_TANK"],
		ai:                 initSimpleTankAi(),
		currentFrameNumber: 0,
	})
	b.spawnEffect("SPAWN", x*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2, y*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2)
}

func (b *battlefield) removeEnemyTank(t *tank) {
	for i := range b.enemies {
		if b.enemies[i] == t {
			b.enemies = append(b.enemies[:i], b.enemies[i+1:]...)
			break
		}
	}
}

func (b *battlefield) shootAsTank(t *tank) {
	newProjectile := &tank{
		centerX:            t.centerX + t.faceX*(t.radius+1),
		centerY:            t.centerY + t.faceY*(t.radius+1),
		faceX:              t.faceX,
		faceY:              t.faceY,
		radius:             4,
		sprites:            projectileAtlaces["BULLET"],
		owner:              t,
		currentFrameNumber: 0,
	}
	b.projectiles = append(b.projectiles, newProjectile)
	t.nextTickToShoot = gameTick + t.stats.shootDelay
}

func (b *battlefield) actForProjectiles() {
	for i := len(b.projectiles) - 1; i >= 0; i-- {
		proj := b.projectiles[i]
		proj.centerX += proj.faceX
		proj.centerY += proj.faceY
		projTx, projTy := b.trueCoordsToTileCoords(proj.centerX, proj.centerY)
		if !b.areTileCoordsValid(projTx, projTy) {
			b.projectiles = append(b.projectiles[:i], b.projectiles[i+1:]...)
			continue
		}
		if b.tiles[projTx][projTy].isImpassable() {
			if b.tiles[projTx][projTy].isDestructible() {
				b.tiles[projTx][projTy].code = TILE_EMPTY
			}
			b.spawnEffect("EXPLOSION", proj.centerX, proj.centerY)
			b.projectiles = append(b.projectiles[:i], b.projectiles[i+1:]...)
			continue
		}
		hitTank := b.getAnotherTankPresentAtTrueCoords(proj.owner, proj.centerX, proj.centerY)
		if hitTank != nil {
			b.removeEnemyTank(hitTank)
			b.spawnEffect("EXPLOSION", proj.centerX, proj.centerY)
			b.projectiles = append(b.projectiles[:i], b.projectiles[i+1:]...)
			continue
		}
	}
}

func (b *battlefield) actForEffects() {
	for i := len(b.effects) - 1; i >= 0; i-- {
		if b.effects[i].canMoveNow() {
			b.effects[i].currentFrameNumber++
			b.effects[i].nextTickToMove = gameTick+tankStatsList["EXPLOSION"].moveDelay
		}
		if b.effects[i].currentFrameNumber >= b.effects[i].sprites.totalFrames {
			b.effects = append(b.effects[:i], b.effects[i+1:]...)
		}
	}
}

func (b *battlefield) getAnotherTankPresentAtTrueCoords(thisTank *tank, x, y int) *tank {
	for _, t := range b.enemies {
		if thisTank == t {
			continue
		}
		tx, ty := t.getCenterCoords()
		tx -= x
		ty -= y
		if tx*tx+ty*ty < t.radius*t.radius {
			return t
		}
	}
	if thisTank != b.playerTank {
		tx, ty := b.playerTank.getCenterCoords()
		tx -= x
		ty -= y
		if tx*tx+ty*ty < b.playerTank.radius*b.playerTank.radius {
			return b.playerTank
		}
	}
	return nil
}

func (b *battlefield) canTankMoveByVector(t *tank, vx, vy int) bool {
	var tx1, ty1, tx2, ty2 int
	diagRadius := t.radius * 6 / 10
	// we need to check "left corner" and "right corner" regarding to the tank
	if vx == 0 {
		tx1, ty1 = b.trueCoordsToTileCoords(t.centerX-diagRadius, t.centerY+vy*t.radius)
		tx2, ty2 = b.trueCoordsToTileCoords(t.centerX+diagRadius, t.centerY+vy*t.radius)
	} else if vy == 0 {
		tx1, ty1 = b.trueCoordsToTileCoords(t.centerX+vx*t.radius, t.centerY-diagRadius)
		tx2, ty2 = b.trueCoordsToTileCoords(t.centerX+vx*t.radius, t.centerY+diagRadius)
	}

	return b.areTileCoordsValid(tx1, ty1) && !b.tiles[tx1][ty1].isImpassable() &&
		b.areTileCoordsValid(tx2, ty2) && !b.tiles[tx2][ty2].isImpassable() &&
		b.getAnotherTankPresentAtTrueCoords(t, t.centerX+vx*t.radius, t.centerY+vy*t.radius) == nil
}
