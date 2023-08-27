package main

func full_adder(a bit, b bit, carry_in bit) (sum bit, carry_out bit) {
	sum = xor_3(a, b, carry_in)
	carry_out = or(and(a, b), and(carry_in, xor(a, b)))
	return sum, carry_out
}

// First 8 bits are result, last 5 are status
func add_8_bit(a [8]bit, b [8]bit, carry_in bit) (out [13]bit) {
	var carry_out bit
	out[7], carry_out = full_adder(a[7], b[7], carry_in)
	out[6], carry_out = full_adder(a[6], b[6], carry_out)
	out[5], carry_out = full_adder(a[5], b[5], carry_out)
	out[4], carry_out = full_adder(a[4], b[4], carry_out)
	out[3], carry_out = full_adder(a[3], b[3], carry_out)
	out[2], carry_out = full_adder(a[2], b[2], carry_out)
	out[1], carry_out = full_adder(a[1], b[1], carry_out)
	out[0], carry_out = full_adder(a[0], b[0], carry_out)
	out[8] = carry_out
	return out
}

// First 8 bits are result, last 5 are status
func sub_8_bit(a [8]bit, b [8]bit, carry_in bit) (out [13]bit) {
	var carry_out bit
	out[7], carry_out = full_adder(a[7], xor(b[7], 1), carry_in)
	out[6], carry_out = full_adder(a[6], xor(b[6], 1), carry_out)
	out[5], carry_out = full_adder(a[5], xor(b[5], 1), carry_out)
	out[4], carry_out = full_adder(a[4], xor(b[4], 1), carry_out)
	out[3], carry_out = full_adder(a[3], xor(b[3], 1), carry_out)
	out[2], carry_out = full_adder(a[2], xor(b[2], 1), carry_out)
	out[1], carry_out = full_adder(a[1], xor(b[1], 1), carry_out)
	out[0], carry_out = full_adder(a[0], xor(b[0], 1), carry_out)
	out[8] = carry_out

	return out
}

func alu(a [8]bit, b [8]bit, opcode [4]bit, carry_in bit) (out [8]bit, status [5]bit) {
	//TODO: carry out bit, zero bit, negative bit, overflow bit, parity bit

	var results [16][13]bit
	results[0] = add_8_bit(a, b, 0)
	results[1] = sub_8_bit(a, b, 1)

	result := mux_4_16(opcode, results)

	out = first_8(result)
	status = last_5(result)

	return out, status
}
