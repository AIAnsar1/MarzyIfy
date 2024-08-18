package DataStore

type IDataStore interface {
	PersistPod(pod Pod, eventType string) error
	PersistService(service Service, eventType string) error
	PersistReplicaSet(rs ReplicaSet, eventType string) error
	PersistDeployment(d Deployment, eventType string) error
	PersistEndpoints(e Endpoints, eventType string) error
	PersistContainer(c Container, eventType string) error
	PersistDaemonSet(ds DaemonSet, eventType string) error
	PersistStatefulSet(ss StatefulSet, eventType string) error
	PersistRequest(request *Request) error
	PersistKafkaEvent(request *KafkaEvent) error
	PersistAliveConnection(trace *AliveConnection) error
}

type DirectionalEvent interface {
	SetFromUID(string)
	SetFromIP(string)
	SetFromType(string)
	SetFromPort(uint16)
	SetToUID(string)
	SetToIP(string)
	SetToType(string)
	SetToPort(uint16)
	ReverseDirection()
}

type ReqInfo [16]interface{}
type ConnInfo [9]interface{}
type TraceInfo [4]interface{}
type KafkaEventInfo [16]interface{}

type LeveledLogger struct {
	LZero zerolog.Logger
}

type NodeExportLogger struct {
	LZero zerolog.Logger
}

type Pod struct {
	UID       string
	Name      string
	Namespace string
	Image     string
	IP        string
	OwnerType string
	OwnerID   string
	OwnerName string
}

type Service struct {
	UID        string
	Name       string
	Namespace  string
	Type       string
	ClusterIP  string
	ClusterIPs []string
	Ports      []struct {
		Name     string `json:"name"`
		Src      int32  `json:"src"`
		Dest     int32  `json:"dest"`
		Protocol string `json:"protocol"`
	}
}

type ReplicaSet struct {
	UID       string
	Name      string
	Namespace string
	OwnerType string
	OwnerID   string
	OwnerName string
	Replicas  int32
}

type DaemonSet struct {
	UID       string
	Name      string
	Namespace string
}

type StatefulSet struct {
	UID       string
	Name      string
	Namespace string
}

type Deployment struct {
	UID       string
	Name      string
	Namespace string
	Replicas  int32
}

type Endpoints struct {
	UID       string
	Name      string
	Namespace string
	Addresses []Address
}

type AddressIP struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	IP        string `json:"ip"`
}

type AddressPort struct {
	Port     int32  `json:"port"`
	Protocol string `json:"protocol"`
	Name     string `json:"name"`
}

type Address struct {
	IPs   []AddressIP   `json:"ips"`
	Ports []AddressPort `json:"ports"`
}

type Container struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	PodUID    string `json:"pod"`
	Image     string `json:"image"`
	Ports     []struct {
		Port     int32  `json:"port"`
		Protocol string `json:"protocol"`
	} `json:"ports"`
}

type AliveConnection struct {
	CheckTime int64
	FromIP    string
	FromType  string
	FromUID   string
	FromPort  uint16
	ToIP      string
	ToType    string
	ToUID     string
	ToPort    uint16
}

type KafkaEvent struct {
	StartTime int64
	Latency   uint64
	FromIP    string
	FromType  string
	FromUID   string
	FromPort  uint16
	ToIP      string
	ToType    string
	ToUID     string
	ToPort    uint16
	Topic     string
	Partition uint32
	Key       string
	Value     string
	Type      string
	Tls       bool
}

type BackendResponse struct {
	Message string `json:"msg"`
	Errors  []struct {
		EventNum int         `json:"event_num"`
		Event    interface{} `json:"event"`
		Error    string      `json:"error"`
	} `json:"errors"`
}

type ReqBackendReponse struct {
	Message string `json:"msg"`
	Errors  []struct {
		EventNum int         `json:"request_num"`
		Event    interface{} `json:"request"`
		Error    string      `json:"errors"`
	} `json:"errors"`
}

type Metadata struct {
	MonitoringID   string `json:"monitoring_id"`
	IdempotencyKey string `json:"idempotency_key"`
	NodeID         string `json:"node_id"`
	MarzyVersion   string `json:"alaz_version"`
}

