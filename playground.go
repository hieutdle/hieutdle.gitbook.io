package main

import "fmt"

func twoSum(nums []int, target int) []int {
	twonums := make([]int, 0, 2)
out:
	for i1, j1 := range nums {
		if i1 > target {
			continue
		} else {
			for i2, j2 := range nums {
				if j1+j2 == target {
					twonums = append(twonums, i1, i2)
					break out
				}
				fmt.Println(i1, j1, i2, j2)
			}
			fmt.Println(i1, j1)
		}
	}
	return twonums
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	twonums := make([]int, 0, 2)
out:
	for i1, j1 := range nums {
		if j1 > target {
			fmt.Println(i1, j1)
			continue

		} else {
			for i2, j2 := range nums {
				if j1+j2 == target {
					twonums = append(twonums, i1, i2)
					break out
				}
				fmt.Println(i1, j1, i2, j2)
			}
			fmt.Println(i1, j1)
		}
	}

}
