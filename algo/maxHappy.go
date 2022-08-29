package main

import (
	"math"
)

type Employee struct {
	Happy int
	Nexts []*Employee
}

func maxHappy1(x *Employee) int {
	if x == nil {
		return 0
	}
	return process1(x, false)
}

func process1(x *Employee, up bool) int {
	if up {
		ans := 0
		for _, next := range x.Nexts {
			ans += process1(next, false)
		}
		return ans
	} else {
		p1 := x.Happy
		p2 := 0
		for _, next := range x.Nexts {
			p1 += process1(next, true)
			p2 += process1(next, false)
		}
		return int(math.Max(float64(p1), float64(p2)))
	}
}

type Info struct {
	Yes int
	No  int
}

func maxHappy2(x *Employee) int {
	allInfo := process2(x)
	return int(math.Max(float64(allInfo.Yes), float64(allInfo.No)))
}

func process2(x *Employee) *Info {
	if x == nil {
		return &Info{
			Yes: 0,
			No:  0,
		}
	}
	yes := x.Happy
	no := 0
	for _, next := range x.Nexts {
		allInfo := process2(next)
		no += int(math.Max(float64(allInfo.Yes), float64(allInfo.No)))
		yes += allInfo.No
	}
	return &Info{
		Yes: yes,
		No:  no,
	}
}
