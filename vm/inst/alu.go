package inst

type ALU struct {
	fields *fields
}

func NewALU() *ALU {
	ret := new(ALU)
	ret.fields = new(fields)
	return ret
}

func (self *ALU) Inst(c Core, inst Inst) {
	self.fields.inst = inst
	opInst(c, self.fields)
}
