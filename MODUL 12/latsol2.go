package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan <= 0 {
		fmt.Println("masukkan bilangan bulat positif: ")
	}
	for bilangan > 0 {
		digit := bilangan % 10
		fmt.Println(digit)
		bilangan /= 10
	}
}