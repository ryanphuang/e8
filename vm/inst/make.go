package inst

func Rinst(s, t, d, funct uint8) Inst {
	ret := uint32(s) << 21
	ret |= uint32(t) << 16
	ret |= uint32(d) << 11
	ret |= uint32(funct)
	return Inst(ret)
}

func RinstShamt(s, t, d, shamt, funct uint8) Inst {
	ret := uint32(Rinst(s, t, d, funct))
	ret |= uint32(shamt) << 6
	return Inst(ret)
}

func Iinst(op, s, t uint8, im uint16) Inst {
	ret := uint32(op) << 26
	ret |= uint32(s) << 21
	ret |= uint32(t) << 16
	ret |= uint32(im)
	return Inst(ret)
}

func Jinst(ad int32) Inst {
	ret := uint32(OpJ) << 26
	ret |= uint32(ad) & 0x3ffffff
	return Inst(ret)
}

func Noop() Inst {
	return Inst(0)
}
