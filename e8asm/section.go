package e8asm

import (
	"fmt"
	"io"
)

type Section struct {
	name  string
	lines []*Line
	// labels []*Label
}

func NewSection(name string) *Section {
	ret := new(Section)
	ret.name = name
	ret.lines = make([]*Line, 0, 1024)
	return ret
}

func (self *Section) Line(s string) error {
	line, e := ParseLine(s)
	if e != nil {
		return e
	}

	self.lines = append(self.lines, line)
	return nil
}

func (self *Section) Label(s string) error {
	panic("todo")
}

func (self *Section) PrintTo(out io.Writer) error {
	for _, line := range self.lines {
		_, e := fmt.Fprintln(out, "%v\n", line)
		if e != nil {
			return e
		}
	}
	return nil
}
