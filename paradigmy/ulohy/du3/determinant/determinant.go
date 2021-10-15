package main

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	N      int
	Values [][]int
}

func remove(mat Matrix, n int) Matrix {
	var rv Matrix
	rv.N = mat.N - 1
	rv.Values = make([][]int, rv.N)

	for i := 1; i <= rv.N; i++ {
		col := 0
		rv.Values[i-1] = make([]int, rv.N)
		for j := 0; j <= rv.N; j++ {
			if j == n {
				col = 1
				continue
			}
			rv.Values[i-1][j-col] = mat.Values[i][j]
		}
	}
	return rv
}

func det(mat Matrix) int64 {

	//Pri vypocte sa spusta N! gorutines cize moze nastat problem s prekrocenim povoleneho mnozstva gorutines

	if mat.N == 2 {
		return int64(mat.Values[0][0]*mat.Values[1][1] - mat.Values[0][1]*mat.Values[1][0])
	}

	stack := make(chan int64, mat.N)

	for i := 0; i < mat.N; i++ {
		con := int64(1)
		if i%2 == 1 {
			con = -1
		}
		go func(col int, con int64) {
			stack <- int64(int64(mat.Values[0][col]) * con * det(remove(mat, col)))
		}(i, con)
	}
	rv := int64(0)
	for i := 0; i < mat.N; i++ {
		rv += <-stack
	}

	return rv

}

func main() {
	var mat Matrix
	mat.N = 10
	arr := make([][]int, mat.N)
	for i := 0; i < mat.N; i++ {
		arr[i] = make([]int, mat.N)
		for j := 0; j < mat.N; j++ {
			arr[i][j] = rand.Intn(10)
		}
	}
	mat.Values = arr
	fmt.Println(arr)
	fmt.Println(det(mat))
}
