package main

func main() {
	//twoSums usage
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
