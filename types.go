package main

//Message struct
type Message struct {
	Action string            `json:"action"`
	Token  string            `json:"token"`
	Params map[string]string `json:"params,omitempty"`
}

//ResponseMessage response base
type ResponseMessage struct {
	Action     string `json:"action"`
	Token      string `json:"token"`
	StatusCode int    `json:"statusCode"`
}

//Stats Response
type Stats struct {
	*ResponseMessage
	CurrentUser      string  `json:"currentUser"`
	MachineName      string  `json:"machineName"`
	OS               string  `json:"os"`
	IP               string  `json:"ip"`
	ProcessorName    string  `json:"processorName"`
	ProcessorCores   int32   `json:"processorCores"`
	ProcessorThreads int     `json:"processorThreads,omitempty"`
	ProcessorClock   float64 `json:"processorClock"`
	NofProcess       uint64  `json:"nofProcess"`
	TotalRAM         uint64  `json:"ramTotal"`
	RAMUsed          uint64  `json:"ramUsed"`
	RAMUsedPercent   float64 `json:"ramUsedPercent"`
	Disks            []Disk  `json:"disks,omitempty"`
	BootTime         uint64  `json:"bootTime,omitempty"`
}

//Disk struct
type Disk struct {
	Letter      string  `json:"letter"`
	Label       string  `json:"label,omitempty"`
	TotalSize   uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

//ScreenShot Response
type ScreenShot struct {
	*ResponseMessage
	Content string `json:"b64img"`
}

//ProcessList Response
type ProcessList struct {
	*ResponseMessage
	ProcessList []Process `json:"processList"`
}

//Process Struct
type Process struct {
	Name           string  `json:"name"`
	PID            int32   `json:"pid"`
	RAMUsed        uint64  `json:"ramUsed"`
	RAMUsedPercent float32 `json:"usedPercent"`
}

//StartProcess Response
type StartProcess struct {
	*ResponseMessage
	PID int `json:"pid"`
}
