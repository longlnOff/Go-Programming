package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	intA := math.MaxInt64
	intA += 1

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println(intA)
	fmt.Println(bigA)
	fmt.Println(math.MaxInt64)
}