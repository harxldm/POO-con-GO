package main

import "fmt"

type numerosf struct {
	n1 float64
	n2 float64
}

type numeros struct {
	n1 int
	n2 int
	n3 numerosf
}

func (res numeros) suma(n1, n2 int) int {
	return n1 + n2

}

func (res numeros) resta(n1, n2 int) int {
	return n1 + n2

}

func main() {
	Go := numeros{
		n1: 1,
		n2: 2,
	}

	fmt.Println(Go.suma(Go.n1, Go.n2))
	fmt.Println(Go.resta(Go.n1, Go.n2))
}
