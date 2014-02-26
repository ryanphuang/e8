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

type Assembler struct {
	In       io.Reader
	Out      io.Writer
	Filename string
}

var _ Locator = new(Assembler)

func (self *Assembler) Locate(s string) (uint32, bool) {
	// TODO
	return 0, false
}

func (self *Assembler) Assemble() error {
	scanner := bufio.NewScanner(self.In)
	sec := NewSection("")
	lineno := 0
	var lastError error
	for scanner.Scan() {
		lineno++
		line := scanner.Text()
		line = trimLine(line)
		if line == "" {
			continue
		}

		if strings.HasSuffix(line, ":") {
			e := sec.Label(line)
			if e != nil {
				fmt.Printf("%d: %v\n", lineno, e)
				lastError = e
			}
		} else {
			e := sec.Line(line, lineno)
			if e != nil {
				fmt.Printf("%d: %v\n", lineno, e)
				lastError = e
			}
		}
	}

	if lastError != nil {
		return lastError
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	sec.FillLabels(self, nil)

	e := sec.CompileTo(self.Out)
	if e != nil {
		return e
	}

	return nil
}
