package istr

import (
	. "github.com/h8liu/e8/vm/inst"
)

var functNames = map[uint8]string{
	FnAdd:  "add",
	FnSub:  "sub",
	FnAnd:  "and",
	FnOr:   "or",
	FnXor:  "xor",
	FnNor:  "nor",
	FnSlt:  "slt",
	FnMul:  "mul",
	FnMulu: "mulu",
	FnDiv:  "div",
	FnDivu: "divu",
	FnMod:  "mod",
	FnModu: "modu",
	FnSll:  "sll",
	FnSrl:  "srl",
	FnSra:  "sra",
	FnSllv: "sllv",
	FnSrav: "srav",
}

func FunctName(f uint8) string { return functNames[f] }

var opNames = map[uint8]string{
	OpBeq:  "beq",
	OpBne:  "bne",
	OpAddi: "addi",
	OpSlti: "slti",
	OpAndi: "andi",
	OpOri:  "ori",
	OpLui:  "lui",
	OpLw:   "lw",
	OpLhs:  "lhs",
	OpLhu:  "lhu",
	OpSw:   "sw",
	OpSh:   "sh",
	OpSb:   "sb",
}

func OpName(op uint8) string { return opNames[op] }
