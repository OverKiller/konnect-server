package main

import (
	"encoding/json"
	"log"

	"github.com/shirou/gopsutil/process"
)

func getProcess(message Message, rm *ResponseMessage) []byte {

	plist, _ := process.Processes()
	processes := make([]Process, 0)

	for _, p := range plist {
		pname, err := p.Name()
		if err != nil {
			log.Println(err)
			continue
		}
		pmem, err := p.MemoryInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		pper, err := p.MemoryPercent()
		if err != nil {
			log.Println(err)
			continue
		}
		processes = append(processes, Process{
			Name:           pname,
			PID:            p.Pid,
			RAMUsed:        pmem.RSS,
			RAMUsedPercent: pper,
		})
	}

	processReponse := &ProcessList{
		ResponseMessage: rm,
		ProcessList:     processes,
	}
	res, err := json.Marshal(processReponse)
	if err != nil {
		return nil
	}
	return res

}
