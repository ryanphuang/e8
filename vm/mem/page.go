package mem

// A general page inteface
type Page interface {
	// Writes a byte at a particular page offset.
	Write(offset uint32, b uint8)

	// Reads a byte at a particular page offset.
	Read(offset uint32) uint8
}

const (
	PageOffset = 12              // Number of bits for page offset
	PageSize   = 1 << PageOffset // 4096 bytes
	PageMask   = PageSize - 1    // Bit mask for page offset
)

// Map from page id to page start address.
func PageStart(i uint32) uint32 { return i << PageOffset }

// Map from address to page id.
func PageId(i uint32) uint32 { return i >> PageOffset }
