package twopointer

import "fmt"

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target{
			return []int{left+1, right+1}
		}else if sum < target{
			left ++
		}else {
			right --
		}
	}

	return []int{0, 0}
}

func TestTwoSum() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
	nums = []int{-1, 0}
	target = -1
	fmt.Println(twoSum(nums, target))
}