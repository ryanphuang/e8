package img

import (
	"fmt"
	"io"
)

type Writer struct {
	io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w}
}

func (self *Writer) Write(addr uint32, bytes []byte) (e error) {
	return Write(self.Writer, addr, bytes)
}

func Write(out io.Writer, addr uint32, bytes []byte) (e error) {
	n := uint64(len(bytes))
	if n > (1 << 31) {
		// this is almost impossible to happen
		return fmt.Errorf("too many bytes")
	}
	if uint64(addr)+n > (1 << 32) {
		return fmt.Errorf("out of memory space")
	}

	header := new(Header)
	header.addr = addr
	header.size = uint32(n)
	if e = header.WriteTo(out); e != nil {
		return e
	}

	if _, e = out.Write(bytes); e != nil {
		return e
	}

	return
}
