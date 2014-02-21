package vm

type DataPage struct {
	bytes []byte
}

var _ Page = new(DataPage)

func NewDataPage() *DataPage {
	ret := new(DataPage)
	ret.bytes = make([]byte, PageSize)
	return ret
}

func (self *DataPage) Read(offset uint32) uint8 {
	return self.bytes[offset]
}

func (self *DataPage) Write(offset uint32, b uint8) {
	self.bytes[offset] = b
}
