package parse

import (
	"strings"
)

func IsIdent(s string) bool {
	if len(s) == 0 {
		return false
	}

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

func IsLabel(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '.' && IsIdent(s[1:]) {
		return true
	}

	dot := strings.Index(s, ".")
	if dot < 0 {
		return IsIdent(s)
	}

	global := s[:dot]
	local := s[dot+1:]
	return IsIdent(global) && IsIdent(local)
}
