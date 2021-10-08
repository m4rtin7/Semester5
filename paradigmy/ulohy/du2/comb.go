package main

import (
	"fmt"
	"math/big"
	"time"
)

func FaktorialMulRange(n int64) *big.Int {
	var x = big.NewInt(0)
	x.MulRange(2, n)
	return x
}

func combBig(n int64, k int64) *big.Int {
	nFact := make(chan *big.Int)
	kFact := make(chan *big.Int)
	nkFact := make(chan *big.Int)
	multNkK := make(chan *big.Int)

	go func() { nFact <- FaktorialMulRange(n) }()
	go func() { kFact <- FaktorialMulRange(k) }()
	go func() { nkFact <- FaktorialMulRange(n - k) }()
	go func() {
		x := <-kFact
		y := <-nkFact
		multNkK <- big.NewInt(0).Mul(x, y)
	}()

	nF := <-nFact
	nkF := <-multNkK

	return big.NewInt(0).Div(nF, nkF)
}

func main() {
	startTime := time.Now()
	res := combBig(100000, 50000)
	fmt.Printf("compute time=%v\n", time.Since(startTime))
	fmt.Printf("dlzka: %d\n", len(res.String()))
}
