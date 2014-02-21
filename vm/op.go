package vm

type Op func(c *Core, inst uint32)

var rops = map[uint32]Op{
	0: func(c *Core, inst uint32) {

	},
}
