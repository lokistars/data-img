package main

import (
	"algorithm/src/leetCode"
	"algorithm/src/lesson"
	"fmt"
)

func main() {

	//lesson.Print(20)
	var arr = []int32{1, 3, 2, 9, 1, 3}
	lesson.InsertionSort(arr)
	fmt.Println(arr)

	leetCode.BubblingSort(arr)
}
