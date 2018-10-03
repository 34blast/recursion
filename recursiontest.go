package main

import (
	"fmt"
)

func main() {
	fmt.Println("recursiontest.main() : starting exectuion")
	fmt.Println()

	fmt.Println("recursiontest.main() : do factorial")
	var idx1 int = 3
	fmt.Printf("Factorial of %d is %d\n", idx1, factorial(idx1))
	fmt.Println()

	fmt.Println("recursiontest.main() : do fibonaci sequence")
	var idx2 int
	var fibonaciMax = 11
	for idx2 = 0; idx2 < fibonaciMax; idx2++ {
		fmt.Printf("%d ", fibonaci(idx2))
	}
	fmt.Println()
	fmt.Println("recursiontest.main() : ending  exectuion")

}

func factorial(i int) int {
	//	fmt.Println("recursiontest.factorial() : entry i = ", i)
	var calculated int
	if i <= 1 {
		calculated = 1
	} else {
		calculated = i * factorial(i-1)
	}
	//	fmt.Println("recursiontest.factorial() : exit i = ", i, ", calculated = ", calculated)
	return calculated
}

func fibonaci(i int) (ret int) {
	// fmt.Println("recursiontest.fibonaci() : entry i = ", i)
	var calculated int
	if i == 0 {
		calculated = 0
	} else if i == 1 {
		calculated = 1
	} else {
		calculated = fibonaci(i-1) + fibonaci(i-2)
	}
	//n  fmt.Println("    recursiontest.fibonaci() : exit i = ", i, ", calculated = ", calculated)
	return calculated
}
