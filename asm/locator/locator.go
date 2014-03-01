package locator

type Locator interface {
	Locate(lab string) (uint32, bool)
}
