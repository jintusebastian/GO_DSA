/* ------------------------LINEAR SEARCH -------------------------------  */

package main

import (
	"fmt"
)

func linearsearch(numlist []int, key int) int {
	for index, item := range numlist {
		if key == item {
			return index
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 89, 23, 4, 567, 23, 56, 78, 12, 90}
	fmt.Println("Value of found at index : ", linearsearch(arr, 90))

}
