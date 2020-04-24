package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

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

//Execute a command.
//Todo: add param to redirect the command output to socket
func startProcess(message Message, rm *ResponseMessage) []byte {
	response := &StartProcess{
		ResponseMessage: rm,
	}
	command := message.Params["cmd"]
	commandList := strings.Split(strings.Trim(command, " "), " ")
	cmd := exec.Command(commandList[0], commandList[1:]...)
	err := cmd.Start()
	if err != nil {
		return nil
	}
	response.PID = cmd.Process.Pid
	res, err := json.Marshal(response)
	if err != nil {
		return nil
	}
	return res
}

//Kill a process by PID
//TODO
func killProcess(message Message, rm *ResponseMessage) []byte {
	return nil
}
