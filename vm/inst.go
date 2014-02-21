package vm

type Inst func(c *Core, fields *Fields)

type Fields struct {
	inst  uint32
	rs    uint8
	rt    uint8
	rd    uint8
	shamt uint8
	im    uint16
}

const (
	Nfunct = 64
	Nop    = 64
)

func opNoop(c *Core, f *Fields) {}

func makeInstList(m map[uint8]Inst, n uint8) []Inst {
	ret := make([]Inst, n)
	for i := range ret {
		ret[i] = opNoopr
	}
	for i, inst := range m {
		ret[i] = inst
	}
	return ret
}

const (
	OpRinst = 0
	OpJ     = 0x02
	OpBeq   = 0x04
	OpBne   = 0x05

	OpAddi = 0x08
	OpSlti = 0x0A
	OpAndi = 0x0C
	OpOri  = 0x0D
	OpLui  = 0x0F

	OpLw  = 0x23
	OpLhs = 0x21
	OpLhu = 0x25
	OpLbs = 0x20
	OpLbu = 0x24
	OpSw  = 0x2B
	OpSh  = 0x29
	OpSb  = 0x28
)

var instList = makeInstList(
	map[uint8]Inst{
		OpRinst: opRinst,
		OpJ:     opJ,
		OpBeq:   opBeq,
		OpBne:   opBne,

		OpAddi: opAddi,
		OpSlti: opSlti,
		OpAndi: opAndi,
		OpOri:  opOri,
		OpLui:  opLui,

		OpLw:  opLw,
		OpLhs: opLhs,
		OpLhu: opLhu,
		OpLbs: opLbs,
		OpLbu: opLbu,
		OpSw:  opSw,
		OpSh:  opSh,
		OpSb:  opSb,
	}, Nop,
)

func opInst(c *Core, f *Fields) {
	op := uint8(f.inst >> 26)
	f.rs = uint8(f.inst>>21) & 0x1f
	f.rt = uint8(f.inst>>16) & 0x1f
	f.im = uint16(f.inst)

	instList[op](c, f)
}

func opRinst(c *Core, f *Fields) {
	f.rd = uint8(f.inst>>11) & 0x1f
	f.shamt = uint8(f.inst>>6) & 0x1f
	funct := uint8(f.inst) & 0x3f

	rInstList[funct](c, f)
}

func opJ(c *Core, f *Fields) {
	pc := c.ReadReg(RegPC)
	c.WriteReg(RegPC, pc+uint32(int32(f.inst<<6)>>4))
}
