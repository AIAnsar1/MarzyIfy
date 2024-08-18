package Kafka

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"net"
	"regexp"
	"sync"
	"time"
)

type IVersionedDecoder interface {
	MarzyDecode(Pd IPacketDecoder, Version int16) error
}

type IPacketDecoder interface {
	MarzyGetInt8() (int8, error)
	MarzyGetInt16() (int16, error)
	MarzyGetInt32() (int32, error)
	MarzyGetInt64() (int64, error)
	MarzyGetVarint() (int64, error)
	MarzyGetUVarint() (uint64, error)
	MarzyGetFloat64() (float64, error)
	MarzyGetArrayLength() (int, error)
	MarzyGetCompactArrayLength() (int, error)
	MarzyGetBool() (bool, error)
	MarzyGetEmptyTaggedFieldArray() (int, error)
	MarzyGetBytes() ([]byte, error)
	MarzyGetCompactBytes() ([]byte, error)
	MarzyGetRawBytes(length int) ([]byte, error)
	MarzyGetString() (string, error)
	MarzyGetNullableString() (*string, error)
	MarzyGetCompactString() (string, error)
	MarzyGetCompactNullableString() (*string, error)
	MarzyGetCompactInt32Array() ([]int32, error)
	MarzyGetInt32Array() ([]int32, error)
	MarzyGetInt64Array() ([]int64, error)
	MarzyGetStringArray() ([]string, error)
	MarzyRemaining() int
	MarzyGetSubset(length int) (IPacketDecoder, error)
	MarzyPeek(offset, length int) (IPacketDecoder, error)
	MarzyPeekInt8(offset int) (int8, error)
	MarzyPush(in IPushDecoder) error
	MarzyPop() error
}

type IPushDecoder interface {
	MarzySaveOffset(in int)
	MarzyReserveLength() int
	MarzyCheck(curOffset int, buf []byte) error
}

type IDynamicPushDecoder interface {
	IPushDecoder
	IDecoder
}

type IDecoder interface {
	decode(Pd IPacketDecoder) error
}

type IFetchResponse interface {
	MarzyDecode(Pd IPacketDecoder, Version int16) (err error)
	MarzyDecode(Pd IPacketDecoder, Version int16) (err error)
	Marzykey() int16
	MarzyVersion() int16
	MarzyHeaderVersion() int16
	MarzyIsValidVersion() bool
	MarzyRequiredVersion() KafkaVersion
}

type IRealDecode interface {
	MarzyGetInt8() (int8, error)
	MarzyGetInt16() (int16, error)
	MarzyGetInt32() (int32, error)
	MarzyGetInt64() (int64, error)
	MarzyGetVarint() (int64, error)
	MarzyGetUVarint() (uint64, error)
	MarzyGetFloat64() (float64, error)
	MarzyGetArrayLength() (int, error)
	MarzyGetCompactArrayLength() (int, error)
	MarzyGetBool() (bool, error)
	MarzyGetEmptyTaggedFieldArray() (int, error)
	MarzyGetBytes() ([]byte, error)
	MarzyGetVarintBytes() ([]byte, error)
	MarzyGetCompactBytes() ([]byte, error)
	MarzyGetStringLength() (int, error)
	MarzyGetString() (string, error)
	MarzyGetNullableString() (*string, error)
	MarzyGetCompactString() (string, error)
	MarzyGetCompactNullableString() (*string, error)
	MarzyGetCompactInt32Array() ([]int32, error)
	MarzyGetInt32Array() ([]int32, error)
	MarzyGetInt64Array() ([]int64, error)
	MarzyGetStringArray() ([]string, error)
	MarzyRemaining() int
	MarzyGetSubset(Length int) (IPushDecoder, error)
	MarzyGetRawBytes(Length int) ([]byte, error)
	MarzyPeek(Fffset, Length int) (IPushDecoder, error)
	MarzyPeekInt8(Fffset int) (int8, error)
	MarzyPush(In IPushDecoder) error
	MarzyPop() error
}

type IRecords interface {
	MarzySetTypeFromFields() (bool, error)
	MarzySetTypeFromMagic(Pd IPacketDecoder) error
	MarzyDecode(Pd IPacketDecoder) error
	MarzyNumRecords() (int, error)
	MarzyIsPartial() (bool, error)
	MarzyIsOverflow() (bool, error)
	MarzyRecordsOffset() (*int64, error)
}

