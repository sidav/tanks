package main

const (
	WEAPON_BULLET = iota
	WEAPON_ROCKET
	WEAPON_PART1
)

type tankWeapon struct {
	code           int
	nexTickToShoot int
}

func (tw *tankWeapon) getStats() *tankWeaponStats {
	return tankWeaponStatsList[tw.code]
}

func (tw *tankWeapon) spendTime() {
	tw.nexTickToShoot = gameTick + tw.getStats().shootDelay
}

type tankWeaponStats struct {
	shootsProjectileOfCode int
	sprites                *spriteAtlas
	shootDelay             int
}

var tankWeaponStatsList map[int]*tankWeaponStats

func initTankWeaponStatsList() {
	tankWeaponStatsList = map[int]*tankWeaponStats{
		WEAPON_BULLET: {
			sprites:                nil,
			shootsProjectileOfCode: PROJ_BULLET,
			shootDelay:             45,
		},
		WEAPON_ROCKET: {
			sprites:                nil,
			shootsProjectileOfCode: PROJ_ROCKET,
			shootDelay:             75,
		},

		WEAPON_PART1: {
			sprites:                weaponAtlaces[rnd.Rand(len(weaponAtlaces))],
			shootsProjectileOfCode: PROJ_BULLET,
			shootDelay:             75,
		},
	}
}
