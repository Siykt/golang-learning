package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		if i2, ok := hash[target-num]; ok {
			return []int{i2, i}
		} else {
			hash[num] = i
		}
	}

	return []int{-1, -1}
}

func main() {
	var nums = []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

	nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))

	nums = []int{3, 3}
	fmt.Println(twoSum(nums, 6))
}
