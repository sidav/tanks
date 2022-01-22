package main

const (
	DESIRED_WALLS         = 75
	DESIRED_ARMORED_WALLS = 40
	DESIRED_WOODS         = 10
	DESIRED_WATER         = 10
)

func (b *battlefield) getRandomEmptyTileCoords(fx, tx, fy, ty int) (int, int) {
	for {
		x, y := rnd.RandInRange(fx, tx), rnd.RandInRange(fy, ty)
		if b.tiles[x][y].code == TILE_EMPTY {
			return x, y
		}
	}
}

func (b *battlefield) init() {
	// todo: REWRITE, add better generator
	b.tiles = make([][]tile, MAP_W)
	for i := range b.tiles {
		b.tiles[i] = make([]tile, MAP_H)
	}

	for i := 0; i < DESIRED_WALLS/2; i++ {
		x, y := b.getRandomEmptyTileCoords(0, MAP_W/2, 0, MAP_H-1)
		b.tiles[x][y].code = TILE_WALL
		b.tiles[MAP_W-x-1][y].code = TILE_WALL
	}

	for i := 0; i < DESIRED_ARMORED_WALLS/2; i++ {
		x, y := b.getRandomEmptyTileCoords(0, MAP_W/2, 0, MAP_H-1)
		b.tiles[x][y].code = TILE_ARMORED
		b.tiles[MAP_W-x-1][y].code = TILE_ARMORED
	}

	for i := 0; i < DESIRED_WOODS/2; i++ {
		x, y := b.getRandomEmptyTileCoords(0, MAP_W/2, 0, MAP_H-1)
		b.tiles[x][y].code = TILE_WOOD
		b.tiles[MAP_W-x-1][y].code = TILE_WOOD
	}

	for i := 0; i < DESIRED_WATER/2; i++ {
		x, y := b.getRandomEmptyTileCoords(0, MAP_W/2, 0, MAP_H-1)
		b.tiles[x][y].code = TILE_WATER
		b.tiles[MAP_W-x-1][y].code = TILE_WATER
	}

	for x := MAP_W/2 - 1; x <= MAP_W/2+1; x++ {
		for y := MAP_H - 2; y <= MAP_H-1; y++ {
			b.tiles[x][y].code = TILE_ARMORED
		}
	}
	b.tiles[MAP_W/2][MAP_H-1].code = TILE_HQ

	b.playerTank = &tank{
		centerX:            MAP_W/2*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2,
		centerY:            (MAP_H-3)*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2,
		faceX:              0,
		faceY:              -1,
		radius:             TILE_SIZE_TRUE / 2,
		sprites:            tankAtlaces["YELLOW_T1_TANK"],
		stats:              tankStatsList["PLAYER_TANK"],
		currentFrameNumber: 0,
	}
	b.tanks = append(b.tanks, b.playerTank)
	b.tiles[MAP_W/2][MAP_H-3].code = TILE_EMPTY

	for i := 0; i < b.initialEnemiesCount; i++ {
		b.spawnTank(0, MAP_W-1, 0, MAP_H-1)
	}
}
