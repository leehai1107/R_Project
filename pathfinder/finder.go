package pathfinder

import (
	"main/entity"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	path      []rl.Vector3
	targetPos rl.Vector3
	playerPos rl.Vector3
	moveSpeed float32
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

func Init(data *entity.Player) {
	targetPos = rl.NewVector3(0, 0, 0)
	playerPos = data.GetModelPosition()
	moveSpeed = 0.5
}

// update handles input and updates the game state.
func Process(camera rl.Camera, player *entity.Player) {
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		handleMouseInput(camera, player.GetModelPosition())
	}

	if len(path) > 0 {
		moveAlongPath(player)
	}
}

// handleMouseInput updates the target position based on mouse input and calculates a new path.
func handleMouseInput(camera rl.Camera, playerPos rl.Vector3) {
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
func moveAlongPath(player *entity.Player) {
	direction := rl.Vector3Subtract(path[0], playerPos)
	distance := rl.Vector3Length(direction)

	if distance > moveSpeed {
		movePlayerAlongPath(direction, player)
	} else {
		path = path[1:]
	}
}

// movePlayerAlongPath moves the player along the given direction with a specified speed.
func movePlayerAlongPath(direction rl.Vector3, player *entity.Player) {
	direction = rl.Vector3Normalize(direction)
	playerPos = rl.Vector3Add(playerPos, rl.Vector3Scale(direction, moveSpeed))
	player.SetModelPosition(playerPos)
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
