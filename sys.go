package e8

type Sys interface {
	Open(c Chan, v uint32)
	Close(c Chan)
	Select() (c Chan, v uint32)

	Page(id uint32) *Page
}
