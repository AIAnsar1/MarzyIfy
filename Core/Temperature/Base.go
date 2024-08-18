package Temperature

import "sync"

type MetricDesc struct {
	desc  *prometheus.Desc
	type_ prometheus.ValueType
}

type GpuCollector struct {
	fieldDesc     map[string]MetricDesc
	gpuDriverDesc *prometheus.Desc

	descs map[string]*Prometheus.Desc

	n  *NvmlDriver
	mu sync.Mutex
}

type NvmlDriver struct {
	initResult  nvml.Return
	deviceCount uint
}

type DeviceInfo struct {
	UUID               string
	PCIBusID           string
	DisplayState       string
	PersistenceMode    string
	Name               *string
	MemoryMiB          *uint64
	PowerW             *uint
	BAR1MiB            *uint64
	PCIBandwidthMBPerS *uint
	CoresClockMHz      *uint
	MemoryClockMHz     *uint
}

type DeviceStatus struct {
	PowerUsageW        *uint
	TemperatureC       *uint
	GPUUtilization     *uint
	MemoryUtilization  *uint
	EncoderUtilization *uint
	DecoderUtilization *uint
	PowerLimit         *uint
	FanCount           *uint
	FanSpeeds          map[int]uint
	BAR1UsedMiB        *uint64
	UsedMemoryMiB      *uint64
	TotalMemoryMiB     *uint64
	FreeMemoryMiB      *uint64
	ECCErrorsL1Cache   *uint64
	ECCErrorsL2Cache   *uint64
	ECCErrorsDevice    *uint64
}

var RootPaths = []string{
	"/proc/1/root/usr/lib",
	"/proc/1/root/lib",
	"/proc/1/root/run/nvidia",
	"proc/1/root/usr/local/nvidia",
	"/proc/1/root/run/nvidia/driver/lib64",
	"proc/1/root/usr/lib64",
	"proc/1/root/lib64",
	"proc/1/root",
}

var (
	Lock           = &sync.Mutex{}
	SingleInstance *NvmlDriver
)
