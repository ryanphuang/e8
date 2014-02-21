package vm

var rInstList = makeInstList(
	map[uint8]Inst{
		0x20: opAdd,
		0x22: opSub,
		0x24: opAnd,
		0x25: opOr,
		0x26: opXor,
		0x27: opNor,
		0x2A: opSlt,

		0x18: opMul,
		0x19: opMulu,
		0x1A: opDiv,
		0x1B: opDivu,
		0x1C: opMod,
		0x1D: opModu,

		0x00: opSll,
		0x02: opSrl,
		0x03: opSra,
		0x04: opSllv,
		0x06: opSrlv,
		0x07: opSrav,
	}, Nfunct,
)

func opNoopr(c *Core, f *Fields) {}

func opAdd(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s+t)
}

func opSub(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s-t)
}

func opMul(c *Core, f *Fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	c.WriteReg(f.rd, uint32(s*t))
}

func opMulu(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s*t)
}

func opDiv(c *Core, f *Fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, uint32(s/t))
	}
}

func opMod(c *Core, f *Fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, uint32(s%t))
	}
}

func opDivu(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, s/t)
	}
}

func opModu(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	if t == 0 {
		c.WriteReg(f.rd, 0)
	} else {
		c.WriteReg(f.rd, s/t)
	}
}

func opAnd(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s&t)
}

func opOr(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s|t)
}

func opXor(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, s^t)
}

func opNor(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, ^(s | t))
}

func opSlt(c *Core, f *Fields) {
	s := int32(c.ReadReg(f.rs))
	t := int32(c.ReadReg(f.rt))
	if s < t {
		c.WriteReg(f.rd, 1)
	} else {
		c.WriteReg(f.rd, 0)
	}
}

func opSll(c *Core, f *Fields) {
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t<<f.shamt)
}

func opSrl(c *Core, f *Fields) {
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t>>f.shamt)
}

func opSra(c *Core, f *Fields) {
	t := int32(c.ReadReg(f.rt))
	c.WriteReg(f.rd, uint32(t>>f.shamt))
}

func opSllv(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t<<s)
}

func opSrlv(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := c.ReadReg(f.rt)
	c.WriteReg(f.rd, t>>s)
}

func opSrav(c *Core, f *Fields) {
	s := c.ReadReg(f.rs)
	t := int32(c.ReadReg(f.rt))
	c.WriteReg(f.rd, uint32(t>>s))
}
