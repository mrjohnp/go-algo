package main

import "fmt"

type Node struct {
	x int
	y int
}

func NewNode(x int, y int) Node {
	return Node{x, y}
}

type Queue struct {
	nodes []Node
}

func (q *Queue) Enqueue(node Node) {
	q.nodes = append(q.nodes, node)

	if DEBUG {
		fmt.Println("LOG::INFO::Queue::Enqueued", node)
	}
}

func (q *Queue) Dequeue() {
	if len(q.nodes) < 1 {
		return
	}

	if DEBUG {
		fmt.Println("LOG::INFO::Queue::Dequeued", q.nodes[0])
	}

	q.nodes = q.nodes[1:]
}

type Graph struct {
	nodes         []Node
	adjacencyList map[Node][]Node
}

const DEBUG bool = true

var Board [][]int = [][]int{{0, 1, 2}, {0, 1, 2}}
var StartPosition Node = NewNode(0, 2)
var EndPostion Node = NewNode(2, 0)

func main() {
	BFS(&Board, &StartPosition, &EndPostion)
}

func BFS(board *[][]int, startPosition *Node, endPosition *Node) {
	var graph Graph
	graph.nodes = createNodes(board)
	graph.adjacencyList = createAdjacencyList(graph.nodes)

	solve(&graph)
}

func createNodes(board *[][]int) []Node {
	var nodes []Node

	for i := 0; i < len((*board)[0]); i++ {
		for j := 0; j < len((*board)[1]); j++ {
			row := (*board)[0][i]
			col := (*board)[1][j]
			nodes = append(nodes, NewNode(row, col))
		}
	}

	return nodes
}

func createAdjacencyList(nodes []Node) map[Node][]Node {
	adjacencyList := make(map[Node][]Node)

	for _, n := range nodes {
		neighbours := getNeighbours(&n)
		adjacencyList[n] = neighbours
	}

	return adjacencyList
}

func getNeighbours(node *Node) []Node {
	var neighbours []Node

	var nb1 Node = NewNode(node.x+1, node.y+2)
	if isValidPosition(nb1.x, nb1.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb1)
	}

	var nb2 Node = NewNode(node.x-1, node.y+2)
	if isValidPosition(nb2.x, nb2.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb2)
	}

	var nb3 Node = NewNode(node.x+1, node.y-2)
	if isValidPosition(nb3.x, nb3.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb3)
	}

	var nb4 Node = NewNode(node.x-1, node.y-2)
	if isValidPosition(nb4.x, nb4.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb4)
	}

	var nb5 Node = NewNode(node.x+2, node.y+1)
	if isValidPosition(nb5.x, nb5.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb5)
	}

	var nb6 Node = NewNode(node.x+2, node.y-1)
	if isValidPosition(nb6.x, nb6.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb6)
	}

	var nb7 Node = NewNode(node.x-2, node.y+1)
	if isValidPosition(nb7.x, nb7.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb7)
	}

	var nb8 Node = NewNode(node.x-2, node.y-1)
	if isValidPosition(nb8.x, nb8.y, len(Board[0])-1) {
		neighbours = append(neighbours, nb8)
	}

	return neighbours
}

func isValidPosition(x int, y int, size int) bool {
	if x >= 0 && x <= size && y >= 0 && y <= size {
		return true
	}

	return false
}

func indexOf(node Node, graph *Graph) int {
	for i, n := range graph.nodes {
		if n == node {
			return i
		}
	}
	return -1
}

func isVisited(index int, visited *[]bool) bool {
	return (*visited)[index]
}

func solve(graph *Graph) {
	var queue Queue = Queue{}
	var visited []bool = make([]bool, len(graph.nodes))
	var prev []Node = make([]Node, len(graph.nodes))

	queue.Enqueue(StartPosition)
	visited[indexOf(StartPosition, graph)] = true

	// fmt.Println("graph", graph.adjacencyList)
	LogMap(graph.adjacencyList)

	for len(queue.nodes) > 0 {
		var currentNode Node = queue.nodes[0]

		queue.Dequeue()

		currentNeighbours := graph.adjacencyList[currentNode]
		visited[indexOf(currentNode, graph)] = true

		for _, nb := range currentNeighbours {
			if !isVisited(indexOf(nb, graph), &visited) {
				queue.Enqueue(nb)
				visited[indexOf(nb, graph)] = true
				prev[indexOf(nb, graph)] = nb
			}
		}
	}

	fmt.Println(visited)
	fmt.Println(prev)
}

func LogMap(list map[Node][]Node) {
	for k, v := range list {
		fmt.Print(k, " >> ")
		fmt.Print(v, "\n")
	}
}
