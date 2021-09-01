package email

import (
	"github.com/shirou/gopsutil/v3/load"
)

func CpuLoadStat() (load1, load5, load15 float64, err error) {
	var avgLoadStat *load.AvgStat
	avgLoadStat, err = load.Avg()
	if err != nil {
		return -1, -1, -1, err
	}
	return avgLoadStat.Load1, avgLoadStat.Load5, avgLoadStat.Load15, nil
}
