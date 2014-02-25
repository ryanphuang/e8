package mem

type Memory struct {
	NoAutoAlloc bool

	pages  map[uint32]Page
	_align *Align
}

var noopPage = NewNoopPage()

func New() *Memory {
	ret := new(Memory)
	ret.pages = make(map[uint32]Page)
	ret._align = new(Align)
	return ret
}

func (self *Memory) Get(addr uint32) Page {
	id := PageId(addr)
	ret := self.pages[PageId(addr)]
	if ret == nil {
		if self.NoAutoAlloc {
			return noopPage
		}
		p := NewPage()
		self.pages[id] = p
		return p
	}

	return ret
}

func (self *Memory) Valid(addr uint32) bool {
	return self.pages[PageId(addr)] != nil
}

func (self *Memory) align(addr uint32) *Align {
	self._align.Page = self.Get(addr)
	return self._align
}

func (self *Memory) WriteU8(addr uint32, value uint8) {
	self.align(addr).WriteU8(addr, value)
}

func (self *Memory) WriteU16(addr uint32, value uint16) {
	self.align(addr).WriteU16(addr, value)
}

func (self *Memory) WriteU32(addr uint32, value uint32) {
	self.align(addr).WriteU32(addr, value)
}

func (self *Memory) ReadU8(addr uint32) uint8 {
	return self.align(addr).ReadU8(addr)
}

func (self *Memory) ReadU16(addr uint32) uint16 {
	return self.align(addr).ReadU16(addr)
}

func (self *Memory) ReadU32(addr uint32) uint32 {
	return self.align(addr).ReadU32(addr)
}

func (self *Memory) Map(addr uint32, page Page) {
	self.pages[PageId(addr)] = page
}

func (self *Memory) Unmap(addr uint32) {
	delete(self.pages, PageId(addr))
}
