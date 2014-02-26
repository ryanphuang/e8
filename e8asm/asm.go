package e8asm

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func trimLine(s string) string {
	i := strings.Index(s, ";")
	if i >= 0 {
		s = s[:i]
	}

	return trim(s)
}

func Assemble(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	sec := NewSection("")
	lineno := 0
	for scanner.Scan() {
		lineno++
		line := scanner.Text()
		line = trimLine(line)
		if strings.HasSuffix(line, ":") {
			// label
		} else {
			e := sec.Line(line)
			if e != nil {
				fmt.Println(lineno, e)
			}
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	fmt.Println(sec)

	return nil
}
