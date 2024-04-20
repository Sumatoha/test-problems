package main

import "fmt"

func main() {
	q3answer := removeDuplicates([]int{1, 2, 3, 4, 3, 2, 4, 8, 12, 34})
	fmt.Println(q3answer)

	q4answer := twoSum([]int{1, 100, 3, 6, 8, 9}, 10)
	fmt.Println(q4answer)

	fmt.Println("DONE")
}

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)

	for i, n := range nums {
		if j, ok := temp[target-n]; ok {
			return []int{i, j}
		}
		temp[n] = i
	}
	return []int{}
}

func removeDuplicates(nums []int) []int {
	res := []int{}
	for _, item := range nums {
		if !contains(res, item) {
			res = append(res, item)
		}
	}
	return res
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
