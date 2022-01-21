package main

import "gorltemplate/fibrandom"

type battlefield struct {
	tiles [][]tile
	playerTank *tank
	enemies []*tank
}

func (b *battlefield) init() {
	// todo: REWRITE, add better generator
	b.tiles = make([][]tile, MAP_SIZE)
	for i := range b.tiles {
		b.tiles[i] = make([]tile, MAP_SIZE)
	}

	rnd := fibrandom.FibRandom{}
	rnd.InitDefault()

	for i := 0; i < 10; i++ {
		x, y := rnd.RandInRange(1, 12), rnd.RandInRange(1,12)
		b.tiles[x][y].impassable = true
		b.tiles[x][y].sprite = tileAtlaces["WALL"]
	}

	b.playerTank = &tank{
		centerX:            TILE_SIZE_TRUE /2,
		centerY:            TILE_SIZE_TRUE /2,
		radius:             8,
		sprites:            tankAtlaces["YELLOW_T1_TANK"],
		currentFrameNumber: 0,
	}

	for i := 0; i < 3; i++ {
		x, y := rnd.RandInRange(6, 12), rnd.RandInRange(0,12)
		for b.tiles[x][y].impassable {
			x, y = rnd.RandInRange(6, 12), rnd.RandInRange(0,12)
		}
		b.enemies = append(b.enemies, &tank{
			centerX:            x*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2,
			centerY:            y*TILE_SIZE_TRUE + TILE_SIZE_TRUE/2,
			radius:             8,
			sprites:            tankAtlaces["RED_T1_TANK"],
			currentFrameNumber: 0,
		})
	}
}
