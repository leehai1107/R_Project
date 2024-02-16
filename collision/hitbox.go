package collision

import rl "github.com/gen2brain/raylib-go/raylib"

type HitBox interface{
  GetHitBox() rl.BoundingBox
  SetHitBox(min rl.Vector3, max rl.Vector3)
  CheckHitBoxes(firstObj,secondObj rl.BoundingBox)bool
}

type hitBox struct {
	boundBox rl.BoundingBox
}

func NewHitBox(min rl.Vector3, max rl.Vector3) HitBox {
	return &hitBox{
    boundBox: rl.NewBoundingBox(min, max),
  }
}

func(h *hitBox) GetHitBox() rl.BoundingBox{
  return h.boundBox
}

func(h *hitBox) SetHitBox(min rl.Vector3, max rl.Vector3){
  h.boundBox = rl.NewBoundingBox(min, max)
}

func(h *hitBox) CheckHitBoxes(firstObj, secondObj rl.BoundingBox) bool{
  return rl.CheckCollisionBoxes(firstObj,secondObj)
}
