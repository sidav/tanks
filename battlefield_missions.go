package main

const (
	MISSION_KILL_ALL = iota
	MISSION_PROTECT_HQ
	MISSION_COLLECT_FLAGS
)

func (b *battlefield) isMissionCompleted() bool {
	switch b.missionType {
	case MISSION_KILL_ALL:
		return gameMap.totalTanksRemainingToSpawn <= 0 && len(gameMap.tanks) == 1 && len(gameMap.effects) == 0
	case MISSION_PROTECT_HQ:
		return gameMap.totalTanksRemainingToSpawn <= 0 && len(gameMap.tanks) == 1 && len(gameMap.effects) == 0
	case MISSION_COLLECT_FLAGS:
		for i := 0; i < MAP_W; i++ {
			for j := 0; j < MAP_H; j++ {
				if b.tiles[i][j].code == TILE_FLAG {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (b *battlefield) PerformMissionSpecificActions() {
	switch b.missionType {
	case MISSION_KILL_ALL:
		return
	case MISSION_PROTECT_HQ:
		return
	case MISSION_COLLECT_FLAGS:
		// remove flag from map if needed
		for _, t := range b.playerTanks {
			if t != nil {
				x, y := t.getCenterCoords()
				x, y = trueCoordsToTileCoords(x, y)
				if b.tiles[x][y].code == TILE_FLAG {
					b.tiles[x][y].code = TILE_EMPTY
				}
			}
		}
		// increase enemy spawn
		if gameTick % 100 == 0 && b.chanceToSpawnEnemyEachTickOneFrom > 5 {
			b.chanceToSpawnEnemyEachTickOneFrom--
		}
	}
	return
}
