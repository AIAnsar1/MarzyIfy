package L7Request

type IConversion interface {
	MarzyString() string
}
type BpfLogMessage struct {
	Level    uint32
	LogMsg   [100]uint8
	FuncName [100]uint8
	Pid      uint32
	Arg1     uint64
	Arg2     uint64
	Arg3     uint64
}

type BpfL7Event struct {
	Fd                  uint64
	WriteTimeNs         uint64
	Pid                 uint32
	Status              uint32
	Duration            uint64
	Protocol            uint8
	Method              uint8
	Padding             uint16
	Payload             [1024]uint8
	PayloadSize         uint32
	PayloadReadComplete uint8
	Failed              uint8
	IsTls               uint8
	_                   [1]byte
	KafkaApiVersion     int16
	_                   [2]byte
	PrepStatementId     uint32
	Saddr               uint32
	Sport               uint16
	_                   [2]byte
	Daddr               uint32
	Dport               uint16
	_                   [6]byte
}
type L7Event struct {
	Fd                  uint64
	Pid                 uint32
	Status              uint32
	Duration            uint64
	Protocol            string
	Tls                 bool
	Method              string
	Payload             [1024]uint8
	PayloadSize         uint32
	PayloadReadComplete bool
	Failed              bool
	WriteTimeNs         uint64
	Tid                 uint32
	Seq                 uint32
	KafkaApiVersion     int16
	MySqlPrepStmtId     uint32
	Saddr               uint32
	Sport               uint16
	Daddr               uint32
	Dport               uint16
	PutBack             bool
}

type L7ProgConfig struct {
	TrafficBpfMapSize  uint32
	L7EventsBpfMapSize uint32
	LogsBpfMapSize     uint32
}

type L7Prog struct {
	Links           map[string]link.Link
	L7Events        *perf.Reader
	L7EventsMapSize uint32
	Logs            *perf.Reader
	LogsMapsSize    uint32
}

type L7ProtocolConversion uint32
type HTTPMethodConversion uint32
type RabbitMQMethodConversion uint32
type PostgresMethodConversion uint32
type KafkaMethodConversion uint32
type MySQLMethodConversion uint32
type Http2MethodConversion uint32
type RedisMethodConversion uint32

const L7_EVENT = "l7_event"

const (
	BPF_L7_PROTOCOL_UNKNOWN = iota
	BPF_L7_PROTOCOL_HTTP
	BPF_L7_PROTOCOL_AMQP
	BPF_L7_PROTOCOL_POSTGRES
	BPF_L7_PROTOCOL_HTTP2
	BPF_L7_PROTOCOL_REDIS
	BPF_L7_PROTOCOL_KAFKA
	BPF_L7_PROTOCOL_MYSQL
	BPF_L7_PROTOCOL_MONGO
)

const (
	L7_PROTOCOL_HTTP     = "HTTP"
	L7_PROTOCOL_HTTP2    = "HTTP2"
	L7_PROTOCOL_AMQP     = "AMQP"
	L7_PROTOCOL_POSTGRES = "POSTGRES"
	L7_PROTOCOL_REDIS    = "REDIS"
	L7_PROTOCOL_KAFKA    = "KAFKA"
	L7_PROTOCOL_MYSQL    = "MYSQL"
	L7_PROTOCOL_MONGO    = "MONGO"
	L7_PROTOCOL_UNKNOWN  = "UNKNOWN"
)

const (
	BPF_METHOD_UNKNOWN = iota
	BPF_METHOD_GET
	BPF_METHOD_POST
	BPF_METHOD_PUT
	BPF_METHOD_PATCH
	BPF_METHOD_DELETE
	BPF_METHOD_HEAD
	BPF_METHOD_CONNECT
	BPF_METHOD_OPTIONS
	BPF_METHOD_TRACE
)

const (
	BPF_HTTP2_METHOD_UNKNOWN = iota
	BPF_HTTP2_METHOD_CLIENT
	BPF_HTTP2_METHOD_SERVER
)

const (
	BPF_AMQP_METHOD_UNKNOWN = iota
	BPF_AMQP_METHOD_PUBLISH
	BPF_AMQP_METHOD_DELIVER
)

const (
	BPF_POSTGRES_METHOD_UNKNOWN = iota
	BPF_POSTGRES_METHOD_STATEMENT_CLOSE_OR_CONN_TERMINATE
	BPF_POSTGRES_METHOD_SIMPLE_QUERY
	BPF_POSTGRES_METHOD_EXTENDED_QUERY

	// BPF_POSTGRES_METHOD_QUERY
	// BPF_POSTGRES_METHOD_EXECUTE
	// BPF_POSTGRES_METHOD_PARSE
	// BPF_POSTGRES_METHOD_BIND
	// BPF_POSTGRES_METHOD_DESCRIBE
	// BPF_POSTGRES_METHOD_SYNC
	// BPF_POSTGRES_METHOD_FLUSH
	// BPF_POSTGRES_METHOD_CONSUME
	// BPF_POSTGRES_METHOD_PARSE_COMPLETE
	// BPF_POSTGRES_METHOD_BIND_COMPLETE
	// BPF_POSTGRES_METHOD_CLOSE_COMPLETE
	// BPF_POSTGRES_METHOD_SYNC_COMPLETE
	// BPF_POSTGRES_METHOD_READY_FOR_QUERY
	//...
)

const (
	BPF_REDIS_METHOD_UNKNOWN = iota
	METHOD_REDIS_COMMAND
	METHOD_REDIS_PUSHED_EVENT
	METHOD_REDIS_PING
)

const (
	BPF_KAFKA_METHOD_UNKNOWN = iota
	METHOD_KAFKA_PRODUCE_REQUEST
	METHOD_KAFKA_FETCH_RESPONSE
)

const (
	BPF_MYSQL_METHOD_UNKNOWN = iota
	METHOD_MYSQL_TEXT_QUERY
	METHOD_MYSQL_PREPARE_STMT
	METHOD_MYSQL_EXEC_STMT
	METHOD_MYSQL_STMT_CLOSE
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

const (
	PUBLISH = "PUBLISH"
	DELIVER = "DELIVER"
)

const (
	CLOSE_OR_TERMINATE = "CLOSE_OR_TERMINATE"
	SIMPLE_QUERY       = "SIMPLE_QUERY"
	EXTENDED_QUERY     = "EXTENDED_QUERY"
)

const (
	CLIENT_FRAME = "CLIENT_FRAME"
	SERVER_FRAME = "SERVER_FRAME"
)

const (
	REDIS_COMMAND      = "COMMAND"
	REDIS_PUSHED_EVENT = "PUSHED_EVENT"
	REDIS_PING         = "PING"
)

const (
	KAFKA_PRODUCE_REQUEST = "PRODUCE_REQUEST"
	KAFKA_FETCH_RESPONSE  = "FETCH_RESPONSE"
)

const (
	MYSQL_TEXT_QUERY   = "TEXT_QUERY"
	MYSQL_PREPARE_STMT = "PREPARE_STMT"
	MYSQL_EXEC_STMT    = "EXEC_STMT"
	MYSQL_STMT_CLOSE   = "STMT_CLOSE"
)

var (
	FirstKernelTime    uint64 = 0
	FirstUserspaceTime uint64 = 0

	defaultConfig *L7ProgConfig = &L7ProgConfig{
		TrafficBpfMapSize:  4096,
		L7EventsBpfMapSize: 4096,
		LogsBpfMapSize:     4,
	}
)
