package pathfinder

import (
	"container/heap"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	path       []rl.Vector3
	targetPos          = rl.NewVector3(0, 0, 0)
	currentPos         = rl.NewVector3(0, 0, 0)
	moveSpeed  float32 = 0.1
)

func GetPath() []rl.Vector3 {
	return path
}

func SetPath(newPath []rl.Vector3) {
	path = newPath
}

// Getters and Setters for targetPos

func GetTargetPos() rl.Vector3 {
	return targetPos
}

func SetTargetPos(newTargetPos rl.Vector3) {
	targetPos = newTargetPos
}

// Getters and Setters for currentPos

func GetcurrentPos() rl.Vector3 {
	return currentPos
}

func SetcurrentPos(newcurrentPos rl.Vector3) {
	currentPos = newcurrentPos
}

// Getters and Setters for moveSpeed

func GetMoveSpeed() float32 {
	return moveSpeed
}

func SetMoveSpeed(newMoveSpeed float32) {
	moveSpeed = newMoveSpeed
}

// Node struct represents a node in the pathfinding grid.
type Node struct {
	position     rl.Vector3
	gCost, hCost float64
	parent       *Node
	index        int // Index in the priority queue
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

func FindPath(target rl.Vector3) {
	path = findPath(currentPos, target)
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
		float32(math.Ceil(float64(pos.X))),
		float32(math.Ceil(float64(pos.Y))),
		float32(math.Ceil(float64(pos.Z))),
	)
	return &Node{position: gridPos}
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
