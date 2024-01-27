package windows

import (
	camera "main/camera"
	cts "main/constants"
	"main/entity"
	"main/pathfinder"
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

	playerData := entity.Init(cts.ModelPath, cts.TexturePath, cts.Position, cts.Scale)
	camera := camera.InitCamera3D()
	pathfinder.Init(&playerData)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraThirdPerson) // Update camera with free camera mode
		if rl.IsKeyDown(rl.KeyZ) {
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
		}
		var x = 0
		x++
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawText(cts.Version, cts.Border, int32(rl.GetScreenHeight())-cts.FontSize*2-cts.Border-cts.Padding, cts.FontSize, rl.Black) // Version
		rl.DrawText(cts.CopyRight, cts.Border, int32(rl.GetScreenHeight())-cts.FontSize-cts.Border, cts.FontSize, rl.Black)             //Copyright

		rl.BeginMode3D(camera)
		world.CreateWorld()
		entity.Process(&playerData, camera)
		pathfinder.Process(camera, &playerData)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		//--------------------------------------------------------------------------------------
		rl.EndDrawing()
	}
	defer entity.CleanUp(playerData)
}
