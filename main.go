package main
import (
	"fmt" 
	"errors"
)

func  main() {
	printMe("Celestial")

	var res int = multipleNum(10, 20)
	fmt.Println(res)

	var devision, mod, err = intDivision(10, 2)

	// ! Using if statement
	// if err != nil {
	// 	fmt.Println(err)
	// } else if mod == 0 {
	// 	fmt.Printf("Result of devision is %v", devision)
	// } else {
	// 	fmt.Printf("Devision: %d, Mod: %d", devision, mod)
	// }

	// ! Using switch statement
	switch {
		case err != nil:
			fmt.Println(err)
		case mod == 0:
			fmt.Printf("Result of devision is %v", devision)
		default:
			fmt.Printf("Devision: %d, Mod: %d", devision, mod)
	}
}

func printMe(fullname string) {
	fmt.Println("Hello, "+fullname+"!")
}

func multipleNum(num1 int, num2 int) int {
	return num2 * num1
}

func intDivision(num1 int, num2 int) (int, int, error) {
	var err error

	if num2 == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var devision int = num1 / num2
	var mod int = num1 % num2

	return devision, mod, err
}