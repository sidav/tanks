package main

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

func (b *battlefield) getTankPresentFromRadius(radius, x, y int) *tank {
	for _, t := range b.tanks {
		tx, ty := t.getCenterCoords()
		if circlesOverlap(x, y, radius, tx, ty, t.getRadius()) {
			return t
		}
	}
	return nil
}

func (b *battlefield) moveTankByVector(t *tank, x, y int) {
	delay := t.getTractionStats().moveDelay
	t.nextTickToMove = gameTick + delay
	tx, ty := trueCoordsToTileCoords(t.centerX, t.centerY)
	if gameMap.tiles[tx][ty].isSlowing() {
		t.nextTickToMove += delay
	}
	speed := b.howFarCanTankMoveByVectorInSingleTick(t, x, y)
	if speed > 0 {
		t.centerX += x * speed
		t.centerY += y * speed
	}
	t.faceX = x
	t.faceY = y
	t.currentFrameNumber++
}

func (b *battlefield) howFarCanTankMoveByVectorInSingleTick(t *tank, vx, vy int) int {
	var tx1, ty1, tx2, ty2 int
	diagRadius := t.getRadius() * 90 / 100
	// we need to check "left corner" and "right corner" regarding to the tank
	for currSpeed := t.getTractionStats().speed; currSpeed > 0; currSpeed-- {
		if vx == 0 {
			tx1, ty1 = trueCoordsToTileCoords(t.centerX-diagRadius, t.centerY+vy*(currSpeed+t.getRadius()))
			tx2, ty2 = trueCoordsToTileCoords(t.centerX+diagRadius, t.centerY+vy*(currSpeed+t.getRadius()))
		} else if vy == 0 {
			tx1, ty1 = trueCoordsToTileCoords(t.centerX+vx*(currSpeed+t.getRadius()), t.centerY-diagRadius)
			tx2, ty2 = trueCoordsToTileCoords(t.centerX+vx*(currSpeed+t.getRadius()), t.centerY+diagRadius)
		}

		if (t.centerX+vx >= t.getRadius()) && (t.centerY+vy >= t.getRadius()) &&
			areTileCoordsValid(tx1, ty1) && !b.tiles[tx1][ty1].isImpassable() &&
			areTileCoordsValid(tx2, ty2) && !b.tiles[tx2][ty2].isImpassable() &&
			b.getAnotherTankPresentAtTrueCoords(t, t.centerX+vx*t.getRadius(), t.centerY+vy*t.getRadius()) == nil {
			return currSpeed
		}
	}
	return 0
}

func (b *battlefield) spawnTankInRect(t *tank, fromx, tox, fromy, toy int) {
	var x, y int
	placeTries := (tox-fromx)*(toy-fromy) + 1
	for try := 0; ; try++ {
		if try == placeTries {
			debugWrite("SPAWNING FAILED")
			return // failure
		}
		x, y = rnd.RandInRange(fromx, tox), rnd.RandInRange(fromy, toy)
		trueX, trueY := tileCoordsToPhysicalCoords(x, y)
		if b.getEffectPresentInRadiusFromTrueCoords(trueX, trueY, TILE_PHYSICAL_SIZE/2+1) != nil {
			continue
		}
		if !b.tiles[x][y].isImpassable() &&
			b.getTankPresentFromRadius(TILE_PHYSICAL_SIZE/2, trueX, trueY) == nil {

			break
		}
	}
	cx, cy := tileCoordsToPhysicalCoords(x, y)
	t.centerX, t.centerY = cx, cy
	b.spawnEffect(EFFECT_SPAWN, cx, cy, t)
}

func (b *battlefield) spawnRandomTankInRect(fromx, tox, fromy, toy int) {
	b.totalTanksRemainingToSpawn--
	tankFaction := rnd.RandInRange(1, b.numFactions-1)
	tankCode := getRandomCode()
	owner := newTank(tankCode, 0, 0, tankFaction)
	owner.ai = initSimpleTankAi()
	b.spawnTankInRect(owner, fromx, tox, fromy, toy)
}

func (b *battlefield) removeTank(t *tank) {
	for i := range b.tanks {
		if b.tanks[i] == t {
			if b.tanks[i].playerControlled {
				for i := len(b.playerTanks) - 1; i >= 0; i-- {
					if b.playerTanks[i] == t {
						b.playerTanks[i] = nil // append(b.playerTanks[:i], b.playerTanks[i+1:]...)
					}
				}
			}
			cx, cy := b.tanks[i].getCenterCoords()
			b.spawnEffect(b.tanks[i].getStats().effectOnDestroy, cx, cy, nil)
			b.tanks = append(b.tanks[:i], b.tanks[i+1:]...)
			break
		}
	}
}

func (b *battlefield) shootAsTank(t *tank) {
	newProjectile := &projectile{
		code:               t.weapons[t.currentWeaponNumber].getStats().shootsProjectileOfCode,
		centerX:            t.centerX + t.faceX*(t.getRadius()+1),
		centerY:            t.centerY + t.faceY*(t.getRadius()+1),
		faceX:              t.faceX,
		faceY:              t.faceY,
		owner:              t,
		currentFrameNumber: 0,
	}
	newProjectile.tickToExpire = gameTick + newProjectile.getStats().duration
	b.projectiles = append(b.projectiles, newProjectile)
	t.weapons[t.currentWeaponNumber].spendTime()
}
