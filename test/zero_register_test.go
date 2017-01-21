package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"go-risc-v-sim/riscv"
)

func TestRegisterZeroReadsZero(t *testing.T) {
	assert := assert.New(t)

	cpu := riscv.NewCpu(0)

	expected := riscv.Register(0)
	assert.Equal(expected, cpu.ReadRegister(0))
}

func TestRegisterZeroWriteDoesNothing(t *testing.T) {
	assert := assert.New(t)

	cpu := riscv.NewCpu(0)

	cpu.WriteRegister(0, riscv.Register(0x1234567812345678))

	expected := riscv.Register(0)
	assert.Equal(expected, cpu.ReadRegister(0))
}