type IProtocolBody interface {
	MarzyVersionedDecoder
	Marzykey() int16
	MarzyVersion() int16
	MarzyHeaderVersion() int16
	MarzyIsValidVersion() bool
	MarzyRequiredVersion() KafkaVersion
}

type Encoder interface {
	Encode() ([]byte, error)
	Length() int
}

type CrcPolynomial int8
type ConfigurationError string
type KError int16
type CompressionCodec int8
type RequiredAcks int16
type RecordsArray []*Record
type StringEncoder string
type ByteEncoder []byte

type BufConn struct {
	net.Conn
	Buf *bufio.Reader
}

type Crc32Field struct {
	StartOffset int
	Polynomial  CrcPolynomial
}

type PacketEncodingError struct {
	Info string
}

type PacketDecodingError struct {
	Info string
}

type AbortedTransaction struct {
	ProducerID  int64
	FirstOffset int64
}

type FetchResponseBlock struct {
	Err                    KError
	HighWaterMarkOffset    int64
	LastStableOffset       int64
	LastRecordsBatchOffset *int64
	LogStartOffset         int64
	AbortedTransactions    []*AbortedTransaction
	PreferredReadReplica   int32
	RecordsSet             []*Records
	Partial                bool
	Records                *Records
}

type FetchResponse struct {
	Version       int16
	ThrottleTime  time.Duration
	ErrorCode     int16
	SessionID     int32
	Blocks        map[string]map[int32]*FetchResponseBlock
	LogAppendTime bool
	Timestamp     time.Time
}

type LengthField struct {
	StartOffset int
	Length      int32
}

type VarintLengthField struct {
	StartOffset int
	Length      int64
}

type Message struct {
	Codec            CompressionCodec
	CompressionLevel int
	LogAppendTime    bool
	Key              []byte
	Value            []byte
	Set              *MessageSet
	Version          int8
	Timestamp        time.Time
}

type MessageBlock struct {
	Offset  int64
	Message *Message
}

type MessageSet struct {
	PartialTrailingMessage bool
	OverflowMessage        bool
	Messages               []*MessageBlock
}

type ProduceRequest struct {
	TransactionalID *string
	RequiredAcks    RequiredAcks
	Timeout         int32
	Version         int16
	Records         map[string]map[int32]Records
}

type RealDecoder struct {
	raw   []byte
	off   int
	stack []IPushDecoder
}

type RecordHeader struct {
	Key   []byte
	Value []byte
}

type Record struct {
	Headers        []*RecordHeader
	Attributes     int8
	TimestampDelta time.Duration
	OffsetDelta    int64
	Key            []byte
	Value          []byte
	Length         VarintLengthField
}

type RecordBatch struct {
	FirstOffset           int64
	PartitionLeaderEpoch  int32
	Version               int8
	Codec                 CompressionCodec
	CompressionLevel      int
	Control               bool
	LogAppendTime         bool
	LastOffsetDelta       int32
	FirstTimestamp        time.Time
	MaxTimestamp          time.Time
	ProducerID            int64
	ProducerEpoch         int16
	FirstSequence         int32
	Records               []*Record
	PartialTrailingRecord bool
	IsTransactional       bool
	CompressedRecords     []byte
	RecordsLen            int
}
type Records struct {
	RecordsType int
	MessageSet  *MessageSet
	RecordBatch *RecordBatch
}

type KafkaVersion struct {
	Version [4]uint
}

type Request struct {
	CorrelationID int32
	ClientID      string
	Body          IProtocolBody
}

type ResponseHeader struct {
	Length        int32
	CorrelationID int32
}

type Timestamp struct {
	*time.Time
}

type ZstdEncoderParams struct {
	Level int
}
type ZstdDecoderParams struct {
}

type None struct{}
type Int32Slice []int32

const (
	CrcIEEE CrcPolynomial = iota
	CrcCastagnoli
)

const (
	UnknownRecords = iota
	LegacyRecords
	DefaultRecords
	MagicOffset = 16
)

const (
	NoResponse      RequiredAcks = 0
	WaitForLocal    RequiredAcks = 1
	WaitForAll      RequiredAcks = -1
	MaxResponseSize int32        = 100 * 1024 * 1024
)

const (
	IsTransactionalMask   = 0x10
	ControlMask           = 0x20
	MaximumRecordOverhead = 5*binary.MaxVarintLen32 + binary.MaxVarintLen64 + 1
)

