package vm

import (
	"fmt"
	"io"
	"os"
)

type Core struct {
	*Registers
	*Memory
	*SysPage
	Stdout io.Writer
	Log    io.Writer

	alu *ALU
}

func NewCore() *Core {
	ret := new(Core)
	ret.Registers = NewRegisters()
	ret.Memory = NewMemory()
	ret.Stdout = os.Stdout

	ret.alu = NewALU()

	return ret
}

func NewVM() *Core {
	ret := NewCore()

	ret.SysPage = NewSysPage()
	ret.Memory.Map(0, ret.SysPage)

	return ret
}

func (self *Core) Step() {
	self.SysPage.ClearError()

	pc := self.IncPC()
	inst := self.Memory.ReadU32(pc)
	if self.Log != nil {
		fmt.Fprintf(self.Log, "%08x: %08x\n", pc, inst)
		self.Registers.PrintTo(self.Log)
	}
	self.alu.Inst(self, inst)

	self.SysPage.FlushStdout(self.Stdout)
}

func (self *Core) Run(n uint32) uint32 {
	for n > 0 {
		self.Step()
		n--

		if self.SysPage.Halted() {
			break
		}
	}

	return n
}

func (self *Core) SetPC(pc uint32) {
	self.Registers.WriteReg(RegPC, pc)
}
