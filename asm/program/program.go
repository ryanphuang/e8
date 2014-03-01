package program

import (
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
	ret.sections = make([]*section.Section, 1024)
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
