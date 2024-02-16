package windows

import (
	camera "main/camera"
	cts "main/constants"
	"main/entity"
	world "main/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Windows interface {
	Init()
	Process()
	Close()
}

type windows struct{}

func NewWindows() Windows {
	return &windows{}
}

func (w *windows) Init() {
	rl.InitWindow(cts.ScreenWidth, cts.ScreenHeight, cts.Title)
	rl.SetTargetFPS(60)
}

func (w *windows) Close() {
	rl.CloseWindow()
}

func (w *windows) Process() {
	playerData := entity.NewPlayer()
	cameraData := camera.NewCamera3D()
	treeData := entity.NewTree()

	for !rl.WindowShouldClose() {
		playerData.HandleCollison(treeData.GetHitBox())
		cameraData.UpdateCamera()
		playerData.KeyboardMovement()
		// playerData.MouseMovement(camera.NewCamera3D().GetCamera())
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawText(cts.Version, cts.Border, int32(rl.GetScreenHeight())-cts.FontSize*2-cts.Border-cts.Padding, cts.FontSize, rl.Black) // Version
		rl.DrawText(cts.CopyRight, cts.Border, int32(rl.GetScreenHeight())-cts.FontSize-cts.Border, cts.FontSize, rl.Black)             // Copyright

		rl.BeginMode3D(cameraData.GetCamera())
		world.CreateWorld()
		playerData.Process()
		treeData.Process()
		playerData.DebugMode(true)
		treeData.DebugMode(true)

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		//--------------------------------------------------------------------------------------
		rl.EndDrawing()
	}
	defer playerData.CleanUp()
	defer treeData.CleanUp()
}
