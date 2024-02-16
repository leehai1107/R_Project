package entity

import (
	"main/collision"
	cts "main/constants"
	"main/model"
	"main/stats"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tree interface{
  Process()
  DebugMode(mode bool)bool
  CleanUp()
  GetHitBox()rl.BoundingBox
}


type tree struct {
	model model.BaseModel
  stat stats.StaticStat
  hitBox collision.HitBox
}

func NewTree() Tree{
return &tree{
    model: model.NewBaseModel(cts.TreeModel,cts.TreeTexture,cts.TreePos,1),
    stat: stats.NewStaticStat(cts.Health, 0,0),
    hitBox: collision.NewHitBox(cts.Vec3Zero, cts.Vec3Zero),
  }
}

func(p *tree) Process(){
	p.model.Process(p.model.GetModel(), p.model.GetPosition(), p.model.GetScale())
}

func (p *tree) DebugMode(mode bool) bool {
	if mode {
		min := rl.NewVector3(p.model.GetPosition().X-1, p.model.GetPosition().Y, p.model.GetPosition().Z-1)
		max := rl.NewVector3(p.model.GetPosition().X+1, p.model.GetPosition().Y+2, p.model.GetPosition().Z+1)
		p.hitBox.SetHitBox(min, max)
    
		rl.DrawBoundingBox(p.hitBox.GetHitBox(), rl.Green)
		return true
	}
	return false
}

func (p *tree) CleanUp(){
  p.model.CleanUp(p.model.GetModel(), p.model.GetTexture())
}

func(p *tree) GetHitBox()rl.BoundingBox{
  return p.hitBox.GetHitBox()
}

