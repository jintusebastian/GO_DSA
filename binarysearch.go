/* ------------------------  BINARY SEARCH -------------------------------  */

package main

import (
	"fmt"
)

func binarysearch(numlist []int, key int) int {
	var (
		startIndex int = 0
		endIndex       = len(numlist) - 1
	)
	for startIndex <= endIndex {
		midIndex := (startIndex + endIndex) / 2
		if key == numlist[midIndex] {
			return midIndex
		}

		if key > numlist[midIndex] {
			startIndex = midIndex + 1
		} else {
			endIndex = midIndex + 1
		}
	}

	return -1
}
func main() {
	item := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fmt.Println("Element found at position : ", binarysearch(item, 14))
}
