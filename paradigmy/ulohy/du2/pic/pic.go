package main

func Pic(dx, dy int) [][]uint8 {
	rv := make([][]uint8, dy)
	for i := range rv {
		rv[i] = make([]uint8, dx)
	}
	return rv
}

func main() {

}
