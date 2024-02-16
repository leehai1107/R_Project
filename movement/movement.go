package movement

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleMovement(pos rl.Vector3) rl.Vector3 {
	/*TODO: Handle movement with keyboard for player */

  if rl.IsKeyDown(rl.KeyJ) {
		pos.Z -= 0.2
	}
	if rl.IsKeyDown(rl.KeyL) {
		pos.Z += 0.2
	}
	if rl.IsKeyDown(rl.KeyI) {
		pos.X -= 0.2
	}
	if rl.IsKeyDown(rl.KeyK) {
		pos.X += 0.2
	}
  return pos
}
