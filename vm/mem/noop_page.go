package mem

type NoopPage struct{}

func NewNoopPage() *NoopPage {
	return new(NoopPage)
}

var _ Page = new(NoopPage)

func (self *NoopPage) Write(offset uint32, b uint8) {}
func (self *NoopPage) Read(offset uint32) uint8     { return 0 }
