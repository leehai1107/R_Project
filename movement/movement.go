package movement

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleMovement(pos rl.Vector3) rl.Vector3 {
	/*TODO: Handle movement with keyboard for player */

  if rl.IsKeyDown(rl.KeyJ) {
		pos.X -= 0.2
	}
	if rl.IsKeyDown(rl.KeyL) {
		pos.X += 0.2
	}
	if rl.IsKeyDown(rl.KeyI) {
		pos.Z -= 0.2
	}
	if rl.IsKeyDown(rl.KeyK) {
		pos.Z += 0.2
	}
  return pos
}
