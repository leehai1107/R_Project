package collision

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PLayerCollision struct {
	shape rl.BoundingBox
	debug bool
}

func DrawCollison(shape rl.BoundingBox) {
	rl.DrawBoundingBox(shape, rl.Green)
}