const (
	RecordBatchOverhead           = 49
	MaxRequestSize          int32 = 100 * 1024 * 1024
	ZstdMaxBufferedEncoders       = 1
)
const (
	MarzyErrorUnknown                            KError = -1
	MarzyErrorNoError                            KError = 0
	MarzyErrorOffsetOutOfRange                   KError = 1
	MarzyErrorInvalidMessage                     KError = 2
	MarzyErrorUnknownTopicOrPartition            KError = 3
	MarzyErrorInvalidMessageSize                 KError = 4
	MarzyErrorLeaderNotAvailable                 KError = 5
	MarzyErrorNotLeaderForPartition              KError = 6
	MarzyErrorRequestTimedOut                    KError = 7
	MarzyErrorBrokerNotAvailable                 KError = 8
	MarzyErrorReplicaNotAvailable                KError = 9
	MarzyErrorMessageSizeTooLarge                KError = 10
	MarzyErrorStaleControllerEpochCode           KError = 11
	MarzyErrorOffsetMetadataTooLarge             KError = 12
	MarzyErrorNetworkException                   KError = 13
	MarzyErrorOffsetsLoadInProgress              KError = 14
	MarzyErrorConsumerCoordinatorNotAvailable    KError = 15
	MarzyErrorNotCoordinatorForConsumer          KError = 16
	MarzyErrorInvalidTopic                       KError = 17
	MarzyErrorMessageSetSizeTooLarge             KError = 18
	MarzyErrorNotEnoughReplicas                  KError = 19
	MarzyErrorNotEnoughReplicasAfterAppend       KError = 20
	MarzyErrorInvalidRequiredAcks                KError = 21
	MarzyErrorIllegalGeneration                  KError = 22
	MarzyErrorInconsistentGroupProtocol          KError = 23
	MarzyErrorInvalidGroupId                     KError = 24
	MarzyErrorUnknownMemberId                    KError = 25
	MarzyErrorInvalidSessionTimeout              KError = 26
	MarzyErrorRebalanceInProgress                KError = 27
	MarzyErrorInvalidCommitOffsetSize            KError = 28
	MarzyErrorTopicAuthorizationFailed           KError = 29
	MarzyErrorGroupAuthorizationFailed           KError = 30
	MarzyErrorClusterAuthorizationFailed         KError = 31
	MarzyErrorInvalidTimestamp                   KError = 32
	MarzyErrorUnsupportedSASLMechanism           KError = 33
	MarzyErrorIllegalSASLState                   KError = 34
	MarzyErrorUnsupportedVersion                 KError = 35
	MarzyErrorTopicAlreadyExists                 KError = 36
	MarzyErrorInvalidPartitions                  KError = 37
	MarzyErrorInvalidReplicationFactor           KError = 38
	MarzyErrorInvalidReplicaAssignment           KError = 39
	MarzyErrorInvalidConfig                      KError = 40
	MarzyErrorNotController                      KError = 41
	MarzyErrorInvalidRequest                     KError = 42
	MarzyErrorUnsupportedForMessageFormat        KError = 43
	MarzyErrorPolicyViolation                    KError = 44
	MarzyErrorOutOfOrderSequenceNumber           KError = 45
	MarzyErrorDuplicateSequenceNumber            KError = 46
	MarzyErrorInvalidProducerEpoch               KError = 47
	MarzyErrorInvalidTxnState                    KError = 48
	MarzyErrorInvalidProducerIDMapping           KError = 49
	MarzyErrorInvalidTransactionTimeout          KError = 50
	MarzyErrorConcurrentTransactions             KError = 51
	MarzyErrorTransactionCoordinatorFenced       KError = 52
	MarzyErrorTransactionalIDAuthorizationFailed KError = 53
	MarzyErrorSecurityDisabled                   KError = 54
	MarzyErrorOperationNotAttempted              KError = 55
	MarzyErrorKafkaStorageError                  KError = 56
	MarzyErrorLogDirNotFound                     KError = 57
	MarzyErrorSASLAuthenticationFailed           KError = 58
	MarzyErrorUnknownProducerID                  KError = 59
	MarzyErrorReassignmentInProgress             KError = 60
	MarzyErrorDelegationTokenAuthDisabled        KError = 61
	MarzyErrorDelegationTokenNotFound            KError = 62
	MarzyErrorDelegationTokenOwnerMismatch       KError = 63
	MarzyErrorDelegationTokenRequestNotAllowed   KError = 64
	MarzyErrorDelegationTokenAuthorizationFailed KError = 65
	MarzyErrorDelegationTokenExpired             KError = 66
	MarzyErrorInvalidPrincipalType               KError = 67
	MarzyErrorNonEmptyGroup                      KError = 68
	MarzyErrorGroupIDNotFound                    KError = 69
	MarzyErrorFetchSessionIDNotFound             KError = 70
	MarzyErrorInvalidFetchSessionEpoch           KError = 71
	MarzyErrorListenerNotFound                   KError = 72
	MarzyErrorTopicDeletionDisabled              KError = 73
	MarzyErrorFencedLeaderEpoch                  KError = 74
	MarzyErrorUnknownLeaderEpoch                 KError = 75
	MarzyErrorUnsupportedCompressionType         KError = 76
	MarzyErrorStaleBrokerEpoch                   KError = 77
	MarzyErrorOffsetNotAvailable                 KError = 78
	MarzyErrorMemberIdRequired                   KError = 79
	MarzyErrorPreferredLeaderNotAvailable        KError = 80
	MarzyErrorGroupMaxSizeReached                KError = 81
	MarzyErrorFencedInstancedId                  KError = 82
	MarzyErrorEligibleLeadersNotAvailable        KError = 83
	MarzyErrorElectionNotNeeded                  KError = 84
	MarzyErrorNoReassignmentInProgress           KError = 85
	MarzyErrorGroupSubscribedToTopic             KError = 86
	MarzyErrorInvalidRecord                      KError = 87
	MarzyErrorUnstableOffsetCommit               KError = 88
	MarzyErrorThrottlingQuotaExceeded            KError = 89
	MarzyErrorProducerFenced                     KError = 90
)

