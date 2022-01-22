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
		b.tiles[x][y].code = TILE_WALL
	}

	for i := 0; i < DESIRED_ARMORED_WALLS; i++ {
		x, y := rnd.RandInRange(0, MAP_W-1), rnd.RandInRange(0, MAP_H-1)
		b.tiles[x][y].code = TILE_ARMORED
	}

	for i := 0; i < DESIRED_WOODS; i++ {
		x, y := rnd.RandInRange(0, MAP_W-1), rnd.RandInRange(0, MAP_H-1)
		b.tiles[x][y].code = TILE_WOOD
	}

	for x := MAP_W/2-1; x <= MAP_W/2+1; x++ {
		for y := MAP_H-2; y <= MAP_H-1; y++ {
			b.tiles[x][y].code = TILE_ARMORED
		}
	}
	b.tiles[MAP_W/2][MAP_H-1].code = TILE_HQ

	b.playerTank = &tank{
		centerX:            MAP_W/2*TILE_SIZE_TRUE + TILE_SIZE_TRUE / 2,
		centerY:            (MAP_H-3)*TILE_SIZE_TRUE + TILE_SIZE_TRUE / 2,
		radius:             TILE_SIZE_TRUE / 2,
		sprites:            tankAtlaces["YELLOW_T1_TANK"],
		stats:              tankStatsList["PLAYER_TANK"],
		currentFrameNumber: 0,
	}
	b.tanks = append(b.tanks, b.playerTank)

	for i := 0; i < b.initialEnemiesCount; i++ {
		b.spawnTank(0, MAP_W-1, 0, MAP_H-1)
	}
}
