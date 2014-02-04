package arch

type Vwriter struct {
	vr    *Vreader
	id    int
	v     uint32
	valid bool
}

func NewVwriter(vr *Vreader) *Vwriter {
	ret := new(Vwriter)
	ret.vr = vr
	ret.id = vr.register(ret)
	return ret
}

func (self *Vwriter) Write(v uint32) {
	wasValid := self.valid
	self.v = v
	self.valid = true
	if !wasValid {
		self.vr.notify(self.id) // soft hint for pulling
	}
}

func (self *Vwriter) Clear() {
	self.valid = false
}

func (self *Vwriter) Fetch() (bool, uint32) {
	if !self.valid {
		return false, 0
	}

	return true, self.v
}
