package inst

func opAdd(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s+t)
}

func opSub(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s-t)
}

func opMul(c Core, f *fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	c.WriteReg(f.rd, uint32(s*t))
}

func opMulu(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s*t)
}

func opDiv(c Core, f *fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, uint32(s/t))
	}
}

func opMod(c Core, f *fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, uint32(s%t))
	}
}

func opDivu(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, s/t)
	}
}

func opModu(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, s/t)
	}
}

func opAnd(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s&t)
}

func opOr(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s|t)
}

func opXor(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s^t)
}

func opNor(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, ^(s | t))
}

func opSlt(c Core, f *fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	if s < t {
		c.WriteReg(f.rd, 1)
	} else {
		c.WriteReg(f.rd, 0)
	}
}

func opSll(c Core, f *fields) {
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t<<f.shamt)
}

func opSrl(c Core, f *fields) {
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t>>f.shamt)
}

func opSra(c Core, f *fields) {
	t := int32(c.ReadReg(f.rt))
	c.WriteReg(f.rd, uint32(t>>f.shamt))
}

func opSllv(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t<<s)
}

func opSrlv(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t>>s)
}

func opSrav(c Core, f *fields) {
	s := c.ReadReg(f.rs)
	t := int32(c.ReadReg(f.rt))
	c.WriteReg(f.rd, uint32(t>>s))
}
