package align

func U16(offset uint32) uint32 { return offset >> 1 << 1 }
func U32(offset uint32) uint32 { return offset >> 2 << 2 }
