package main

import "fmt"

func main() {
	var sample string = "123"

	sample2 := "234"

	var input string

	const number int8 = 1

	var number2 uint8 = 255

	fmt.Printf("Hello World, %v, %v, %v, %v", sample, number, sample2, number2)

	fmt.Println("")

	fmt.Printf("The type of number is %T", number)

	fmt.Println("")

	fmt.Scan(&input)

	fmt.Println("The input is: ")

	fmt.Println(input)

	fmt.Println("The address of input is: ")

	fmt.Println(&input)
}
