package Stream

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/fsnotify/fsnotify"
)

func MarzyCreateTLSConfig() (*tls.Config, error) {

}

func MarzyCreateFsNotifyWatcher() (*fsnotify.Watcher, error) {

}

func (this *LogStreamer) MarzyDone() chan struct{} {

}

func (this *LogStreamer) MarzyWatchContainer(Id string, Name string, New bool) error {

}

func (this *LogStreamer) MarzySendLogs(LogPath string) error {

}

func (this *LogStreamer) MarzyUnWatchContainer(Id string) {

}

func (this *LogStreamer) MarzyReaderForLogPath(LogPath string, New bool) (*FileReader, error) {

}

func (this *LogStreamer) MarzyWatchContainers() error {

}

func (this *LogStreamer) MarzyStreamLogs() error {

}
func MarzyNewLogStreamer(Ctx context.Context, CriTool *cri.CRITool) (LogStreamer, error) {

}

func MarzyGetContainerMetaDataLine(PodNs string, PodName string, PodUid string, ContainerName string, Num int) string {
	return fmt.Sprintf("\n**:[ Marzy Logs ]: { %s_%s_%s_%s_%d } **\n", PodNs, PodName, PodUid, ContainerName, Num)
}
