package model

import rl "github.com/gen2brain/raylib-go/raylib"

type BaseModel struct {
	B_Model    rl.Model
	B_Texture  rl.Texture2D
	B_Position rl.Vector3
	B_Scale    float32
}

func Init(ModelPath string, TexturePath string, Posistion rl.Vector3, Scale float32) BaseModel {
	var Data BaseModel
	Data.B_Model = rl.LoadModel(ModelPath)
	Data.B_Texture = rl.LoadTexture(TexturePath)
	rl.SetMaterialTexture(Data.B_Model.Materials, rl.MapDiffuse, Data.B_Texture) // Set map diffuse texture
	Data.B_Scale = Scale
	return Data
}

func Process(Data BaseModel) {
	rl.DrawModel(Data.B_Model, Data.B_Position, Data.B_Scale, rl.White)
}

func CleanUp(Data BaseModel) {
	rl.UnloadTexture(Data.B_Texture)
	rl.UnloadModel(Data.B_Model)
}
