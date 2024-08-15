package Aggregator

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
)

type IData interface {
	Run()
	ProcessK8S()
	ProccessEbpfProc(Ctx context.Context)
	ProccessEbpfTCP(Ctx context.Context)
	ProccessEbpf(Ctx context.Context)
	GetRateLimiterForPid(Pid uint32) *Rate.Limiter
	ProccessExec(D *Proc.ProcEvent)
	ProccessExit(Pid uint32)
	SignalTlsAttachment(Pid uint32)
	ProcessTcpConnect(Ctx context.Context, D *TcpState.TcpConnectEvent)
	ProcessHttp2Frames()
	GetPodWithIP(Address string) (types.UID, bool)
	GetSvcWithIP(Address string) (types.UID, bool)
	SetFromToV2(AddressPair *AddressPair, D *L7Req.L7Event, Event datastore.DirectionalEvent, HostHeader string) error
	SetFromTo(SkInfo *SockInfo, D *L7Req.L7Event, Event datastore.DirectionalEvent, HostHeader string) error
	GetConnKey(Pid uint32, Fd uint64) string
	DecodeKafkaPayload(D *L7Req.L7Event) ([]*KafkaMessage, error)
	ProcessHttp2Event(D *L7Req.L7Event)
	ProcessKafkaEvent(Ctx context.Context, D *L7Req.L7Event)
	ProcessAmqpEvent(Ctx context.Context, D *L7Req.L7Event)
	ProcessRedisEvent(Ctx context.Context, D *L7Req.L7Event)
	AdvertiseDebugData()
	ProcessHttpEvent(Ctx context.Context, D *L7Req.L7Event)
	ProcessMongoEvent(Ctx context.Context, D *L7Req.L7Event)
	ProcessMySQLEvent(Ctx context.Context, D *L7Req.L7Event)
	ProcessPostgresEvent(Ctx context.Context, D *L7Req.L7Event)
	ProcessL7(Ctx context.Context, D *L7Req.L7Event)
	FindRelatedSocket(Ctx context.Context, D *L7Req.L7Event) (*SockInfo, error)
	ParseMySQLCommand(D *L7Req.L7Event) (string, error)
	ParsePostgresCommand(D *L7Req.L7Event) (string, error)
	ParseMongoEvent(D *L7Req.L7Event) (string, error)
	GetPgStmtKey(Pid uint32, Fd uint64, StmtName string) string
	SendOpenConnection(Sl *SocketLine)
	ClearSocketLines(Ctx context.Context)
}

type ICluster interface {
	SignalSocketMapCreation(pid uint32)
	HandleSocketMapCreation()
	ClearProc(pid uint32)
}
type ClusterInfo struct {
	K8Smu                 sync.RWMutex
	PodIPToPodUid         map[string]types.UID `json:"podIPToPodUid"`
	ServiceIPToServiceUid map[string]types.UID `json:"serviceIPToServiceUid"`
	SocketMaps            []*SocketMap
	SocketMapsmu          sync.Mutex
	MuIndex               atomic.Uint64
	MuArray               []*sync.RWMutex
	SignalChan            chan uint32
}

type Aggregator struct {
	Ctx                 context.Context
	Ct                  *Cri.CRITool
	K8sChan             chan interface{}
	EbpfChan            chan interface{}
	EbpfProcChan        chan interface{}
	EbpfTcpChan         chan interface{}
	TlsAttachSignalChan chan uint32
	ClusterInfo         *ClusterInfo
	Ds                  DataStore.DataStore
	H2Mu                sync.RWMutex
	H2Ch                chan *L7Req.L7Event
	H2Frames            map[string]*FrameArrival
	H2ParserMu          sync.RWMutex
	H2Parsers           map[string]*Http2Parser
	PgStmtsMu           sync.RWMutex
	PgStmts             map[string]string
	MySqlStmtsMu        sync.RWMutex
	MySqlStmts          map[string]string
	LiveProcessesMu     sync.RWMutex
	LiveProcesses       map[uint32]struct{}
	RateLimiters        map[uint32]*Rate.Limiter
	RateLimitMu         sync.RWMutex
}

type Http2Parser struct {
	ClientHpackDecoder *Hpack.Decoder
	ServerHpackDecoder *Hpack.Decoder
}

type FrameArrival struct {
	ClientHeadersFrameArrived bool
	ServerHeadersFrameArrived bool
	ServerDataFrameArrived    bool
	Event                     *L7Req.L7Event
	Req                       *DataStore.Request
	StatusCode                uint32
	GrpcStatus                uint32
}

type KafkaMessage struct {
	TopicName string
	Partition int32
	Key       string
	Value     string
	Type      string
}

type AddressPair struct {
	Saddr string `json:"saddr"`
	Sport uint16 `json:"sport"`
	Daddr string `json:"daddr"`
	Dport uint16 `json:"dport"`
}

type SockInfo struct {
	Pid   uint32 `json:"pid"`
	Fd    uint64 `json:"fd"`
	Saddr string `json:"saddr"`
	Sport uint16 `json:"sport"`
	Daddr string `json:"daddr"`
	Dport uint16 `json:"dport"`
}

type SocketMap struct {
	Mu             *sync.RWMutex
	Pid            uint32
	M              map[uint64]*SocketLine `json:"fdToSockLine"`
	WaitingFds     chan uint64
	ProcessedFds   map[uint64]struct{}
	ProcessedFdsmu sync.RWMutex
	CloseCh        chan struct{}
	Ctx            context.Context
}

type TimestampedSocket struct {
	Timestamp uint64
	LastMatch uint64
	SockInfo  *SockInfo
}

type SocketLine struct {
	Mu     sync.RWMutex
	Pid    uint32
	Fd     uint64
	Values []*TimestampedSocket
	Ctx    context.Context
}

type sock struct {
	Pid uint32
	Fd  uint64
	TcpSocket
}

type TcpSocket struct {
	Inode  string
	SAddr  Netaddr.IPPort
	DAddr  NetAddress.IPPort
	Listen bool
}

type Fd struct {
	Fd          uint64
	Dest        string
	SocketInode string
}

const (
	POD              = "pod"
	SVC              = "service"
	OUTBOUND         = "outbound"
	KAFKA            = "kafka"
	REDIS            = "redis"
	ADD              = "ADD"
	UPDATE           = "UPDATE"
	DELETE           = "DELETE"
	StateEstablished = "01"
	StateListen      = "0A"
)

var (
	RetryInterval     = 20 * time.Millisecond
	AttemptLimit      = 3
	DefaultExpiration = 5 * time.Minute
	PurgeTime         = 10 * time.Minute
	ReverseDnsCache   *cache.Cache
	Re                *regexp.Regexp
	MaxPid            int
	MongoOpCompressed uint32 = 2012
	MongoOpMsg        uint32 = 2013
)
