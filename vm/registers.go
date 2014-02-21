package vm

type Registers struct {
	regs  []uint32
	fregs []float64
}

const (
	Nreg  = 32
	Nfreg = Nreg
	RegPC = Nreg - 1
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

	if a == 0 {
		self.regs[0] = 0
	} else if a == RegPC {
		self.regs[RegPC] = alignU32(self.regs[Nreg-1])
	}
}

func (self *Registers) WriteFreg(a uint8, v float64) {
	self.fregs[a] = v
}

func (self *Registers) IncPC() {
	self.regs[RegPC] += 4
}
