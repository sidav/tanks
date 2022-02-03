package main

const (
	EFFECT_EXPLOSION = iota
	EFFECT_BIG_EXPLOSION
	EFFECT_SPAWN

	BONUS_HELM
	BONUS_CLOCK
	BONUS_SHOVEL
	BONUS_STAR
	BONUS_GRENADE
	BONUS_TANK
	BONUS_GUN
)

// Thing is a projectile or special effect
type event struct {
	code             int
	centerX, centerY int
	//faceX, faceY       int
	currentFrameNumber int
	nextTickToMove     int
	tickToExpire       int // when event.duration ends
	owner              *tank
}

func (p *event) getStats() *eventStats {
	return eventStatsList[p.code]
}

func (p *event) getRadius() int {
	return p.getStats().damage
}

func (p *event) canMoveNow() bool {
	return gameTick >= p.nextTickToMove
}

func (p *event) getCenterCoords() (int, int) {
	return p.centerX, p.centerY
}

type eventStats struct {
	moveDelay       int
	duration        int // how many frames does it live
	sprites         *spriteAtlas
	radius, speed   int
	damage          int
	canDestroyArmor bool
}

var eventStatsList map[int]*eventStats

func initEventsStatsList() {
	eventStatsList = map[int]*eventStats{
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

		BONUS_HELM: {
			sprites:   bonusAtlaces[BONUS_HELM],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},
		BONUS_CLOCK: {
			sprites:   bonusAtlaces[BONUS_CLOCK],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},
		BONUS_SHOVEL: {
			sprites:   bonusAtlaces[BONUS_SHOVEL],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},
		BONUS_STAR: {
			sprites:   bonusAtlaces[BONUS_STAR],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},
		BONUS_GRENADE: {
			sprites:   bonusAtlaces[BONUS_GRENADE],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},
		BONUS_TANK: {
			sprites:   bonusAtlaces[BONUS_TANK],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},
		BONUS_GUN: {
			sprites:   bonusAtlaces[BONUS_GUN],
			radius:    halfPhysicalTileSize() - 1,
			moveDelay: 7,
		},
	}
}
