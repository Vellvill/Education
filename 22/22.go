package main

import (
	"fmt"
	"math/big"
)

type bigInts struct {
	a *big.Int
	b *big.Int
}

func main() {
	ints := &bigInts{
		a: new(big.Int),
		b: new(big.Int),
	}
	ints.a, ints.b = ints.a.SetInt64(922337203685477580), ints.b.SetInt64(922337203685477580)
	fmt.Println(ints.add(), ints.div(), ints.sub(), ints.multiple())
}

func (b *bigInts) add() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

func (b *bigInts) multiple() *big.Int {
	return new(big.Int).Mul(b.a, b.b)
}

func (b *bigInts) sub() *big.Int {
	return new(big.Int).Sub(b.a, b.b)
}

func (b *bigInts) div() *big.Int {
	return new(big.Int).Quo(b.a, b.b)
}
