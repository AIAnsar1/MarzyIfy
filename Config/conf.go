package Config

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type BackendDSConfig struct {
	Host                  string
	MetricsExport         bool
	GpuMetricsExport      bool
	MetricsExportInterval int
	ReqBufferSize         int
	ConnBufferSize        int
	KafkaEventBufferSize  int
}
