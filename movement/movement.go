package movement

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleMovement(pos rl.Vector3, speed float32) rl.Vector3 {
	/*TODO: Handle movement with keyboard for player */

  if rl.IsKeyDown(rl.KeyJ) {
		pos.Z -= speed
	}
	if rl.IsKeyDown(rl.KeyL) {
		pos.Z += speed
	}
	if rl.IsKeyDown(rl.KeyI) {
		pos.X -= speed
	}
	if rl.IsKeyDown(rl.KeyK) {
		pos.X += speed
	}
  return pos
}
