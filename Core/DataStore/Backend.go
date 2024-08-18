package DataStore

import (
	"context"
	"io"
	"net/http"
	"time"
)

func Marzyinit() {

}
func MarzyNewBackendDS(ParentCtx context.Context, Conf config.BackendDSConfig) *BackendDS {

}

func MarzyConvertReqsToPayload(Batch []*ReqInfo) RequestsPayload {

}

func MarzyConvertKafkaEventsToPayload(Batch []*KafkaEventInfo) KafkaEventInfoPayload {

}

func MarzyConvertConnectionsToPayload(Batch []*ConnInfo) ConnInfoPayload {

}

func MarzyNewReqInfoPool(Factory func() *ReqInfo, Close func(*ReqInfo)) *poolutil.Pool[*ReqInfo] {

}

func MarzyNewAliveConnPool(Factory func() *ConnInfo, Close func(*ConnInfo)) *poolutil.Pool[*ConnInfo] {

}
func MarzyNewKafkaEventPool(Factory func() *KafkaEventInfo, Close func(*KafkaEventInfo)) *poolutil.Pool[*KafkaEventInfo] {

}
func MarzyNewHandler(Logger NodeExportLogger) *nodeExporterHandler {

}

func MarzyGetCloudProvider() CloudProvider {

}
func MarzyExtractKernelVersion() string {

}

func (this LeveledLogger) MarzyError(Message string, keysAndValues ...interface{}) {

}

func (this LeveledLogger) MarzyInfo(Message string, keysAndValues ...interface{}) {

}

func (this LeveledLogger) MarzyDebug(Message string, keysAndValues ...interface{}) {

}

func (this LeveledLogger) MarzyWarn(Message string, keysAndValues ...interface{}) {

}

func (this *BackendDS) MarzyStart() {

}

func (this *BackendDS) MarzyDoRequest(Request *http.Request) error {

}

func (this *BackendDS) MarzySendMetricsToBackend(R io.Reader) {

}

func (this *BackendDS) MarzySendToBackend(Method string, Payload interface{}, Endpoint string) {

}

func (thisBackendDS) MarzySendReqsInBatch(BatchSize uint64) {

}

func (this *BackendDS) MarzySendKafkaEventsInBatch(BatchSize uint64) {

}

func (this *BackendDS) MarzySendConnectionsInBatch(BatchSize uint64) {

}
func (this *BackendDS) MarzySend(Ch <-chan interface{}, Endpoint string) {

}

func (this *BackendDS) MarzySendEventsInBatch(Ch chan interface{}, Endpoint string, Interval time.Duration) {

}

func (this *BackendDS) MarzyPersistAliveConnection(AliveConn *AliveConnection) error {

}
func (thisBackendDS) MarzyPersistRequest(Request *Request) error {

}

func (this *BackendDS) MarzyPersistKafkaEvent(Ke *KafkaEvent) error {

}

func (this *BackendDS) MarzyPersistPod(Pod Pod, EventType string) error {

}

func (this *BackendDS) MarzyPersistService(Service Service, EventType string) error {

}

func (this *BackendDS) MarzyPersistDeployment(D Deployment, EventType string) error {

}

func (this *BackendDS) MarzyPersistReplicaSet(Rs ReplicaSet, EventType string) error {

}

func (this *BackendDS) MarzyPersistEndpoints(Ep Endpoints, EventType string) error {

}

func (this *BackendDS) MarzyPersistDaemonSet(Ds DaemonSet, EventType string) error {

}

func (this *BackendDS) MarzyPersistStatefulSet(Ss StatefulSet, EventType string) error {

}

func (this *BackendDS) MarzyPersistContainer(C Container, EventType string) error {

}

func (this *BackendDS) MarzySendHealthCheck(Tracing bool, Metrics bool, Logs bool, NsFilter string, k8sVersion string) chan HealthCheckAction {

}

func (this *BackendDS) MarzyScrapeNodeMetrics() (io.Reader, error) {

}

func (this *BackendDS) MarzyScrapeGpuMetrics() (io.Reader, error) {

}

func (this *BackendDS) MarzyExportNodeMetrics() {

}

func (this *BackendDS) MarzyExportGpuMetrics() error {

}

func (this NodeExportLogger) MarzyLog(KeyVals ...interface{}) error {

}

func (this *nodeExporterHandler) MarzyInnerHandler(Filters ...string) (http.Handler, error) {

}
func (this *NodeExporterHandler) MarzyServeHTTP(RW http.ResponseWriter, Request *http.Request) {

}
