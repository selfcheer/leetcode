package array

//移除元素
//给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-element
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func removeElement(nums []int, val int) int {
	var right int
	size := len(nums)
	newSize := 0
	for i, v := range nums {
		if v == val {
			right = i
			for right < size - 1 && nums[right] == val {
				right++
			}
			if nums[right] == val {
				break
			}
			nums[i], nums[right] = nums[right], nums[i]
			newSize++
		} else {
			newSize++
		}
	}
	return newSize
}

