package argfmt

import (
	"fmt"

	"github.com/h8liu/e8/asm/parse"
	"github.com/h8liu/e8/vm/inst"
)

type Format struct {
	s string
}

func New(s string) *Format {
	ret := new(Format)
	ret.s = s
	return ret
}

func ef(s string, args ...interface{}) (uint32, string, error) {
	return 0, "", fmt.Errorf(s, args...)
}

func argName(i int) string {
	if i == 0 {
		return "first arg"
	}
	if i == 1 {
		return "second arg"
	}
	if i == 2 {
		return "third arg"
	}
	if i == 3 {
		return "forth arg"
	}
	return fmt.Sprintf("arg %d", i+1)
}

func (self *Format) Parse(base uint32, s string) (uint32, string, error) {
	args := parseArgs(s)

	if len(args) != len(self.s) {
		return ef("invalid arg count")
	}

	ret := base
	label := ""

	for i, r := range self.s {
		arg := args[i]

		switch r {
		case 's', 't', 'd':
			_reg, valid := parseReg(arg)
			if !valid {
				return ef("%s is not a register", argName(i))
			}
			reg := uint32(_reg)

			if r == 's' {
				ret |= reg << inst.RsShift
			} else if r == 't' {
				ret |= reg << inst.RtShift
			} else if r == 'd' {
				ret |= reg << inst.RdShift
			}
		case 'S':
			shamt, valid := parseShamt(arg)
			if !valid {
				return ef("%s is not a shamt", argName(i))
			}

			ret |= uint32(shamt) << inst.ShamtMask
		case 'u':
			if parse.IsIdent(arg) {
				label = arg
			} else {
				imu, valid := parseImu(arg)
				if !valid {
					return ef("%s is not an unsigned immediate", argName(i))
				}

				ret |= uint32(imu)
			}
		case 'i':
			if parse.IsIdent(arg) {
				label = arg
			} else {
				ims, valid := parseIms(arg)
				if !valid {
					return ef("%s is not a signed immediate", argName(i))
				}
				ret |= uint32(uint16(ims))
			}
		case 'a':
			ims, rs, valid := parseAddr(arg)
			if !valid {
				return ef("%s is not a valid address", argName(i))
			}

			ret |= uint32(uint16(ims))
			ret |= uint32(rs) << inst.RsShift
		case 'l':
			if !parse.IsIdent(arg) {
				return ef("invalid label")
			}

			label = arg
		default:
			panic("bug")
		}
	}

	return ret, label, nil
}
