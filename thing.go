package main

const (
	PROJ_BULLET = iota
	PROJ_ROCKET
	PROJ_LIGHTNING

	EFFECT_EXPLOSION
	EFFECT_BIG_EXPLOSION
	EFFECT_SPAWN
)

// Thing is a projectile or special effect
type thing struct {
	code               int
	centerX, centerY   int
	faceX, faceY       int
	currentFrameNumber int
	nextTickToMove     int
	tickToExpire       int // when thing.duration ends
	owner              *tank
	markedToRemove     bool
}

func (p *thing) getStats() *thingStats {
	return projStatsList[p.code]
}

func (p *thing) getRadius() int {
	return p.getStats().damage
}

func (p *thing) canMoveNow() bool {
	return gameTick >= p.nextTickToMove
}

func (p *thing) getCenterCoords() (int, int) {
	return p.centerX, p.centerY
}

type thingStats struct {
	moveDelay                      int
	duration                       int // how many frames does it live
	sprites                        *spriteAtlas
	effectOnDestroy                int
	radius, speed, acceleratesEach int
	damage                         int
	canDestroyArmor                bool
}

var projStatsList map[int]*thingStats

func initProjectileStatsList() {
	projStatsList = map[int]*thingStats{
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
			acceleratesEach: 60,
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

		EFFECT_EXPLOSION: {
			sprites:   effectAtlaces[EFFECT_EXPLOSION],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 3,
			duration:  20,
		},
		EFFECT_BIG_EXPLOSION: {
			sprites:   effectAtlaces[EFFECT_BIG_EXPLOSION],
			radius:    TILE_PHYSICAL_SIZE,
			moveDelay: 5,
			duration:  25,
		},
		EFFECT_SPAWN: {
			sprites:   effectAtlaces[EFFECT_SPAWN],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
			duration:  60,
		},
	}
}
