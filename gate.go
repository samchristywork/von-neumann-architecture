package main

func xor(a bit, b bit) (out bit) {
	return a ^ b
}

func and(a bit, b bit) (out bit) {
	return a & b
}

func or(a bit, b bit) (out bit) {
	return a | b
}

func xor_3(a bit, b bit, c bit) (out bit) {
	return xor(xor(a, b), c)
}
