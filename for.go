package main

import "fmt"

func main(){
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Practice With For Loops")
	fmt.Println()

	i := 1
	for i <= 3 {
		fmt.Println("i =", i)
		i++
	}

	fmt.Println()

	for j:=3; j<=9; j++ {
		fmt.Println("j =", j)
	}

	fmt.Println()

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println()

	for n:=0; n<=5; n++ {
		if n%2==0 {
			continue
		}
		fmt.Println("n =", n)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}