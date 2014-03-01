package args

import (
	"bytes"
	"strings"
)

var fieldNames = map[string]rune{
	"rs":    's',
	"rt":    't',
	"rd":    'd',
	"shamt": 'S',
	"imu":   'u',
	"ims":   'i',
	"addr":  'a',
	"label": 'l',
}

func Compile(s string) string {
	ret := new(bytes.Buffer)

	args := strings.Split(strings.ToLower(s), ",")
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		r, found := fieldNames[arg]
		if !found {
			panic("invalid field")
		}

		ret.WriteRune(r)
	}

	return ret.String()
}
