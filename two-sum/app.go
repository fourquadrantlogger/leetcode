package two_sum

func TwoSum(nums []int, target int) (result []int) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for ia, a := range nums {
		b := target - a
		if ib, have := m[b]; have && ia != ib {
			return []int{ia, ib}
		}
	}
	return
}
