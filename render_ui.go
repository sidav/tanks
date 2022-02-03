package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

func (r *renderer) renderUI(playerNum int, b *battlefield, t *tank) {
	posX := int32(playerNum * r.viewportW)
	rl.DrawRectangleGradientV(posX, WINDOW_H-2*TEXT_MARGIN-TEXT_SIZE, int32(r.viewportW), TEXT_MARGIN*2+TEXT_SIZE,
		color.RGBA{
			R: 64,
			G: 64,
			B: 64,
			A: 255,
		},
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 64,
		},
	)
	if !gameOver {
		posX = int32(playerNum*r.viewportW) / 2
		if r.viewportW < WINDOW_W {
			posX *= 2
		}
		uiStr := fmt.Sprintf("HP %d/%d Enemies %d", t.hitpoints, t.getBodyStats().maxHp, b.totalTanksRemainingToSpawn)
		rl.DrawText(uiStr, posX, WINDOW_H-TEXT_MARGIN-TEXT_SIZE, TEXT_SIZE,
			color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
	}

	if b.numPlayers > 1 && !r.doesLevelFitInScreenHorizontally() {
		separatorWidth := int32(TILE_PHYSICAL_SIZE / 2)
		rl.DrawRectangleGradientV(WINDOW_W/2-separatorWidth/2, 0, separatorWidth, WINDOW_H,
			color.RGBA{
				R: 32,
				G: 32,
				B: 32,
				A: 255,
			},
			color.RGBA{
				R: 64,
				G: 64,
				B: 64,
				A: 255,
			},
		)
	}
}
