package main

import (
	"math"
	"math/big"
)

func main() {
	a := new(big.Int).SetUint64(uint64(math.MaxUint64))
	b := new(big.Int).SetUint64(uint64(math.MaxUint64))
	c := new(big.Int).Mul(a, b)

	var result uint64

	println(c.BitLen())

	// check for overflow
	if c.BitLen() > 64 {
		result = uint64(math.MaxUint64)
	}

	result = c.Uint64()

	println(result)
}
