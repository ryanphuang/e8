package inst

type codeMap struct {
	names map[uint8]string
	codes map[string]uint8
}

func reverseMap(m map[uint8]string) map[string]uint8 {
	ret := make(map[string]uint8)
	for k, v := range m {
		ret[v] = k
	}
	return ret
}

func newCodeMap(m map[uint8]string) *codeMap {
	ret := new(codeMap)
	ret.names = m
	ret.codes = reverseMap(m)
	return ret
}

func (self *codeMap) Code(s string) uint8 { return self.codes[s] }
func (self *codeMap) Name(c uint8) string { return self.names[c] }

var (
	functMap = newCodeMap(map[uint8]string{
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
	})

	opMap = newCodeMap(map[uint8]string{
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
		OpLbs:  "lbs",
		OpLbu:  "lbu",
		OpSw:   "sw",
		OpSh:   "sh",
		OpSb:   "sb",
		OpJ:    "j",
	})
)

func FunctName(f uint8) string { return functMap.Name(f) }
func OpName(op uint8) string   { return opMap.Name(op) }

func FunctCode(f string) uint8 { return functMap.Code(f) }
func OpCode(op string) uint8   { return opMap.Code(op) }
