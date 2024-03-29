package entity

import (
	"fmt"
	"main/collision"
	cts "main/constants"
	"main/model"
	"main/movement"
	f "main/pathfinder"
	"main/picker"
	"main/stats"

	// "main/stats"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player interface {
	Process()
	KeyboardMovement()
	MouseMovement(camera rl.Camera)
	DebugMode(mode bool) bool
	moveObjectAlongPath(direction rl.Vector3)
	moveAlongPath()
	CleanUp()
	HandleCollison(obj rl.BoundingBox)
}

type player struct {
	model  model.BaseModel
	stat   stats.StaticStat
	hitBox collision.HitBox
}

// NewPlayer creates a new instance of Player with initial values
func NewPlayer() Player {
	return &player{
		model:  model.NewBaseModel(cts.ModelPath, cts.TexturePath, cts.Position, cts.Scale),
		stat:   stats.NewStaticStat(cts.Health, cts.Mana, cts.MoveSpeed),
		hitBox: collision.NewHitBox(cts.Vec3Zero, cts.Vec3Zero),
	}
}

func (p *player) Process() {
	p.model.Process(p.model.GetModel(), p.model.GetPosition(), p.model.GetScale())
}

func (p *player) MouseMovement(camera rl.Camera) {
	var (
		g0 = rl.NewVector3(p.model.GetPosition().X-cts.GridSize, 0.0, p.model.GetPosition().Z-cts.GridSize)
		g1 = rl.NewVector3(p.model.GetPosition().X-cts.GridSize, 0.0, p.model.GetPosition().Z+cts.GridSize)
		g2 = rl.NewVector3(p.model.GetPosition().X+cts.GridSize, 0.0, p.model.GetPosition().Z+cts.GridSize)
		g3 = rl.NewVector3(p.model.GetPosition().X+cts.GridSize, 0.0, p.model.GetPosition().Z-cts.GridSize)
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

func (p *player) KeyboardMovement() {
	// TODO: KeyboardMovement() here
	p.model.SetPosition(movement.HandleMovement(p.model.GetPosition(),p.stat.GetSpeed()))
}

// moveAlongPath moves the player along the calculated path.
func (p *player) moveAlongPath() {
	direction := rl.Vector3Subtract(f.GetPath()[0], f.GetcurrentPos())
	distance := rl.Vector3Length(direction)

	if distance > f.GetMoveSpeed() {
		p.moveObjectAlongPath(direction)
	} else {
		f.SetPath(f.GetPath()[1:])
	}
}

func (p *player) moveObjectAlongPath(direction rl.Vector3) {
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
	p.model.SetPosition(f.GetcurrentPos())
}

func (p *player) DebugMode(mode bool) bool {
	if mode {
		min := rl.NewVector3(p.model.GetPosition().X-1, p.model.GetPosition().Y, p.model.GetPosition().Z-1)
		max := rl.NewVector3(p.model.GetPosition().X+1, p.model.GetPosition().Y+2, p.model.GetPosition().Z+1)
		p.hitBox.SetHitBox(min, max)

		rl.DrawBoundingBox(p.hitBox.GetHitBox(), rl.Green)
		return true
	}
	return false
}

func (p *player) CleanUp() {
	p.model.CleanUp(p.model.GetModel(), p.model.GetTexture())
}

//TODO: implement collison for player and tree

func (p *player) HandleCollison(obj rl.BoundingBox) {
	if p.hitBox.CheckHitBoxes(p.hitBox.GetHitBox(), obj) {
		fmt.Println("On Hit!")
	}
}
