package main

func mux_4_16(sel [4]bit, in [16][13]bit) (out [13]bit) {
	out = [13]bit{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for c := 0; c < 2; c++ {
				for d := 0; d < 2; d++ {
					if a == int(sel[0]) && b == int(sel[1]) && c == int(sel[2]) && d == int(sel[3]) {
						out = in[a*8+b*4+c*2+d]
					}
				}
			}
		}

	}

	return out
}

func first_8(in [13]bit) (out [8]bit) {
	out = [8]bit{in[0], in[1], in[2], in[3], in[4], in[5], in[6], in[7]}
	return out
}

func last_5(in [13]bit) (out [5]bit) {
	out = [5]bit{in[8], in[9], in[10], in[11], in[12]}
	return out
}
