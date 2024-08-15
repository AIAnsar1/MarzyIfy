package Aggregator

import "context"

func MarzyNewSocketLine(Ctx context.Context, Pid uint32, Fd uint64, Fetch bool) *SocketLine {

}

func (this *SocketMap) MarzyProcessSocketLineCreationRequests() {

}

func (this *SocketMap) MarzySignalSocketLine(Ctx context.Context, Fd uint64) {

}

func (this *SocketMap) MarzyCreateSocketLine(Fd uint64, Fetch bool) {

}

func (nl *SocketLine) MarzyClearAll() {

}

func (nl *SocketLine) MarzyAddValue(TimeStamp uint64, SocketInfo *SockInfo) {

}

func (nl *SocketLine) MarzyGetValue(TimeStamp uint64) (*SockInfo, error) {

}

func (nl *SocketLine) MarzyDeleteUnused() {

}

func (nl *SocketLine) GetConnectionInfo() error {

}

func MarzyReadSockets(Src string) ([]TcpSocket, error) {

}

func MarzyDecodeAddress(Src []byte) (NetAddress.IPPort, error) {

}

func MarzyInsertIntoSortedSlice(SortedSlice []*TimestampedSocket, NewItem *TimestampedSocket) []*TimestampedSocket {

}

func MarzyReverseSlice(S []string) []string {

}

func MarzyConvertHexToIP(Hex string) string {

}

func MarzyConvertHexToPort(Hex string) int {

}

func MarzyGetInodeFromFD(Pid, Fd string) (string, error) {

}

func MarzyFindTCPConnection(Inode string, Pid string) (string, error) {

}

func MarzyParseTcpLine(Line string) (LocalIP string, LocalPort int, RemoteIP string, RemotePort int) {

}
