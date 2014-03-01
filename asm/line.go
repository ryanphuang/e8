package asm

import (
	"fmt"

	"github.com/h8liu/e8/istr"
	"github.com/h8liu/e8/vm/inst"
)

type Line struct {
	in     inst.Inst
	label  string
	lineno int
}

func newLine(in inst.Inst) *Line {
	ret := new(Line)
	ret.in = in
	return ret
}

func (self *Line) Op() uint8    { return self.in.Op() }
func (self *Line) IsJump() bool { return self.in.Op() == inst.OpJ }
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

var (
	r3Insts = map[string]uint8{
		"add": inst.FnAdd,
		"sub": inst.FnSub,
		"and": inst.FnAnd,
		"or":  inst.FnOr,
		"xor": inst.FnXor,
		"nor": inst.FnXor,
		"slt": inst.FnSlt,

		"mul":  inst.FnMul,
		"mulu": inst.FnMulu,
		"div":  inst.FnDiv,
		"divu": inst.FnDivu,
		"mod":  inst.FnMod,
		"modu": inst.FnModu,
	}

	r3rInsts = map[string]uint8{
		"sllv": inst.FnSllv,
		"srlv": inst.FnSrlv,
		"srav": inst.FnSrav,
	}

	r3sInsts = map[string]uint8{
		"sll": inst.FnSll,
		"srl": inst.FnSrl,
		"sra": inst.FnSra,
	}

	i3aInsts = map[string]uint8{
		"lw":  inst.OpLw,
		"lhs": inst.OpLhs,
		"lhu": inst.OpLhu,
		"lbs": inst.OpLbs,
		"lbu": inst.OpLbu,
		"sw":  inst.OpSw,
		"sh":  inst.OpSh,
		"sb":  inst.OpSb,
	}

	i3sInsts = map[string]uint8{
		"addi": inst.OpAddi,
		"slti": inst.OpSlti,
	}

	i3uInsts = map[string]uint8{
		"andi": inst.OpAndi,
		"ori":  inst.OpOri,
	}

	i2Insts = map[string]uint8{
		"lui": inst.OpLui,
	}

	bInsts = map[string]uint8{
		"bne": inst.OpBne,
		"beq": inst.OpBeq,
	}

	parseDispatch = []*struct {
		ops map[string]uint8
		fn  func(uint8, string) (*Line, error)
	}{
		{bInsts, bLine},
		{i3sInsts, i3sLine},
		{i3uInsts, i3uLine},
		{i3aInsts, i3aLine},
		{i2Insts, i2Line},
		{r3Insts, r3Line},
		{r3rInsts, r3rLine},
		{r3sInsts, r3sLine},
	}
)

func ParseLine(s string) (*Line, error) {
	s = trim(s)
	op, args := opSplit(s)
	op = lower(op)

	if op == "j" {
		return jLine(inst.OpJ, args)
	}

	for _, p := range parseDispatch {
		if code, found := p.ops[op]; found {
			return p.fn(code, args)
		}
	}

	return nil, fmt.Errorf("invalid op")
}

func jLine(code uint8, args string) (*Line, error) {
	if !isIdent(args) {
		return lef("invalid label")
	}

	ret := newLine(inst.Jinst(0))
	ret.label = args

	return ret, nil
}

func bLine(code uint8, args string) (*Line, error) {
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
	if !isIdent(label) {
		return lef("third field is not a label")
	}

	ret := newLine(inst.Iinst(code, rs, rt, 0))
	ret.label = label

	return ret, nil
}

func i3sLine(code uint8, args string) (*Line, error) {
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

	ret := newLine(inst.Iinst(code, rs, rt, im))
	return ret, nil
}

func i3uLine(code uint8, args string) (*Line, error) {
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

	ret := newLine(inst.Iinst(code, rs, rt, im))
	return ret, nil
}

func i3aLine(code uint8, args string) (*Line, error) {
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

	ret := newLine(inst.Iinst(code, rs, rt, im))
	return ret, nil
}

func i2Line(code uint8, args string) (*Line, error) {
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

	ret := newLine(inst.Iinst(code, 0, rt, im))
	return ret, nil
}

func r3Line(code uint8, args string) (*Line, error) {
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

	ret := newLine(inst.Rinst(rs, rt, rd, code))
	return ret, nil
}

func r3rLine(code uint8, args string) (*Line, error) {
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

	ret := newLine(inst.Rinst(rs, rt, rd, code))
	return ret, nil
}

func r3sLine(code uint8, args string) (*Line, error) {
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

	ret := newLine(inst.RinstShamt(0, rt, rd, shamt, code))
	return ret, nil
}
