package picker

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Process(camera rl.Camera, g0, g1, g2, g3 rl.Vector3) rl.RayCollision {
	ray := rl.GetMouseRay(rl.GetMousePosition(), camera)
	rayHit := rl.GetRayCollisionQuad(ray, g0, g1, g2, g3)
	return rayHit
}
