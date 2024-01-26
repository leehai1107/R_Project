package terrain

import (
	cts "main/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreateDemoTerrain() {
	rl.DrawGrid(cts.Slices, cts.Spacing)
}
