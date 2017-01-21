package main

import (
	"fmt"
	"go-risc-v-sim/riscv"
)

func main() {
	var cpu = riscv.NewCpu(10*riscv.KB)
	fmt.Println(cpu);
}

