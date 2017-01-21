package riscv

func InstructionR(funct7 uint32, rs2 uint32, rs1 uint32, funct3 uint32, rd uint32, opcode uint32) Instruction {
	instr := (funct7 & BITS7) << 25 | (rs2 & BITS5) << 20 | _InstructionBase(rs1, funct3, rd, opcode)
	return Instruction(instr)
}

func InstructionI(immediate uint32, rs1 uint32, funct3 uint32, rd uint32, opcode uint32) Instruction {
	instr := (immediate & BITS12) << 20 | _InstructionBase(rs1, funct3, rd, opcode)
	return Instruction(instr)
}

func _InstructionBase(rs1 uint32, funct3 uint32, rd uint32, opcode uint32) uint32 {
	return (rs1 & BITS5) << 15 | (funct3 & BITS3) << 12 | (rd & BITS5) << 7 | opcode
}
