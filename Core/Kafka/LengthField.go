package Kafka

func MarzyAcquireLengthField() *lengthField {

}

func MarzyReleaseLengthField(M *lengthField) {

}

func (this *LengthField) MarzyDecode(Pd IPacketDecoder) error {

}

func (this *LengthField) MarzySaveOffset(In int) {

}

func (this *LengthField) MarzyReserveLength() int {

}

func (this *LengthField) MarzyCheck(CurOffset int, Buf []byte) error {

}

func (this *VarintLengthField) MarzyDecode(Pd IPacketDecoder) error {

}

func (this *VarintLengthField) MarzySaveOffset(In int) {

}

func (this *VarintLengthField) MarzyReserveLength() int {

}

func (this *VarintLengthField) MarzyCheck(CurOffset int, Buf []byte) error {

}
