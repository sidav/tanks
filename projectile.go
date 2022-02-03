package main

const (
	PROJ_BULLET = iota
	PROJ_ROCKET
	PROJ_LIGHTNING
	PROJ_BIG
	PROJ_ANNIHILATOR
)

type projectile struct {
	code               int
	centerX, centerY   int
	faceX, faceY       int
	currentFrameNumber int
	nextTickToMove     int
	tickToExpire       int // when projectile.duration ends
	owner              *tank
	markedToRemove     bool
}

func (p *projectile) getStats() *projectileStats {
	return projStatsList[p.code]
}

func (p *projectile) getRadius() int {
	return p.getStats().damage
}

func (p *projectile) canMoveNow() bool {
	return gameTick >= p.nextTickToMove
}

func (p *projectile) getCenterCoords() (int, int) {
	return p.centerX, p.centerY
}

type projectileStats struct {
	moveDelay                      int
	duration                       int // how many frames does it live
	sprites                        *spriteAtlas
	effectOnDestroy                int
	radius, speed, acceleratesEach int
	damage                         int
	canDestroyArmor                bool
}

var projStatsList map[int]*projectileStats

func initProjectileStatsList() {
	projStatsList = map[int]*projectileStats{
		PROJ_BULLET: {
			sprites:         projectileAtlaces[PROJ_BULLET],
			damage:          1,
			speed:           4,
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
			duration:        200,
		},
		PROJ_ROCKET: {
			sprites:         projectileAtlaces[PROJ_ROCKET],
			damage:          2,
			speed:           1,
			acceleratesEach: 45,
			duration:        200,
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
			canDestroyArmor: true,
		},
		PROJ_LIGHTNING: {
			sprites:         projectileAtlaces[PROJ_LIGHTNING],
			damage:          1,
			speed:           3,
			duration:        200,
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
		},
		PROJ_BIG: {
			sprites:         projectileAtlaces[PROJ_BIG],
			damage:          3,
			speed:           1,
			duration:        500,
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 2,
		},
		PROJ_ANNIHILATOR: {
			sprites:         projectileAtlaces[PROJ_ANNIHILATOR],
			damage:          1,
			speed:           2,
			duration:        200,
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
		},
	}
}
