package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera interface {
	UpdateCamera()
  GetCamera() rl.Camera3D
}

type camera struct {
	camera3D rl.Camera3D
}

func NewCamera3D() Camera {
	return &camera{
		camera3D: rl.Camera3D{
			Position:   rl.NewVector3(10.0, 10.0, 10.0),
			Target:     rl.NewVector3(0.0, 0.0, 0.0),
			Up:         rl.NewVector3(0.0, 1.0, 0.0),
			Fovy:       45.0,
			Projection: rl.CameraPerspective,
		},
	}
}

func (c *camera) UpdateCamera() {
	rl.UpdateCamera(&c.camera3D, rl.CameraFree) // Update camera with free camera mode
	if rl.IsKeyDown(rl.KeyZ) {
		c.camera3D.Target = rl.NewVector3(0.0, 0.0, 0.0)
	}
}

func (c *camera) GetCamera() rl.Camera3D{
  return c.camera3D
}
