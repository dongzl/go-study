package _078

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/subsets/
var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	helper(nums, 0, []int{})
	return result
}

func helper(nums []int, k int, path []int) {
	if len(nums) == k {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	helper(nums, k+1, path)
	path = append(path, nums[k])
	helper(nums, k+1, path)
	path = path[:len(path)-1]
}

func TestName(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}
