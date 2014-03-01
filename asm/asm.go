package asm

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/h8liu/e8/asm/program"
)

func trim(s string) string { return strings.TrimSpace(s) }

func trimEndl(s string) string {
	for {
		n := len(s)
		if n > 0 && s[n-1] == '\n' {
			s = s[:n-1]
		} else {
			break
		}
	}

	return s
}

func trimLine(s string) (string, string) {
	i := strings.Index(s, ";")
	var c string
	if i >= 0 {
		s = s[:i]
		c = trimEndl(s[i:])
	}

	return trim(s), c
}

type Assembler struct {
	In       io.Reader
	Out      io.Writer
	Err      io.Writer
	Filename string
}

func (self *Assembler) Assemble() error {
	prog := program.New()

	scanner := bufio.NewScanner(self.In)
	sec := prog.NewSection("")
	lineNo := 0
	var lastError error
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		line, _ = trimLine(line) // TODO: record comments for fmt tools
		if line == "" {
			continue
		}

		if strings.HasSuffix(line, ":") {
			e := sec.AddLabel(line)
			if e != nil {
				fmt.Printf("%s:%d: %v\n", self.Filename, lineNo, e)
				lastError = e
			}
		} else {
			e := sec.AddLine(line, lineNo)
			if e != nil {
				fmt.Printf("%s:%d: %v\n", self.Filename, lineNo, e)
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

	e := prog.FillLabels(self.Err)
	if e != nil {
		return e
	}

	e = prog.CompileTo(self.Out)
	if e != nil {
		return e
	}

	return nil
}
