package e8asm

import (
	"fmt"
	"io"
	"strings"
)

type Section struct {
	Name     string
	Start    uint32
	lines    []*Line
	labelMap map[string]*Label
	labels   []*Label
}

func NewSection(name string) *Section {
	ret := new(Section)
	ret.Name = name
	ret.lines = make([]*Line, 0, 1024)
	ret.labelMap = make(map[string]*Label)
	ret.labels = make([]*Label, 0, 1024)
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
	if !strings.HasSuffix(s, ":") {
		return ef("not a label")
	}

	s = s[:len(s)-1]
	if !isIdent(s) {
		return ef("not a valid label")
	}

	if self.labelMap[s] != nil {
		return ef("redefined label: %s", s)
	}

	label := &Label{len(self.lines), s}
	self.labelMap[s] = label
	self.labels = append(self.labels, label)
	return nil
}

func (self *Section) PrintTo(out io.Writer) error {
	labIndex := 0
	for i, line := range self.lines {
		for labIndex < len(self.labels) {
			lab := self.labels[labIndex]
			if lab.index <= i {
				fmt.Fprintf(out, "%s:\n", lab.name)
				labIndex++
			} else {
				break
			}
		}

		_, e := fmt.Fprintf(out, "\t%v\n", line)
		if e != nil {
			return e
		}
	}
	return nil
}

func (self *Section) LocateLabel(lab string) (uint32, bool) {
	label, found := self.labelMap[lab]
	if !found {
		return 0, false
	}

	return uint32(label.index<<2) + self.Start, true
}

func (self *Section) CompileTo(out io.Writer) {
	panic("todo")
}
