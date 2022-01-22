package main

type battlefield struct {
	tiles [][]tile

	playerTank *tank

	maxTanksOnMap                     int
	initialEnemiesCount               int
	totalTanksRemainingToSpawn        int
	chanceToSpawnEnemyEachTickOneFrom int
	numFactions                       int
	tanks                             []*tank

	projectiles []*tank // haha, projectiles are tanks. TODO: refactor
	effects     []*tank // haha, effecrs are too. TODO: refactor
}

func (b *battlefield) spawnEffect(code int, cx, cy int, owner *tank) {
	b.effects = append(b.effects, &tank{
		centerX:        cx,
		centerY:        cy,
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
		if b.effects[i].currentFrameNumber >= b.effects[i].getSpritesAtlas().totalFrames {
			if b.effects[i].code == EFFECT_SPAWN {
				b.tanks = append(b.tanks, b.effects[i].owner)
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
	b.totalTanksRemainingToSpawn--
	tankFaction := rnd.RandInRange(1, b.numFactions-1)
	tankCode := -1
	switch tankFaction {
	case 1:
		tankCode = TANK_T2
	case 2:
		tankCode = TANK_T3
	case 3:
		tankCode = TANK_T4
	default:
		tankCode = TANK_T5
	}
	owner := &tank{
		code:               tankCode,
		centerX:            x*TILE_PHYSICAL_SIZE + TILE_PHYSICAL_SIZE/2,
		centerY:            y*TILE_PHYSICAL_SIZE + TILE_PHYSICAL_SIZE/2,
		ai:                 initSimpleTankAi(),
		faction:            tankFaction,
		currentFrameNumber: 0,
	}
	b.spawnEffect(EFFECT_SPAWN, x*TILE_PHYSICAL_SIZE+TILE_PHYSICAL_SIZE/2, y*TILE_PHYSICAL_SIZE+TILE_PHYSICAL_SIZE/2, owner)
}

func (b *battlefield) removeTank(t *tank) {
	for i := range b.tanks {
		if b.tanks[i] == t {
			if b.tanks[i] == b.playerTank {
				b.playerTank = nil
			}
			cx, cy := b.tanks[i].getCenterCoords()
			b.spawnEffect(EFFECT_EXPLOSION, cx, cy, nil)
			b.tanks = append(b.tanks[:i], b.tanks[i+1:]...)
			break
		}
	}
}

func (b *battlefield) shootAsTank(t *tank) {
	newProjectile := &tank{
		code:               PROJ_BULLET,
		centerX:            t.centerX + t.faceX*(t.getRadius()+1),
		centerY:            t.centerY + t.faceY*(t.getRadius()+1),
		faceX:              t.faceX,
		faceY:              t.faceY,
		owner:              t,
		currentFrameNumber: 0,
	}
	b.projectiles = append(b.projectiles, newProjectile)
	t.nextTickToShoot = gameTick + t.getStats().shootDelay
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
			b.spawnEffect(EFFECT_EXPLOSION, proj.centerX, proj.centerY, nil)
			continue
		}

		// check if we hit another projectile
		for _, p := range b.projectiles {
			if p == proj {
				continue
			}
			if circlesOverlap(proj.centerX, proj.centerY, proj.getRadius(), p.centerX, p.centerY, p.getRadius()) {
				proj.markedToRemove = true
				p.markedToRemove = true
				continue
			}
		}

		if b.tiles[projTx][projTy].stopsProjectiles() {
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
		if circlesOverlap(x, y, r, tx, ty, t.getRadius()) {
			return t
		}
	}
	return nil
}

func (b *battlefield) getAnotherTankPresentAtTrueCoords(thisTank *tank, x, y int) *tank {
	r := 0
	if thisTank != nil {
		r = thisTank.getRadius()
	}
	for _, t := range b.tanks {
		if thisTank == t {
			continue
		}
		tx, ty := t.getCenterCoords()
		if circlesOverlap(x, y, r, tx, ty, t.getRadius()) {
			return t
		}
	}
	return nil
}

func (b *battlefield) canTankMoveByVector(t *tank, vx, vy int) bool {
	var tx1, ty1, tx2, ty2 int
	diagRadius := t.getRadius() * 65 / 100
	// we need to check "left corner" and "right corner" regarding to the tank
	if vx == 0 {
		tx1, ty1 = trueCoordsToTileCoords(t.centerX-diagRadius, t.centerY+vy*t.getRadius())
		tx2, ty2 = trueCoordsToTileCoords(t.centerX+diagRadius, t.centerY+vy*t.getRadius())
	} else if vy == 0 {
		tx1, ty1 = trueCoordsToTileCoords(t.centerX+vx*t.getRadius(), t.centerY-diagRadius)
		tx2, ty2 = trueCoordsToTileCoords(t.centerX+vx*t.getRadius(), t.centerY+diagRadius)
	}

	return (t.centerX+vx >= t.getRadius()) && (t.centerY+vy >= t.getRadius()) &&
		areTileCoordsValid(tx1, ty1) && !b.tiles[tx1][ty1].isImpassable() &&
		areTileCoordsValid(tx2, ty2) && !b.tiles[tx2][ty2].isImpassable() &&
		b.getAnotherTankPresentAtTrueCoords(t, t.centerX+vx*t.getRadius(), t.centerY+vy*t.getRadius()) == nil
}
