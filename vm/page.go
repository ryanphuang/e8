package vm

type Page interface {
	Write(offset uint32, b uint8)
	Read(offset uint32) uint8
}
