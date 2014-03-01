package asm

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/h8liu/e8/asm/locator"
	"github.com/h8liu/e8/asm/section"
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
	Filename string

	sections   []*section.Section
	sectionMap map[string]*section.Section
}

var _ locator.Locator = new(Assembler)

func (self *Assembler) Locate(s string) (uint32, bool) {
	sec := self.sectionMap[s]
	if sec == nil {
		return 0, false
	}

	return sec.Start, true
}

func (self *Assembler) LocateData(s string) (uint32, bool) {
	panic("todo")
}

func (self *Assembler) Assemble() error {
	self.sections = make([]*section.Section, 0, 1024)
	self.sectionMap = make(map[string]*section.Section)

	scanner := bufio.NewScanner(self.In)
	sec := section.New("")
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

	sec.FillLabels(self, nil)

	e := sec.CompileTo(self.Out)
	if e != nil {
		return e
	}

	return nil
}
