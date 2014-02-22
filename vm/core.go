package vm

import (
	"fmt"
	"io"
	"os"

	"github.com/h8liu/e8/vm/inst"
	"github.com/h8liu/e8/vm/mem"
)

type Core struct {
	*Registers
	*mem.Memory
	*SysPage
	Stdout io.Writer
	Log    io.Writer

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

	ret.SysPage = NewSysPage()
	ret.Memory.Map(0, ret.SysPage)

	return ret
}

func (self *Core) Step() {
	self.SysPage.ClearError()

	pc := self.IncPC()
	in := self.Memory.ReadU32(pc)
	if self.Log != nil {
		fmt.Fprintf(self.Log, "%08x: %08x\n", pc, in)
		self.Registers.PrintTo(self.Log)
	}
	self.alu.Inst(self, inst.Inst(in))

	self.SysPage.FlushStdout(self.Stdout)
}

func (self *Core) Run(n uint32) uint32 {
	i := uint32(0)
	for i < n {
		self.Step()
		i++

		if self.SysPage.Halted() {
			break
		}
	}

	return i
}

func (self *Core) SetPC(pc uint32) {
	self.Registers.WriteReg(inst.RegPC, pc)
}
