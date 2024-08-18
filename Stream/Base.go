package Stream

import (
	"bufio"
	"context"
	"errors"
	"github.com/fsnotify/fsnotify"
	"net"
	"sync"
)

type PoolConnection struct {
	net.Conn
	Mu       sync.RWMutex
	C        *ChannelPool
	Unusable bool
	Tls      bool
}

type ChannelPool struct {
	Mu          sync.RWMutex
	Connections chan *PoolConnection
	Factory     Factory
	Tls         bool
}

type LogStreamer struct {
	CriTool                *Cri.CRITool
	ConnectPool            *ChannelPool
	Watcher                *fsnotify.Watcher
	ReaderMapMu            sync.RWMutex
	LogPathToFile          map[string]*FileReader
	MetaMapMu              sync.RWMutex
	LogPathToContainerMeta map[string]*FileReader
	IdtoPathMu             sync.RWMutex
	ContainerIdToLogPath   map[string]string
	Ctx                    context.Context
	Done                   chan struct{}
	MaxConnection          int
}

type FileReader struct {
	Mu sync.Mutex
	*bufio.Reader
}

var (
	ErrorClosed = errors.New("[ ERROR ]: Pool Is Closed!")
)

type Factory func() (net.Conn, error)
