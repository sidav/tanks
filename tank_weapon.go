package main

const (
	WEAPON_BULLET = iota
	WEAPON_BULLET_FASTFIRE
	WEAPON_ROCKET
	WEAPON_LIGHTNING
	WEAPON_BIG
	WEAPON_ANNIHILATOR
	WEAPON_PART1
)

func getRandomWeaponCode() int {
	return rnd.RandInRange(WEAPON_BULLET_FASTFIRE, WEAPON_ANNIHILATOR)
}

type tankWeapon struct {
	code           int
	ammo           int
	nexTickToShoot int
}

func newWeapon(code int) *tankWeapon {
	return &tankWeapon{
		code: code,
		ammo: tankWeaponStatsList[code].defaultAmmo,
	}
}

func (tw *tankWeapon) getStats() *tankWeaponStats {
	return tankWeaponStatsList[tw.code]
}

func (tw *tankWeapon) spendTime() {
	tw.nexTickToShoot = gameTick + tw.getStats().shootDelay
}

type tankWeaponStats struct {
	name                   string
	shootsProjectileOfCode int
	sprites                *spriteAtlas
	defaultAmmo            int
	shootDelay             int
}

var tankWeaponStatsList map[int]*tankWeaponStats

func initTankWeaponStatsList() {
	tankWeaponStatsList = map[int]*tankWeaponStats{
		WEAPON_BULLET: {
			name:                   "Cannon",
			sprites:                nil,
			shootsProjectileOfCode: PROJ_BULLET,
			defaultAmmo:            999,
			shootDelay:             45,
		},
		WEAPON_BULLET_FASTFIRE: {
			name:                   "Rotary cannon",
			sprites:                nil,
			shootsProjectileOfCode: PROJ_BULLET,
			defaultAmmo:            99,
			shootDelay:             5,
		},
		WEAPON_ROCKET: {
			name:                   "Rocket",
			sprites:                nil,
			shootsProjectileOfCode: PROJ_ROCKET,
			defaultAmmo:            20,
			shootDelay:             75,
		},
		WEAPON_LIGHTNING: {
			name:                   "Tesla",
			sprites:                nil,
			shootsProjectileOfCode: PROJ_LIGHTNING,
			defaultAmmo:            30,
			shootDelay:             65,
		},
		WEAPON_BIG: {
			name:                   "Big Bertha",
			sprites:                nil,
			shootsProjectileOfCode: PROJ_BIG,
			defaultAmmo:            10,
			shootDelay:             110,
		},
		WEAPON_ANNIHILATOR: {
			name:                   "Death Ray",
			sprites:                nil,
			shootsProjectileOfCode: PROJ_ANNIHILATOR,
			defaultAmmo:            5,
			shootDelay:             110,
		},

		WEAPON_PART1: {
			sprites:                weaponAtlaces[rnd.Rand(len(weaponAtlaces))],
			shootsProjectileOfCode: PROJ_BULLET,
			shootDelay:             75,
		},
	}
}
