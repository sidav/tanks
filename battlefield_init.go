package main

func (b *battlefield) init(desiredWalls, desiredArmoredWalls, desiredWoods, desiredWater, desiredIce, numPlayers int) {
	// todo: REWRITE, add better generator
	b.tiles = make([][]tile, MAP_W)
	for i := range b.tiles {
		b.tiles[i] = make([]tile, MAP_H)
	}

	b.placeTilesRandomSymmetric(TILE_WALL, desiredWalls)
	b.placeTilesRandomSymmetric(TILE_ARMORED, desiredArmoredWalls)
	b.placeTilesRandomSymmetric(TILE_WATER, desiredWater)
	b.placeTilesRandomSymmetric(TILE_WOOD, desiredWoods)
	b.placeTilesRandomSymmetric(TILE_ICE, desiredIce)
	b.clearTilesForTanksSpawnIfNeeded(b.initialEnemiesCount)

	for x := MAP_W/2 - 1; x <= MAP_W/2+1; x++ {
		for y := MAP_H - 2; y <= MAP_H-1; y++ {
			b.tiles[x][y].code = TILE_ARMORED
		}
	}
	b.tiles[MAP_W/2][MAP_H-1].code = TILE_HQ

	b.numPlayers = numPlayers
	if numPlayers == 2 {
		cx, cy := tileCoordsToPhysicalCoords(MAP_W/2+1, MAP_H-3)
		player1 := newTank(TANK_PLAYER1, cx, cy, 0)
		player1.playerControlled = true
		cx, cy = tileCoordsToPhysicalCoords(MAP_W/2-1, MAP_H-3)
		player2 := newTank(TANK_PLAYER2, cx, cy, 0)
		player2.playerControlled = true
		b.playerTanks = append(b.playerTanks, player1)
		b.tanks = append(b.tanks, player1)
		b.playerTanks = append(b.playerTanks, player2)
		b.tanks = append(b.tanks, player2)

		b.tiles[MAP_W/2-1][MAP_H-3].code = TILE_EMPTY
		b.tiles[MAP_W/2+1][MAP_H-3].code = TILE_EMPTY
	} else {
		cx, cy := tileCoordsToPhysicalCoords(MAP_W/2, MAP_H-3)
		player := newTank(TANK_PLAYER1, cx, cy, 0)
		player.playerControlled = true
		b.playerTanks = append(b.playerTanks, player)
		b.tiles[MAP_W/2][MAP_H-3].code = TILE_EMPTY
		b.tanks = append(b.tanks, player)
	}

	for i := 0; i < b.initialEnemiesCount; i++ {
		b.spawnRandomTankInRect(0, MAP_W-1, 0, MAP_H-1)
	}
}

func (b *battlefield) getRandomEmptyTileCoords(fx, tx, fy, ty int) (int, int) {
	for tries := 0; tries < MAP_W*MAP_H*2; tries++ {
		x, y := rnd.RandInRange(fx, tx), rnd.RandInRange(fy, ty)
		if b.tiles[x][y].code == TILE_EMPTY {
			return x, y
		}
	}
	return 0, 0
}

func (b *battlefield) placeTilesRandomSymmetric(tileCode, tileCount int) {
	if tileCount > MAP_W*MAP_H {
		tileCount = MAP_W * MAP_H
	}
	for i := 0; i < tileCount/2; i++ {
		x, y := b.getRandomEmptyTileCoords(0, MAP_W/2, 0, MAP_H-1)
		b.tiles[x][y].code = tileCode
		b.tiles[MAP_W-x-1][y].code = tileCode
	}
}

func (b *battlefield) clearTilesForTanksSpawnIfNeeded(count int) {
	debugWritef("COUNT %d", count)
	if count > MAP_W*MAP_H {
		count = MAP_W * MAP_H
	}
	if count == 0 {
		count = 1
	}
	for x := 0; x < MAP_W; x++ {
		for y := 0; y < MAP_H; y++ {
			if !b.tiles[x][y].isImpassable() {
				count--
			}
		}
	}
	debugWritef("COUNT %d", count)
	for i := 0; i < count; i++ {
		x, y := rnd.RandInRange(1, MAP_W-2), rnd.RandInRange(1, MAP_H-2)
		b.tiles[x][y].code = TILE_EMPTY
	}
}
