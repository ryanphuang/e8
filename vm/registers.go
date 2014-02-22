package vm

import (
	"fmt"
	"io"

	"github.com/h8liu/e8/vm/align"
	"github.com/h8liu/e8/vm/inst"
)

type Registers struct {
	regs  []uint32
	fregs []float64
}

func NewRegisters() *Registers {
	ret := new(Registers)

	ret.regs = make([]uint32, inst.Nreg)
	ret.fregs = make([]float64, inst.Nfreg)

	return ret
}

func (self *Registers) ReadReg(a uint8) uint32   { return self.regs[a] }
func (self *Registers) ReadFreg(a uint8) float64 { return self.fregs[a] }

func (self *Registers) WriteReg(a uint8, v uint32) {
	self.regs[a] = v

	if a == 0 {
		self.regs[0] = 0
	} else if a == inst.RegPC {
		self.regs[inst.RegPC] = align.U32(self.regs[inst.RegPC])
	}
}

func (self *Registers) WriteFreg(a uint8, v float64) {
	self.fregs[a] = v
}

func (self *Registers) IncPC() uint32 {
	ret := self.regs[inst.RegPC]
	self.regs[inst.RegPC] += 4
	return ret
}

func (self *Registers) PrintTo(w io.Writer) {
	for i := uint8(0); i < inst.Nreg; i++ {
		fmt.Fprintf(w, "$%02d:%08x", i, self.ReadReg(i))
		if (i+1)%4 == 0 {
			fmt.Fprintln(w)
		} else {
			fmt.Fprint(w, " ")
		}
	}
}
