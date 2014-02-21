package vm

type Memory struct {
	pages map[uint32]Page
	align *Align
}

var noopPage = NewNoopPage()

func NewMemory() *Memory {
	ret := new(Memory)
	ret.pages = make(map[uint32]Page)
	ret.align = new(Align)
	return ret
}

func (self *Memory) Get(addr uint32) Page {
	ret := self.pages[addr >> pageOffset]
	if ret == nil {
		return noopPage
	}
	return ret
}

func (self *Memory) Align(addr uint32) *Align {
	self.align.Page = self.Get(addr)
	return self.align
}

func (self *Memory) WriteU8(addr uint32, value uint8) {
	self.Align(addr).WriteU8(addr, value)
}

func (self *Memory) WriteU16(addr uint32, value uint16) {
	self.Align(addr).WriteU16(addr, value)
}

func (self *Memory) WriteU32(addr uint32, value uint32) {
	self.Align(addr).WriteU32(addr, value)
}

func (self *Memory) ReadU8(addr uint32) uint8 {
	return self.Align(addr).ReadU8(addr)
}

func (self *Memory) ReadU16(addr uint32) uint16 {
	return self.Align(addr).ReadU16(addr)
}

func (self *Memory) ReadU32(addr uint32) uint32 {
	return self.Align(addr).ReadU32(addr)
}

func (self *Memory) Map(addr uint32, page Page) {
	self.pages[addr >> pageOffset] = page
}

func (self *Memory) Unmap(addr uint32) {
	delete(self.pages, addr >> pageOffset)
}
