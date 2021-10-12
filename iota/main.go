package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
	w
)

const v = iota

func main() {
	fmt.Printf("x=%d\ny=%d\nz=%d\nw=%d\nv=%d\n", x, y, z, w, v)
}
