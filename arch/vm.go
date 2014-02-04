package arch

type Vm struct {
	Daemon bool

	vsends map[uint32]*Vwriter
	vrecvs map[uint32]*Vreader
	// psends map[uint32]*Pwriter
	// precvs map[uint32]*Preader
}

func (self *Vm) Run() {
	panic("todo")
}
