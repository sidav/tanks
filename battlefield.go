package main

type battlefield struct {
	tiles [][]tile
	playerTank *tank
	enemies []*tank
}

func (b *battlefield) areTileCoordsValid(tx, ty int) bool {
	return tx >= 0 && tx < MAP_SIZE && ty >= 0 && ty < MAP_SIZE
}

func (b *battlefield) trueCoordsToTileCoords(tx, ty int) (int, int) {
	return tx / TILE_SIZE_TRUE, ty / TILE_SIZE_TRUE
}

func (b *battlefield) isAnotherTankPresentAtTrueCoords(thisTank *tank, x, y int) bool {
	for _, t := range b.enemies {
		if thisTank == t {
			continue
		}
		tx, ty := t.getCenterCoords()
		tx -= x
		ty -= y
		if tx*tx + ty*ty < t.radius * t.radius {
			return true
		}
	}
	if thisTank != b.playerTank {
		tx, ty := b.playerTank.getCenterCoords()
		tx -= x
		ty -= y
		if tx*tx + ty*ty < b.playerTank.radius * b.playerTank.radius {
			return true
		}
	}
	return false
}

func (b *battlefield) canTankMoveByVector(t *tank, vx, vy int) bool {
	var tx1, ty1, tx2, ty2 int
	diagRadius := t.radius*6/10
	// we need to check "left corner" and "right corner" regarding to the tank
	if vx == 0 {
		tx1, ty1 = b.trueCoordsToTileCoords(t.centerX - diagRadius, t.centerY + vy*t.radius)
		tx2, ty2 = b.trueCoordsToTileCoords(t.centerX + diagRadius, t.centerY + vy*t.radius)
	} else if vy == 0 {
		tx1, ty1 = b.trueCoordsToTileCoords(t.centerX + vx*t.radius, t.centerY - diagRadius)
		tx2, ty2 = b.trueCoordsToTileCoords(t.centerX + vx*t.radius, t.centerY + diagRadius)
	}

	return b.areTileCoordsValid(tx1, ty1) && !b.tiles[tx1][ty1].impassable &&
		b.areTileCoordsValid(tx2, ty2) && !b.tiles[tx2][ty2].impassable &&
		!b.isAnotherTankPresentAtTrueCoords(t, t.centerX + vx*t.radius, t.centerY + vy*t.radius)
}
