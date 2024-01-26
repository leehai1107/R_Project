package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func InitCamera3D() rl.Camera3D {
	// Define the camera to look into our 3d world
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective
	return camera
}

func UpdateCamera(camera *rl.Camera) {

}
