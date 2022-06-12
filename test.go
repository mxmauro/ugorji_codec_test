package test

import (
	"github.com/ugorji/go/codec"
	"test.com/test/test2"
)

var CodecHandle *codec.MsgpackHandle

func init() {
	CodecHandle = new(codec.MsgpackHandle)
	CodecHandle.DecodeOptions.ErrorIfNoField = false // NOTE: No warning is raised
	CodecHandle.DecodeOptions.ErrorIfNoArrayExpand = true
	CodecHandle.EncodeOptions.Canonical = true
	CodecHandle.EncodeOptions.RecursiveEmptyCheck = true
	CodecHandle.WriteExt = true
	CodecHandle.PositiveIntUnsigned = true
	CodecHandle.EncodeOptions.Raw = true
}

type B struct {
	_struct struct{} `codec:",omitempty"`
	A1      test2.A  `codec:"a1"`
}

type C struct {
	_struct struct{} `codec:",omitempty"`
	A1      *test2.A `codec:"a1"`
}

// EncodeMsgPack returns a msgpack-encoded byte buffer for a given object.
func EncodeMsgPack(obj interface{}) []byte {
	var b []byte

	enc := codec.NewEncoderBytes(&b, CodecHandle)
	enc.MustEncode(obj)
	return b
}
