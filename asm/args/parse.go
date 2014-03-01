package args

import (
	"math"
	"strconv"
	"strings"

	"github.com/h8liu/e8/vm/inst"
)

func parseArgs(args string) []string {
	ret := strings.Split(args, ",")
	for i, s := range ret {
		ret[i] = strings.TrimSpace(s)
	}
	return ret
}

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 32)
}

func parseReg(s string) (uint8, bool) {
	if len(s) < 2 {
		return 0, false
	}
	s = strings.ToLower(s)

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
