package main

import (
	"fmt"
	"math/big"
	"time"
)

func matrixMulti(x [][]*big.Int, y [][]*big.Int) [][]*big.Int {
	rv := [][]*big.Int{
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0)},
	}
	i00 := make(chan *big.Int)
	i01 := make(chan *big.Int)
	i10 := make(chan *big.Int)
	i11 := make(chan *big.Int)

	go func() {
		i00 <- big.NewInt(0).Add(big.NewInt(0).Mul(x[0][0], y[0][0]), big.NewInt(0).Mul(x[1][0], y[0][1]))
	}()
	go func() {
		i01 <- big.NewInt(0).Add(big.NewInt(0).Mul(x[0][0], y[1][0]), big.NewInt(0).Mul(x[1][0], y[1][1]))
	}()
	go func() {
		i10 <- big.NewInt(0).Add(big.NewInt(0).Mul(x[0][1], y[0][0]), big.NewInt(0).Mul(x[1][1], y[0][1]))
	}()
	go func() {
		i11 <- big.NewInt(0).Add(big.NewInt(0).Mul(x[0][1], y[1][0]), big.NewInt(0).Mul(x[1][1], y[1][1]))
	}()

	rv[0][0] = <-i00
	rv[0][1] = <-i01
	rv[1][0] = <-i10
	rv[1][1] = <-i11

	return rv

}

func fib(n int64) [][]*big.Int {
	mat := [][]*big.Int{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}
	if n == 1 {
		return mat
	}

	num := int64(1)

	for 2*num <= n {
		mat = matrixMulti(mat, mat)
		num *= 2
	}

	if num != n {
		return matrixMulti(mat, fib(n-num))
	}
	return mat

}

func main() {
	startTime := time.Now()
	n := int64(1234567890) //sem zadaj N
	n -= 1
	res := fib(n)[0][0]
	fmt.Printf("compute time=%v\n", time.Since(startTime))
	fmt.Printf("dlzka: %d\n", len(res.String()))
	fmt.Print("======================================")
	fmt.Print(res)
}
