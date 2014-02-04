package arch

type Vmm struct {
	vms []*Vm
}

func NewVmm() *Vmm {
	panic("todo")
}

func (self *Vmm) AddVm(core Core) *Vm {
	panic("todo")
}

// run until cycles runs out for all non-daemon vms
func (self *Vmm) Run() {
	for {
		for _, vm := range self.vms {
			vm.Run()
			// TODO: check stop condition
		}

		// handle messages
		panic("todo")
	}
}
