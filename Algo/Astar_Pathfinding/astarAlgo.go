package main

import (
	"fmt"
	"math"
)

type astarNode struct {
	xcorodinate int
	ycorodinate int
	closed      bool
	visited     bool
	gcost       float32
	hcost       float32
	prev        *astarNode
}

func (node *astarNode) setHCost(target *astarNode) {
	node.hcost = float32((math.Abs(float64(node.xcorodinate) - float64(target.xcorodinate))) + (math.Abs(float64(node.ycorodinate) - float64(target.ycorodinate))))
}

func (node *astarNode) fcost() float32 {
	return node.gcost + node.hcost
}

// Interface to impleamnet any type of queue you want, custom priority queue is used here.
type queue interface {
	Init(size int)
	Push(data interface{}, priority float32) (bool, error)
	Pop() *interface{}
}

type Astar struct {
	Grid    [][]*astarNode
	Openset queue
	start   *astarNode
	end     *astarNode
}

func (algo *Astar) Init(rows int, cols int) {
	algo.start = nil
	algo.end = nil
	algo.Grid = make([][]*astarNode, rows)

	for i := range algo.Grid {
		algo.Grid[i] = make([]*astarNode, cols)
		for j := range algo.Grid[i] {
			algo.Grid[i][j] = &astarNode{
				xcorodinate: i,
				ycorodinate: j,
			}
		}
	}
	algo.Openset = &PriorityQueue{}
	// Setting max size of openset/priorityQ(based on slicce) for better perfomance
	algo.Openset.Init(rows * cols)
}

func (algo *Astar) SetStart(set [2]int) {
	algo.start = algo.Grid[set[0]][set[1]]
}

func (algo *Astar) SetEnd(set [2]int) {
	algo.end = algo.Grid[set[0]][set[1]]
}

func (algo *Astar) Config(walls *[][]int, weights *[][]int, hMutliplier int) {
	if len(*walls) <= len(algo.Grid) && len((*walls)[0]) <= len(algo.Grid[0]) {
		for i := 0; i < len(*walls); i++ {

		}
	}
}

func (algo *Astar) findAdjacentNeigbours(node *astarNode) []*astarNode {
	var nodeList []*astarNode

	rows := len(algo.Grid)
	cols := len(algo.Grid[0])

	x := node.xcorodinate
	y := node.ycorodinate

	if x-1 >= 0 && !algo.Grid[x-1][y].closed {
		nodeList = append(nodeList, algo.Grid[x-1][y])
	}
	if x+1 < rows && !algo.Grid[x+1][y].closed {
		nodeList = append(nodeList, algo.Grid[x+1][y])
	}
	if y-1 >= 0 && !algo.Grid[x][y-1].closed {
		nodeList = append(nodeList, algo.Grid[x][y-1])
	}
	if y+1 < cols && !algo.Grid[x][y+1].closed {
		nodeList = append(nodeList, algo.Grid[x][y+1])
	}

	return nodeList
}

func (algo *Astar) isCoordinateOpen(x int, y int) bool {
	//return algo.Grid[x][y].closed &&
}

func (algo *Astar) addOpenSet(fromNode *astarNode, nodeList []*astarNode) {
	tempGCost := fromNode.gcost + 1
	for i := range nodeList {
		if nodeList[i] != nil {
			if nodeList[i].visited && tempGCost < nodeList[i].gcost {
				nodeList[i].prev = fromNode
				nodeList[i].gcost = tempGCost
				fmt.Printf("Updated %v: \n", *nodeList[i])

			} else {
				nodeList[i].setHCost(algo.end)
				nodeList[i].gcost = tempGCost
				nodeList[i].prev = fromNode

				algo.Openset.Push(nodeList[i], nodeList[i].fcost())
				nodeList[i].visited = true
				fmt.Printf("Pushed %v and visited %v: \n", *nodeList[i], nodeList[i].visited)

			}
		}
	}
}

func (algo *Astar) Run() ([]*astarNode, error) {
	if algo.start == nil || algo.end == nil {
		return nil, fmt.Errorf("did not set start and end, use setstart/setend methods")
	}
	var path []*astarNode
	algo.Openset.Push(algo.start, algo.start.fcost())
	for {
		temp := *(algo.Openset.Pop())
		var next *astarNode = temp.(*astarNode)

		next.closed = true

		fmt.Printf("\nPopped %v and closed %v: \n", *next, next.closed)

		path = append(path, next)

		if next == algo.end {
			break
		}
		nodeList := algo.findAdjacentNeigbours(next)
		algo.addOpenSet(next, nodeList)
	}
	return path, nil
}
