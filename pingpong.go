package e8

type PingPong struct{}

func (self *PingPong) run(sys Sys) {
	valid, c, v := sys.Select()
	if !valid {
		if !sys.IsOpen(Pout) {
			sys.Open(Pin, 0)
		}
		return
	}

	switch c {
	case Pin: // has some input
		sys.Open(Pout, v) // write the value now
	case Pout: // output just got out
		sys.Open(Pin, 0) // read again
	}
}

func (self *PingPong) Run(sys Sys, step int) int {
	if step == 0 {
		return 0
	}

	self.run(sys)
	return step - 1
}
