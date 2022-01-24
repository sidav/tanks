package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

func handleSinglePlayer() {
	currTank := gameMap.playerTanks[0]
	if currTank == nil {
		return
	}
	if currTank.canMoveNow() {
		if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
			gameMap.moveTankByVector(currTank, 1, 0)
		} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
			gameMap.moveTankByVector(currTank, -1, 0)
		} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
			gameMap.moveTankByVector(currTank, 0, -1)
		} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
			gameMap.moveTankByVector(currTank, 0, 1)
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
			if currTank.canMoveNow() {
				if rl.IsKeyDown(rl.KeyD) {
					gameMap.moveTankByVector(currTank, 1, 0)
				} else if rl.IsKeyDown(rl.KeyA) {
					gameMap.moveTankByVector(currTank, -1, 0)
				} else if rl.IsKeyDown(rl.KeyW) {
					gameMap.moveTankByVector(currTank, 0, -1)
				} else if rl.IsKeyDown(rl.KeyS) {
					gameMap.moveTankByVector(currTank, 0, 1)
				}
			}
			if currTank.canShootNow() {
				if rl.IsKeyDown(rl.KeyF) {
					gameMap.shootAsTank(currTank)
				}
			}
		} else if i == 1 {
			if currTank.canMoveNow() {
				if rl.IsKeyDown(rl.KeyRight) {
					gameMap.moveTankByVector(currTank, 1, 0)
				} else if rl.IsKeyDown(rl.KeyLeft) {
					gameMap.moveTankByVector(currTank, -1, 0)
				} else if rl.IsKeyDown(rl.KeyUp) {
					gameMap.moveTankByVector(currTank, 0, -1)
				} else if rl.IsKeyDown(rl.KeyDown) {
					gameMap.moveTankByVector(currTank, 0, 1)
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
