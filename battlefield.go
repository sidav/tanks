package main

import "gorltemplate/fibrandom"

const MAPSIZE = 13

type battlefield struct {
	tiles [][]tile
	playerTank *tank
	enemies []*tank
}

func (b *battlefield) init() {
	// todo: REWRITE, add better generator
	b.tiles = make([][]tile, MAPSIZE)
	for i := range b.tiles {
		b.tiles[i] = make([]tile, MAPSIZE)
	}

	rnd := fibrandom.FibRandom{}
	rnd.InitDefault()

	for i := 0; i < 10; i++ {
		x, y := rnd.RandInRange(1, 12), rnd.RandInRange(1,12)
		b.tiles[x][y].impassable = true
		b.tiles[x][y].sprite = tileAtlaces["WALL"]
	}

	b.playerTank = &tank{
		centerX:            32,
		centerY:            32,
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
			centerX:            x*48+24,
			centerY:            y*48+24,
			radius:             8,
			sprites:            tankAtlaces["RED_T1_TANK"],
			currentFrameNumber: 0,
		})
	}
}
