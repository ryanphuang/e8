package cores

import (
	"github.com/h8liu/e8/arch"
)

type PingPong struct{}

var _ arch.Core = new(PingPong)

func (self *PingPong) run(sys arch.Sys) {
	valid, c, v := sys.Select()
	if !valid {
		if !sys.IsOpen(arch.Vout) {
			sys.Open(arch.Vin, 0)
		}
		return
	}

	// handle some transitioning
	switch c {
	case arch.Vin: // received some input
		sys.Open(arch.Vout, v) // write the value now
	case arch.Vout: // output just sent out
		sys.Open(arch.Vin, 0) // open for read
	}
}

func (self *PingPong) Run(sys arch.Sys, step int) int {
	if step == 0 {
		return 0
	}

	self.run(sys)
	return step - 1
}
