package main

const (
	TILE_EMPTY = iota
	TILE_WALL
	TILE_ARMORED
	TILE_WOOD
	TILE_WATER
	TILE_ICE
	TILE_HQ
)

type tile struct {
	code int
	damageTaken int
}

func (t *tile) isImpassable() bool {
	return tilesDictionary[t.code].impassable
}

func (t *tile) stopsProjectiles() bool {
	return tilesDictionary[t.code].stopsBullets
}

func (t *tile) isSlowing() bool {
	return tilesDictionary[t.code].slows
}

func (t *tile) isDestructible() bool {
	return tilesDictionary[t.code].destructible
}

func (t *tile) getSpritesAtlas() *horizSpriteAtlas {
	return tilesDictionary[t.code].sprite
}

type tileStats struct {
	sprite       *horizSpriteAtlas
	maxDamage 	 int
	impassable   bool
	stopsBullets bool
	slows 		 bool
	destructible bool
}

var tilesDictionary map[int]tileStats

func initTileDictionary() {
	tilesDictionary = map[int]tileStats{
		TILE_EMPTY:
		{
			sprite: nil,
			impassable:   false,
			destructible: false,
		},
		TILE_WALL:
		{
			sprite: tileAtlaces["WALL"],
			impassable:   true,
			stopsBullets: true,
			destructible: true,
		},
		TILE_ARMORED:
		{
			sprite: tileAtlaces["ARMORED_WALL"],
			impassable:   true,
			stopsBullets: true,
			destructible: false,
		},
		TILE_WOOD:
		{
			sprite: tileAtlaces["WOOD"],
			impassable:   false,
			destructible: false,
		},
		TILE_ICE:
		{
			sprite: tileAtlaces["ICE"],
			impassable:   false,
			stopsBullets: false,
			destructible: false,
			slows: true,
		},
		TILE_WATER:
		{
			sprite: tileAtlaces["WATER"],
			impassable:   true,
			stopsBullets: false,
			destructible: false,
		},
		TILE_HQ:
		{
			sprite:
			tileAtlaces["HQ"],
			impassable:   true,
			destructible: true,
		},
	}
}
