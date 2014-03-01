package program

import (
	"io"

	"github.com/h8liu/e8/asm/locator"
	"github.com/h8liu/e8/asm/section"
)

type Program struct {
	sections   []*section.Section
	sectionMap map[string]*section.Section
}

var _ locator.Locator = new(Program)

func New() *Program {
	ret := new(Program)
	ret.sections = make([]*section.Section, 0, 1024)
	ret.sectionMap = make(map[string]*section.Section)
	return ret
}

func (self *Program) Locate(s string) (uint32, bool) {
	sec := self.sectionMap[s]
	if sec == nil {
		return 0, false
	}

	return sec.Start, true
}

func (self *Program) NewSection(name string) *section.Section {
	ret := section.New(name)
	self.sections = append(self.sections, ret)
	self.sectionMap[name] = ret
	return ret
}

func (self *Program) CompileTo(out io.Writer) error {
	for _, sec := range self.sections {
		e := sec.CompileTo(out)
		if e != nil {
			return e
		}
	}

	return nil
}

func (self *Program) FillLabels(err io.Writer) error {
	for _, sec := range self.sections {
		e := sec.FillLabels(self, err)
		if e != nil {
			return e
		}
	}

	return nil
}
