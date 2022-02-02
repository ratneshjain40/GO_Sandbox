package main

import (
	"errors"
)

type QueueNode struct {
	Data     interface{}
	Priority float32
}

type PriorityQueue struct {
	Length int
	Queue  []*QueueNode
}

//<----------------------- QUEUE METHODS ----------------------->

func (q *PriorityQueue) Init(size int) {
	q.Queue = make([]*QueueNode, size)
	q.Length = 0
}

func (q *PriorityQueue) Push(data interface{}, priority float32) (bool, error) {
	if len(q.Queue) == 0 {
		return false, errors.New("queue not initialized")
	}

	var node *QueueNode = &QueueNode{Data: data, Priority: priority}

	// Binary Search in sorted array for finding index to insert at
	index, err := q.searchInsertPosition(0, q.Length-1, node)

	if !(err == nil) {
		return false, err
	}

	q.Queue = append(q.Queue[:index+1], q.Queue[index:]...)
	q.Queue[index] = node
	q.Length = q.Length + 1

	return true, nil
}

func (q *PriorityQueue) searchInsertPosition(lowerlimit int, upperlimit int, node *QueueNode) (int, error) {
	if len(q.Queue) == 0 {
		return -1, errors.New("queue not initialized")
	}

	for lowerlimit <= upperlimit {
		var mid int = (lowerlimit + upperlimit) / 2

		if q.Queue[mid].Priority < (*node).Priority {
			lowerlimit = mid + 1
		} else {
			upperlimit = mid - 1
		}
	}

	return lowerlimit, nil
}

func (q *PriorityQueue) Pop() *interface{} {
	node := q.Queue[0]
	q.Queue = q.Queue[1:]
	q.Length = q.Length - 1
	return &node.Data
}
