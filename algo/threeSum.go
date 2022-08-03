package algo



// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if nums[len(nums)-1] < 0 || nums[0] > 0 {
		return res
	}

	lastI := math.MinInt
	for i := 0; i < len(nums)-2; i++ {
		if lastI == nums[i] {
			continue
		}
		lastJ := math.MinInt
		for j := i + 1; j < len(nums)-1; {
			if nums[j] == lastJ {
				j++
				continue
			}
			tmp := 0 - nums[i] - nums[j]
			left := j + 1
			right := len(nums) - 1
           
            for left <= right {
                 mid := left + (right - left) >> 1
                if nums[mid] == tmp {
                    res = append(res, []int{nums[i], nums[j], tmp})
                    break
                } else if nums[mid] < tmp {
                    left = mid + 1
                } else {
                    right = mid -1 
                }
            }

			lastJ = nums[j]
			j++
		}
		lastI = nums[i]
	}
	return res
}
