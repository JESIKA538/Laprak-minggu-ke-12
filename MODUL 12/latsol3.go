package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	if x < y || y <= 0 {
		fmt.Println("input tidak valid")
		fmt.Scan(&x, &y)
	}
	result := 0
	temp := x
	for temp >= y {
		temp -= y
		result++
	}
	fmt.Println(result)
}
