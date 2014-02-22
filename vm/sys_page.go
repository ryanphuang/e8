package vm

import (
	"io"
)

type SysPage struct {
	AddrError bool
	Halt      bool
	HaltValue uint8

	stdin  chan byte
	stdout chan byte
}

var _ Page = new(SysPage)

func NewSysPage() *SysPage {
	ret := new(SysPage)
	ret.stdin = make(chan byte, 32)
	ret.stdout = make(chan byte, 32)

	return ret
}

func (self *SysPage) Halted() bool {
	if self == nil {
		return false
	}

	return self.Halt
}

func (self *SysPage) ClearError() {
	if self == nil {
		return
	}
	self.AddrError = false
	self.Halt = false
}

func (self *SysPage) addrError() {
	self.AddrError = true
	self.Halt = true
	self.HaltValue = 0xff
}

func (self *SysPage) Read(offset uint32) uint8 {
	if offset < 4 {
		self.addrError()
		return 0
	}

	switch offset {
	case 5: // stdout ready
		if len(self.stdout) < cap(self.stdout) {
			return 0 // ready
		}
		return 1 // busy
	case 6: // stdin ready
		if len(self.stdin) > 0 {
			return 0
		}
		return 1 // invalid
	case 7: // stdin value
		if len(self.stdin) > 0 {
			return <-self.stdin
		}
		return 0
	default:
		return 0
	}
}

func (self *SysPage) Write(offset uint32, b uint8) {
	if offset < 4 {
		self.addrError()
		return
	}

	switch offset {
	case 4: // halt
		self.Halt = true
		self.HaltValue = b
	case 5: // stdout
		if len(self.stdout) < cap(self.stdout) {
			self.stdout <- b
		}
	}
}

func (self *SysPage) FlushStdout(w io.Writer) {
	if self == nil {
		return
	}

	for len(self.stdout) > 0 {
		b := <-self.stdout
		w.Write([]byte{b})
	}
}
