package bin

import (
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/vm"
	"github.com/h8liu/e8/vm/mem"
)

func Make(in io.Reader) (*vm.Core, error) {
	c := vm.New()
	e := LoadInto(c, in)
	if e != nil {
		return nil, e
	}
	return c, nil
}

func Open(path string) (*vm.Core, error) {
	fin, e := os.Open(path)
	if e != nil {
		return nil, e
	}

	defer fin.Close()

	return Make(fin)
}

func LoadInto(c *vm.Core, in io.Reader) error {
	header := new(Header)
	var p mem.Page
	cur := uint32(0)

	for {
		e := header.ReadIn(in)
		if e == io.EOF {
			return nil
		}
		if e != nil {
			return e
		}

		buf := make([]byte, header.size)
		_, e = io.ReadFull(in, buf)
		if e != nil {
			return e
		}

		for i, b := range buf {
			addr := header.addr + uint32(i)
			id := mem.PageId(addr)
			if id == 0 {
				return fmt.Errorf("attempt to map system page")
			}

			if cur == 0 || cur != id {
				cur = id
				p = c.Get(cur)
				if p == nil {
					p = mem.NewPage()
					c.Map(cur, p)
				}
			}

			p.Write(addr&mem.PageMask, b)
		}
	}

	return nil
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
