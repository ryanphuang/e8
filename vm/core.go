package vm

import (
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/vm/inst"
	"github.com/h8liu/e8/vm/mem"
)

type registers struct{ *Registers }
type memory struct{ *mem.Memory }

type Core struct {
	registers
	memory

	Stdout io.Writer
	Log    io.Writer

	sys *SysPage
	alu *inst.ALU
}

var _ inst.Core = new(Core)

func NewCore() *Core {
	ret := new(Core)
	ret.Registers = NewRegisters()
	ret.Memory = mem.New()
	ret.Stdout = os.Stdout

	ret.alu = inst.NewALU()

	return ret
}

func New() *Core {
	ret := NewCore()

	ret.sys = NewSysPage()
	ret.Memory.Map(0, ret.sys)

	return ret
}

func (self *Core) Step() {
	self.sys.ClearError()

	pc := self.IncPC()
	u32 := self.Memory.ReadU32(pc)
	in := inst.Inst(u32)
	if self.Log != nil {
		fmt.Fprintf(self.Log, "%08x: %08x   %v", pc, u32, in)
		if in.Op() != inst.OpJ {
			rs := in.Rs()
			rt := in.Rt()
			rsv := self.ReadReg(rs)
			rtv := self.ReadReg(rt)
			fmt.Fprintf(self.Log, "  ; $%d=%d(%08x) $%d=%d(%08x)",
				rs, rsv, rsv, rt, rtv, rtv)
		}
		fmt.Fprintf(self.Log, "\n")
		// self.Registers.PrintTo(self.Log)
	}
	self.alu.Inst(self, in)

	self.sys.FlushStdout(self.Stdout)
}

func (self *Core) Run(n int) int {
	i := 0
	for i < n {
		self.Step()
		i++

		if self.sys.Halted() {
			break
		}
	}

	return i
}

func (self *Core) SetPC(pc uint32) {
	self.Registers.WriteReg(inst.RegPC, pc)
}

func (self *Core) Halted() bool     { return self.sys.Halted() }
func (self *Core) AddrError() bool  { return self.sys.AddrError }
func (self *Core) HaltValue() uint8 { return self.sys.HaltValue }
func (self *Core) RIP() bool {
	return self.Halted() && self.HaltValue() == 0 && !self.AddrError()
}
