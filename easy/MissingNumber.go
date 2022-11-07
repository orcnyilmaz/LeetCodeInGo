package easy

func MissingNumber(nums []int) int {
	length := len(nums)
	sum := (length * (length + 1)) / 2

	for i := 0; i < length; i++ {
		sum -= nums[i]
	}

	return sum
}
