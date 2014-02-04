package e8

type Vreader struct {
	writers map[int]*Vwriter

	seq int
}

func NewVreader() *Vreader {
	ret := new(Vreader)
	ret.writers = make(map[int]*Vwriter)
	return ret
}

func (self *Vreader) nextSeq() int {
	for {
		ret := self.seq
		self.seq++
		if self.writers[ret] == nil {
			return ret
		}
	}
}

func (self *Vreader) register(vw *Vwriter) int {
	seq := self.nextSeq()
	self.writers[seq] = vw
	return seq
}

func (self *Vreader) notify(i int) {
	panic("todo")
}
