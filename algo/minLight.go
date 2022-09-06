package algo

import (
	"math"
)

func minLight1(road string) int {
	if len(road) == 0 {
		return 0
	}
	lights := make(map[int]struct{})
	return process(road, 0, &lights)
}

// str[index....]位置，自由选择放灯还是不放灯
// str[0..index-1]位置呢？已经做完决定了，那些放了灯的位置，存在lights里
// 要求选出能照亮所有.的方案，并且在这些有效的方案中，返回最少需要几个灯

func process(road string, index int, lights *map[int]struct{}) int {
	if index == len(road) {
		for i := 0; i < len(road); i++ {
			if road[i] != 'X' {
				_, ok1 := (*lights)[i-1]
				_, ok2 := (*lights)[i]
				_, ok3 := (*lights)[i+1]
				if !ok1 && !ok2 && !ok3 {
					return math.MaxInt
				}
			}
		}
		return len(*lights)
	} else {
		no := process(road, index+1, lights)
		yes := math.MaxInt
		if road[index] != 'X' {
			(*lights)[index] = struct{}{}
			yes = process(road, index+1, lights)
			delete(*lights, index)
		}
		if no < yes {
			return no
		}
		return yes
	}
}

func minLight2(road string) int {
	i := 0
	light := 0
	for i < len(road) {
		if road[i] == 'X' {
			i++
		} else {
			light++
			if i+1 == len(road) {
				break
			} else {
				if road[i+1] == 'X' {
					i = i + 2
				} else {
					i = i + 3
				}
			}
		}
	}
	return light
}

func minLight3(road string) int {
	cur := 0
	light := 0

	for i := 0; i < len(road); i++ {
		if road[i] == 'X' {
			light += (cur + 2) / 3
			cur = 0
		} else {
			cur++
		}
	}
	light += (cur + 2) / 3
	return light
}
