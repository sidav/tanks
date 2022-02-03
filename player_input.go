package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WEAPON_CHANGE_DELAY = 20

func listenPlayerInput() {
	if gameOver {
		if rl.IsKeyDown(rl.KeyEscape) {
			gameIsRunning = false
		}
		return
	}
	if gameMap.numPlayers < 2 {
		handleSinglePlayer()
	} else {
		handleAllPlayers()
	}
	if rl.IsKeyDown(rl.KeyEscape) {
		gameIsRunning = false
	}
}

var lastKeyPressed int32

func handleSinglePlayer() {
	currTank := gameMap.playerTanks[0]
	if currTank == nil {
		return
	}
	currKeyPressed := rl.GetKeyPressed()
	movementKeysToConsider := []int32{rl.KeyRight, rl.KeyLeft, rl.KeyUp, rl.KeyDown, rl.KeyW, rl.KeyS, rl.KeyA, rl.KeyD}
	for _, k := range movementKeysToConsider {
		if currKeyPressed == k || rl.IsKeyReleased(lastKeyPressed) {
			lastKeyPressed = currKeyPressed
			break
		}
	}

	if lastKeyPressed == rl.KeyRight || lastKeyPressed == rl.KeyD {
		gameMap.moveTankByVector(currTank, 1, 0)
	} else if lastKeyPressed == rl.KeyLeft || lastKeyPressed == rl.KeyA {
		gameMap.moveTankByVector(currTank, -1, 0)
	} else if lastKeyPressed == rl.KeyUp || lastKeyPressed == rl.KeyW {
		gameMap.moveTankByVector(currTank, 0, -1)
	} else if lastKeyPressed == rl.KeyDown || lastKeyPressed == rl.KeyS {
		gameMap.moveTankByVector(currTank, 0, 1)
	}

	if currTank.canMoveNow() {
		if rl.IsKeyDown(rl.KeyQ) || rl.IsKeyDown(rl.KeyRightShift) {
			currTank.currentWeaponNumber = (currTank.currentWeaponNumber + 1) % len(currTank.weapons)
			currTank.nextTickToMove = gameTick + WEAPON_CHANGE_DELAY
		}
	}
	if currTank.canShootNow() {
		if rl.IsKeyDown(rl.KeySpace) || rl.IsKeyDown(rl.KeyF) {
			gameMap.shootAsTank(currTank)
		}
	}
}

func handleAllPlayers() {
	for i := 0; i < len(gameMap.playerTanks); i++ {
		currTank := gameMap.playerTanks[i]
		if currTank == nil {
			continue
		}
		if i == 0 {
			if rl.IsKeyDown(rl.KeyD) {
				gameMap.moveTankByVector(currTank, 1, 0)
			} else if rl.IsKeyDown(rl.KeyA) {
				gameMap.moveTankByVector(currTank, -1, 0)
			} else if rl.IsKeyDown(rl.KeyW) {
				gameMap.moveTankByVector(currTank, 0, -1)
			} else if rl.IsKeyDown(rl.KeyS) {
				gameMap.moveTankByVector(currTank, 0, 1)
			}
			if currTank.canMoveNow() {
				if rl.IsKeyDown(rl.KeyQ) {
					currTank.currentWeaponNumber = (currTank.currentWeaponNumber + 1) % len(currTank.weapons)
					currTank.nextTickToMove = gameTick + WEAPON_CHANGE_DELAY
				}
			}
			if currTank.canShootNow() {
				if rl.IsKeyDown(rl.KeyF) {
					gameMap.shootAsTank(currTank)
				}
			}
		} else if i == 1 {
			if rl.IsKeyDown(rl.KeyRight) {
				gameMap.moveTankByVector(currTank, 1, 0)
			} else if rl.IsKeyDown(rl.KeyLeft) {
				gameMap.moveTankByVector(currTank, -1, 0)
			} else if rl.IsKeyDown(rl.KeyUp) {
				gameMap.moveTankByVector(currTank, 0, -1)
			} else if rl.IsKeyDown(rl.KeyDown) {
				gameMap.moveTankByVector(currTank, 0, 1)
			}
			if currTank.canMoveNow() {
				if rl.IsKeyDown(rl.KeyRightShift) {
					currTank.currentWeaponNumber = (currTank.currentWeaponNumber + 1) % len(currTank.weapons)
					currTank.nextTickToMove = gameTick + WEAPON_CHANGE_DELAY
				}
			}
			if currTank.canShootNow() {
				if rl.IsKeyDown(rl.KeySpace) {
					gameMap.shootAsTank(currTank)
				}
			}
		}
	}
}
