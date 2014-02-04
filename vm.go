package e8

type Vm struct {
	vsends map[uint32]*Vwriter
	vrecvs map[uint32]*Vreader
	// psends map[uint32]*Pwriter
	// precvs map[uint32]*Preader
}
