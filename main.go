package main

import "fmt"

func factorial(n int) {
	if(n < 1){
		fmt.Println("Invalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++{
		result *= i
	}
	fmt.Println(n, "-", result)
}

func main() {
	fmt.Println("Run app")

	for i := 1; i <= 10; i++ {
		go factorial(i)
	}

	fmt.Scanln()
	fmt.Println("The End")
}
