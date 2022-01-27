package main

import (
	"github.com/sidav/golibrl/random/additive_random"
	"time"
)

var gameTick = 0
var gameWon = false
var gameOver = false
var rnd additive_random.FibRandom
var render renderer

func runGame() {
	tickStart := time.Now()
	debugWritef("%d: START ", gameTick)
	listenPlayerInput()

	debugWrite("PROJS ")
	gameMap.actForProjectiles()

	debugWrite("EFFECTS ")
	gameMap.actForEffects()

	debugWrite("AI {")
	for i := range gameMap.tanks {
		if gameMap.tanks[i].playerControlled {
			continue
		}
		gameMap.actAiForTank(gameMap.tanks[i])
	}
	debugWrite("} ")

	debugWrite("SPAWN ")
	if len(gameMap.tanks) < gameMap.maxTanksOnMap && rnd.OneChanceFrom(gameMap.chanceToSpawnEnemyEachTickOneFrom) && gameMap.totalTanksRemainingToSpawn > 0 {
		gameMap.spawnRandomTankInRect(0, MAP_W-1, 0, MAP_H-MAP_H/5)
	}

	gameOver = true
	for i := range gameMap.playerTanks {
		if gameMap.playerTanks[i] != nil {
			gameOver = false
		}
	}
	if gameOver {
		gameMap.maxTanksOnMap = MAP_W*MAP_H/2
		gameMap.totalTanksRemainingToSpawn = MAP_W*MAP_H/2
		gameMap.chanceToSpawnEnemyEachTickOneFrom = 5
	}
	gameWon = gameWon || gameMap.isMissionCompleted()
	if gameWon {
		for i := 0; i < 3; i++ {
			gameMap.spawnEffect(EFFECT_EXPLOSION, rnd.Rand(WINDOW_W), rnd.Rand(WINDOW_H), nil)
			gameMap.spawnEffect(EFFECT_BIG_EXPLOSION, rnd.Rand(WINDOW_W), rnd.Rand(WINDOW_H), nil)
		}
	}
	gameMap.PerformMissionSpecificActions()

	logicDuration := time.Since(tickStart)

	debugWrite("RENDER ")
	renderStart := time.Now()
	render.renderBattlefield(gameMap)
	renderDuration := time.Since(renderStart)

	debugWrite("TURN FINISHED. \n")
	gameTick++
	debugWritef("TICK DURATION %v (LOGIC %v, RENDER %v) - approx. %v FPS\n", time.Since(tickStart), logicDuration, renderDuration, int(time.Second/time.Since(tickStart)))
}
