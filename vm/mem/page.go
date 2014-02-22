package mem

type Page interface {
	Write(offset uint32, b uint8)
	Read(offset uint32) uint8
}

const (
	pageOffset = 12
	PageSize   = 1 << pageOffset
	pageMask   = PageSize - 1
)

func PageStart(i uint32) uint32 { return i << pageOffset }