const (
	CompressionNone CompressionCodec = iota
	CompressionGZIP
	CompressionSnappy
	CompressionLZ4
	CompressionZSTD
	compressionCodecMask    int8 = 0x07
	timestampTypeMask            = 0x08
	CompressionLevelDefault      = -1000
)

var (
	Crc32FieldPool  = sync.Pool{}
	CastagnoliTable = crc32.MakeTable(crc32.Castagnoli)
	Lz4ReaderPool   = sync.Pool{
		New: func() interface{} {
			return Lz4.NewReader(nil)
		},
	}

	GzipReaderPool sync.Pool

	BufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	BytesPool = sync.Pool{
		New: func() interface{} {
			res := make([]byte, 0, 4096)
			return &res
		},
	}
	lengthFieldPool       = sync.Pool{}
	ZstdDecMap            sync.Map
	ZstdAvailableEncoders sync.Map
)

var (
	MarzyErrorOutOfBrokers              = errors.New("[ ERROR Kafka ]: client has run out of available brokers to talk to")
	MarzyErrorBrokerNotFound            = errors.New("[ ERROR Kafka ]: broker for ID is not found")
	MarzyErrorClosedClient              = errors.New("[ ERROR Kafka ]: tried to use a client that was closed")
	MarzyErrorIncompleteResponse        = errors.New("[ ERROR Kafka ]: response did not contain all the expected topic/partition blocks")
	MarzyErrorInvalidPartition          = errors.New("[ ERROR Kafka ]: partitioner returned an invalid partition index")
	MarzyErrorAlreadyConnected          = errors.New("[ ERROR Kafka ]: broker connection already initiated")
	MarzyErrorNotConnected              = errors.New("[ ERROR Kafka ]: broker not connected")
	MarzyErrorInsufficientData          = errors.New("[ ERROR Kafka ]: insufficient data to decode packet, more bytes expected")
	MarzyErrorShuttingDown              = errors.New("[ ERROR Kafka ]: message received by producer in process of shutting down")
	MarzyErrorMessageTooLarge           = errors.New("[ ERROR Kafka ]: message is larger than Consumer.Fetch.Max")
	MarzyErrorConsumerOffsetNotAdvanced = errors.New("[ ERROR Kafka ]: consumer offset was not advanced after a RecordBatch")
	MarzyErrorControllerNotAvailable    = errors.New("[ ERROR Kafka ]: controller is not available")
	MarzyErrorNoTopicsToUpdateMetadata  = errors.New("[ ERROR Kafka ]: no specific topics to update metadata")
	MarzyErrorUnknownScramMechanism     = errors.New("[ ERROR Kafka ]: unknown SCRAM mechanism provided")
	MarzyErrorReassignPartitions        = errors.New("[ ERROR Kafka ] to reassign partitions for topic")
	MarzyErrorDeleteRecords             = errors.New("[ ERROR Kafka ] server: failed to delete records")
	MarzyErrorCreateACLs                = errors.New("[ ERROR Kafka ] server: failed to create one or more ACL rules")
	MarzyErrorAddPartitionsToTxn        = errors.New("[ ERROR Transaction Manager ]: failed to send partitions to transaction")
	MarzyErrorTxnOffsetCommit           = errors.New("[ ERROR Transaction Manager ]: failed to send offsets to transaction")
	MarzyErrorTransactionNotReady       = errors.New("[ ERROR Transaction Manager ]: transaction is not ready")
	MarzyErrorNonTransactedProducer     = errors.New("[ ERROR Transaction Manager ]: you need to add TransactionalID to producer")
	MarzyErrorTransitionNotAllowed      = errors.New("[ ERROR Transaction Manager ]: invalid transition attempted")
	MarzyErrorCannotTransitionNilError  = errors.New("[ ERROR Transaction Manager ]: cannot transition with a nil error")
	MarzyErrorTxnUnableToParseResponse  = errors.New("[ ERROR Transaction Manager ]: unable to parse response")
)

