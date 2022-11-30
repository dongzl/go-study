package _077

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/combinations/
var result [][]int

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	helper(n, k, 1, []int{})
	return result
}

func helper(n, k, step int, path []int) {
	if len(path) == k {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	if step == n+1 {
		return
	}
	helper(n, k, step+1, path)
	path = append(path, step)
	helper(n, k, step+1, path)
	path = path[:len(path)-1]
}

func TestName(t *testing.T) {
	fmt.Println(combine(4, 2))
}
