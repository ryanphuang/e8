package line

import (
	"fmt"
	"strings"

	"github.com/h8liu/e8/asm/args"
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

// To setup label fields
func (self *Line) J(addr int32) { self.in = inst.Jinst(addr) }
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

	return self.in.String() // TODO: use your own format
}

var dispatch = func() map[string]string {
	ret := make(map[string]string)
	o := func(f string, cs ...string) {
		f = args.Compile(f)
		for _, c := range cs {
			ret[c] = f
		}
	}

	o("rd, rs, rt", "add", "sub", "and", "or", "xor", "nor", "slt")
	o("rd, rs, rt", "mul", "mulu", "div", "divu", "mod", "modu")
	o("rd, rt, rs", "sllv", "srlv", "srav")
	o("rd, rt, shamt", "sll", "srl", "sra")
	o("rt, addr", "lw", "lhs", "lhu", "lbs", "lbu")
	o("rt, addr", "sw", "shu", "sb")
	o("rt, rs, ims", "addi", "slti")
	o("rt, rs, imu", "andi", "ori")
	o("rt, imu", "lui")
	o("rs, rt, label", "bne", "beq")
	o("label", "j")

	return ret
}()

func Parse(s string) (*Line, error) {
	s = strings.TrimSpace(s)
	op, a := opSplit(s)
	op = strings.ToLower(op)

	f, found := dispatch[op]
	if !found {
		return nil, fmt.Errorf("invalid op")
	}

	base := uint32(inst.OpCode(op)) << inst.OpShift
	base |= uint32(inst.FunctCode(op)) & inst.FunctMask

	in, lab, e := args.Parse(f, a, base)
	if e != nil {
		return nil, e
	}

	return &Line{inst.Inst(in), lab, 0}, nil
}