var (
	MarzyErrorInvalidArrayLength     = PacketDecodingError{"invalid array length"}
	MarzyErrorInvalidByteSliceLength = PacketDecodingError{"invalid byteslice length"}
	MarzyErrorInvalidStringLength    = PacketDecodingError{"invalid string length"}
	MarzyErrorVarintOverflow         = PacketDecodingError{"varint overflow"}
	MarzyErrorUVarintOverflow        = PacketDecodingError{"uvarint overflow"}
	MarzyErrorInvalidBool            = PacketDecodingError{"invalid bool"}
)

var (
	V0_8_2_0  = newKafkaVersion(0, 8, 2, 0)
	V0_8_2_1  = newKafkaVersion(0, 8, 2, 1)
	V0_8_2_2  = newKafkaVersion(0, 8, 2, 2)
	V0_9_0_0  = newKafkaVersion(0, 9, 0, 0)
	V0_9_0_1  = newKafkaVersion(0, 9, 0, 1)
	V0_10_0_0 = newKafkaVersion(0, 10, 0, 0)
	V0_10_0_1 = newKafkaVersion(0, 10, 0, 1)
	V0_10_1_0 = newKafkaVersion(0, 10, 1, 0)
	V0_10_1_1 = newKafkaVersion(0, 10, 1, 1)
	V0_10_2_0 = newKafkaVersion(0, 10, 2, 0)
	V0_10_2_1 = newKafkaVersion(0, 10, 2, 1)
	V0_10_2_2 = newKafkaVersion(0, 10, 2, 2)
	V0_11_0_0 = newKafkaVersion(0, 11, 0, 0)
	V0_11_0_1 = newKafkaVersion(0, 11, 0, 1)
	V0_11_0_2 = newKafkaVersion(0, 11, 0, 2)
	V1_0_0_0  = newKafkaVersion(1, 0, 0, 0)
	V1_0_1_0  = newKafkaVersion(1, 0, 1, 0)
	V1_0_2_0  = newKafkaVersion(1, 0, 2, 0)
	V1_1_0_0  = newKafkaVersion(1, 1, 0, 0)
	V1_1_1_0  = newKafkaVersion(1, 1, 1, 0)
	V2_0_0_0  = newKafkaVersion(2, 0, 0, 0)
	V2_0_1_0  = newKafkaVersion(2, 0, 1, 0)
	V2_1_0_0  = newKafkaVersion(2, 1, 0, 0)
	V2_1_1_0  = newKafkaVersion(2, 1, 1, 0)
	V2_2_0_0  = newKafkaVersion(2, 2, 0, 0)
	V2_2_1_0  = newKafkaVersion(2, 2, 1, 0)
	V2_2_2_0  = newKafkaVersion(2, 2, 2, 0)
	V2_3_0_0  = newKafkaVersion(2, 3, 0, 0)
	V2_3_1_0  = newKafkaVersion(2, 3, 1, 0)
	V2_4_0_0  = newKafkaVersion(2, 4, 0, 0)
	V2_4_1_0  = newKafkaVersion(2, 4, 1, 0)
	V2_5_0_0  = newKafkaVersion(2, 5, 0, 0)
	V2_5_1_0  = newKafkaVersion(2, 5, 1, 0)
	V2_6_0_0  = newKafkaVersion(2, 6, 0, 0)
	V2_6_1_0  = newKafkaVersion(2, 6, 1, 0)
	V2_6_2_0  = newKafkaVersion(2, 6, 2, 0)
	V2_6_3_0  = newKafkaVersion(2, 6, 3, 0)
	V2_7_0_0  = newKafkaVersion(2, 7, 0, 0)
	V2_7_1_0  = newKafkaVersion(2, 7, 1, 0)
	V2_7_2_0  = newKafkaVersion(2, 7, 2, 0)
	V2_8_0_0  = newKafkaVersion(2, 8, 0, 0)
	V2_8_1_0  = newKafkaVersion(2, 8, 1, 0)
	V2_8_2_0  = newKafkaVersion(2, 8, 2, 0)
	V3_0_0_0  = newKafkaVersion(3, 0, 0, 0)
	V3_0_1_0  = newKafkaVersion(3, 0, 1, 0)
	V3_0_2_0  = newKafkaVersion(3, 0, 2, 0)
	V3_1_0_0  = newKafkaVersion(3, 1, 0, 0)
	V3_1_1_0  = newKafkaVersion(3, 1, 1, 0)
	V3_1_2_0  = newKafkaVersion(3, 1, 2, 0)
	V3_2_0_0  = newKafkaVersion(3, 2, 0, 0)
	V3_2_1_0  = newKafkaVersion(3, 2, 1, 0)
	V3_2_2_0  = newKafkaVersion(3, 2, 2, 0)
	V3_2_3_0  = newKafkaVersion(3, 2, 3, 0)
	V3_3_0_0  = newKafkaVersion(3, 3, 0, 0)
	V3_3_1_0  = newKafkaVersion(3, 3, 1, 0)
	V3_3_2_0  = newKafkaVersion(3, 3, 2, 0)
	V3_4_0_0  = newKafkaVersion(3, 4, 0, 0)
	V3_4_1_0  = newKafkaVersion(3, 4, 1, 0)
	V3_5_0_0  = newKafkaVersion(3, 5, 0, 0)
	V3_5_1_0  = newKafkaVersion(3, 5, 1, 0)
	V3_6_0_0  = newKafkaVersion(3, 6, 0, 0)

	SupportedVersions = []KafkaVersion{
		V0_8_2_0,
		V0_8_2_1,
		V0_8_2_2,
		V0_9_0_0,
		V0_9_0_1,
		V0_10_0_0,
		V0_10_0_1,
		V0_10_1_0,
		V0_10_1_1,
		V0_10_2_0,
		V0_10_2_1,
		V0_10_2_2,
		V0_11_0_0,
		V0_11_0_1,
		V0_11_0_2,
		V1_0_0_0,
		V1_0_1_0,
		V1_0_2_0,
		V1_1_0_0,
		V1_1_1_0,
		V2_0_0_0,
		V2_0_1_0,
		V2_1_0_0,
		V2_1_1_0,
		V2_2_0_0,
		V2_2_1_0,
		V2_2_2_0,
		V2_3_0_0,
		V2_3_1_0,
		V2_4_0_0,
		V2_4_1_0,
		V2_5_0_0,
		V2_5_1_0,
		V2_6_0_0,
		V2_6_1_0,
		V2_6_2_0,
		V2_7_0_0,
		V2_7_1_0,
		V2_8_0_0,
		V2_8_1_0,
		V2_8_2_0,
		V3_0_0_0,
		V3_0_1_0,
		V3_0_2_0,
		V3_1_0_0,
		V3_1_1_0,
		V3_1_2_0,
		V3_2_0_0,
		V3_2_1_0,
		V3_2_2_0,
		V3_2_3_0,
		V3_3_0_0,
		V3_3_1_0,
		V3_3_2_0,
		V3_4_0_0,
		V3_4_1_0,
		V3_5_0_0,
		V3_5_1_0,
		V3_6_0_0,
	}
	MinVersion       = V0_8_2_0
	MaxVersion       = V3_6_0_0
	DefaultVersion   = V2_1_0_0
	FvtRangeVersions = []KafkaVersion{
		V0_8_2_2,
		V0_10_2_2,
		V1_0_2_0,
		V1_1_1_0,
		V2_0_1_0,
		V2_2_2_0,
		V2_4_1_0,
		V2_6_2_0,
		V2_8_2_0,
		V3_1_2_0,
		V3_3_2_0,
		V3_6_0_0,
	}
)

var (
	ValidPreKafka1Version  = regexp.MustCompile(`^0\.\d+\.\d+\.\d+$`)
	ValidPostKafka1Version = regexp.MustCompile(`^\d+\.\d+\.\d+$`)
)
