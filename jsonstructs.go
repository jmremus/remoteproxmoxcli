package main

type Data struct {
	Data []Dir `json:"data"`
}

type Dir struct {
	Subdir string `json:"subdir"`
}

type NodeData struct {
	Data []NodeInfo `json:"data"`
}

type NodeInfo struct {
	Ssl_fingerprint string  `json:"ssl_fingerprint"`
	Node            string  `json:"node"`
	Ntype           string  `json:"type"`
	Uptime          int64   `json:"uptime"`
	Maxdisk         int64   `json:"maxdisk"`
	Disk            int64   `json:"disk"`
	Cpu             float32 `json:"cpu"`
	Level           string  `json:"level"`
	Mem             int64   `json:"mem"`
	Status          string  `json:"status"`
	Maxcpu          int16   `json:"maxcpu"`
	Id              string  `json:"id"`
	Maxmem          int64   `json:"maxmem"`
}

type VmData struct {
	Data []VmInfo `json:"data"`
}

type VmInfo struct {
	Uptime    int64   `json:"uptime"`
	Diskread  int64   `json:"diskread"`
	Maxdisk   int64   `json:"maxdisk"`
	Vmid      int16   `json:"vmid"`
	Netin     int16   `json:"netin"`
	Maxmem    int64   `json:"maxmem"`
	Name      string  `json:"Name"`
	Mem       int64   `json:"mem"`
	Status    string  `json:"status"`
	Netout    int64   `json:"netout"`
	Disk      int16   `json:"disk"`
	Diskwrite int64   `json:"diskwrite"`
	Cpus      int16   `json:"cpus"`
	Cpu       float32 `json:"cpu"`
}
