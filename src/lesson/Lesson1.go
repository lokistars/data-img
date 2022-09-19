package lesson

import "fmt"

func SelectionSort(arr []int32) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func Print() {
	fmt.Printf("测试")
}
