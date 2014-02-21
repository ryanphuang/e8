package vm

type Core struct {
	*Registers
	*Memory
	*State
}
