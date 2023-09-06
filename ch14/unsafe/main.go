package main

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

func main() {

}

func dataFromBytesUnsafe(b [16]byte) data {
	d := *(*data)(unsafe.Pointer(&b))
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	return d
}

var isLE bool

func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = xb[0] == 0x00
}

func dataFromBytes(b [16]byte) data {
	d := data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

func bytesFromData(d data) [16]byte {
	out := [16]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func bytesFromDataUnsafe(d data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	b := *(*[16]byte)(unsafe.Pointer(&d))
	return b
}

type data struct {
	Value  uint32
	Label  [10]byte
	Active bool
}
