package e8

type Page struct {
	bytes []byte
}

const PageSize = 4096

func NewPage() *Page {
	ret := new(Page)
	ret.bytes = make([]byte, PageSize)
	return ret
}
