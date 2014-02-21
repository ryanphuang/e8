package vm

type Align struct {
	Page
}

const (
	pageOffset = 12
	PageSize   = 1 << pageOffset
	pageMask   = PageSize - 1
)

func maskOffset(offset uint32) uint32 { return offset & pageMask }
func alignU16(offset uint32) uint32   { return offset >> 1 << 1 }
func alignU32(offset uint32) uint32   { return offset >> 2 << 2 }

func offset8(offset uint32) uint32 {
	return maskOffset(offset)
}

func offset16(offset uint32) uint32 {
	return alignU16(maskOffset(offset))
}

func offset32(offset uint32) uint32 {
	return alignU32(maskOffset(offset))
}

func (self *Align) WriteU8(offset uint32, value uint8) {
	self.writeU8(offset8(offset), value)
}

func (self *Align) WriteU16(offset uint32, value uint16) {
	self.writeU16(offset16(offset), value)
}

func (self *Align) WriteU32(offset uint32, value uint32) {
	self.writeU32(offset32(offset), value)
}

func (self *Align) ReadU8(offset uint32) uint8 {
	return self.readU8(offset8(offset))
}

func (self *Align) ReadU16(offset uint32) uint16 {
	return self.readU16(offset16(offset))
}

func (self *Align) ReadU32(offset uint32) uint32 {
	return self.readU32(offset32(offset))
}

func (self *Align) writeU8(offset uint32, value uint8) {
	self.Page.Write(offset, value)
}

func (self *Align) writeU16(offset uint32, value uint16) {
	self.Page.Write(offset, uint8(value))
	self.Page.Write(offset+1, uint8(value>>8))
}

func (self *Align) writeU32(offset uint32, value uint32) {
	self.Page.Write(offset, uint8(value))
	self.Page.Write(offset+1, uint8(value>>8))
	self.Page.Write(offset+2, uint8(value>>16))
	self.Page.Write(offset+3, uint8(value>>24))
}

func (self *Align) readU8(offset uint32) uint8 {
	return self.Page.Read(offset)
}

func (self *Align) readU16(offset uint32) uint16 {
	ret := uint16(self.Page.Read(offset))
	ret |= uint16(self.Page.Read(offset+1)) << 8
	return ret
}

func (self *Align) readU32(offset uint32) uint32 {
	ret := uint32(self.Page.Read(offset))
	ret |= uint32(self.Page.Read(offset+1)) << 8
	ret |= uint32(self.Page.Read(offset+2)) << 16
	ret |= uint32(self.Page.Read(offset+3)) << 24
	return ret
}
