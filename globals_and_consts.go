package main

const (
	SPRITE_SCALE_FACTOR      = 4.0
	TILE_SIZE_IN_PIXELS      = 16*SPRITE_SCALE_FACTOR
	TILE_SIZE_TRUE           = 16
	PIXEL_TO_GAMECOORD_RATIO = TILE_SIZE_IN_PIXELS/TILE_SIZE_TRUE

	MAP_W = 21
	MAP_H = 15
	NUM_FACTIONS = 4

	WINDOW_W = MAP_W*TILE_SIZE_IN_PIXELS
	WINDOW_H = MAP_H*TILE_SIZE_IN_PIXELS+TILE_SIZE_IN_PIXELS
	TEXT_SIZE = TILE_SIZE_IN_PIXELS/2
)

func areTileCoordsValid(tx, ty int) bool {
	return tx >= 0 && tx < MAP_W && ty >= 0 && ty < MAP_H
}

func trueCoordsToTileCoords(tx, ty int) (int, int) {
	return tx / TILE_SIZE_TRUE, ty / TILE_SIZE_TRUE
}

func circlesOverlap(x1, y1, r1, x2, y2, r2 int) bool {
	tx := x2-x1
	ty := y2-y1
	r := r1+r2

	if tx*tx+ty*ty < r*r {
		return true
	}
	return false
}
