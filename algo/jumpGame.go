package algo


func jump(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	// position := len(nums) - 1
	// step := 0
	// for position > 0 {
	//     for i:=0;i<position;i++ {
	//         if nums[i]+i >= position {
	//             step++
	//             position = i
	//             break
	//         }
	//     }
	// }
	// return step
	position := 0
	end := len(nums)-1
	step := 0

	for position < end {
		if nums[position] >= end - position {
			step++
			break
		}
		maxPosition := position
		nextPosition := position
		for i := position+1; i <= nums[position]+position;i++ {
			if i+nums[i] > maxPosition {
				maxPosition = i+nums[i]
				nextPosition = i
			}
		}
		position = nextPosition
		step++
	}
	return step
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
