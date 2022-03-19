package median_of_two_sorted_arrays

func Mid(nums []int) (mid float64) {
	if len(nums)%2 == 0 {
		mid = float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2.0
	} else {
		mid = float64(nums[len(nums)/2])
	}
	return
}

//找到2个数组中，第k个数，0<=k<m+n

func findMedianSortedArrays(nums1 []int, nums2 []int) (mid float64) {

	return
}
