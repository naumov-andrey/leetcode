func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]byte, 0)
	cur := x
	for cur != 0 {
		nums = append(nums, byte(cur%10))
		cur /= 10
	}

	cur = x
	for i := len(nums) - 1; cur != 0; i-- {
		if nums[i] != byte(cur%10) {
			return false
		}
		cur /= 10
	}

	return true
}
