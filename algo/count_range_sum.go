package algo

// https://leetcode.cn/problems/count-of-range-sum/
func countRangeSum(nums []int, lower int, upper int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i-1]
	}
	return process(sums, 0, len(nums)-1, lower, upper)
}

func process(sums []int, L int, R int, lower int, upper int) int {
	if L == R {
		if sums[L] >= lower && sums[L] <= upper {
			return 1
		}
		return 0
	}
	M := L + (R-L)>>1
	return process(sums, L, M, lower, upper) + process(sums, M+1, R, lower, upper) + merge(sums, L, M, R, lower, upper)
}

func merge(sums []int, L int, M int, R int, lower int, upper int) int {
	windowL, windowR := L, L
	ans := 0
	for i := M + 1; i <= R; i++ {
		for windowR <= M && sums[windowR] <= sums[i]-lower {
			windowR++
		}
		for windowL <= M && sums[windowL] < sums[i]-upper {
			windowL++
		}
		ans += windowR - windowL
	}
	p1, p2 := L, M+1
	help := make([]int, R-L+1)
	index := 0
	for p1 <= M && p2 <= R {
		if sums[p1] <= sums[p2] {
			help[index] = sums[p1]
			p1++
		} else {
			help[index] = sums[p2]
			p2++
		}
		index++
	}
	for p1 <= M {
		help[index] = sums[p1]
		index++
		p1++
	}
	for p2 <= R {
		help[index] = sums[p2]
		index++
		p2++
	}
	for i := 0; i < len(help); i++ {
		sums[L+i] = help[i]
	}
	return ans
}
