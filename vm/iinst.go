package vm

func memAddr(c *Core, f *fields) uint32 {
	s := c.ReadReg(f.rs)
	return s + signExt(f.im)
}

func signExt(im uint16) uint32 {
	return uint32(int32(int16(im)))
}

func unsignExt(im uint16) uint32 {
	return uint32(im)
}

func signExt8(a uint8) uint32 {
	return uint32(int32(int8(a)))
}

func unsignExt8(a uint8) uint32 {
	return uint32(a)
}

func opAddi(c *Core, f *fields) {
	s := c.ReadReg(f.rs)
	c.WriteReg(f.rt, s+signExt(f.im))
}

func opLw(c *Core, f *fields) {
	addr := memAddr(c, f)
	c.WriteReg(f.rt, c.ReadU32(addr))
}

func opLhs(c *Core, f *fields) {
	addr := memAddr(c, f)
	c.WriteReg(f.rt, signExt(c.ReadU16(addr)))
}

func opLhu(c *Core, f *fields) {
	addr := memAddr(c, f)
	c.WriteReg(f.rt, unsignExt(c.ReadU16(addr)))
}

func opLbs(c *Core, f *fields) {
	addr := memAddr(c, f)
	c.WriteReg(f.rt, signExt8(c.ReadU8(addr)))
}

func opLbu(c *Core, f *fields) {
	addr := memAddr(c, f)
	c.WriteReg(f.rt, unsignExt8(c.ReadU8(addr)))
}

func opSw(c *Core, f *fields) {
	t := c.ReadReg(f.rt)
	addr := memAddr(c, f)
	c.WriteU32(addr, t)
}

func opSh(c *Core, f *fields) {
	t := uint16(c.ReadReg(f.rt))
	addr := memAddr(c, f)
	c.WriteU16(addr, t)
}

func opSb(c *Core, f *fields) {
	t := uint8(c.ReadReg(f.rt))
	addr := memAddr(c, f)
	c.WriteU8(addr, t)
}

func opLui(c *Core, f *fields) {
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rt, (t&0xffff)|(unsignExt(f.im)<<16))
}

func opAndi(c *Core, f *fields) {
	s := c.ReadReg(f.rs)
	c.WriteReg(f.rt, s&unsignExt(f.im))
}

func opOri(c *Core, f *fields) {
	s := c.ReadReg(f.rs)
	c.WriteReg(f.rt, s|unsignExt(f.im))
}

func opSlti(c *Core, f *fields) {
	s := int32(c.ReadReg(f.rs))
	if s < int32(signExt(f.im)) {
		c.WriteReg(f.rt, 1)
	} else {
		c.WriteReg(f.rt, 0)
	}
}

func opBeq(c *Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	if s == t {
		pc := c.ReadReg(RegPC)
		c.WriteReg(RegPC, pc+(signExt(f.im)<<2))
	}
}

func opBne(c *Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	if s != t {
		pc := c.ReadReg(RegPC)
		c.WriteReg(RegPC, pc+(signExt(f.im)<<2))
	}
}
