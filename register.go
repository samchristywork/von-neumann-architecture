package main

import (
	"fmt"
)

type dFlipFlop struct {
	d       bit
	clk     bit
	prevClk bit
	q       bit
}

func risingEdge(prev bit, current bit) (out bit) {
	if prev == 0 && current == 1 {
		return 1
	}

	return 0
}

func (ff *dFlipFlop) update() {
	if risingEdge(ff.prevClk, ff.clk) == 1 {
		ff.q = ff.d
	}

	ff.prevClk = ff.clk
}

type register8bit struct {
	bits [8]dFlipFlop
}

func (reg *register8bit) load(data [8]bit, clk bit) {
	for i := 0; i < 8; i++ {
		reg.bits[i].d = data[i]
		reg.bits[i].clk = clk
		reg.bits[i].update()
	}
}

func (reg *register8bit) read() (data [8]bit) {
	for i := 0; i < 8; i++ {
		data[i] = reg.bits[i].q
	}
	return data
}

func (reg *register8bit) print() {
	for i := 0; i < 8; i++ {
		fmt.Printf("%v", reg.bits[i].q)
	}
	fmt.Printf("\n")
}
