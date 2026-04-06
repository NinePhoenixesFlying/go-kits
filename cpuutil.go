package cpuutil
// package main

import (
	"time"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
)

// func main(){
// 	s, _ := GetCPUCount()
// 	fmt.Println(s)

// 	u, _ :=  GetCPUUsage()
// 	fmt.Println(u)
// }
// GetCPUUsage 获取当前CPU总利用率（百分比，0~100）
// 内部自动采样 500ms，返回总核平均使用率
func GetCPUUsage() (float64, error) {
	// 采样 500ms 计算CPU利用率
	percents, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		return 0, err
	}
	if len(percents) == 0 {
		return 0, nil
	}
	return percents[0], nil
}

// GetPerCPUUsage 获取每个CPU核心的利用率
func GetPerCPUUsage() ([]float64, error) {
	return cpu.Percent(500*time.Millisecond, true)
}

// GetCPUInfo 获取CPU核心数、型号等信息
func GetCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// GetCPUCount 获取CPU逻辑核心数
func GetCPUCount() (int, error) {
	return cpu.Counts(true)
}