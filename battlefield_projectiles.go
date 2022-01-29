package main

func (b *battlefield) actForProjectiles() {
	for i := len(b.projectiles) - 1; i >= 0; i-- {
		proj := b.projectiles[i]
		if proj.tickToExpire <= gameTick {
			proj.markedToRemove = true
		}
		speed := proj.getStats().speed
		if proj.getStats().acceleratesEach > 0 {
			speed = speed + (gameTick - proj.tickToExpire + proj.getStats().duration)/proj.getStats().acceleratesEach
		}
		proj.centerX += proj.faceX * speed
		proj.centerY += proj.faceY * speed
		proj.currentFrameNumber++
		projTx, projTy := trueCoordsToTileCoords(proj.centerX, proj.centerY)
		if proj.markedToRemove || !areTileCoordsValid(projTx, projTy) ||
			proj.centerX+proj.faceX <= 0 || proj.centerY+proj.faceY <= 0 {

			b.projectiles = append(b.projectiles[:i], b.projectiles[i+1:]...)
			b.spawnEffect(proj.getStats().effectOnDestroy, proj.centerX, proj.centerY, nil)
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

		// check if we hit wall
		if b.tiles[projTx][projTy].stopsProjectiles() {
			b.dealDamageToTile(projTx, projTy, proj.getStats().damage, proj.getStats().canDestroyArmor)
			proj.markedToRemove = true
			continue
		}

		// check if we hit tank
		hitTank := b.getAnotherTankPresentAtTrueCoords(proj.owner, proj.centerX, proj.centerY)
		if hitTank != nil {
			proj.markedToRemove = true
			if proj.owner == nil || proj.owner.faction != hitTank.faction {
				b.dealDamageToTank(hitTank, proj.getStats().damage)
			}
			continue
		}
	}
}
