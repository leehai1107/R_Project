package entity

import (
	"main/constants"
	"main/model"
	"main/stats"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Model model.BaseModel
	Stat  stats.StaticStat
}

// NewPlayer creates a new instance of Player with initial values
func NewPlayer(initialModel model.BaseModel, initialStat stats.StaticStat) *Player {
	return &Player{
		Model: initialModel,
		Stat:  initialStat,
	}
}

// SetModel allows you to set the Model from outside the package
func (p *Player) SetModel(newModel model.BaseModel) {
	p.Model = newModel
}

// GetModel allows you to get the Model from outside the package
func (p *Player) GetModel() model.BaseModel {
	return p.Model
}

// SetStat allows you to set the Stat from outside the package
func (p *Player) SetStat(newStat stats.StaticStat) {
	p.Stat = newStat
}

// GetStat allows you to get the Stat from outside the package
func (p *Player) GetStat() stats.StaticStat {
	return p.Stat
}

// SetModelPosition allows you to set the Model.Position from outside the package
func (p *Player) SetModelPosition(newPosition rl.Vector3) {
	p.Model.B_Position = newPosition
}

// GetModelPosition allows you to get the Model.Position from outside the package
func (p *Player) GetModelPosition() rl.Vector3 {
	return p.Model.B_Position
}

func Init(modelPath string, texturePath string, posistion rl.Vector3, scale float32) (data Player) {
	data.Model = initModel(modelPath, texturePath, posistion, scale)
	data.Stat = *initStat()
	// Return an empty slice if the model failed to load
	return data
}

func initModel(modelPath string, texturePath string, posistion rl.Vector3, scale float32) model.BaseModel {
	data := model.Init(modelPath, texturePath, posistion, scale)
	return data
}

func initStat() (data *stats.StaticStat) {
	data = stats.NewStaticStat(constants.Health, constants.Mana, constants.Speed)
	return data
}

func Process(data *Player, camera rl.Camera) {
	rl.DrawModel(data.Model.B_Model, data.Model.B_Position, data.Model.B_Scale, rl.White)
}

func CleanUp(data Player) {
	rl.UnloadModel(data.Model.B_Model)
	rl.UnloadTexture(data.Model.B_Texture)
}
