package Net

import (
	"context"
	"regexp"
	"sync"
)

type IProgram interface {
	MarzyAttach()
	MarzyInitMaps()
	MarzyConsume(ctx context.Context, ch chan interface{})
	MarzyClose()
}

type IBpfEvent interface {
	MarzyType() string
}

type ICollector interface {
	MarzyDone() chan struct{}
	MarzyEbpfEvents() chan interface{}
	MarzyEbpfProcEvents() chan interface{}
	MarzyEbpfTcpEvents() chan interface{}
	MarzyTlsAttachQueue() chan uint32
	MarzyLoad()
	MarzyInit()
	MarzyListenEvents()
	MarzyClose()
	MarzyAttachUpRobesForEncrypted()
	MarzyAttachGoTlsUpRobesOnProcess(Procfs string, Pid uint32) []error
	MarzyAttachSslUpRobesOnProcess(Procfs string, Pid uint32) []error
	MarzyAttachSSlUpRobes(Pid uint32, ExecutablePath string, Version string) error
}

type TcpStateConversion uint32

type BpfTcpEvent struct {
	Fd        uint64
	Timestamp uint64
	Type      uint32
	Pid       uint32
	SPort     uint16
	DPort     uint16
	SAddr     [16]byte
	DAddr     [16]byte
}

type TcpConnectEvent struct {
	Fd        uint64
	Timestamp uint64
	Type_     string
	Pid       uint32
	SPort     uint16
	DPort     uint16
	SAddr     string
	DAddr     string
}

type TcpStateConfig struct {
	BpfMapSize uint32
}

type TcpStateProg struct {
	Links             map[string]link.Link
	TcpConnectMapSize uint32
	TcpConnectEvents  *perf.Reader
	ContainerPidMap   *ebpf.Map
}

type EbpfCollector struct {
	Ctx                 context.Context
	Done                chan struct{}
	EbpfEvents          chan interface{}
	EbpfProcEvents      chan interface{}
	EbpfTcpEvents       chan interface{}
	TlsAttachQueue      chan uint32
	BpfPrograms         map[string]IProgram
	SslWriteUprobes     map[uint32]link.Link
	SslReadEnterUprobes map[uint32]link.Link
	SslReadURetprobes   map[uint32]link.Link
	GoTlsWriteUprobes   map[uint32]link.Link
	GoTlsReadUprobes    map[uint32]link.Link
	GoTlsReadUretprobes map[uint32][]link.Link
	ProbesMu            sync.Mutex
	TlsPidMap           map[uint32]struct{}
	Mu                  sync.Mutex
	Ct                  *cri.CRITool
}

type SslLib struct {
	Path    string
	Version string
}

const (
	BPF_EVENT_TCP_ESTABLISHED = iota + 1
	BPF_EVENT_TCP_CONNECT_FAILED
	BPF_EVENT_TCP_LISTEN
	BPF_EVENT_TCP_LISTEN_CLOSED
	BPF_EVENT_TCP_CLOSED
)

const (
	EVENT_TCP_ESTABLISHED    = "EVENT_TCP_ESTABLISHED"
	EVENT_TCP_CONNECT_FAILED = "EVENT_TCP_CONNECT_FAILED"
	EVENT_TCP_LISTEN         = "EVENT_TCP_LISTEN"
	EVENT_TCP_LISTEN_CLOSED  = "EVENT_TCP_LISTEN_CLOSED"
	EVENT_TCP_CLOSED         = "EVENT_TCP_CLOSED"
)

const (
	GoTlsWriteSymbol = "crypto/tls.(*Conn).Write"
	GoTlsReadSymbol  = "crypto/tls.(*Conn).Read"
	ExeMaxSizeInMB   = 200
)

const mapKey uint32 = 0

var (
	TcpState      *TcpStateConfig
	DefaultConfig *TcpStateConfig = &TcpStateConfig{
		BpfMapSize: 64,
	}
	LibSSLRegex string = `.*libssl(?P<AdjacentVersion>\d)*-*.*\.so\.*(?P<SuffixVersion>[0-9\.]+)*.*`
	Re          *regexp.Regexp
)
