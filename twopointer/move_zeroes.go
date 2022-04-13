package twopointer

import "fmt"

func moveZeroes(nums []int)  {
	n := len(nums)
	slow, fast := 0, 0

	for fast < n {
		if nums[fast] != 0  {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow ++
		}
		fast ++
	}
}

func TestMoveZeroes() {
	nums := []int{0,1,0,3,12}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}