package riscv

func (cpu *Cpu) addi(rd RegisterNum, rsVal Register, imm Word) {
	cpu.WriteRegister(rd, rsVal + Register(imm))
}

func (cpu *Cpu) add(rd RegisterNum, rs1Val Register, rs2Val Register) {
	cpu.WriteRegister(rd, rs1Val + rs2Val)
}

func (cpu *Cpu) sub(rd RegisterNum, rs1Val Register, rs2Val Register) {
	cpu.WriteRegister(rd, rs1Val - rs2Val)
}
