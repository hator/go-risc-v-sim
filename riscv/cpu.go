package riscv

import "fmt"

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
)

const (
	BITS3 = 0x07	// Lowest 3 bits on
	BITS5 = 0x1F	// Lowest 5 bits on
	BITS7 = 0x7F	// Lowest 7 bits on
	BITS12 = 0x0FFF	// Lowest 12 bits on
)

type usize uint64

type Word uint64

type Instruction Word
type Opcode Word

type Register Word
type RegisterNum uint8

const (
	SR_S	Register = 0x00000001
	SR_PS	Register = 0x00000002
	SR_EI	Register = 0x00000004
	SR_PEI	Register = 0x00000008
	SR_EF	Register = 0x00000010
	// TODO To be continued...
)

const (
	ADD = 0
	SUB7 = 0x20
)

type Memory []Word

type Cpu struct {
	pc Register
	status_reg Register
	gen_reg [31]Register
	memory Memory
}

func NewCpu(memorySize usize) *Cpu {
	return &Cpu {
		0x2000, // PC
		0x0000, // Status reg
		[31]Register{},
		make(Memory, memorySize, memorySize),
	}
}

func (cpu Cpu) ReadRegister(register RegisterNum) Register {
	if(register > 32) {
		panic(fmt.Sprintf("Register %d beyond 32 register count", register))
	}

	if(register == 0) {
		return 0
	}

	return cpu.gen_reg[register-1]
}

func (cpu *Cpu) WriteRegister(register RegisterNum, value Register) {
	if(register > 32) {
		panic(fmt.Sprintf("Register %d beyond 32 register count", register))
	}

	if(register != 0) {
		cpu.gen_reg[register-1] = value
	}
}

func (cpu *Cpu) RunInstruction(instr Instruction) {
	opcode := instr.GetOpcode()

	// Decode all parameters as a hardware processor would do
	funct3 := instr.GetFunct3()
	funct7 := instr.GetFunct7()
	rd := instr.GetRd()
	rs1 := instr.GetRs1()
	rs2 := instr.GetRs2()
	immI := instr.GetImmI()

	fmt.Printf("Decoded instruction:\n\tfunc %b\n\trd %b\n\trs1 %b\n\trs2 %b\n\timmI %b\n", funct3, rd, rs1, rs2, immI)

	switch opcode {
	case 0x13: // Immediate integer arithmetic
		rs1Val := cpu.ReadRegister(rs1)
		switch funct3 {
		case ADD:
			fmt.Println("addi")
			cpu.addi(rd, rs1Val, immI)
		}
	case 0x33: // Reg to reg integer arithmetic
		rs1Val := cpu.ReadRegister(rs1)
		rs2Val := cpu.ReadRegister(rs2)
		switch funct3 {
		case ADD:
			switch funct7 {
			case ADD:
				fmt.Println("add")
				cpu.add(rd, rs1Val, rs2Val)
			case SUB7:
				fmt.Println("sub")
				cpu.sub(rd, rs1Val, rs2Val)
			default:
				panic("Invalid instruction, expecting funct7 to be either Add or Sub for Add func3")
				// TODO throw cpu hardware exception
			}
		}
	}
}
