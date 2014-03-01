package section

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/h8liu/e8/asm/line"
	"github.com/h8liu/e8/asm/locator"
	"github.com/h8liu/e8/asm/parse"
)

func ef(s string, args ...interface{}) error {
	return fmt.Errorf(s, args...)
}

type Section struct {
	Name      string
	Start     uint32
	lines     []*line.Line
	labelMap  map[string]*label
	labels    []*label
	curGlobal string
}

func New(name string) *Section {
	ret := new(Section)
	ret.Name = name
	ret.lines = make([]*line.Line, 0, 1024)
	ret.labelMap = make(map[string]*label)
	ret.labels = make([]*label, 0, 1024)

	return ret
}

func (self *Section) Nline() int   { return len(self.lines) }
func (self *Section) Size() uint32 { return uint32(self.Nline() << 2) }

func (self *Section) AddLine(s string, lineNo int) error {
	line, e := line.Parse(s)
	if e != nil {
		return e
	}
	line.LineNo = lineNo
	line.Scope = self.curGlobal

	self.lines = append(self.lines, line)
	return nil
}

func (self *Section) AddLabel(s string) error {
	if strings.HasSuffix(s, ":") {
		s = s[:len(s)-1]
	}

	if len(s) == 0 {
		return ef("empty label")
	}

	if s[0] == '.' {
		// local label
		s = s[1:]
		if !parse.IsIdent(s) {
			return ef("invalid local label")
		}

		s = self.curGlobal + "." + s
	} else {
		// global label
		if !parse.IsIdent(s) {
			return ef("invalid global label")
		}

		self.curGlobal = s
	}

	if self.labelMap[s] != nil {
		return ef("redefined label: %s", s)
	}

	label := &label{len(self.lines), s}
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

func (self *Section) Locate(lab string) (uint32, bool) {
	label, found := self.labelMap[lab]
	if !found {
		return 0, false
	}

	return uint32(label.index<<2) + self.Start, true
}

func pu32(buf []byte, i uint32) {
	buf[0] = uint8(i)
	buf[1] = uint8(i >> 8)
	buf[2] = uint8(i >> 16)
	buf[3] = uint8(i >> 24)
}

func (self *Section) CompileTo(out io.Writer) error {
	buf := make([]byte, 4)

	for _, line := range self.lines {
		pu32(buf, line.U32())
		_, e := out.Write(buf)
		if e != nil {
			return e
		}
	}

	return nil
}

func (self *Section) FillLabels(loc locator.Locator, err io.Writer) error {
	for i, line := range self.lines {
		curPos := self.Start + uint32(i<<2) + 4
		if line.IsJump() {
			label := line.Label()
			pos, found := self.Locate(label)
			if !found {
				pos, found = loc.Locate(label)
			}
			offset := int32(pos-curPos) >> 2
			if offset > (0x1<<25)-1 {
				panic("jump out of range") // TODO
			} else if offset < -(0x1 << 25) {
				panic("jump out of range") // TODO
			}

			line.J(offset)
		} else if line.IsBranch() {
			label := line.Label()
			pos, found := self.Locate(label)
			if !found {
				pos, found = loc.Locate(label)
			}
			offset := int32(pos-curPos) >> 2
			if offset > math.MaxInt16 {
				panic("branch out of range") // TODO
			} else if offset < math.MinInt16 {
				panic("branch out of range") // TODO
			}

			line.Ims(int16(offset))
		}
	}

	return nil
}
