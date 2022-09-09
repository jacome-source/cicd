package main

import "fmt"

func main() {
	fmt.Println(soma(1, 2, 3))
	fmt.Println(multiplica(1, 2, 3))
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
