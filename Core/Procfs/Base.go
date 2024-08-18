package Procfs

type ProcEventConversion uint32

type PEvent struct {
	Pid   uint32
	Type_ uint8
	_     [3]byte
}

type ProcEvent struct {
	Pid   uint32
	Type_ string
}

type ProcProgConfig struct {
	ProcEventsMapSize uint32
}

type ProcProg struct {
	Links             map[string]link.Link
	ProcEvents        *perf.Reader
	ProcEventsMapSize uint32
	ContainerPidMap   *ebpf.Map
}

var DefaultConfig *ProcProgConfig = &ProcProgConfig{
	ProcEventsMapSize: 16,
}

const (
	EVENT_PROC_EXEC     = "EVENT_PROC_EXEC"
	EVENT_PROC_EXIT     = "EVENT_PROC_EXIT"
	BPF_EVENT_PROC_EXEC = iota + 1
	BPF_EVENT_PROC_EXIT
	PROC_EVENT = "proc_event"
)
