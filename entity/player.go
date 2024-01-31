package entity

import (
	cts "main/constants"
	"main/model"
	f "main/pathfinder"
	"main/picker"
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
		stat:  *stats.NewStaticStat(cts.Health, cts.Mana, cts.Speed),
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

func (p *Player) Movement(camera rl.Camera) {
	var (
		g0 = rl.NewVector3(p.GetModelPosition().X-cts.GridSize, 0.0, p.GetModelPosition().Z-cts.GridSize)
		g1 = rl.NewVector3(p.GetModelPosition().X-cts.GridSize, 0.0, p.GetModelPosition().Z+cts.GridSize)
		g2 = rl.NewVector3(p.GetModelPosition().X+cts.GridSize, 0.0, p.GetModelPosition().Z+cts.GridSize)
		g3 = rl.NewVector3(p.GetModelPosition().X+cts.GridSize, 0.0, p.GetModelPosition().Z-cts.GridSize)
	)

	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		picker := picker.Process(camera, g0, g1, g2, g3)
		if picker.Hit {
			f.SetTargetPos(picker.Point)
			f.FindPath(f.GetTargetPos())
		}
	}
	if len(f.GetPath()) > 0 {
		p.moveAlongPath()
	}
}

// moveAlongPath moves the player along the calculated path.
func (p *Player) moveAlongPath() {
	direction := rl.Vector3Subtract(f.GetPath()[0], f.GetcurrentPos())
	distance := rl.Vector3Length(direction)

	if distance > f.GetMoveSpeed() {
		p.moveObjectAlongPath(direction)
	} else {
		f.SetPath(f.GetPath()[1:])
	}

}

func (p *Player) moveObjectAlongPath(direction rl.Vector3) {
	direction = rl.Vector3Normalize(direction)
	f.SetcurrentPos(rl.Vector3Add(f.GetcurrentPos(), rl.Vector3Scale(direction, f.GetMoveSpeed())))

	// Smoothly interpolate between path points for smoother movement
	if len(f.GetPath()) > 1 {
		directionToNextPoint := rl.Vector3Subtract(f.GetPath()[0], f.GetcurrentPos())
		directionToNextPoint = rl.Vector3Normalize(directionToNextPoint)
		f.SetcurrentPos(rl.Vector3Add(f.GetcurrentPos(), rl.Vector3Scale(directionToNextPoint, f.GetMoveSpeed())))

		// Check if reached the next point
		distanceToNextPoint := rl.Vector3Distance(f.GetcurrentPos(), f.GetPath()[0])
		if distanceToNextPoint < f.GetMoveSpeed() {
			f.SetPath(f.GetPath()[1:])
		}
	}
	p.SetModelPosition(f.GetcurrentPos())
}

func CleanUp(data *Player) {
	rl.UnloadModel(data.model.B_Model)
	rl.UnloadTexture(data.model.B_Texture)
}
