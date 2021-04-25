func sortArray(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}

	return sort(nums, 0, len(nums) - 1)
}

func sort(nums []int, left, right int) []int {
    if right - left < 1 {
        return nums
    }

	pivot, i := right, left

	for index := left; index < pivot; index++ {
		if nums[index] < nums[pivot] {
			swap(nums, i, index)
            i++
		}
	}
    swap(nums, i, pivot)
    
    sort(nums, left, i - 1)
    sort(nums, i + 1, right)

	return nums
}

func swap(arr []int, first, second int) {
	tmp := arr[first]
	arr[first] = arr[second]
	arr[second] = tmp
}
