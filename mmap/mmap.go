package mmap

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
				if !c.Valid(addr) {
					p = mem.NewPage()
					c.Map(addr, p)
				}
			}

			p.Write(addr&mem.PageMask, b)
		}
	}

	return nil
}
