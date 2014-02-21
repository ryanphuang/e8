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

	// TODO: map core pages

	return ret
}

func (self *Core) Run(n uint32) uint32 {
	for n > 0 {
		pc := self.IncPC()
		self.fields.inst = self.Memory.ReadU32(pc)
		opInst(self, self.fields)
		n--

		// TODO: check pausing condition
	}

	return n
}
