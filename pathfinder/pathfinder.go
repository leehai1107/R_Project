package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
	moveSpeed    = 0.05
)

var (
	playerPos = rl.NewVector3(2.5, 0, 2.5)
	targetPos = rl.NewVector3(0, 0, 0)
	path      []rl.Vector3
)

// RayHitInfo represents information about a ray hit.
type RayHitInfo struct {
	Hit      bool
	Position rl.Vector3
	Normal   rl.Vector3
	Distance float32
}

// Node represents a node in the grid for A* pathfinding
type Node struct {
	Position     rl.Vector3
	GCost, HCost float64
	Parent       *Node
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "3D Pathfinding")

	camera := rl.Camera{
		Position:   rl.NewVector3(5, 4, 5),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)
		if rl.IsKeyDown(rl.KeyZ) {
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
		}
		update(camera)
		draw(camera)
	}

	rl.CloseWindow()
}

func update(camera rl.Camera) {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		handleMouseInput(camera)
	}

	if len(path) > 0 {
		moveAlongPath()
	}
}

func handleMouseInput(camera rl.Camera) {
	ray := rl.GetMouseRay(rl.GetMousePosition(), camera)
	rayHit := getRayHitInfo(ray)

	if rayHit.Hit {
		targetPos = rayHit.Position
		path = findPath(playerPos, targetPos)
	}
}

func moveAlongPath() {
	direction := rl.Vector3Subtract(path[0], playerPos)
	distance := rl.Vector3Length(direction)

	if distance > moveSpeed {
		direction = rl.Vector3Normalize(direction)
		playerPos = rl.Vector3Add(playerPos, rl.Vector3Scale(direction, moveSpeed))
	} else {
		path = path[1:]
	}
}

func getRayHitInfo(ray rl.Ray) RayHitInfo {
	groundNormal := rl.NewVector3(0, 1, 0)
	groundDistance := 0.0

	denom := rl.Vector3DotProduct(ray.Direction, groundNormal)
	epsilon := math.Nextafter(0, 1)
	if math.Abs(float64(denom)) > epsilon {
		t := -(rl.Vector3DotProduct(ray.Position, groundNormal) + float32(groundDistance)) / denom
		hitPoint := rl.Vector3Add(ray.Position, rl.Vector3Scale(ray.Direction, t))
		return RayHitInfo{Hit: true, Position: hitPoint, Normal: groundNormal, Distance: t}
	}

	return RayHitInfo{Hit: false}
}

func draw(camera rl.Camera) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)
	{
		rl.DrawSphere(playerPos, 0.2, rl.Red)
		rl.DrawSphere(targetPos, 0.2, rl.Green)
		rl.DrawGrid(100, 1)

		if len(path) > 1 {
			drawPath()
		}
	}
	rl.EndMode3D()

	rl.DrawText("Click to set pathfinding target", 10, 10, 20, rl.DarkGray)

	rl.EndDrawing()
}

func drawPath() {
	for i := 0; i < len(path)-1; i++ {
		rl.DrawLine3D(path[i], path[i+1], rl.DarkGray)
	}
}

func findPath(start, target rl.Vector3) []rl.Vector3 {
	startNode := getNodeFromWorldPos(start)
	targetNode := getNodeFromWorldPos(target)

	openSet := make(map[rl.Vector3]*Node)
	closedSet := make(map[rl.Vector3]*Node)

	openSet[startNode.Position] = startNode

	for len(openSet) > 0 {
		current := getCurrentNode(openSet)

		delete(openSet, current.Position)
		closedSet[current.Position] = current

		if current.Position == targetNode.Position {
			return reconstructPath(current)
		}

		neighbors := getNeighbors(current)
		for _, neighbor := range neighbors {
			updateNeighbor(neighbor, current, targetNode, openSet, closedSet)
		}
	}

	return []rl.Vector3{}
}

func getCurrentNode(openSet map[rl.Vector3]*Node) *Node {
	var current *Node
	for _, node := range openSet {
		if current == nil || (node.GCost+node.HCost) < (current.GCost+current.HCost) {
			current = node
		}
	}
	return current
}

func reconstructPath(current *Node) []rl.Vector3 {
	path := make([]rl.Vector3, 0)
	for current != nil {
		path = append([]rl.Vector3{current.Position}, path...)
		current = current.Parent
	}
	return path
}

func updateNeighbor(neighbor *Node, current, targetNode *Node, openSet, closedSet map[rl.Vector3]*Node) {
	if closedSet[neighbor.Position] != nil {
		return
	}

	tentativeGCost := float32(current.GCost) + rl.Vector3Distance(current.Position, neighbor.Position)

	if openSet[neighbor.Position] == nil || tentativeGCost < float32(openSet[neighbor.Position].GCost) {
		neighbor.GCost = float64(tentativeGCost)
		neighbor.HCost = float64(rl.Vector3Distance(neighbor.Position, targetNode.Position))
		neighbor.Parent = current

		openSet[neighbor.Position] = neighbor
	}
}

func getNodeFromWorldPos(pos rl.Vector3) *Node {
	gridPos := rl.NewVector3(float32(math.Floor(float64(pos.X))),
		float32(math.Floor(float64(pos.Y))),
		float32(math.Floor(float64(pos.Z))))
	return &Node{Position: gridPos}
}

func getNeighbors(node *Node) []*Node {
	neighbors := make([]*Node, 0)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				neighborPos := rl.NewVector3(node.Position.X+float32(x), node.Position.Y+float32(y), node.Position.Z+float32(z))
				neighbors = append(neighbors, &Node{Position: neighborPos})
			}
		}
	}
	return neighbors
}
