package main

type battlefield struct {
	tiles [][]tile

	numPlayers  int
	playerTanks []*tank

	missionType                       int
	maxTanksOnMap                     int
	initialEnemiesCount               int
	totalTanksRemainingToSpawn        int
	chanceToSpawnEnemyEachTickOneFrom int
	numFactions                       int
	tanks                             []*tank

	projectiles []*tank // haha, projectiles are tanks. TODO: refactor
	effects     []*tank // haha, effects are too. TODO: refactor
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
		if b.effects[i].currentFrameNumber > b.effects[i].getStats().frameChangesForEffect {
			if b.effects[i].code == EFFECT_SPAWN {
				b.tanks = append(b.tanks, b.effects[i].owner)
			}
			b.effects = append(b.effects[:i], b.effects[i+1:]...)
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

func (b *battlefield) dealDamageToTile(projTx, projTy, damage int) {
	if b.tiles[projTx][projTy].isDestructible() {
		b.tiles[projTx][projTy].damageTaken += damage
		if b.tiles[projTx][projTy].damageTaken >= b.tiles[projTx][projTy].getMaxDamageTaken() {
			b.tiles[projTx][projTy].code = TILE_EMPTY
		}
	}
}

func (b *battlefield) dealDamageToTank(t *tank, damage int) {
	t.hitpoints -= damage
	if t.hitpoints <= 0 {
		b.removeTank(t)
	}
}
