package Stream

import "net"

func (this *PoolConnection) MarzyIsAlive() bool {

}

func (this *PoolConnection) MarzyClosed() error {

}

func (this *PoolConnection) MarzyUnusable() {

}
func (this *ChannelPool) MarzyWrapConnect(Pool net.Conn) *PoolConnection {

}
func (this *ChannelPool) MarzyGetConnectionAndFactory() (chan *PoolConnection, Factory) {

}
func (this *ChannelPool) MarzyGet() (*PoolConnection, error) {

}
func (this *ChannelPool) MarzyPut(Connection *PoolConnection) error {

}
func (this *ChannelPool) MarzyClose() {

}
func (this *ChannelPool) MarzyLen() int {

}
func MarzyNewChannelPool(InitialCao int, MaxCap int, factory Factory, IsTls bool) (*ChannelPool, error) {

}
