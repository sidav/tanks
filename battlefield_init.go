package main

func (b *battlefield) init() {
	// todo: REWRITE, add better generator
	b.tiles = make([][]tile, MAP_SIZE)
	for i := range b.tiles {
		b.tiles[i] = make([]tile, MAP_SIZE)
	}

	for i := 0; i < 25; i++ {
		x, y := rnd.RandInRange(1, 12), rnd.RandInRange(1, 12)
		b.tiles[x][y].impassable = true
		b.tiles[x][y].sprite = tileAtlaces["WALL"]
	}

	b.playerTank = &tank{
		centerX:            TILE_SIZE_TRUE / 2,
		centerY:            TILE_SIZE_TRUE / 2,
		radius:             TILE_SIZE_TRUE / 2,
		sprites:            tankAtlaces["YELLOW_T1_TANK"],
		stats:              tankStatsList["PLAYER_TANK"],
		currentFrameNumber: 0,
	}

	for i := 0; i < b.initialEnemiesCount; i++ {
		b.spawnEnemyTank()
	}
}
