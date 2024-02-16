package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func ScaleVector3(vec3 rl.Vector3, scale float32) rl.Vector3{
	return rl.NewVector3(vec3.X*scale, vec3.Y*scale, vec3.Z*scale)
}
