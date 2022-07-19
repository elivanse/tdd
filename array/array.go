package main

func Sum(numbers [5]int) int {
	sum := 0
	// con for como no usamos la var i
	// le decimos al roedor que de vueltas
	// hasta el tamano del array
	// que es fijo en go
	for _, number := range numbers {
		sum += number
	}
	return sum
}
