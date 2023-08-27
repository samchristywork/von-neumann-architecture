package main

func decimel_to_byte(dec int) (out [8]bit) {
	for i := 7; i >= 0; i-- {
		out[i] = bit(dec % 2)
		dec /= 2
	}

	return out
}

func byte_to_decimel(in [8]bit) (out int) {
	out = 0
	for i := 0; i < 8; i++ {
		out *= 2
		out += int(in[i])
	}

	return out
}

func decimel_to_nibble(dec int) (out [4]bit) {
	for i := 3; i >= 0; i-- {
		out[i] = bit(dec % 2)
		dec /= 2
	}

	return out
}

func match(a [8]bit, b [8]bit) (out bit) {
	out = 1
	for i := 0; i < 8; i++ {
		if a[i] != b[i] {
			out = 0
		}
	}

	return out
}

func isZero(a [8]bit) (out bit) {
	out = 1
	for i := 0; i < 8; i++ {
		if a[i] != 0 {
			out = 0
		}
	}

	return out
}
