package main

import "fmt"

func main() {
	nums := [11]int{}
	for i := 0; i < 11; i++ {
		nums[i] = i
	}

	for i := range nums {
		if nums[i]%2 == 0 {
			fmt.Println(nums[i], "is even")
		} else {
			fmt.Println(nums[i], "is odd")
		}
	}
}
