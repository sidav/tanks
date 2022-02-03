package main

import (
	"github.com/sidav/golibrl/random"
)

const TICKS_TO_EXPIRE_BONUS = 600

func (b *battlefield) spawnRandomBonus() {
	if rnd.OneChanceFrom(b.bonusSpawnPeriod) {
		x, y := b.getRandomEmptyTileCoords(0, MAP_W-1, 0, MAP_H-1)
		x, y = tileCoordsToPhysicalCoords(x, y)
		bonus := random.RandInRange(BONUS_HELM, BONUS_GUN)
		newBonus := &event{
			code:               bonus,
			centerX:            x,
			centerY:            y,
			currentFrameNumber: 0,
			tickToExpire:       gameTick + TICKS_TO_EXPIRE_BONUS,
			owner:              nil,
		}
		b.bonuses = append(b.bonuses, newBonus)
	}
}

func (b *battlefield) iterateBonuses() {
	for bindex := len(b.bonuses) - 1; bindex >= 0; bindex-- {
		bon := b.bonuses[bindex]

		if bon.tickToExpire <= gameTick {
			b.bonuses = append(b.bonuses[:bindex], b.bonuses[bindex+1:]...)
			continue
		}

		tx, ty := bon.getCenterCoords()
		picker := b.getTankPresentFromRadius(bon.getRadius(), tx, ty)
		if picker != nil {
			b.applyBonusEffect(picker, bon.code)
			b.bonuses = append(b.bonuses[:bindex], b.bonuses[bindex+1:]...)
		}
	}
}

func (b *battlefield) applyBonusEffect(picker *tank, bonusCode int) {
	switch bonusCode {
	case BONUS_HELM:
		newCode := picker.getStats().codeOfNextArmorUpgrade
		if newCode == 0 {
			newCode = getRandomCode()
		}
		picker.code = newCode
		picker.hitpoints = picker.getBodyStats().maxHp
	case BONUS_CLOCK:
		for _, t := range b.tanks {
			if t.faction != picker.faction {
				t.nextTickToMove = gameTick + 500
			}
		}
	case BONUS_SHOVEL:
		b.placeTilesRandomSymmetric(TILE_WALL, 2)
	case BONUS_STAR:
		for i := 0; i < 2; i++ {
			t := newTank(TANK_T6, 0, 0, picker.faction)
			t.ai = initSimpleTankAi()
			b.spawnTankInRect(t, 0, MAP_W-1, 0, MAP_H-1)
		}
	case BONUS_GRENADE:
		for _, t := range b.tanks {
			if t.faction != picker.faction && !t.playerControlled {
				b.dealDamageToTank(t, 10)
				cx, cy := t.getCenterCoords()
				b.spawnEffect(EFFECT_EXPLOSION, cx, cy, picker)
			}
		}
	case BONUS_TANK:
		t := newTank(TANK_T4, 0, 0, picker.faction)
		t.ai = initSimpleTankAi()
		b.spawnTankInRect(t, 0, MAP_W-1, 0, MAP_H-1)
	case BONUS_GUN:
		weaponCode := getRandomWeaponCode()
		if picker.playerControlled {
			for _, w := range picker.weapons {
				if w.code == weaponCode {
					w.ammo += w.getStats().defaultAmmo
					return
				}
			}
			picker.weapons = append(picker.weapons, newWeapon(weaponCode))
		} else {
			picker.weapons[0].code = weaponCode
		}
	}
}
