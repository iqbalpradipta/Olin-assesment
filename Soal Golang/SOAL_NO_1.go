package main

import "fmt"

func bilangan(nums []int, target int) []int {
	var dump []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				dump = append(dump, i, j)
				return dump
			}
		}
	}
	return dump
}

func main1() {
	var nums = []int{2, 7, 11, 15}
	target := 9

	fmt.Println(bilangan(nums, target))
}