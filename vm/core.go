package vm

type Core struct {
	*Registers
	*Memory
	*State

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

	// TODO: map vm system pages

	return ret
}

func (self *Core) Step() {
	pc := self.IncPC()
	self.fields.inst = self.Memory.ReadU32(pc)
	opInst(self, self.fields)
}

func (self *Core) Run(n uint32) uint32 {
	for n > 0 {
		self.Step()
		n--

		// TODO: check pausing condition
	}

	return n
}

func (self *Core) SetPC(pc uint32) {
	self.Registers.WriteReg(RegPC, pc)
}
