package inst

type fields struct {
	inst  uint32
	rs    uint8
	rt    uint8
	rd    uint8
	shamt uint8
	im    uint16
}

const (
	Nfunct = 64
	Nop    = 64
)
