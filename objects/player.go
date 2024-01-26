package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	P_Model    rl.Model
	P_Texture  rl.Texture2D
	P_Position rl.Vector3
	P_Scale    float32
}

func Init(ModelPath string, TexturePath string, Posistion rl.Vector3, Scale float32) Player {
	var Player Player
	Player.P_Model = rl.LoadModel(ModelPath)
	Player.P_Texture = rl.LoadTexture(TexturePath)
	rl.SetMaterialTexture(Player.P_Model.Materials, rl.MapDiffuse, Player.P_Texture) // Set map diffuse texture
	Player.P_Scale = Scale
	return Player
}

func Process(Player Player) {
	rl.DrawModel(Player.P_Model, Player.P_Position, Player.P_Scale, rl.White)
}

func CleanUp(Player Player) {
	rl.UnloadTexture(Player.P_Texture)
	rl.UnloadModel(Player.P_Model)
}
