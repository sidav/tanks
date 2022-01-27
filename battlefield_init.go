package main

func (b *battlefield) init(desiredWalls, desiredArmoredWalls, desiredWoods, desiredWater, desiredIce, numPlayers, missionType int) {
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

	b.numPlayers = numPlayers
	b.initMission(missionType)

	for i := 0; i < b.initialEnemiesCount; i++ {
		b.spawnRandomTankInRect(0, MAP_W-1, 0, MAP_H-1)
	}
}

func (b *battlefield) initMission(missionType int) {
	b.missionType = missionType
	switch missionType {
	case MISSION_KILL_ALL:
		b.placePlayers(b.numPlayers, MAP_W/2, MAP_H-1, 0, MAP_H-1, MAP_W-1, MAP_H-1)
	case MISSION_PROTECT_HQ:
		b.placePlayers(b.numPlayers, MAP_W/2, MAP_H-3, MAP_W/2-1, MAP_H-3, MAP_W/2+1, MAP_H-3)
		b.placeTilesInRect(TILE_EMPTY, MAP_W/2-2, MAP_H-3, 5, 3)
		b.placeTilesInRect(TILE_ARMORED, MAP_W/2-1, MAP_H-2, 3, 2)
		// b.tiles[MAP_W/2][MAP_H-1].code = TILE_HQ
	case MISSION_COLLECT_FLAGS:
		b.placePlayers(b.numPlayers, MAP_W/2, MAP_H-3, MAP_W/2-1, MAP_H-3, MAP_W/2+1, MAP_H-3)
		for i := 0; i < 5*b.numPlayers; i++ {
			x, y := rnd.RandInRange(0, MAP_W-1), rnd.RandInRange(0, MAP_H-3)
			b.tiles[x][y].code = TILE_FLAG
		}
	}
}

func (b *battlefield) placePlayers(numPlayers, x, y, x1, y1, x2, y2 int) {
	if numPlayers == 1 {
		b.clearImpassableTilesInCrossFrom(x, y)
		x, y = tileCoordsToPhysicalCoords(x, y)
		player := newTank(TANK_PLAYER1, x, y, 0)
		player.faceY = -1
		player.playerControlled = true
		b.playerTanks = append(b.playerTanks, player)
		b.tanks = append(b.tanks, player)
	} else {
		b.clearImpassableTilesInCrossFrom(x1, y1)
		b.clearImpassableTilesInCrossFrom(x2, y2)
		x1, y1 = tileCoordsToPhysicalCoords(x1, y1)
		player1 := newTank(TANK_PLAYER1, x1, y1, 0)
		player1.faceY = -1
		player1.playerControlled = true
		x2, y2 = tileCoordsToPhysicalCoords(x2, y2)
		player2 := newTank(TANK_PLAYER2, x2, y2, 0)
		player2.faceY = -1
		player2.playerControlled = true
		b.playerTanks = append(b.playerTanks, player1)
		b.tanks = append(b.tanks, player1)
		b.playerTanks = append(b.playerTanks, player2)
		b.tanks = append(b.tanks, player2)
	}
}

func (b *battlefield) clearImpassableTilesInCrossFrom(x, y int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i*j == 0 && areTileCoordsValid(x+i, y+j) {
				if b.tiles[x+i][y+j].isImpassable() && !b.tiles[x+i][y+j].isDestructible() {
					b.tiles[x+i][y+j].code = TILE_EMPTY
				}
				b.tiles[x][y].code = TILE_EMPTY
			}
		}
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

func (b *battlefield) placeTilesRandomSymmetric(tileCode, tilePercent int) {
	if tilePercent > 100 {
		panic("Wrong tile percent")
	}
	tileCount := (MAP_W*MAP_H)*tilePercent/100
	if tileCount > MAP_W*MAP_H {
		tileCount = MAP_W * MAP_H
	}
	for i := 0; i < tileCount/2; i++ {
		x, y := b.getRandomEmptyTileCoords(0, MAP_W/2, 0, MAP_H-1)
		b.tiles[x][y].code = tileCode
		b.tiles[MAP_W-x-1][y].code = tileCode
	}
}

func (b *battlefield) placeTilesInRect(tileCode, x, y, w, h int) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			b.tiles[i][j].code = tileCode
		}
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
