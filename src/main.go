package main

import (
	"algorithm/src/leetCode"
	"fmt"
)

func main() {

	//lesson.Print(20)
	var arr = []int32{1, 3, 2, 9, 1, 3}
	//lesson.InsertionSort(arr)
	leetCode.BubblingSort(arr)

	fmt.Println(arr)
}