type HealthCheckPayload struct {
	Metadata Metadata `json:"metadata"`
	Info     struct {
		TracingEnabled  bool   `json:"tracing"`
		MetricsEnabled  bool   `json:"metrics"`
		LogsEnabled     bool   `json:"logs"`
		NamespaceFilter string `json:"namespace_filter"`
	} `json:"alaz_info"`
	Telemetry struct {
		KernelVersion string `json:"kernel_version"`
		K8sVersion    string `json:"k8s_version"`
		CloudProvider string `json:"cloud_provider"`
	} `json:"telemetry"`
}

type EventPayload struct {
	Metadata Metadata      `json:"metadata"`
	Events   []interface{} `json:"events"`
}

type PodEvent struct {
	UID       string `json:"uid"`
	EventType string `json:"event_type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	IP        string `json:"ip"`
	OwnerType string `json:"owner_type"`
	OwnerName string `json:"owner_name"`
	OwnerID   string `json:"owner_id"`
}

type SvcEvent struct {
	UID        string   `json:"uid"`
	EventType  string   `json:"event_type"`
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	Type       string   `json:"type"`
	ClusterIPs []string `json:"cluster_ips"`
	Ports      []struct {
		Name     string `json:"name"`
		Src      int32  `json:"src"`
		Dest     int32  `json:"dest"`
		Protocol string `json:"protocol"`
	} `json:"ports"`
}

type RsEvent struct {
	UID       string `json:"uid"`
	EventType string `json:"event_type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas  int32  `json:"replicas"`
	OwnerType string `json:"owner_type"`
	OwnerName string `json:"owner_name"`
	OwnerID   string `json:"owner_id"`
}

type DsEvent struct {
	UID       string `json:"uid"`
	EventType string `json:"event_type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type SsEvent struct {
	UID       string `json:"uid"`
	EventType string `json:"event_type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type DepEvent struct {
	UID       string `json:"uid"`
	EventType string `json:"event_type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas  int32  `json:"replicas"`
}

type EpEvent struct {
	UID       string    `json:"uid"`
	EventType string    `json:"event_type"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Addresses []Address `json:"addresses"`
}

type ContainerEvent struct {
	UID       string `json:"uid"`
	EventType string `json:"event_type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
	Image     string `json:"image"`
	Ports     []struct {
		Port     int32  `json:"port"`
		Protocol string `json:"protocol"`
	} `json:"ports"`
}

type RequestsPayload struct {
	Metadata Metadata   `json:"metadata"`
	Requests []*ReqInfo `json:"requests"`
}

type ConnInfoPayload struct {
	Metadata    Metadata    `json:"metadata"`
	Connections []*ConnInfo `json:"connections"`
}

type TracePayload struct {
	Metadata Metadata     `json:"metadata"`
	Traces   []*TraceInfo `json:"traffic"`
}

type KafkaEventInfoPayload struct {
	Metadata    Metadata          `json:"metadata"`
	KafkaEvents []*KafkaEventInfo `json:"kafka_events"`
}

const (
	PodEndpoint                             = "/pod/"
	SvcEndpoint                             = "/svc/"
	RsEndpoint                              = "/replicaset/"
	DepEndpoint                             = "/deployment/"
	EpEndpoint                              = "/endpoint/"
	ContainerEndpoint                       = "/container/"
	DsEndpoint                              = "/daemonset/"
	SsEndpoint                              = "/statefulset/"
	ReqEndpoint                             = "/requests/"
	ConnEndpoint                            = "/connections/"
	KafkaEventEndpoint                      = "/events/kafka/"
	HealthCheckEndpoint                     = "/healthcheck/"
	CloudProviderAWS          CloudProvider = "AWS"
	CloudProviderGCP          CloudProvider = "GCP"
	CloudProviderAzure        CloudProvider = "Azure"
	CloudProviderDigitalOcean CloudProvider = "DigitalOcean"
	CloudProviderUnknown      CloudProvider = ""
)

var (
	MonitoringID        string
	NodeID              string
	Tag                 string
	kernelVersion       string
	CloudProvider       CloudProvider
	ResourceBatchSize   int64 = 1000
	InnerMetricsPort    int   = 8182
	InnerGpuMetricsPort int   = 8183
)
