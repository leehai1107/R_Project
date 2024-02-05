package main

import (
	"container/heap"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Constants
const (
	screenWidth  = 800
	screenHeight = 450
	moveSpeed    = 0.05
	gridSize     = 50
)

// Variables
var (
	playerPos = rl.NewVector3(2.5, 0, 2.5)
	targetPos = rl.NewVector3(0, 0, 0)
	path      []rl.Vector3

	g0 = rl.NewVector3(-15.0, 0.0, -15.0)
	g1 = rl.NewVector3(-15.0, 0.0, 15.0)
	g2 = rl.NewVector3(15.0, 0.0, 15.0)
	g3 = rl.NewVector3(15.0, 0.0, -15.0)
)

// Node struct represents a node in the pathfinding grid.
type Node struct {
	Position     rl.Vector3
	GCost, HCost float64
	Parent       *Node
	Index        int // Index in the priority queue
}

// PriorityQueue is a min heap for nodes based on total cost.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].GCost+pq[i].HCost < pq[j].GCost+pq[j].HCost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].Index, pq[j].Index = i, j }
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// main function
func main() {
	rl.InitWindow(screenWidth, screenHeight, "3D Pathfinding")

	camera := initCamera()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraMode(rl.CameraOrthographic))
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
	rayHit := rl.GetRayCollisionQuad(ray, g0, g1, g2, g3)

	if rayHit.Hit && !rl.Vector3Equals(rayHit.Point, playerPos) {
		targetPos = rayHit.Point
		path = findPath(playerPos, targetPos)
	}
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

	// Smoothly interpolate between path points for smoother movement
	if len(path) > 1 {
		directionToNextPoint := rl.Vector3Subtract(path[0], playerPos)
		directionToNextPoint = rl.Vector3Normalize(directionToNextPoint)
		playerPos = rl.Vector3Add(playerPos, rl.Vector3Scale(directionToNextPoint, moveSpeed))

		// Check if reached the next point
		distanceToNextPoint := rl.Vector3Distance(playerPos, path[0])
		if distanceToNextPoint < moveSpeed {
			path = path[1:] // Remove the reached point from the path
		}
	}
}

// draw renders the game entities and path.
func draw(camera rl.Camera) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)
	{
		drawEntities()

		rl.DrawGrid(gridSize, 10)

		if len(path) > 1 {
			drawPath()
		}
	}
	rl.EndMode3D()
	rl.DrawFPS(30, 30)

	rl.DrawText("Click to set pathfinding target", 10, 10, 20, rl.DarkGray)

	rl.EndDrawing()
}

// drawEntities renders the player and target entities.
func drawEntities() {
	rl.DrawSphere(playerPos, 0.2, rl.Red)
	rl.DrawSphere(targetPos, 0.2, rl.Green)
}

// drawPath renders the path as a series of connected lines.
func drawPath() {
	for i := 0; i < len(path)-1; i++ {
		rl.DrawLine3D(path[i], path[i+1], rl.DarkGray)
	}
}

// getCurrentNode finds the node with the lowest cost in the open set using a priority queue.
func getCurrentNode(openSet *PriorityQueue) *Node {
	if openSet.Len() == 0 {
		return nil
	}
	return heap.Pop(openSet).(*Node)
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
		float32(math.Ceil(float64(pos.X))),
		float32(math.Ceil(float64(pos.Y))),
		float32(math.Ceil(float64(pos.Z))),
	)
	return &Node{Position: gridPos}
}

// getNeighbors returns the neighboring nodes of a given node.
func getNeighbors(node *Node) []*Node {
	neighbors := make([]*Node, 0)

	for x := -0.5; x <= 0.5; x += 0.5 {
		for y := -0.5; y <= 0.5; y += 0.5 {
			for z := -0.5; z <= 0.5; z += 0.5 {
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

// heuristicEuclidean calculates the Euclidean distance between two 3D points.
func heuristicEuclidean(a, b rl.Vector3) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	dz := float64(a.Z - b.Z)
	return (dx*dx + dy*dy + dz*dz)
}

// findPath performs A* pathfinding to find a path from start to target using a priority queue.
func findPath(start, target rl.Vector3) []rl.Vector3 {
	startNode := getNodeFromWorldPos(start)
	targetNode := getNodeFromWorldPos(target)

	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)
	closedSet := make(map[rl.Vector3]*Node)

	heap.Push(&openSet, startNode)

	for openSet.Len() > 0 {
		current := getCurrentNode(&openSet)

		closedSet[current.Position] = current

		if current.Position == targetNode.Position {
			return reconstructPath(current)
		}

		neighbors := getNeighbors(current)
		for _, neighbor := range neighbors {
			updateNeighbor(neighbor, current, targetNode, &openSet, closedSet)
		}
	}

	return []rl.Vector3{}
}

// updateNeighbor updates the neighbor's cost and parent if a shorter path is found using a priority queue.
func updateNeighbor(neighbor *Node, current, targetNode *Node, openSet *PriorityQueue, closedSet map[rl.Vector3]*Node) {
	if closedSet[neighbor.Position] != nil {
		return
	}

	tentativeGCost := float64(current.GCost) + float64(rl.Vector3Distance(current.Position, neighbor.Position))

	if openSetContains(openSet, neighbor) && tentativeGCost >= float64(neighbor.GCost) {
		return
	}

	neighbor.GCost = tentativeGCost
	neighbor.HCost = heuristicEuclidean(neighbor.Position, targetNode.Position)
	neighbor.Parent = current

	if !openSetContains(openSet, neighbor) {
		heap.Push(openSet, neighbor)
	}
}

// openSetContains checks if the priority queue (open set) contains a node.
func openSetContains(openSet *PriorityQueue, node *Node) bool {
	for _, n := range *openSet {
		if n.Position == node.Position {
			return true
		}
	}
	return false
}
