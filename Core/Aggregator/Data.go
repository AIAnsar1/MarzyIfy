package Aggregator

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

func init() {

}

func NewAggregator(ParentCtx context.Context, Ct *Cri.CRITool, K8sChan chan interface{}, Events chan interface{}, ProcEvents chan interface{}, TcpEvents chan interface{},  TlsAttachSignalChan chan uint32, Ds datastore.DataStore) *Aggregator {

}

func (this *Aggregator) Run() {

}

func (this *Aggregator) ProcessK8S() {

}

func (this *Aggregator) ProccessEbpfProc(Ctx context.Context) {

}

func (this *Aggregator) ProccessEbpfTCP(Ctx context.Context) {

}

func (this *Aggregator) ProccessEbpf(Ctx context.Context) {

}

func (this *Aggregator) GetRateLimiterForPid(Pid uint32) *Rate.Limiter {

}

func (this *Aggregator) ProccessExec(D *Proc.ProcEvent) {

}
func (this *Aggregator) ProccessExit(Pid uint32) {

}
func (this *Aggregator) SignalTlsAttachment(Pid uint32) {

}
func (this *Aggregator) ProcessTcpConnect(Ctx context.Context, D *TcpState.TcpConnectEvent) {

}
func (this *Aggregator) ProcessHttp2Frames() {

}
func (this *Aggregator) GetPodWithIP(Address string) (types.UID, bool) {

}
func (this *Aggregator) GetSvcWithIP(Address string) (types.UID, bool) {

}
func (this *Aggregator) SetFromToV2(AddressPair *AddressPair, D *L7Req.L7Event, Event datastore.DirectionalEvent, HostHeader string) error {

}
func (this *Aggregator) SetFromTo(SkInfo *SockInfo, D *L7Req.L7Event, Event datastore.DirectionalEvent, HostHeader string) error {

}
func (this *Aggregator)  GetConnKey(Pid uint32, Fd uint64) string {

}
func (this *Aggregator) DecodeKafkaPayload(D *L7Req.L7Event) ([]*KafkaMessage, error) {

}
func (this *Aggregator) ProcessHttp2Event(D *L7Req.L7Event) {

}
func (this *Aggregator ProcessKafkaEvent(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator) ProcessAmqpEvent(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator) ProcessRedisEvent(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator) AdvertiseDebugData() {

}
func (this *Aggregator) ProcessHttpEvent(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator) ProcessMongoEvent(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator) ProcessMySQLEvent(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator) ProcessPostgresEvent(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator) ProcessL7(Ctx context.Context, D *L7Req.L7Event) {

}
func (this *Aggregator)  FindRelatedSocket(Ctx context.Context, D *L7Req.L7Event) (*SockInfo, error) {

}
func (this *Aggregator) ParseMySQLCommand(D *L7Req.L7Event) (string, error) {

}
func (this *Aggregator) ParsePostgresCommand(D *L7Req.L7Event) (string, error) {

}
func (this *Aggregator) ParseMongoEvent(D *L7Req.L7Event) (string, error) {

}
func (this *Aggregator) GetPgStmtKey(Pid uint32, Fd uint64, StmtName string) string {

}

func (this *Aggregator) SendOpenConnection(Sl *SocketLine) {

}
func (this *Aggregator) ClearSocketLines(Ctx context.Context) {

}

func ParseHttpPayload(Request string) (Method string, Path string, HttpVersion string, HostHeader string) {

}

func GetHostNameFromIP(IPAddress string) (string, error) {

}

func ContainsSQLKeywords(input string) bool {

}
func GetPidMax() (int, error) {

}

func ConvertKernelTimeToUserspaceTime(WriteTime uint64) uint64 {

}

func ConvertUserTimeToKernelTime(Now uint64) uint64 {

}

func IntToIPv4(ipaddr uint32) Net.IP {

}
func ExtractAddressPair(D *L7Req.L7Event) *AddressPair {

}
