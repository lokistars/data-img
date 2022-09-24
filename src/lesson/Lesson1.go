package lesson

import "fmt"

// 打印一个数的32位
func Print(num int32) {
	for i := 31; i >= 0; i-- {
		// num & (1 << i)
		fmt.Print((num >> i) & 1)
	}
}

func BubblingSort(arr []int32) {
	// 控制循环几次 每次都要从头开始，所以i,j 必须为0
	for i := 0; i < len(arr)-1; i++ {
		// 控制排好一次需要几次,减去i，是因为i从1开始表示有多少已经排好了
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

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
func SelectionSorts(arr []int32) {
	for i := 0; i < len(arr)-1; i++ {
		var minIndex = i
		// 从下一个开始进行遍历
		for j := i + 1; j < len(arr); j++ {
			// 当遇见比选择数小的位置时,将选择的数设为当前最小的数的位置
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func InsertionSort(arr []int32) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1 //下标
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}
