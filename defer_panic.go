package main

import (
	"fmt"
)

func endApp() {
	if message := recover(); message != nil {
		fmt.Println("error dengan message:", message)
		//os.Exit(1)
	}
	fmt.Println("Aplikasi selesai")
}

func divide(number1 int, number2 int) float64 {
	defer endApp()
	if number2 == 0 {
		panic("Dilarang dibagi dengan 0")
	}
	return float64(number1) / float64(number2)
}
