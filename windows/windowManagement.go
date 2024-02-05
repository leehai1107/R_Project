package windows

import (
	camera "main/camera"
	cts "main/constants"
	"main/entity"
	world "main/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	rl.InitWindow(cts.ScreenWidth, cts.ScreenHeight, cts.Title)
	rl.SetTargetFPS(60)

}

func Close() {
	rl.CloseWindow()
}

func Process() {

	playerData := entity.NewPlayer(cts.ModelPath, cts.TexturePath, cts.Position, cts.Scale)
	camera := camera.InitCamera3D()

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraThirdPerson) // Update camera with free camera mode
		if rl.IsKeyDown(rl.KeyZ) {
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
		}
		playerData.Movement(camera)
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawText(cts.Version, cts.Border, int32(rl.GetScreenHeight())-cts.FontSize*2-cts.Border-cts.Padding, cts.FontSize, rl.Black) // Version
		rl.DrawText(cts.CopyRight, cts.Border, int32(rl.GetScreenHeight())-cts.FontSize-cts.Border, cts.FontSize, rl.Black)             //Copyright

		rl.BeginMode3D(camera)
		world.CreateWorld()
		entity.Process(playerData)
    playerData.DrawCollision()

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		//--------------------------------------------------------------------------------------
		rl.EndDrawing()
	}
	defer entity.CleanUp(playerData)
}
