package model

import rl "github.com/gen2brain/raylib-go/raylib"

type BaseModel interface {
	SetModel(newModel rl.Model)
	GetModel() rl.Model
	SetTexture(newTexture rl.Texture2D)
	GetTexture() rl.Texture2D
	SetPosition(newPosition rl.Vector3)
	GetPosition() rl.Vector3
	SetScale(newScale float32)
	GetScale() float32
	Process(model rl.Model, position rl.Vector3, scale float32)
	CleanUp(model rl.Model, texture rl.Texture2D)
}

// NewBaseModel creates a new instance of BaseModel with initial values
func NewBaseModel(modelPath string, texturePath string, posistion rl.Vector3, scale float32) BaseModel {
	model := rl.LoadModel(modelPath)
	texture := rl.LoadTexture(texturePath)
	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texture)
	return &baseModel{
		b_Model:    model,
		b_Texture:  texture,
		b_Position: posistion,
		b_Scale:    scale,
	}
}

type baseModel struct {
	b_Model    rl.Model
	b_Texture  rl.Texture2D
	b_Position rl.Vector3
	b_Scale    float32
}

// SetModel allows you to set the B_Model from outside the package
func (bm *baseModel) SetModel(newModel rl.Model) {
	bm.b_Model = newModel
}

// GetModel allows you to get the B_Model from outside the package
func (bm *baseModel) GetModel() rl.Model {
	return bm.b_Model
}

// SetTexture allows you to set the B_Texture from outside the package
func (bm *baseModel) SetTexture(newTexture rl.Texture2D) {
	bm.b_Texture = newTexture
}

// GetTexture allows you to get the B_Texture from outside the package
func (bm *baseModel) GetTexture() rl.Texture2D {
	return bm.b_Texture
}

// SetPosition allows you to set the B_Position from outside the package
func (bm *baseModel) SetPosition(newPosition rl.Vector3) {
	bm.b_Position = newPosition
}

// GetPosition allows you to get the B_Position from outside the package
func (bm *baseModel) GetPosition() rl.Vector3 {
	return bm.b_Position
}

// SetScale allows you to set the B_Scale from outside the package
func (bm *baseModel) SetScale(newScale float32) {
	bm.b_Scale = newScale
}

// GetScale allows you to get the B_Scale from outside the package
func (bm *baseModel) GetScale() float32 {
	return bm.b_Scale
}

func (bm *baseModel) Process(model rl.Model, position rl.Vector3, scale float32) {
	rl.DrawModel(model, position, scale, rl.White)
}

func (bm *baseModel) CleanUp(model rl.Model, texture rl.Texture2D) {
	rl.UnloadTexture(texture)
	rl.UnloadModel(model)
}
