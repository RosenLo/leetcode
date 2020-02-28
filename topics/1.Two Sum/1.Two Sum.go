package main

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

// original
func twoSum(nums []int, target int) []int {
	for j := 0; j < len(nums); j++ {
		for k := j + 1; k < len(nums); k++ {
			if nums[j]+nums[k] == target {
				return []int{j, k}
			}
		}
	}
	return nil
}

// Approach 2: Two-pass Hash Table
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		m[val] = idx
	}
	for idx1 := 0; idx1 < len(nums); idx1++ {
		idx2, exists := m[target-nums[idx1]]
		if exists && idx1 != idx2 {
			return []int{idx1, idx2}
		}
	}
	return nil
}

// Approach 3: One-pass Hash Table
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	for idx1, val := range nums {
		idx2, exists := m[target-val]
		if exists && idx1 != idx2 {
			return []int{idx2, idx1}
		}
		m[val] = idx1
	}
	return nil
}
