package inst

type Core interface {
	// register operations
	WriteReg(a uint8, v uint32)
	ReadReg(a uint8) uint32

	// memory operations
	WriteU8(addr uint32, v uint8)
	WriteU16(addr uint32, v uint16)
	WriteU32(addr uint32, v uint32)

	ReadU8(addr uint32) uint8
	ReadU16(addr uint32) uint16
	ReadU32(addr uint32) uint32
}

const (
	Nreg  = 32
	Nfreg = Nreg
	RegPC = Nreg - 1
)
