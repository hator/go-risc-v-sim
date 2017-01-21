package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"go-risc-v-sim/riscv"
)

func TestLoadImmediateToReg(t *testing.T) {
	assert := assert.New(t)

	// Given
	var value uint32 = 0xf0f
	var register uint32 = 1
	cpu := riscv.NewCpu(0)

	// load value to register
	var addFunct3 uint32 = 0
	var addIOpcode uint32 = 0x13
	addImmInstr := riscv.InstructionI(value, 0, addFunct3, register, addIOpcode)

	// When
	cpu.RunInstruction(addImmInstr)

	// Then
	expected := riscv.Register(value)
	regNumber := riscv.RegisterNum(register)
	assert.Equal(expected, cpu.ReadRegister(regNumber))
}

func TestAddRegToImmediate(t *testing.T) {
	assert := assert.New(t)

	// Given
	var regValue uint32 = 2
	var immediate uint32 = 1
	var register uint32 = 1
	cpu := riscv.NewCpu(0)

	// load value to register
	_LoadToRegister(cpu, register, regValue)

	var addFunct3 uint32 = 0
	var addIOpcode uint32 = 0x13
	addImmToRegInstr := riscv.InstructionI(immediate, register, addFunct3, register, addIOpcode)
	cpu.RunInstruction(addImmToRegInstr)

	expected := riscv.Register(regValue + immediate)
	regNumber := riscv.RegisterNum(register)
	assert.Equal(expected, cpu.ReadRegister(regNumber))
}

func TestAddRegToReg(t *testing.T) {
	assert := assert.New(t)

	// Given
	var regA uint32 = 1
	var regAValue uint32 = 123
	var regB uint32 = 2
	var regBValue uint32 = 321
	var regC uint32 = 3
	cpu := riscv.NewCpu(0)

	// load value to register
	_LoadToRegister(cpu, regA, regAValue)
	_LoadToRegister(cpu, regB, regBValue)

	var addFunct3 uint32 = 0
	var addOpcode uint32 = 0x33
	addImmToRegInstr := riscv.InstructionR(0, regB, regA, addFunct3, regC, addOpcode)
	cpu.RunInstruction(addImmToRegInstr)

	expected := riscv.Register(regAValue + regBValue)
	regNumber := riscv.RegisterNum(regC)
	assert.Equal(expected, cpu.ReadRegister(regNumber))
}

func TestSub(t *testing.T) {
	assert := assert.New(t)

	// Given
	var regA uint32 = 1
	var regAValue uint32 = 123
	var regB uint32 = 2
	var regBValue uint32 = 321
	var regC uint32 = 3
	cpu := riscv.NewCpu(0)

	// load value to register
	_LoadToRegister(cpu, regA, regAValue)
	_LoadToRegister(cpu, regB, regBValue)

	var subFunc7 uint32 = 0x20
	var subFunct3 uint32 = 0
	var subOpcode uint32 = 0x33
	addImmToRegInstr := riscv.InstructionR(subFunc7, regB, regA, subFunct3, regC, subOpcode)
	cpu.RunInstruction(addImmToRegInstr)

	expected := riscv.Register(uint64(regAValue) - uint64(regBValue))
	regNumber := riscv.RegisterNum(regC)
	assert.Equal(expected, cpu.ReadRegister(regNumber))
}

func _LoadToRegister(cpu *riscv.Cpu, register uint32, value uint32) {
	loadToRegInstr := riscv.InstructionI(value, 0, 0, register, 0x13)
	cpu.RunInstruction(loadToRegInstr)
}
