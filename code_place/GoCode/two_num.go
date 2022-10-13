package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int, 0)
	for i, item := range nums {
		tmpTarget := target - item

		if index, ok := tmpMap[tmpTarget]; ok {
			return []int{index, i}
		}

		tmpMap[item] = i
	}

	return nil
}

func main() {

	arr1 := []int{2, 7, 11, 15}
	arr2 := []int{3, 2, 4}
	arr3 := []int{3, 3}
	fmt.Printf("%v\n", twoSum(arr1, 9))
	fmt.Printf("%v\n", twoSum(arr2, 6))
	fmt.Printf("%v\n", twoSum(arr3, 6))
}
