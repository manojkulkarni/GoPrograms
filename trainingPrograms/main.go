package main

import "fmt"

func CheckNumberIfPrime(CheckNumber int) bool {
	if CheckNumber <= 1 {
		return false
	}
	for num := 2; num <= CheckNumber-1; num++ {
		if CheckNumber%num == 0 {
			return false
		}
	}
	return true
}

func GetNumberFactorial(CheckNumber int) int {
	if CheckNumber < 0 {
		fmt.Println("Why nonsense thing? Factorial of negative number, try something else.")
		return 0
	}
	factorialNumber := 1
	for num := 1; num <= CheckNumber; num++ {
		factorialNumber = factorialNumber * num
	}
	return factorialNumber
}

func main() {
	numbersForPrimeCheck := []int{0, 1, 2, 5, 7, 8, 11, 15, 17, 18, 19, 20, 21}
	for _, num := range numbersForPrimeCheck {
		fmt.Printf("Is %d Prime ? Ans : %t \n", num, CheckNumberIfPrime(num))
	}
	fmt.Println("******************************************************")
	numbersToGetFactorial := []int{0, 1, 2, 5, -4, 7, 3, 6, 9, 8}
	for _, num := range numbersToGetFactorial {
		fmt.Printf("%d factorial = %d \n", num, GetNumberFactorial(num))
	}
}
