package inst

type Inst uint32

func (i Inst) U32() uint32 { return uint32(i) }

type instFunc func(c Core, fields *fields)

func makeInstList(m map[uint8]instFunc, n uint8) []instFunc {
	ret := make([]instFunc, n)
	for i := range ret {
		ret[i] = opNoop
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
	map[uint8]instFunc{
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

const (
	FnAdd = 0x20
	FnSub = 0x22
	FnAnd = 0x24
	FnOr  = 0x25
	FnXor = 0x26
	FnNor = 0x27
	FnSlt = 0x2A

	FnMul  = 0x18
	FnMulu = 0x19
	FnDiv  = 0x1A
	FnDivu = 0x1B
	FnMod  = 0x1C
	FnModu = 0x1D

	FnSll  = 0x00
	FnSrl  = 0x02
	FnSra  = 0x03
	FnSllv = 0x04
	FnSrlv = 0x06
	FnSrav = 0x07
)

var rInstList = makeInstList(
	map[uint8]instFunc{
		FnAdd: opAdd,
		FnSub: opSub,
		FnAnd: opAnd,
		FnOr:  opOr,
		FnXor: opXor,
		FnNor: opNor,
		FnSlt: opSlt,

		FnMul:  opMul,
		FnMulu: opMulu,
		FnDiv:  opDiv,
		FnDivu: opDivu,
		FnMod:  opMod,
		FnModu: opModu,

		FnSll:  opSll,
		FnSrl:  opSrl,
		FnSra:  opSra,
		FnSllv: opSllv,
		FnSrlv: opSrlv,
		FnSrav: opSrav,
	}, Nfunct,
)

func opInst(c Core, f *fields) {
	op := uint8(f.inst >> 26)
	f.rs = uint8(f.inst>>21) & 0x1f
	f.rt = uint8(f.inst>>16) & 0x1f
	f.im = uint16(f.inst)

	instList[op](c, f)
}

func opRinst(c Core, f *fields) {
	f.rd = uint8(f.inst>>11) & 0x1f
	f.shamt = uint8(f.inst>>6) & 0x1f
	funct := uint8(f.inst) & 0x3f

	rInstList[funct](c, f)
}

func opJ(c Core, f *fields) {
	pc := c.ReadReg(RegPC)
	c.WriteReg(RegPC, pc+uint32(int32(f.inst<<6)>>4))
}

func opNoop(c Core, f *fields) {}
