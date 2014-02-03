package e8

type Chan uint32

const (
	IsRecv = 0x1
	IsPage = 0x2
)

func (c Chan) IsRecv() bool  { return c&IsRecv != 0 }
func (c Chan) IsSend() bool  { return c&IsRecv == 0 }
func (c Chan) IsPage() bool  { return c&IsPage != 0 }
func (c Chan) IsValue() bool { return c&IsPage == 0 }

func Vsend(i uint32) Chan { return Chan(i << 2) }
func Vrecv(i uint32) Chan { return Chan((i << 2) | IsRecv) }
func Psend(i uint32) Chan { return Chan((i << 2) | IsPage) }
func Precv(i uint32) Chan { return Chan((i << 2) | IsPage | IsRecv) }

const (
	Pin  = IsPage | IsRecv
	Pout = IsPage
	Vin  = IsRecv
	Vout = 0
)
