// func rotate(nums []int, k int) {
//     n := len(nums)
//     k = k % n

//     original := make([]int, n)
//     copy(original, nums)

//     for i := range nums {
//         nums[i] = original[(i + n - k) % n]
//     }
// }

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
