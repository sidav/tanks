package main

type battlefield struct {
	tiles [][]tile

	playerTank *tank

	MaxTanksOnMap                     int
	initialEnemiesCount               int
	totalTanksRemainingToSpawn        int
	chanceToSpawnEnemyEachTickOneFrom int
	numFactions                       int
	tanks                             []*tank

	projectiles []*tank // haha, projectiles are tanks. TODO: refactor
	effects     []*tank // haha, effecrs are too. TODO: refactor
}

func (b *battlefield) spawnEffect(code string, cx, cy int, owner *tank) {
	b.effects = append(b.effects, &tank{
		centerX:        cx,
		centerY:        cy,
		sprites:        effectAtlaces[code],
		nextTickToMove: gameTick + tankStatsList[code].moveDelay,
		owner:          owner,
		code:           code,
	})
}

func (b *battlefield) actForEffects() {
	for i := len(b.effects) - 1; i >= 0; i-- {
		if b.effects[i].canMoveNow() {
			b.effects[i].currentFrameNumber++
			b.effects[i].nextTickToMove = gameTick + tankStatsList[b.effects[i].code].moveDelay
		}
		if b.effects[i].currentFrameNumber >= b.effects[i].sprites.totalFrames {
			if b.effects[i].code == "SPAWN" {
				b.tanks = append(b.tanks, b.effects[i].owner)
				b.totalTanksRemainingToSpawn--
			}
			b.effects = append(b.effects[:i], b.effects[i+1:]...)
		}
	}
}

func (b *battlefield) spawnTank(fromx, tox, fromy, toy int) {
	var x, y int
	for {
		x, y = rnd.RandInRange(fromx, tox), rnd.RandInRange(fromy, toy)
		trueX, trueY := tileCoordsToPhysicalCoords(x, y)
		if b.getEffectPresentInRadiusFromTrueCoords(trueX, trueY, TILE_PHYSICAL_SIZE/2+1) != nil {
			continue
		}
		if !b.tiles[x][y].isImpassable() &&
			b.getAnotherTankPresentAtTrueCoords(nil, trueX, trueY) == nil {

			break
		}
	}
	tankFaction := rnd.RandInRange(1, b.numFactions-1)
	tankCode := ""
	switch tankFaction {
	case 1:
		tankCode = "GRAY_T1_TANK"
	case 2:
		tankCode = "GREEN_T1_TANK"
	case 3:
		tankCode = "RED_T1_TANK"
	}
	owner := &tank{
		centerX:            x*TILE_PHYSICAL_SIZE + TILE_PHYSICAL_SIZE/2,
		centerY:            y*TILE_PHYSICAL_SIZE + TILE_PHYSICAL_SIZE/2,
		radius:             TILE_PHYSICAL_SIZE/2 - 1,
		sprites:            tankAtlaces[tankCode],
		stats:              tankStatsList[tankCode],
		ai:                 initSimpleTankAi(),
		faction:            tankFaction,
		currentFrameNumber: 0,
	}
	b.spawnEffect("SPAWN", x*TILE_PHYSICAL_SIZE+TILE_PHYSICAL_SIZE/2, y*TILE_PHYSICAL_SIZE+TILE_PHYSICAL_SIZE/2, owner)
}

func (b *battlefield) removeTank(t *tank) {
	for i := range b.tanks {
		if b.tanks[i] == t {
			if b.tanks[i] == b.playerTank {
				b.playerTank = nil
			}
			cx, cy := b.tanks[i].getCenterCoords()
			b.spawnEffect("EXPLOSION", cx, cy, nil)
			b.tanks = append(b.tanks[:i], b.tanks[i+1:]...)
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
		radius:             2,
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
		projTx, projTy := trueCoordsToTileCoords(proj.centerX, proj.centerY)
		if proj.markedToRemove || !areTileCoordsValid(projTx, projTy) ||
			proj.centerX+proj.faceX <= 0 || proj.centerY+proj.faceY <= 0 {

			b.projectiles = append(b.projectiles[:i], b.projectiles[i+1:]...)
			b.spawnEffect("EXPLOSION", proj.centerX, proj.centerY, nil)
			continue
		}

		// check if we hit another projectile
		for _, p := range b.projectiles {
			if p == proj {
				continue
			}
			if circlesOverlap(proj.centerX, proj.centerY, proj.radius, p.centerX, p.centerY, p.radius) {
				proj.markedToRemove = true
				p.markedToRemove = true
				continue
			}
		}

		if b.tiles[projTx][projTy].isImpassable() {
			if b.tiles[projTx][projTy].isDestructible() {
				b.tiles[projTx][projTy].code = TILE_EMPTY
			}
			proj.markedToRemove = true
			continue
		}
		hitTank := b.getAnotherTankPresentAtTrueCoords(proj.owner, proj.centerX, proj.centerY)
		if hitTank != nil {
			b.removeTank(hitTank)
			proj.markedToRemove = true
			continue
		}
	}
}

func (b *battlefield) getEffectPresentInRadiusFromTrueCoords(x, y, r int) *tank {
	for _, t := range b.effects {
		tx, ty := t.getCenterCoords()
		if circlesOverlap(x, y, r, tx, ty, t.radius) {
			return t
		}
	}
	return nil
}

func (b *battlefield) getAnotherTankPresentAtTrueCoords(thisTank *tank, x, y int) *tank {
	r := 0
	if thisTank != nil {
		r = thisTank.radius
	}
	for _, t := range b.tanks {
		if thisTank == t {
			continue
		}
		tx, ty := t.getCenterCoords()
		if circlesOverlap(x, y, r, tx, ty, t.radius) {
			return t
		}
	}
	return nil
}

func (b *battlefield) canTankMoveByVector(t *tank, vx, vy int) bool {
	var tx1, ty1, tx2, ty2 int
	diagRadius := t.radius * 65 / 100
	// we need to check "left corner" and "right corner" regarding to the tank
	if vx == 0 {
		tx1, ty1 = trueCoordsToTileCoords(t.centerX-diagRadius, t.centerY+vy*t.radius)
		tx2, ty2 = trueCoordsToTileCoords(t.centerX+diagRadius, t.centerY+vy*t.radius)
	} else if vy == 0 {
		tx1, ty1 = trueCoordsToTileCoords(t.centerX+vx*t.radius, t.centerY-diagRadius)
		tx2, ty2 = trueCoordsToTileCoords(t.centerX+vx*t.radius, t.centerY+diagRadius)
	}

	return (t.centerX+vx >= t.radius) && (t.centerY+vy >= t.radius) &&
		areTileCoordsValid(tx1, ty1) && !b.tiles[tx1][ty1].isImpassable() &&
		areTileCoordsValid(tx2, ty2) && !b.tiles[tx2][ty2].isImpassable() &&
		b.getAnotherTankPresentAtTrueCoords(t, t.centerX+vx*t.radius, t.centerY+vy*t.radius) == nil
}
