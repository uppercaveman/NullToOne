package main

import "fmt"

/*
一个数组中，两个数的和等于10的，数组下标，找到一组即可
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	t := make(map[int]int)
	for k, v := range nums {
		data := target - v
		i, ok := t[data]
		if ok {
			return []int{i, k}
		}
		t[v] = k
	}
	return nil
}
