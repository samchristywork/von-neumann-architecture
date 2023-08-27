package main

import (
	"fmt"
)

const (
	UNKNOWN = iota
	LOAD_REG_A
	LOAD_REG_A_C
	LOAD_REG_B
	PRINT_REG_A
	PRINT_REG_B
	PRINT_REG_C
	ADD
	SUB
	SET_PC
	JZ
)

func control() {
	regA := register8bit{}
	regB := register8bit{}
	regC := register8bit{}

	for i := 0; i < len(input); i++ {
	}
}
