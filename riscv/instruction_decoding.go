package riscv

func (instruction Instruction) GetOpcode() Opcode {
	return Opcode(instruction & BITS7)	// instr[6:0]
}

func (instruction Instruction) GetRd() RegisterNum {
	return RegisterNum((instruction >> 7) & BITS5)	// instr[11:7]
}

func (instruction Instruction) GetFunct3() RegisterNum {
	return RegisterNum((instruction >> 12) & BITS3)	// instr[14:12]
}

func (instruction Instruction) GetFunct7() RegisterNum {
	return RegisterNum((instruction >> 25) & BITS7)	// instr[31:25]
}

func (instruction Instruction) GetRs1() RegisterNum {
	return RegisterNum((instruction >> 15) & BITS5)	// instr[19:15]
}

func (instruction Instruction) GetRs2() RegisterNum {
	return RegisterNum((instruction >> 20) & BITS5)	// instr[24:20]
}

func (instruction Instruction) GetImmI() Word {
	return Word((instruction >> 20) & BITS12)	// instr[31:20]
}
