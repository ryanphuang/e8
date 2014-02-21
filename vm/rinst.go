package vm

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
	map[uint8]Inst{
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
