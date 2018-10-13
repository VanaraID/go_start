package main

import "fmt"

func tambah(a, b int) int {
	return a + b
}

func bagi(a, b int) int {
	return a / b
}

func main() {
	var a, b int

	fmt.Print("Enter a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter b: ")
	fmt.Scanln(&b)

	d := bagi(a, b)
	c := tambah(a, b)
	fmt.Println("Hasil tambah: ", c)
	fmt.Println("Hasil bagi: ", d)
}
