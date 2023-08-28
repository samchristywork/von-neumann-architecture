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
		instruction := decimel_to_byte(int(input[i]))

		var lookahead [8]bit
		if i+1 < len(input) {
			lookahead = decimel_to_byte(int(input[i+1]))
		}

		if match(instruction, decimel_to_byte(LOAD_REG_A)) == 1 { // Load reg A
			regA.load(lookahead, 0)
			regA.load(lookahead, 1)
			i++
		} else if match(instruction, decimel_to_byte(LOAD_REG_A_C)) == 1 { // Load reg C into reg A
			regA.load(regC.read(), 0)
			regA.load(regC.read(), 1)
		} else if match(instruction, decimel_to_byte(LOAD_REG_B)) == 1 { // Load reg B
			regB.load(lookahead, 0)
			regB.load(lookahead, 1)
			i++
		} else if match(instruction, decimel_to_byte(ADD)) == 1 { // Add reg A and B into reg C
			out, _ := alu(regA.read(), regB.read(), decimel_to_nibble(ALU_ADD), 0)
			regC.load(out, 0)
			regC.load(out, 1)
		} else if match(instruction, decimel_to_byte(SUB)) == 1 { // Subtract reg B from A into reg C
			out, _ := alu(regA.read(), regB.read(), decimel_to_nibble(ALU_SUB), 0)
			regC.load(out, 0)
			regC.load(out, 1)
		} else if match(instruction, decimel_to_byte(JZ)) == 1 { // Jump if zero
			if isZero(regC.read()) == 1 {
				i = byte_to_decimel(lookahead) - 1
			}
			i++
		} else if match(instruction, decimel_to_byte(SET_PC)) == 1 { // Set program counter
			i = byte_to_decimel(lookahead) - 1
		} else {
			fmt.Printf("Unknown instruction: %v\n", instruction)
		}
	}
}
