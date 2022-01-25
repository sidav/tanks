package main

const (
	TANK_PLAYER1 = iota
	TANK_PLAYER2
	TANK_T1
	TANK_T2
	TANK_T3
	TANK_T4
	TANK_T5
	TANK_T6
	TANK_T7
	TANK_T8

	PROJ_BULLET
	PROJ_ROCKET

	EFFECT_EXPLOSION
	EFFECT_BIG_EXPLOSION
	EFFECT_SPAWN
)

func getRandomCode() int {
	return rnd.RandInRange(TANK_T1, TANK_T8)
}

type tankStats struct {
	shootsProjectileOfCode int
	effectOnDestroy        int
	speed                  int

	sprites *spriteAtlas

	radius int

	moveDelay  int
	shootDelay int
}

var tankStatsList map[int]*tankStats

func initTankStatsList() {
	tankStatsList = map[int]*tankStats{

		/////// PROJECTILES
		PROJ_BULLET: {
			sprites:         projectileAtlaces[PROJ_BULLET],
			speed:           4,
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
		},
		PROJ_ROCKET: {
			sprites:         projectileAtlaces[PROJ_ROCKET],
			speed:           2,
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
		},

		/////// EFFECTS
		EFFECT_EXPLOSION: {
			sprites:   effectAtlaces[EFFECT_EXPLOSION],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 3,
		},
		EFFECT_BIG_EXPLOSION: {
			sprites:   effectAtlaces[EFFECT_BIG_EXPLOSION],
			radius:    TILE_PHYSICAL_SIZE,
			moveDelay: 5,
		},
		EFFECT_SPAWN: {
			sprites:   effectAtlaces[EFFECT_SPAWN],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},

		/////// TANKS
		TANK_PLAYER1: {
			sprites: tankAtlaces[TANK_T1],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_BIG_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  2,
			speed:      2,
			shootDelay: 45,
		},
		TANK_PLAYER2: {
			sprites: tankAtlaces[TANK_T6],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_BIG_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  2,
			speed:      2,
			shootDelay: 45,
		},
		TANK_T1: {
			sprites: tankAtlaces[TANK_T1],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  6,
			speed:      2,
			shootDelay: 40,
		},
		TANK_T2: {
			sprites: tankAtlaces[TANK_T2],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  8,
			speed:      2,
			shootDelay: 50,
		},
		TANK_T3: {
			sprites: tankAtlaces[TANK_T3],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  3,
			speed:      2,
			shootDelay: 60,
		},
		TANK_T4: {
			sprites: tankAtlaces[TANK_T4],

			shootsProjectileOfCode: PROJ_ROCKET,
			effectOnDestroy:        EFFECT_BIG_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  10,
			speed:      1,
			shootDelay: 20,
		},
		TANK_T5: {
			sprites: tankAtlaces[TANK_T5],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  1,
			speed:      1,
			shootDelay: 60,
		},
		TANK_T6: {
			sprites: tankAtlaces[TANK_T6],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  2,
			speed:      3,
			shootDelay: 45,
		},
		TANK_T7: {
			sprites: tankAtlaces[TANK_T7],

			shootsProjectileOfCode: PROJ_BULLET,
			effectOnDestroy:        EFFECT_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  5,
			speed:      1,
			shootDelay: 45,
		},
		TANK_T8: {
			sprites: tankAtlaces[TANK_T8],

			shootsProjectileOfCode: PROJ_ROCKET,
			effectOnDestroy:        EFFECT_BIG_EXPLOSION,

			radius:     halfPhysicalTileSize() - 2,
			moveDelay:  10,
			speed:      1,
			shootDelay: 55,
		},
	}
}
