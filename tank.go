package main

type tank struct {
	playerControlled bool

	centerX, centerY   int
	faceX, faceY       int
	currentFrameNumber int

	faction int

	owner          *tank
	markedToRemove bool
	code           int

	nextTickToMove int

	hitpoints           int
	weapons             []*tankWeapon
	currentWeaponNumber int

	ai *tankAi
}

func newTank(code, x, y, faction int) *tank {
	t := &tank{code: code, centerX: x, centerY: y, faction: faction}
	for i := 0; i < len(t.getStats().weaponCodes); i++ {
		t.weapons = append(t.weapons, &tankWeapon{code: t.getStats().weaponCodes[i]})
	}
	t.hitpoints = t.getBodyStats().maxHp
	return t
}

func (t *tank) getCenterCoords() (int, int) {
	return t.centerX, t.centerY
}

func (t *tank) isAtCenterOfTile() bool {
	xInTile := t.centerX % TILE_PHYSICAL_SIZE
	yInTile := t.centerY % TILE_PHYSICAL_SIZE
	precision := TILE_PHYSICAL_SIZE / 10
	if abs(halfPhysicalTileSize()-xInTile) <= precision && abs(halfPhysicalTileSize()-yInTile) <= precision {
		return true
	}
	return false
}

func (t *tank) getRadius() int {
	return tankStatsList[t.code].radius
}

func (t *tank) canMoveNow() bool {
	return gameTick >= t.nextTickToMove
}

func (t *tank) canShootNow() bool {
	return gameTick >= t.weapons[t.currentWeaponNumber].nexTickToShoot
}

func (t *tank) getStats() *tankStats {
	return tankStatsList[t.code]
}

func (t *tank) getTractionStats() *tankTractionStats {
	return tankTractionStatsList[t.getStats().tractionCode]
}

func (t *tank) getBodyStats() *tankBodyStats {
	return tankBodyStatsList[t.getStats().bodyCode]
}

func (t *tank) getSpritesAtlas() *spriteAtlas {
	//debugWritef("ATLAS{%v}", t.code)
	return t.getStats().sprites
}
