package main

import rl "github.com/gen2brain/raylib-go/raylib"

func renderTank(t *tank) {
	x, y := t.getTopLeftCoordForDraw()
	rl.DrawTextureRec(
		t.sprites.atlas,
		t.getCurrentSpriteRect(),
		rl.Vector2{
			X: x,
			Y: y,
		},
		rl.White,
	)
}

func renderTile(t *tile) {

}
