package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Constants
const (
	screenWidth  = 800
	screenHeight = 450
	moveSpeed    = 0.05
	chunkSize    = 50
)

// Variables
var (
	playerPos = rl.NewVector3(2.5, 0, 2.5)
	targetPos = rl.NewVector3(0, 0, 0)
	path      []rl.Vector3
)

// RayHitInfo struct represents information about a ray hit.
type RayHitInfo struct {
	Hit      bool
	Position rl.Vector3
	Normal   rl.Vector3
	Distance float32
}

// Node struct represents a node in the pathfinding grid.
type Node struct {
	Position     rl.Vector3
	GCost, HCost float64
	Parent       *Node
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "3D Pathfinding")

	camera := initCamera()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)
		if rl.IsKeyDown(rl.KeyZ) {
			resetCameraTarget(&camera)
		}
		update(camera)
		draw(camera)
	}

	rl.CloseWindow()
}

// initCamera initializes and returns a 3D camera.
func initCamera() rl.Camera {
	return rl.Camera{
		Position:   rl.NewVector3(5, 4, 5),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}
}

// resetCameraTarget resets the camera target to the default position.
func resetCameraTarget(camera *rl.Camera) {
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
}

// update handles input and updates the game state.
func update(camera rl.Camera) {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		handleMouseInput(camera)
	}

	if len(path) > 0 {
		moveAlongPath()
	}
}

// handleMouseInput updates the target position based on mouse input and calculates a new path.
func handleMouseInput(camera rl.Camera) {
	ray := rl.GetMouseRay(rl.GetMousePosition(), camera)
	rayHit := getRayHitInfo(ray)

	if rayHit.Hit {
		targetPos = rayHit.Position
		path = findPath(playerPos, targetPos)
	}
}

// getRayHitInfo calculates information about a ray hit on the ground.
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

// moveAlongPath moves the player along the calculated path.
func moveAlongPath() {
	direction := rl.Vector3Subtract(path[0], playerPos)
	distance := rl.Vector3Length(direction)

	if distance > moveSpeed {
		movePlayerAlongPath(direction)
	} else {
		path = path[1:]
	}
}

// movePlayerAlongPath moves the player along the given direction with a specified speed.
func movePlayerAlongPath(direction rl.Vector3) {
	direction = rl.Vector3Normalize(direction)
	playerPos = rl.Vector3Add(playerPos, rl.Vector3Scale(direction, moveSpeed))
}

// draw renders the game entities and path.
func draw(camera rl.Camera) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)
	{
		drawEntities()

		chunkCenter := getChunkCenter(playerPos)
		drawChunk(chunkCenter)

		if len(path) > 1 {
			drawPath()
		}
	}
	rl.EndMode3D()

	rl.DrawText("Click to set pathfinding target", 10, 10, 20, rl.DarkGray)

	rl.EndDrawing()
}

// drawEntities renders the player and target entities.
func drawEntities() {
	rl.DrawSphere(playerPos, 0.2, rl.Red)
	rl.DrawSphere(targetPos, 0.2, rl.Green)
}

// drawChunk renders the chunk and its bounding box.
func drawChunk(center rl.Vector3) {
	halfSize := float32(chunkSize) / 2.0
	min := rl.NewVector3(center.X-halfSize, center.Y-halfSize, center.Z-halfSize)
	max := rl.NewVector3(center.X+halfSize, center.Y+halfSize, center.Z+halfSize)

	rl.DrawGrid(chunkSize, 1)
	rl.DrawBoundingBox(rl.NewBoundingBox(min, max), rl.DarkGray)
}

// drawPath renders the path as a series of connected lines.
func drawPath() {
	for i := 0; i < len(path)-1; i++ {
		rl.DrawLine3D(path[i], path[i+1], rl.DarkGray)
	}
}

// getCurrentNode finds the node with the lowest cost in the open set.
func getCurrentNode(openSet map[rl.Vector3]*Node) *Node {
	var current *Node
	for _, node := range openSet {
		if current == nil || (node.GCost+node.HCost) < (current.GCost+current.HCost) {
			current = node
		}
	}
	return current
}

// reconstructPath reconstructs the path by traversing the parent pointers.
func reconstructPath(current *Node) []rl.Vector3 {
	path := make([]rl.Vector3, 0)
	for current != nil {
		path = append([]rl.Vector3{current.Position}, path...)
		current = current.Parent
	}
	return path
}

// getNodeFromWorldPos converts a world position to a grid node.
func getNodeFromWorldPos(pos rl.Vector3) *Node {
	gridPos := rl.NewVector3(
		float32(math.Floor(float64(pos.X))),
		float32(math.Floor(float64(pos.Y))),
		float32(math.Floor(float64(pos.Z))),
	)
	return &Node{Position: gridPos}
}

// getNeighbors returns the neighboring nodes of a given node.
func getNeighbors(node *Node) []*Node {
	neighbors := make([]*Node, 0)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				neighborPos := rl.NewVector3(
					node.Position.X+float32(x),
					node.Position.Y+float32(y),
					node.Position.Z+float32(z),
				)
				neighbors = append(neighbors, &Node{Position: neighborPos})
			}
		}
	}
	return neighbors
}

// getChunkCenter calculates the center of the chunk containing a given position.
func getChunkCenter(pos rl.Vector3) rl.Vector3 {
	chunkX := float32(math.Floor(float64(pos.X / chunkSize)))
	chunkY := float32(math.Floor(float64(pos.Y / chunkSize)))
	chunkZ := float32(math.Floor(float64(pos.Z / chunkSize)))

	centerX := chunkX*chunkSize + float32(chunkSize)/2
	centerY := chunkY*chunkSize + float32(chunkSize)/2
	centerZ := chunkZ*chunkSize + float32(chunkSize)/2

	return rl.NewVector3(centerX, centerY, centerZ)
}

// heuristicEuclidean calculates the Euclidean distance between two 3D points.
func heuristicEuclidean(a, b rl.Vector3) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	dz := float64(a.Z - b.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// findPath performs A* pathfinding to find a path from start to target.
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

// updateNeighbor updates the neighbor's cost and parent if a shorter path is found.
func updateNeighbor(neighbor *Node, current, targetNode *Node, openSet, closedSet map[rl.Vector3]*Node) {
	if closedSet[neighbor.Position] != nil {
		return
	}

	tentativeGCost := float64(current.GCost) + float64(rl.Vector3Distance(current.Position, neighbor.Position))

	if openSet[neighbor.Position] == nil || tentativeGCost < float64(openSet[neighbor.Position].GCost) {
		neighbor.GCost = tentativeGCost
		neighbor.HCost = heuristicEuclidean(neighbor.Position, targetNode.Position)
		neighbor.Parent = current

		openSet[neighbor.Position] = neighbor
	}
}
