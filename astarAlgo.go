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

func (algo *Astar) SetStart(x int, y int) {
	algo.start = algo.Grid[x][y]
}

func (algo *Astar) SetEnd(x int, y int) {
	algo.end = algo.Grid[x][y]
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

func (algo *Astar) addOpenSet(fromNode *astarNode, nodeList []*astarNode) {
	tempGCost := fromNode.gcost + 1
	for i := range nodeList {
		if nodeList[i] != nil {
			if nodeList[i].visited && tempGCost < nodeList[i].gcost {
				nodeList[i].prev = fromNode
				nodeList[i].gcost = tempGCost
			} else {
				nodeList[i].setHCost(algo.end)
				nodeList[i].gcost = tempGCost
				nodeList[i].prev = fromNode

				algo.Openset.Push(nodeList[i], nodeList[i].fcost())
				fmt.Printf("Pushed %v: \n", *nodeList[i])

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
		fmt.Printf("\nPopped %v: \n", *next)

		path = append(path, next)

		if next == algo.end {
			break
		}
		nodeList := algo.findAdjacentNeigbours(next)
		algo.addOpenSet(next, nodeList)
	}
	return path, nil
}
