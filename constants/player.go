package constants

import rl "github.com/gen2brain/raylib-go/raylib"

const ModelPath string = "res/models/obj/bunny.obj"
const TexturePath string = "res/models/texture/grass.png"

var Position rl.Vector3 = rl.NewVector3(0.0, 0.0, 0.0) // Set model position
const Scale float32 = 5

var Health = 100
var Mana = 100
var Speed = 5
var MoveSpeed float32 = 0.2
