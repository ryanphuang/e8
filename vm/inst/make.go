package inst

func Rinst(s, t, d, funct uint8) uint32 {
	ret := uint32(s) << 21
	ret |= uint32(t) << 16
	ret |= uint32(d) << 11
	ret |= uint32(funct)
	return ret
}

func RinstShamt(s, t, d, shamt, funct uint8) uint32 {
	ret := Rinst(s, t, d, funct)
	ret |= uint32(shamt) << 6
	return ret
}

func Iinst(op, s, t uint8, im uint16) uint32 {
	ret := uint32(op) << 26
	ret |= uint32(s) << 21
	ret |= uint32(t) << 16
	ret |= uint32(im)
	return ret
}

func Jinst(ad int32) uint32 {
	ret := uint32(OpJ) << 26
	ret |= uint32(ad) & 0x3ffffff
	return ret
}
