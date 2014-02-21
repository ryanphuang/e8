package vm

type Registers struct {
	pc    uint32
	regs  []uint32
	fregs []float64
}

const (
	Nreg  = 8
	Nfreg = Nreg
)

func NewRegisters() *Registers {
	ret := new(Registers)

	ret.regs = make([]uint32, Nreg)
	ret.fregs = make([]float64, Nfreg)

	return ret
}

func (self *Registers) ReadReg(a uint8) uint32   { return self.regs[a] }
func (self *Registers) ReadFreg(a uint8) float64 { return self.fregs[a] }

func (self *Registers) WriteReg(a uint8, v uint32) {
	self.regs[a] = v
	self.regs[0] = 0
}

func (self *Registers) WriteFreg(a uint8, v float64) {
	self.fregs[a] = v
}

func (self *Registers) PC() uint32 { return self.pc }
func (self *Registers) SetPC(pc uint32) { self.pc = alignU32(pc) }
