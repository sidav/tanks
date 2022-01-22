package main

const (
	DESIRED_WALLS = 40
	DESIRED_ARMORED_WALLS = 15
	DESIRED_WOODS = 10
)

func (b *battlefield) init() {
	// todo: REWRITE, add better generator
	b.tiles = make([][]tile, MAP_W)
	for i := range b.tiles {
		b.tiles[i] = make([]tile, MAP_H)
	}

	for i := 0; i < DESIRED_WALLS; i++ {
		x, y := rnd.RandInRange(0, MAP_W-1), rnd.RandInRange(0, MAP_H-1)
		b.tiles[x][y].impassable = true
		b.tiles[x][y].sprite = tileAtlaces["WALL"]
		b.tiles[x][y].destructible = true
	}

	for i := 0; i < DESIRED_ARMORED_WALLS; i++ {
		x, y := rnd.RandInRange(0, MAP_W-1), rnd.RandInRange(0, MAP_H-1)
		b.tiles[x][y].impassable = true
		b.tiles[x][y].sprite = tileAtlaces["ARMORED_WALL"]
	}

	for i := 0; i < DESIRED_WOODS; i++ {
		x, y := rnd.RandInRange(0, MAP_W-1), rnd.RandInRange(0, MAP_H-1)
		b.tiles[x][y].sprite = tileAtlaces["WOOD"]
	}

	for x := MAP_W/2-1; x <= MAP_W/2+1; x++ {
		for y := MAP_H-2; y <= MAP_H-1; y++ {
			b.tiles[x][y].destructible = false
			b.tiles[x][y].impassable = true
			b.tiles[x][y].sprite = tileAtlaces["ARMORED_WALL"]
		}
	}
	b.tiles[MAP_W/2][MAP_H-1].destructible = true
	b.tiles[MAP_W/2][MAP_H-1].sprite = tileAtlaces["HQ"]

	b.playerTank = &tank{
		centerX:            MAP_W/2*TILE_SIZE_TRUE + TILE_SIZE_TRUE / 2,
		centerY:            (MAP_H-3)*TILE_SIZE_TRUE + TILE_SIZE_TRUE / 2,
		radius:             TILE_SIZE_TRUE / 2,
		sprites:            tankAtlaces["YELLOW_T1_TANK"],
		stats:              tankStatsList["PLAYER_TANK"],
		currentFrameNumber: 0,
	}

	for i := 0; i < b.initialEnemiesCount; i++ {
		b.spawnEnemyTank()
	}
}
