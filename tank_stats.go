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
	TANK_GENERATED

	PROJ_BULLET
	PROJ_ROCKET
	PROJ_LIGHTNING

	EFFECT_EXPLOSION
	EFFECT_BIG_EXPLOSION
	EFFECT_SPAWN
)

func getRandomCode() int {
	return rnd.RandInRange(TANK_T1, TANK_T8) // PROJ_BULLET-1)
}

type tankStats struct {
	tractionCode int
	bodyCode     int
	weaponCodes  []int

	effectOnDestroy  int
	sprites          *spriteAtlas
	radius           int
	speed, moveDelay int
	shootDelay       int

	frameChangesForEffect int
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
		PROJ_LIGHTNING: {
			sprites:         projectileAtlaces[PROJ_LIGHTNING],
			speed:           3,
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
		},

		/////// EFFECTS
		EFFECT_EXPLOSION: {
			sprites:               effectAtlaces[EFFECT_EXPLOSION],
			radius:                halfPhysicalTileSize() - 1,
			moveDelay:             3,
			frameChangesForEffect: 7,
		},
		EFFECT_BIG_EXPLOSION: {
			sprites:               effectAtlaces[EFFECT_BIG_EXPLOSION],
			radius:                TILE_PHYSICAL_SIZE,
			moveDelay:             5,
			frameChangesForEffect: 3,
		},
		EFFECT_SPAWN: {
			sprites:               effectAtlaces[EFFECT_SPAWN],
			radius:                halfPhysicalTileSize() - 1,
			moveDelay:             7,
			frameChangesForEffect: 8,
		},

		/////// TANKS
		TANK_PLAYER1: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:     TBODY_TANK1,
			tractionCode: TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_PLAYER2: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:     TBODY_TANK1,
			tractionCode: TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_T1: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK1,
			tractionCode: TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_T2: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK2,
			tractionCode: TTRACTION_DEFAULT_SLOW,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_T3: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK3,
			tractionCode: TTRACTION_DEFAULT_SLOW,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_T4: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK4,
			tractionCode: TTRACTION_DEFAULT_SLOWEST,
			weaponCodes:  []int{WEAPON_ROCKET},
		},
		TANK_T5: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK5,
			tractionCode: TTRACTION_DEFAULT_FAST,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_T6: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK6,
			tractionCode: TTRACTION_DEFAULT_FASTEST,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_T7: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK7,
			tractionCode: TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_T8: {
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,

			bodyCode:     TBODY_TANK8,
			tractionCode: TTRACTION_DEFAULT_SLOWEST,
			weaponCodes:  []int{WEAPON_LIGHTNING},
		},

		TANK_GENERATED: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 3,
			bodyCode:        TBODY_PART1,
			tractionCode:    TTRACTION_PART1,
			weaponCodes:     []int{WEAPON_PART1},
		},
	}
}
