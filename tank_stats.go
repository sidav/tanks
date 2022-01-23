package main

const (
	TANK_PLAYER = iota
	TANK_T1
	TANK_T2
	TANK_T3
	TANK_T4
	TANK_T5

	PROJ_BULLET
	PROJ_ROCKET

	EFFECT_EXPLOSION
	EFFECT_BIG_EXPLOSION
	EFFECT_SPAWN
)

type tankStats struct {
	sprites *horizSpriteAtlas

	radius int

	moveDelay  int
	shootDelay int
}

var tankStatsList map[int]*tankStats

func initTankStatsList() {
	tankStatsList = map[int]*tankStats{
		PROJ_BULLET:
		{
			sprites:
			projectileAtlaces[PROJ_BULLET],
			radius: 2,
		},
		PROJ_ROCKET:
		{
			sprites:
			projectileAtlaces[PROJ_ROCKET],
			radius: 2,
		},

		EFFECT_EXPLOSION:
		{
			sprites:   effectAtlaces[EFFECT_EXPLOSION],
			radius:    TILE_PHYSICAL_SIZE/2 - 1,
			moveDelay: 3,
		},
		EFFECT_BIG_EXPLOSION:
		{
			sprites:   effectAtlaces[EFFECT_BIG_EXPLOSION],
			radius:    TILE_PHYSICAL_SIZE,
			moveDelay: 3,
		},
		EFFECT_SPAWN:
		{
			sprites:   effectAtlaces[EFFECT_SPAWN],
			radius:    TILE_PHYSICAL_SIZE/2 - 1,
			moveDelay: 7,
		},

		TANK_PLAYER:
		{
			sprites:    tankAtlaces[TANK_T1],
			radius:     TILE_PHYSICAL_SIZE/2 - 2,
			moveDelay:  2,
			shootDelay: 25,
		},
		TANK_T1:
		{
			sprites:    tankAtlaces[TANK_T1],
			radius:     TILE_PHYSICAL_SIZE/2 - 2,
			moveDelay:  6,
			shootDelay: 40,
		},
		TANK_T2:
		{
			sprites:    tankAtlaces[TANK_T2],
			radius:     TILE_PHYSICAL_SIZE/2 - 2,
			moveDelay:  8,
			shootDelay: 50,
		},
		TANK_T3:
		{
			sprites:    tankAtlaces[TANK_T3],
			radius:     TILE_PHYSICAL_SIZE/2 - 2,
			moveDelay:  3,
			shootDelay: 60,
		},
		TANK_T4:
		{
			sprites:    tankAtlaces[TANK_T4],
			radius:     TILE_PHYSICAL_SIZE/2 - 2,
			moveDelay:  10,
			shootDelay: 20,
		},
		TANK_T5:
		{
			sprites:    tankAtlaces[TANK_T5],
			radius:     TILE_PHYSICAL_SIZE/2 - 2,
			moveDelay:  2,
			shootDelay: 60,
		},
	}
}
