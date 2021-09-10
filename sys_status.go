package email

import (
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
)

func CpuLoadStat() (load1, load5, load15 float64, err error) {
	var avgLoadStat *load.AvgStat
	avgLoadStat, err = load.Avg()
	if err != nil {
		return -1, -1, -1, err
	}
	return avgLoadStat.Load1, avgLoadStat.Load5, avgLoadStat.Load15, nil
}

func DiskUsage() (float64, error) {
	dir, err := os.Getwd()
	if err != nil {
		return -1, err
	}
	usage, err := disk.Usage(dir)
	if err != nil {
		return -1, err
	}
	return usage.UsedPercent, nil
}

func MemUsage() (float64, error) {
	vMemStat, err := mem.VirtualMemory()
	if err != nil {
		return -1, err
	}
	return vMemStat.UsedPercent, nil
}
