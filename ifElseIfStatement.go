package main

import "fmt"

func activity3() {
	var input int
	fmt.Println("Input a number!")
	fmt.Scanln(&input)
	a := input%5 == 0
	b := input%6 == 0
	if a && b {
		fmt.Println("This number is divisible by 5 and 6")
	} else {
		fmt.Println("This number is not divisible by 5 and 6")
	}

}
func main() {
	activity3()
}
