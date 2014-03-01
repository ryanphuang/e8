package line

import (
	"fmt"

	"github.com/h8liu/e8/asm/parse"
	"github.com/h8liu/e8/istr"
	"github.com/h8liu/e8/vm/inst"
)

type Line struct {
	in     inst.Inst
	label  string
	LineNo int
}

func NewLine(in inst.Inst) *Line {
	ret := new(Line)
	ret.in = in
	return ret
}

func (self *Line) Set(in inst.Inst) { self.in = in }
func (self *Line) J(addr int32)     { self.in = inst.Jinst(addr) }
func (self *Line) Ims(im int16) {
	in := self.in.U32() & 0xffff0000
	in |= uint32(uint16(im))
	self.in = inst.Inst(in)
}

func (self *Line) U32() uint32   { return self.in.U32() }
func (self *Line) Label() string { return self.label }
func (self *Line) Op() uint8     { return self.in.Op() }
func (self *Line) IsJump() bool  { return self.in.Op() == inst.OpJ }
func (self *Line) IsBranch() bool {
	op := self.in.Op()
	return op == inst.OpBne || op == inst.OpBeq
}

func (self *Line) String() string {
	if self.label != "" {
		op := self.in.Op()
		if op == inst.OpJ {
			return fmt.Sprintf("j %s", self.label)
		} else if op == inst.OpBne {
			return fmt.Sprintf("bne $%d, $%d, %s",
				self.in.Rs(), self.in.Rt(), self.label,
			)
		} else if op == inst.OpBeq {
			return fmt.Sprintf("beq $%d, $%d, %s",
				self.in.Rs(), self.in.Rt(), self.label,
			)
		}
	}

	return istr.String(self.in)
}

type parseFunc func(c, args string) (*Line, error)

var dispatch = func() map[string]parseFunc {
	ret := make(map[string]parseFunc)
	bind := func(f parseFunc, cs ...string) {
		for _, c := range cs {
			ret[c] = f
		}
	}

	bind(r3Line, "add", "sub", "and", "or", "xor", "nor", "slt")
	bind(r3Line, "mul", "mulu", "div", "divu", "mod", "modu")
	bind(r3rLine, "sllv", "srlv", "srav")
	bind(r3sLine, "sll", "srl", "sra")

	bind(i3aLine, "lw", "lhs", "lhu", "lbs", "lbu")
	bind(i3aLine, "sw", "sh", "sb")
	bind(i3sLine, "addi", "slti")
	bind(i3uLine, "andi", "ori")
	bind(i2Line, "lui")

	bind(bLine, "bne", "beq")

	bind(jLine, "j")

	return ret
}()

func ParseLine(s string) (*Line, error) {
	s = trim(s)
	op, args := opSplit(s)
	op = lower(op)

	f := dispatch[op]
	if f == nil {
		return nil, fmt.Errorf("invalid op")
	}

	return f(op, args)
}

func jLine(_, args string) (*Line, error) {
	if !parse.IsIdent(args) {
		return lef("invalid label")
	}

	ret := NewLine(inst.Jinst(0))
	ret.label = args

	return ret, nil
}

func bLine(c, args string) (*Line, error) {
	code := istr.OpCode(c)

	fs := fields(args)
	if len(fs) != 3 {
		return lef("invalid field count")
	}

	rs, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}
	rt, valid := parseReg(fs[1])
	if !valid {
		return lef("second field not register")
	}

	label := fs[2]
	if !parse.IsIdent(label) {
		return lef("third field is not a label")
	}

	ret := NewLine(inst.Iinst(code, rs, rt, 0))
	ret.label = label

	return ret, nil
}

func i3sLine(c, args string) (*Line, error) {
	code := istr.OpCode(c)

	fs := fields(args)
	if len(fs) != 3 {
		return lef("invalid field count")
	}

	rt, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}
	rs, valid := parseReg(fs[1])
	if !valid {
		return lef("second field not register")
	}

	im, valid := parseIms(fs[2])
	if !valid {
		return lef("third field not a signed immediate")
	}

	ret := NewLine(inst.Iinst(code, rs, rt, im))
	return ret, nil
}

func i3uLine(c, args string) (*Line, error) {
	code := istr.OpCode(c)

	fs := fields(args)
	if len(fs) != 3 {
		return lef("invalid field count")
	}

	rt, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}
	rs, valid := parseReg(fs[1])
	if !valid {
		return lef("second field not register")
	}

	im, valid := parseImu(fs[2])
	if !valid {
		return lef("third field not an unsigned immediate")
	}

	ret := NewLine(inst.Iinst(code, rs, rt, im))
	return ret, nil
}

func i3aLine(c, args string) (*Line, error) {
	code := istr.OpCode(c)

	fs := fields(args)
	if len(fs) != 2 {
		return lef("invalid field count")
	}

	rt, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}

	im, rs, valid := parseAddr(fs[1])
	if !valid {
		return lef("second field not an address")
	}

	ret := NewLine(inst.Iinst(code, rs, rt, im))
	return ret, nil
}

func i2Line(c, args string) (*Line, error) {
	code := istr.OpCode(c)

	fs := fields(args)
	if len(fs) != 2 {
		return lef("invalid field count")
	}

	rt, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}

	im, valid := parseIms(fs[2])
	if !valid {
		return lef("second field not a signed immediate")
	}

	ret := NewLine(inst.Iinst(code, 0, rt, im))
	return ret, nil
}

func r3Line(c, args string) (*Line, error) {
	code := istr.FunctCode(c)

	fs := fields(args)
	if len(fs) != 3 {
		return lef("invalid field count")
	}

	rd, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}
	rs, valid := parseReg(fs[1])
	if !valid {
		return lef("second field not register")
	}
	rt, valid := parseReg(fs[2])
	if !valid {
		return lef("third field not register")
	}

	ret := NewLine(inst.Rinst(rs, rt, rd, code))
	return ret, nil
}

func r3rLine(c, args string) (*Line, error) {
	code := istr.FunctCode(c)

	fs := fields(args)
	if len(fs) != 3 {
		return lef("invalid field count")
	}

	rd, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}
	rt, valid := parseReg(fs[1])
	if !valid {
		return lef("second field not register")
	}
	rs, valid := parseReg(fs[2])
	if !valid {
		return lef("third field not register")
	}

	ret := NewLine(inst.Rinst(rs, rt, rd, code))
	return ret, nil
}

func r3sLine(c, args string) (*Line, error) {
	code := istr.FunctCode(c)

	fs := fields(args)
	if len(fs) != 3 {
		return lef("invalid field count")
	}

	rd, valid := parseReg(fs[0])
	if !valid {
		return lef("first field not register")
	}
	rt, valid := parseReg(fs[1])
	if !valid {
		return lef("second field not register")
	}
	shamt, valid := parseShamt(fs[2])
	if !valid {
		return lef("third field not shamt")
	}

	ret := NewLine(inst.RinstShamt(0, rt, rd, shamt, code))
	return ret, nil
}
