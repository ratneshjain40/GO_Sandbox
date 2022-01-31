package main

import (
	"fmt"
)

func main() {
	//newQueue := PriorityQueue{}
	//newQueue.Init(10)
	//fmt.Printf("Queue Initialized %v\n", newQueue)
	//newQueue.Push(10, 12)
	//newQueue.Push(11, 11)
	//newQueue.Push(12, 12)
	//fmt.Printf("Popped   : %v\n", *(newQueue.Pop()))
	//fmt.Print("Elements : ")
	//for i := 0; i < newQueue.Length; i++ {
	//	fmt.Printf("%v ", newQueue.Queue[i])
	//}

	fmt.Print("\nGrid : \n")
	algo := Astar{}
	algo.Init(100, 100)
	algo.SetStart(1, 1)
	algo.SetEnd(99, 78)
	path, err := algo.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Path :\n")
	for i := 0; i < len(path); i++ {
		fmt.Printf("%v \n", *path[i])
	}

}
