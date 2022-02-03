package main

const (
	TANK_PLAYER1 = iota
	TANK_PLAYER1_WU
	TANK_PLAYER1_AU
	TANK_PLAYER1_AWU
	TANK_PLAYER2
	TANK_PLAYER2_WU
	TANK_PLAYER2_AU
	TANK_PLAYER2_AWU
	TANK_T1
	TANK_T2
	TANK_T3
	TANK_T4
	TANK_T5
	TANK_T6
	TANK_T7
	TANK_T8
	TANK_GENERATED
)

func getRandomCode() int {
	return rnd.RandInRange(TANK_T1, TANK_T8) // PROJ_BULLET-1)
}

type tankStats struct {
	damageAsProjectile int

	tractionCode int
	bodyCode     int
	weaponCodes  []int

	effectOnDestroy int
	radius          int

	codeWhenWeaponUpgraded, codeWhenArmorUpgraded int // temp
}

var tankStatsList map[int]*tankStats

func initTankStatsList() {
	tankStatsList = map[int]*tankStats{
		TANK_PLAYER1: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:               TBODY_PLAYER1,
			tractionCode:           TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:            []int{WEAPON_BULLET},
			codeWhenWeaponUpgraded: TANK_PLAYER1_WU,
			codeWhenArmorUpgraded:  TANK_PLAYER1_AU,
		},
		TANK_PLAYER1_WU: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:               TBODY_PLAYER1_WU,
			tractionCode:           TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:            []int{WEAPON_BULLET},
			codeWhenWeaponUpgraded: TANK_PLAYER1_WU,
			codeWhenArmorUpgraded:  TANK_PLAYER1_AWU,
		},
		TANK_PLAYER1_AU: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:               TBODY_PLAYER1_AU,
			tractionCode:           TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:            []int{WEAPON_BULLET},
			codeWhenWeaponUpgraded: TANK_PLAYER1_AWU,
			codeWhenArmorUpgraded:  TANK_PLAYER1_AU,
		},
		TANK_PLAYER1_AWU: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:               TBODY_PLAYER1_AWU,
			tractionCode:           TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:            []int{WEAPON_ROCKET},
			codeWhenWeaponUpgraded: TANK_PLAYER1_AWU,
			codeWhenArmorUpgraded:  TANK_PLAYER1_AWU,
		},
		TANK_PLAYER2: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:     TBODY_PLAYER2,
			tractionCode: TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_PLAYER2_WU: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:     TBODY_PLAYER2_WU,
			tractionCode: TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_PLAYER2_AU: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:     TBODY_PLAYER2_AU,
			tractionCode: TTRACTION_DEFAULT_FAST,
			weaponCodes:  []int{WEAPON_BULLET},
		},
		TANK_PLAYER2_AWU: {
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          halfPhysicalTileSize() - 2,

			bodyCode:     TBODY_PLAYER2_AWU,
			tractionCode: TTRACTION_DEFAULT_MEDIUM,
			weaponCodes:  []int{WEAPON_ROCKET},
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
