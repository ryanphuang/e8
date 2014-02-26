package mem

// A paged memory structure
type Memory struct {
	// Weather automatically allocate a new page on page miss
	NoAutoAlloc bool

	pages  map[uint32]Page
	_align *Align
}

var noopPage = NewNoopPage()

// Create an empty memory space
func New() *Memory {
	ret := new(Memory)
	ret.pages = make(map[uint32]Page)
	ret._align = new(Align)
	return ret
}

// Fetch the page associated with the address.
// If the page is missing, and auto allocation is on, a new page
// will be auto allocated
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

// Check if a page exists for the address
func (self *Memory) Check(addr uint32) bool {
	return self.pages[PageId(addr)] != nil
}

func (self *Memory) align(addr uint32) *Align {
	self._align.Page = self.Get(addr)
	return self._align
}

// Write a byte at addr.
// If the page is missing, and auto allocation is off, this is an noop.
func (self *Memory) WriteU8(addr uint32, value uint8) {
	self.align(addr).WriteU8(addr, value)
}

// Write a half word at addr, the address will be automatically aligned down.
// Byte order is little endian, and the lower bytes will be written first.
// If the page is missing and auto allocation is off, this is a noop.
func (self *Memory) WriteU16(addr uint32, value uint16) {
	self.align(addr).WriteU16(addr, value)
}

// Write a word at addr, the address will be automatically aligned down.
// Byte order is little endian, and the lower bytes will be written first.
// If the page is missing and auto allocation is off, this is a noop.
func (self *Memory) WriteU32(addr uint32, value uint32) {
	self.align(addr).WriteU32(addr, value)
}

// Read a byte at addr.
// If the page is missing and auto allocation is off, 0 is returned.
func (self *Memory) ReadU8(addr uint32) uint8 {
	return self.align(addr).ReadU8(addr)
}

// Read a half word at addr.
// Byte order is little endian, and the lower bytes will be read first.
// If the page is missing and auto allocation is off, 0 is returned.
func (self *Memory) ReadU16(addr uint32) uint16 {
	return self.align(addr).ReadU16(addr)
}

// Read a word at addr.
// Byte order is little endian, and the lower bytes will be read first.
// If the page is missing and auto allocation is off, 0 is returned.
func (self *Memory) ReadU32(addr uint32) uint32 {
	return self.align(addr).ReadU32(addr)
}

// Map a page for the address, the address will be auto aligned down to
// page boundaries.
func (self *Memory) Map(addr uint32, page Page) {
	self.pages[PageId(addr)] = page
}

// Unmap a page for the address. The unmapped page is returned.
func (self *Memory) Unmap(addr uint32) Page {
	id := PageId(addr)
	ret := self.pages[id]
	delete(self.pages, id)
	return ret
}
