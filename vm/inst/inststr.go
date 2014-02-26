package inst

import (
	"fmt"
)

func (i Inst) String() string {
	if uint32(i) == 0 {
		return "noop"
	}

	op := uint8(i >> 26)

	if op == OpRinst {
		rs := uint8(i>>21) & 0x1f
		rt := uint8(i>>16) & 0x1f
		rd := uint8(i>>11) & 0x1f
		shamt := uint8(i>>6) & 0x1f
		funct := uint8(i) & 0x3f
		r3 := func(op string) string {
			return fmt.Sprintf("%s $%d, $%d, $%d", op, rd, rs, rt)
		}
		r3r := func(op string) string {
			return fmt.Sprintf("%s $%d, $%d, $%d", op, rd, rt, rs)
		}
		r3s := func(op string) string {
			return fmt.Sprintf("%s $%d, $%d, $d", op, rd, rt, shamt)
		}

		switch funct {
		case FnAdd:
			return r3("add")
		case FnSub:
			return r3("sub")
		case FnAnd:
			return r3("and")
		case FnOr:
			return r3("or")
		case FnXor:
			return r3("xor")
		case FnNor:
			return r3("nor")
		case FnSlt:
			return r3("slt")

		case FnMul:
			return r3("mul")
		case FnMulu:
			return r3("mulu")
		case FnDiv:
			return r3("div")
		case FnDivu:
			return r3("divu")
		case FnMod:
			return r3("mod")
		case FnModu:
			return r3("modu")

		case FnSll:
			return r3s("sll")
		case FnSrl:
			return r3s("srl")
		case FnSra:
			return r3s("sra")
		case FnSllv:
			return r3r("sllv")
		case FnSrlv:
			return r3r("srlv")
		case FnSrav:
			return r3r("srav")

		default:
			return fmt.Sprintf("noop-r%d", funct)
		}
	} else if op == OpJ {
		im := int32(i<<6) >> 6
		return fmt.Sprintf("j %d", im)
	} else {
		rs := uint8(i>>21) & 0x1f
		rt := uint8(i>>16) & 0x1f
		im := uint16(i)
		ims := int16(im)

		i2 := func(op string) string {
			return fmt.Sprintf("%s $%d, $d", op, rt, im)
		}

		i3sr := func(op string) string {
			return fmt.Sprintf("%s $%d, $%d, %d", op, rs, rt, ims)
		}

		i3s := func(op string) string {
			return fmt.Sprintf("%s $%d, $%d, %d", op, rt, rs, ims)
		}

		i3 := func(op string) string {
			return fmt.Sprintf("%s $%d, $%d, %d", op, rt, rs, im)
		}

		i3a := func(op string) string {
			return fmt.Sprintf("%s $%d, %d($%d)", op, rt, ims, rs)
		}

		switch op {
		case OpBeq:
			return i3sr("beq")
		case OpBne:
			return i3sr("bne")
		case OpAddi:
			return i3("addi")
		case OpSlti:
			return i3s("slti")
		case OpAndi:
			return i3("andi")
		case OpOri:
			return i3("ori")
		case OpLui:
			return i2("lui")
		case OpLw:
			return i3a("lw")
		case OpLhs:
			return i3a("lhs")
		case OpLhu:
			return i3a("lhu")
		case OpLbs:
			return i3a("lbs")
		case OpLbu:
			return i3a("lbu")
		case OpSw:
			return i3a("sw")
		case OpSh:
			return i3a("sh")
		case OpSb:
			return i3a("sb")
		}
	}

	return fmt.Sprintf("noop-%d", op)
}
