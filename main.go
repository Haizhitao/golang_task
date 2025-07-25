package main

import (
	"fmt"
	"sort"
	"strconv"
)

func singleNumber(nums []int) int {
	_map := map[int]int{}
	for _, v := range nums {
		_map[v]++
	}
	for k, v_ := range _map {
		if v_ == 1 {
			return k
		}
	}
	return 0
}

func isPalindromeBak(x int) bool {
	x_str := strconv.Itoa(x)
	x_str_len := len(x_str)
	var y string
	for i := x_str_len - 1; i >= 0; i-- {
		y += string(x_str[i])
	}
	z, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}
	return z == x
}

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	xStrLen := len(xStr)
	for i := 0; i < xStrLen/2; i++ {
		if xStr[i] != xStr[xStrLen-i-1] {
			return false
		}
	}
	return true
}

func isValid(s string) bool {
	_map := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) != 0 && stack[len(stack)-1] == _map[v] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	r := ""
	if len(strs) == 1 {
		return strs[0]
	}
outter:
	for i := 0; ; i++ {
		if len(strs[0]) < i+1 {
			break
		}
		c := strs[0][0 : i+1]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 {
				break outter
			}
			if c != strs[j][0:i+1] {
				break outter
			} else if j == len(strs)-1 {
				r = c
			}
		}
	}
	return r
}

func plusOneBak(digits []int) []int {
	r := []int{}
	x := ""
	for _, v := range digits {
		x += strconv.Itoa(v)
	}
	y, _ := strconv.Atoi(x)
	y += 1
	z := strconv.Itoa(y)
	for _, v := range z {
		r = append(r, int(v-'0'))
	}
	return r
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func removeDuplicatesBak(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var arr []int
	for _, v := range nums {
		b := true
		for _, v2 := range arr {
			if v == v2 {
				b = false
				break
			}
		}
		if b {
			arr = append(arr, v)
		}
	}
	fmt.Println(arr)
	return len(arr)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	r := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lastR := r[len(r)-1]
		current := intervals[i]
		if current[0] <= lastR[1] {
			if current[1] > lastR[1] {
				lastR[1] = current[1]
			}
		} else {
			r = append(r, current)
		}
	}
	return r
}

func twoSum(nums []int, target int) []int {
	var r []int
	var map_ = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		j, exists := map_[diff]
		if exists {
			return []int{j, i}
		}
		map_[nums[i]] = i
	}
	return r
}

func main() {
	//两数之和
	//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
	nums := []int{0, 4, 3, 0}
	target := 0
	fmt.Println(twoSum(nums, target))

	//合并区间
	//intervals := [][]int{{1, 3}, {2, 6}, {7, 9}, {8, 12}, {3, 7}, {4, 6}}
	//fmt.Println(merge(intervals))

	//删除有序数组中的重复项
	//a := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8, 8, 9}
	//fmt.Println(removeDuplicates(a))

	//加一
	//a := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	//fmt.Println(plusOne(a))

	//查找字符串数组中的最长公共前缀
	//strs := []string{"dog"}
	//fmt.Println(longestCommonPrefix(strs))

	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
	//a := "([{])}"
	//fmt.Println(isValid(a))

	//回文数
	//a := 121
	//b := isPalindrome(a)
	//fmt.Println(b)

	//只出现一次的数字
	//nums := []int{2, 2, 1}
	//r := singleNumber(nums)
	//fmt.Println(r)
}
