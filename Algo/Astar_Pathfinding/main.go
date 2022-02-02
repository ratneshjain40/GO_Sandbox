package main

import (
	"fmt"
)

func main() {
	start := [2]int{0, 0}
	end := [2]int{11, 18}

	fmt.Print("\nGrid : \n")

	algo := Astar{}
	algo.Init(23, 22)
	algo.SetStart(start)
	algo.SetEnd(end)
	wallList := [][2]int{{10, 18}}
	algo.SetWalls(&wallList)
	path, err := algo.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Path :\n")
	for i := 0; i < len(path); i++ {
		fmt.Printf("%v \n", *path[i])
	}

	fmt.Print("\nResetting.....")

	algo.Reset()

	algo.SetStart(start)
	algo.SetEnd(end)
	pathnew, err := algo.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Path :\n")
	for i := 0; i < len(pathnew); i++ {
		fmt.Printf("%v \n", *pathnew[i])
	}

}
