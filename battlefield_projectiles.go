package main

const CHANGE_PROJECTILE_FRAME_EACH = 3

func (b *battlefield) actForProjectiles() {
	for i := len(b.projectiles) - 1; i >= 0; i-- {
		proj := b.projectiles[i]
		if proj.tickToExpire <= gameTick {
			proj.hitsRemaining = 0
		}
		speed := proj.getStats().speed
		if proj.getStats().acceleratesEach > 0 {
			speed = speed + (gameTick-proj.tickToExpire+proj.getStats().duration)/proj.getStats().acceleratesEach
		}
		proj.centerX += proj.faceX * speed
		proj.centerY += proj.faceY * speed
		if gameTick%CHANGE_PROJECTILE_FRAME_EACH == 0 {
			proj.currentFrameNumber++
		}
		projTx, projTy := trueCoordsToTileCoords(proj.centerX, proj.centerY)
		if proj.hitsRemaining <= 0 || !areTileCoordsValid(projTx, projTy) ||
			proj.centerX+proj.faceX <= 0 || proj.centerY+proj.faceY <= 0 {

			if proj.hitsRemaining > 0 {
				b.spawnEffect(proj.getStats().effectOnDestroy, proj.centerX, proj.centerY, nil)
			}
			b.projectiles = append(b.projectiles[:i], b.projectiles[i+1:]...)
			continue
		}

		// check if we hit another projectile
		for _, p := range b.projectiles {
			if p == proj {
				continue
			}
			if circlesOverlap(proj.centerX, proj.centerY, proj.getRadius(), p.centerX, p.centerY, p.getRadius()) {
				proj.hitsRemaining--
				b.spawnEffect(proj.getStats().effectOnDestroy, proj.centerX, proj.centerY, nil)
				p.hitsRemaining--
				continue
			}
		}

		// check if we hit wall
		if b.tiles[projTx][projTy].stopsProjectiles() {
			b.dealDamageToTile(projTx, projTy, proj.getStats().damage, proj.getStats().canDestroyArmor)
			b.spawnEffect(proj.getStats().effectOnDestroy, proj.centerX, proj.centerY, nil)
			proj.hitsRemaining--
			continue
		}

		// check if we hit tank
		hitTank := b.getAnotherTankPresentAtTrueCoords(proj.owner, proj.centerX, proj.centerY)
		if hitTank != nil {
			proj.hitsRemaining--
			b.spawnEffect(proj.getStats().effectOnDestroy, proj.centerX, proj.centerY, nil)
			if proj.owner == nil || proj.owner.faction != hitTank.faction {
				b.dealDamageToTank(hitTank, proj.getStats().damage)
			}
			continue
		}
	}
}
