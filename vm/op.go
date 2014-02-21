package vm

type Op func(inst uint32, reg *Registers, mem *Memory, state *State)
