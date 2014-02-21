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

var instList = makeInstList(
	map[uint8]Inst{
		0x00: opRinst,
		0x02: opJ,
		0x04: opBeq,
		0x05: opBne,

		0x08: opAddi,
		0x0A: opSlti,
		0x0C: opAndi,
		0x0D: opOri,
		0x0F: opLui,

		0x23: opLw,
		0x21: opLhs,
		0x25: opLhu,
		0x20: opLbs,
		0x24: opLbu,
		0x2B: opSw,
		0x29: opSh,
		0x28: opSb,
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
