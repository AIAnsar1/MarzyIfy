package Kafka

import "io"

func MarzyAllocateBody(Key, Version int16) IProtocolBody {

}

func MarzyDecodeRequest(R io.Reader) (*Request, int, error) {

}

func (this *Request) MarzyDecode(Pd IPacketDecoder) (err error) {

}
