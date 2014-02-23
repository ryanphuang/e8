package img

import (
	"io"
)

type Header struct {
	addr uint32
	size uint32
}

func u32(buf []byte) uint32 {
	ret := uint32(buf[0])
	ret |= uint32(buf[1]) << 8
	ret |= uint32(buf[2]) << 16
	ret |= uint32(buf[3]) << 24
	return ret
}

func pu32(buf []byte, i uint32) {
	buf[0] = uint8(i)
	buf[1] = uint8(i >> 8)
	buf[2] = uint8(i >> 16)
	buf[3] = uint8(i >> 24)
}

func (self *Header) ReadIn(in io.Reader) error {
	buf := make([]byte, 8)
	_, e := io.ReadFull(in, buf)
	if e != nil {
		return e
	}

	self.addr = u32(buf[0:4])
	self.size = u32(buf[4:8])
	return nil
}

func (self *Header) WriteTo(out io.Writer) error {
	buf := make([]byte, 8)
	pu32(buf[0:4], self.addr)
	pu32(buf[4:8], self.size)

	_, e := out.Write(buf)
	return e
}
