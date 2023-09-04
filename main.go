package main

type bit int

var input = []byte{
	LOAD_REG_A, 100,
	LOAD_REG_B, 1, // START
	SUB,
	PRINT_REG_C,
	LOAD_REG_A_C,
	JZ, 11, // END
	SET_PC, 2, // START
	PRINT_REG_A, // END
	PRINT_REG_A,
	PRINT_REG_A,
	PRINT_REG_A,
	PRINT_REG_A,
	PRINT_REG_A,
}

func main() {
	control()
}
