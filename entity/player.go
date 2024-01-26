package entity

import (
	"main/model"
	"main/stats"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Model model.BaseModel
	Stat  stats.StaticStat
}

func Init(modelPath string, texturePath string, posistion rl.Vector3, scale float32) (data Player) {
	data.Model.B_Model = rl.LoadModel(modelPath)
	data.Model.B_Texture = rl.LoadTexture(texturePath)
	rl.SetMaterialTexture(data.Model.B_Model.Materials, rl.MapDiffuse, data.Model.B_Texture) // Set map diffuse texture
	data.Model.B_Position = posistion
	data.Model.B_Scale = scale
	data.Stat.Health = 100
	data.Stat.Mana = 100
	data.Stat.Speed = 300

	// Return an empty slice if the model failed to load
	return data
}

func Process(data Player) {
	rl.DrawModel(data.Model.B_Model, data.Model.B_Position, data.Model.B_Scale, rl.White)
}

func CleanUp(data Player) {
	rl.UnloadModel(data.Model.B_Model)
	rl.UnloadTexture(data.Model.B_Texture)
}
