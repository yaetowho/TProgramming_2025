package main

import (
	"fmt"
	"math"
)

func main() {
	//   fmt.Println("Журавлев Сергей Викторович")
	//   fmt.Println("Значения функции arcsin(x^a) + arccos(x^b):", calculate(2, 3, 0.08))
	fmt.Println("Задача А:")
	taskA()
}

func calculate(a, b, x float64) float64 {
	arcsin := math.Asin(math.Pow(x, a))
	arccos := math.Acos(math.Pow(x, b))
	result := arcsin + arccos
	return result
}

func taskA() {
	xn, xe, dx := 0.11, 0.36, 0.05

	for x := xn; x < xe; x += dx {
		fmt.Println(calculate(2.0, 3.0, x))
	}
}
