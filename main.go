package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Журавлев Сергей Викторович")
	fmt.Println("Значения функции arcsin(x^a) + arccos(x^b):")
	fmt.Println(calculate(2, 3, 0.8))
}

func calculate(a, b, x float64) (float64 float64) {
	arcsin := math.Asin(math.Pow(x, a))
	arccos := math.Acos(math.Pow(x, b))
	result := arcsin + arccos

	return result
}
