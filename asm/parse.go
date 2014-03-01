package asm

import (
	"math"
	"strconv"
	"strings"

	"github.com/h8liu/e8/vm/inst"
)

func trim(s string) string  { return strings.TrimSpace(s) }
func lower(s string) string { return strings.ToLower(s) }

func fields(args string) []string {
	ret := strings.Split(args, ",")
	for i, s := range ret {
		ret[i] = trim(s)
	}
	return ret
}

func opSplit(s string) (op, args string) {
	firstSpace := strings.IndexAny(s, " \t")
	if firstSpace < 0 {
		op = s
	} else {
		op = s[:firstSpace]
		args = trim(s[firstSpace:])
	}

	return
}

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 32)
}

func parseReg(s string) (uint8, bool) {
	if len(s) < 2 {
		return 0, false
	}
	s = lower(s)

	if s == "pc" {
		return inst.RegPC, true
	}

	if s[0] != '$' && s[0] != 'r' {
		return 0, false
	}

	n, e := parseInt(s[1:])
	if e != nil {
		return 0, false
	}
	if n < 0 {
		return 0, false
	}
	if n >= inst.Nreg {
		return 0, false
	}

	return uint8(n), true
}

func parseShamt(s string) (uint8, bool) {
	n, e := parseInt(s)
	if e != nil {
		return 0, false
	}
	if n < 0 {
		return 0, false
	}
	if n >= 32 {
		return 0, false
	}
	return uint8(n), true
}

func parseIms(s string) (uint16, bool) {
	n, e := parseInt(s)
	if e != nil {
		return 0, false
	}
	if n < math.MinInt16 {
		return 0, false
	}
	if n > math.MaxInt16 {
		return 0, false
	}
	return uint16(int16(n)), true
}

func parseImu(s string) (uint16, bool) {
	n, e := parseInt(s)
	if e != nil {
		return 0, false
	}
	if n < 0 {
		return 0, false
	}
	if n > math.MaxUint16 {
		return 0, false
	}
	return uint16(n), true
}

func parseAddr(s string) (im uint16, rs uint8, valid bool) {
	ns := len(s)
	if s[ns-1] != ')' {
		// bare signed im
		im, valid = parseIms(s)
		if !valid {
			return 0, 0, false
		}
		return im, 0, true
	}
	sep := strings.Index(s, "(")
	if sep < 0 {
		return 0, 0, false
	}

	imStr := s[:sep]
	regStr := s[sep+1 : ns-1]
	im, valid = parseIms(imStr)
	if !valid {
		return 0, 0, false
	}
	rs, valid = parseReg(regStr)
	if !valid {
		return 0, 0, false
	}

	return im, rs, true
}

func isIdent(s string) bool {
	for i, c := range s {
		if c == '_' {
			continue
		}
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			continue
		}
		if c >= '0' && c <= '9' {
			if i == 0 {
				return false
			}
			continue
		}
		return false
	}

	return true
}
