package entity

import (
	"main/constants"
	"main/model"
	"main/stats"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	model model.BaseModel
	stat  stats.StaticStat
}

// NewPlayer creates a new instance of Player with initial values
func NewPlayer(modelPath string, texturePath string, posistion rl.Vector3, scale float32) *Player {
	data := Player{
		model: model.Init(modelPath, texturePath, posistion, scale),
		stat:  *stats.NewStaticStat(constants.Health, constants.Mana, constants.Speed),
	}

	return &data
}

// SetModel allows you to set the Model from outside the package
func (p *Player) SetModel(newModel model.BaseModel) {
	p.model = newModel
}

// GetModel allows you to get the Model from outside the package
func (p *Player) GetModel() model.BaseModel {
	return p.model
}

// SetStat allows you to set the Stat from outside the package
func (p *Player) SetStat(newStat stats.StaticStat) {
	p.stat = newStat
}

// GetStat allows you to get the Stat from outside the package
func (p *Player) GetStat() stats.StaticStat {
	return p.stat
}

// SetModelPosition allows you to set the Model.Position from outside the package
func (p *Player) SetModelPosition(newPosition rl.Vector3) {
	p.model.B_Position = newPosition
}

// GetModelPosition allows you to get the Model.Position from outside the package
func (p *Player) GetModelPosition() rl.Vector3 {
	return p.model.B_Position
}

func Process(data *Player, camera rl.Camera) {
	// rl.DrawModelWires(data.GetModel().B_Model, data.GetModel().B_Position, data.GetModel().B_Scale, rl.Black)
	rl.DrawModel(data.GetModel().B_Model, data.GetModel().B_Position, data.GetModel().B_Scale, rl.White)
}

func CleanUp(data *Player) {
	rl.UnloadModel(data.model.B_Model)
	rl.UnloadTexture(data.model.B_Texture)
}
