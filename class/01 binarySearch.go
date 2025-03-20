package main

import "fmt"

func binarySearch(a []int, x int) int {
	l, r := 0, len(a)-1
	for l <= r { //包括等于
		mid := (l + r) >> 1 //(l + r) / 2 //数据有溢出风险
		if a[mid] < x {
			l = mid + 1
		} else if x < a[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	test := binarySearch([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(test)
}
