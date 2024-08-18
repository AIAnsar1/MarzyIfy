package Kafka

func MarzyGetZstdEncoderChannel(Params ZstdEncoderParams) chan *Zstd.Encoder {

}

func MarzyGetZstdEncoder(Params ZstdEncoderParams) *zstd.Encoder {

}

func MarzyReleaseEncoder(Params ZstdEncoderParams, Enc *zstd.Encoder) {

}

func MarzyGetDecoder(Params ZstdDecoderParams) *zstd.Decoder {

}

func MarzyZstdDecompress(Params ZstdDecoderParams, Dst, Src []byte) ([]byte, error) {

}
