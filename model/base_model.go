package model

import rl "github.com/gen2brain/raylib-go/raylib"

type BaseModel struct {
	B_Model    rl.Model
	B_Texture  rl.Texture2D
	B_Position rl.Vector3
	B_Scale    float32
}

// NewBaseModel creates a new instance of BaseModel with initial values
func NewBaseModel(initialModel rl.Model, initialTexture rl.Texture2D, initialPosition rl.Vector3, initialScale float32) *BaseModel {
	return &BaseModel{
		B_Model:    initialModel,
		B_Texture:  initialTexture,
		B_Position: initialPosition,
		B_Scale:    initialScale,
	}
}

// SetModel allows you to set the B_Model from outside the package
func (bm *BaseModel) SetModel(newModel rl.Model) {
	bm.B_Model = newModel
}

// GetModel allows you to get the B_Model from outside the package
func (bm *BaseModel) GetModel() rl.Model {
	return bm.B_Model
}

// SetTexture allows you to set the B_Texture from outside the package
func (bm *BaseModel) SetTexture(newTexture rl.Texture2D) {
	bm.B_Texture = newTexture
}

// GetTexture allows you to get the B_Texture from outside the package
func (bm *BaseModel) GetTexture() rl.Texture2D {
	return bm.B_Texture
}

// SetPosition allows you to set the B_Position from outside the package
func (bm *BaseModel) SetPosition(newPosition rl.Vector3) {
	bm.B_Position = newPosition
}

// GetPosition allows you to get the B_Position from outside the package
func (bm *BaseModel) GetPosition() rl.Vector3 {
	return bm.B_Position
}

// SetScale allows you to set the B_Scale from outside the package
func (bm *BaseModel) SetScale(newScale float32) {
	bm.B_Scale = newScale
}

// GetScale allows you to get the B_Scale from outside the package
func (bm *BaseModel) GetScale() float32 {
	return bm.B_Scale
}

func Init(ModelPath string, TexturePath string, Posistion rl.Vector3, Scale float32) (data BaseModel) {
	data.B_Model = rl.LoadModel(ModelPath)
	data.B_Texture = rl.LoadTexture(TexturePath)
	rl.SetMaterialTexture(data.B_Model.Materials, rl.MapDiffuse, data.B_Texture) // Set map diffuse texture
	data.B_Scale = Scale
	return data
}

func Process(Data BaseModel) {
	rl.DrawModel(Data.B_Model, Data.B_Position, Data.B_Scale, rl.White)
}

func CleanUp(Data BaseModel) {
	rl.UnloadTexture(Data.B_Texture)
	rl.UnloadModel(Data.B_Model)
}
