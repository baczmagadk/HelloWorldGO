package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println("Practice With Constants")
	fmt.Println()

	fmt.Println("s =", s)

	const n float64 = 500000000
	fmt.Println("n =", n)

	const d float64 = 3e20 / n

	fmt.Println("3e20 / n =", d)
	fmt.Println("int64(d) =", int64(d))

	fmt.Println("math.Sin(n) =", math.Sin(n))

}