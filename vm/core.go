package vm

import (
	"os"
)

type Core struct {
	*Registers
	*Memory
	*SysPage

	fields *Fields
}

func NewCore() *Core {
	ret := new(Core)
	ret.Registers = NewRegisters()
	ret.Memory = NewMemory()
	ret.fields = new(Fields)

	return ret
}

func NewVm() *Core {
	ret := NewCore()

	ret.SysPage = NewSysPage()
	ret.Memory.Map(0, ret.SysPage)

	return ret
}

func (self *Core) Step() {
	self.SysPage.ClearError()

	pc := self.IncPC()
	self.fields.inst = self.Memory.ReadU32(pc)
	opInst(self, self.fields)

	self.SysPage.FlushStdout(os.Stdout)
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
