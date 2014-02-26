package e8asm

import (
	"github.com/h8liu/e8/vm/inst"
	"strings"
)

type Line struct {
	in     inst.Inst
	label  string
	isJump bool
}

func newLine(in inst.Inst) *Line {
	ret := new(Line)
	ret.in = in
	return ret
}

func ParseLine(s string) *Line {
	s = strings.TrimSpace(s)

	// TODO:

	return newLine(inst.Noop())
}
