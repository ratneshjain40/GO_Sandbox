package main

import (
	"fmt"
)

func main() {
	start := [2]int{0, 0}
	end := [2]int{41, 13}
	wallList := [][2]int{{10, 18}, {10, 17}, {10, 14}, {10, 12}}

	algo := Astar{}
	algo.Init(50, 50)
	algo.SetStart(start)
	algo.SetEnd(end)
	algo.SetWalls(&wallList)

	Algo(&algo)

	fmt.Print("\nResetting.....\n")

	algo.Reset()

	algo.SetStart(start)
	algo.SetEnd(end)

	bufferedAlgo(&algo)

}

func Algo(algo *Astar) {
	path, err := algo.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Path :\n")
	for i := 0; i < len(path); i++ {
		fmt.Printf("%v \n", *path[i])
	}
}

func bufferedAlgo(algo *Astar) {

	fmt.Print("Buffered Path :\n")

	i := 50

	for {

		pathnew, reached, err := algo.BufferedRun(i)

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(pathnew); i++ {
			fmt.Printf("%v \n", *pathnew[i])
		}

		if reached {
			fmt.Print("Reached.....")
			break
		}

	}
}
