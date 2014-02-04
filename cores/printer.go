package cores

import (
	"fmt"
	"github.com/h8liu/e8/arch"
)

type Printer struct{}

var _ arch.Core = new(Printer)

func (self *Printer) run(sys arch.Sys) {
	valid, c, v := sys.Select()
	if !valid {
		return
	}

	if c == arch.Vin {
		fmt.Println(v)
		sys.Open(arch.Vin, v)
	}
}

func (self *Printer) Run(sys arch.Sys, step int) int {
	if step == 0 {
		return 0
	}

	self.run(sys)
	return step - 1
}
