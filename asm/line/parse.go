package line

import (
	"strings"
)

func opSplit(s string) (op, args string) {
	firstSpace := strings.IndexAny(s, " \t")
	if firstSpace < 0 {
		op = s
	} else {
		op = s[:firstSpace]
		args = strings.TrimSpace(s[firstSpace:])
	}

	return
}
