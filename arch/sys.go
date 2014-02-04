package arch

type Sys interface {
	Open(c Chan, v uint32)
	Close(c Chan)
	IsOpen(c Chan) bool
	Select() (b bool, c Chan, v uint32)

	Page(id uint32) *Page

	Halt()
}
