package main

import "fmt"

func binarySearchL(a []int, x int) int {
	l, r := 0, len(a)-1
	candidate := -1
	for l <= r {
		mid := (l + r) >> 1
		if a[mid] < x {
			l = mid + 1
		} else if x < a[mid] {
			r = mid - 1
		} else {
			candidate = mid
			r = mid - 1
		}
	}
	return candidate
}

func main() {
	test := binarySearchL([]int{1, 2, 3, 3, 4, 5}, 3)
	fmt.Println(test)
}
