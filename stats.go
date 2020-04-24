package main

import (
	"encoding/json"
	"fmt"
	"os/user"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

//Returns the basic information of the device
func getStats(message Message, rm *ResponseMessage) []byte {
	currentUser, _ := user.Current()
	cpuInfo, _ := cpu.Info()
	hostInfo, _ := host.Info()
	memInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Partitions(true)
	disks := make([]Disk, 0)

	for _, partition := range diskInfo {
		d, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		disks = append(disks, Disk{
			Letter:      partition.Device,
			Label:       partition.Mountpoint,
			TotalSize:   d.Total,
			Free:        d.Free,
			Used:        d.Used,
			UsedPercent: d.UsedPercent,
		})
	}

	stats := &Stats{
		ResponseMessage: rm,
		CurrentUser:     fmt.Sprintf("%s (%s)", currentUser.Name, currentUser.Username),
		MachineName:     hostInfo.Hostname,
		OS:              hostInfo.Platform,
		ProcessorName:   cpuInfo[0].ModelName,
		ProcessorCores:  cpuInfo[0].Cores,
		ProcessorClock:  cpuInfo[0].Mhz,
		NofProcess:      hostInfo.Procs,
		BootTime:        hostInfo.BootTime,
		TotalRAM:        memInfo.Total,
		RAMUsed:         memInfo.Used,
		RAMUsedPercent:  memInfo.UsedPercent,
		Disks:           disks,
	} //TODO: IP, ProcessorThreads

	res, err := json.Marshal(stats)
	if err != nil {
		return nil
	}
	return res
}
