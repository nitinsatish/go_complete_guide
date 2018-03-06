package main

import "fmt"

func main() {
	list1 := []int{2, 5, 6, 8, 12, 23, 15, 1, 33, 54}

	for _, ele := range list1 {
		if ele%2 == 0 {
			fmt.Printf("%d is even\n", ele)
		} else {
			fmt.Printf("%d is odd\n", ele)
		}
	}
}
