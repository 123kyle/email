package code17

import (
	"math"
)

// only for no-negative value

func radixSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	digit := maxBits(arr)
	radix := 10
	help := make([]int, len(arr))
	L, R := 0, len(arr)-1

	for d := 1; d <= digit; d++ {
		counter := make([]int, radix) // [0..9]
		for i := L; i <= R; i++ {
			j := getDigit(arr[i], d)
			counter[j]++
		}

		for i := 1; i < radix; i++ {
			counter[i] += counter[i-1]
		}

		for i := R; i >= L; i-- {
			j := getDigit(arr[i], d)
			help[counter[j]-1] = arr[i]
			counter[j]--
		}
		for i, j := L, 0; i <= R; i++ {
			arr[i] = help[j]
			j++
		}
	}

}

func getDigit(x int, d int) int {
	// (x / ((int) Math.pow(10, d - 1))) % 10
	return (x / (int)(math.Pow(10.0, float64(d-1)))) % 10
}

func maxBits(arr []int) int {
	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	res := 0
	for max != 0 {
		res++
		max /= 10
	}
	return res
}
