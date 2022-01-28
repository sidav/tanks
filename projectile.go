package main

const (
	PROJ_BULLET = iota
	PROJ_ROCKET
	PROJ_LIGHTNING
)

type projectile struct {
	code               int
	centerX, centerY   int
	faceX, faceY       int
	currentFrameNumber int
	owner              *tank
	markedToRemove     bool
}

func (p *projectile) getStats() *projectileStats {
	return projStatsList[p.code]
}

func (p *projectile) getRadius() int {
	return p.getStats().damage
}

type projectileStats struct {
	sprites         *spriteAtlas
	effectOnDestroy int
	radius, speed   int
	damage          int
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
		},
		PROJ_ROCKET: {
			sprites:         projectileAtlaces[PROJ_ROCKET],
			damage:          2,
			speed:           2,
			effectOnDestroy: EFFECT_BIG_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
		},
		PROJ_LIGHTNING: {
			sprites:         projectileAtlaces[PROJ_LIGHTNING],
			damage:          1,
			speed:           3,
			effectOnDestroy: EFFECT_EXPLOSION,
			radius:          TILE_PHYSICAL_SIZE / 8,
		},
	}
}
