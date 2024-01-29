package pathfinder

import (
	"container/heap"
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
	hit      bool
	position rl.Vector3
	normal   rl.Vector3
	distance float32
}

// Node struct represents a node in the pathfinding grid.
type Node struct {
	position     rl.Vector3
	gCost, hCost float64
	parent       *Node
	index        int // Index in the priority queue
}

func Init(data *entity.Player) {
	targetPos = rl.NewVector3(0, 0, 0)
	playerPos = data.GetModelPosition()
	moveSpeed = 0.1
}

// PriorityQueue is a min heap for nodes based on total cost.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return (pq[i].gCost + pq[i].hCost) < (pq[j].gCost + pq[j].hCost)
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
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

	if rayHit.hit {
		targetPos = rayHit.position
		path = findPath(playerPos, targetPos)
	}
}

// getRayHitInfo calculates information about a ray hit on the ground.
func getRayHitInfo(ray rl.Ray) RayHitInfo {
	groundNormal := rl.NewVector3(0, 0.01, 0)
	groundDistance := 0.0

	denom := rl.Vector3DotProduct(ray.Direction, groundNormal)
	epsilon := math.Nextafter(0, 1)
	if math.Abs(float64(denom)) > epsilon {
		t := -(rl.Vector3DotProduct(ray.Position, groundNormal) + float32(groundDistance)) / denom
		hitPoint := rl.Vector3Add(ray.Position, rl.Vector3Scale(ray.Direction, t))
		return RayHitInfo{hit: true, position: hitPoint, normal: groundNormal, distance: t}
	}

	return RayHitInfo{hit: false}
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

func movePlayerAlongPath(direction rl.Vector3, player *entity.Player) {
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
			path = path[1:]
		}
	}
	player.SetModelPosition(playerPos)
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
		path = append([]rl.Vector3{current.position}, path...)
		current = current.parent
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
	return &Node{position: gridPos}
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
					node.position.X+float32(x),
					node.position.Y+float32(y),
					node.position.Z+float32(z),
				)
				neighbors = append(neighbors, &Node{position: neighborPos})
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

		closedSet[current.position] = current

		if current.position == targetNode.position {
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
	if closedSet[neighbor.position] != nil {
		return
	}

	tentativeGCost := float64(current.gCost) + float64(rl.Vector3Distance(current.position, neighbor.position))

	if openSetContains(openSet, neighbor) && tentativeGCost >= float64(neighbor.gCost) {
		return
	}

	neighbor.gCost = tentativeGCost
	neighbor.hCost = heuristicEuclidean(neighbor.position, targetNode.position)
	neighbor.parent = current

	if !openSetContains(openSet, neighbor) {
		heap.Push(openSet, neighbor)
	}
}

// openSetContains checks if the priority queue (open set) contains a node.
func openSetContains(openSet *PriorityQueue, node *Node) bool {
	for _, n := range *openSet {
		if n.position == node.position {
			return true
		}
	}
	return false
}
