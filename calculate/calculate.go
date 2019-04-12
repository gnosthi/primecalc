package calculate

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Ensure that the passed argument is an actual number.
func checkIfNumber(a string) (bool, error) {
	if _, err := strconv.Atoi(a); err != nil {
		err = fmt.Errorf("%s is not a number...skipping", a)
		return false, err
	}
	return true, nil
}

func checkifPositive(n int64) (bool, error) {
	if n <= 0 {
		err := fmt.Errorf("%v is not a positive integer...skipping", n)
		return false, err
	}
	return true, nil
}

func isPrimePrint(a string) {
	fmt.Printf("%s : prime\n", a)
}

func isNotPrime(a string, divisor1 int64, divisor2 int64) {
	fmt.Printf("%s : not prime : %v * %v = %s\n", a, divisor1, divisor2, a)
}

func getDivisor(n int64, i int64) int64 {
	return int64(n / i)
}

func getRemainder(n int64, i int64) int64 {
	return int64(n % i)
}

// Calculate takes n as unsigned int and returns whether n is prime
func Calculate(n int64) (isPrime bool, divisor1, divisor2 int64, err error) {
	if n <= 0 {
		return false, 0, 0, nil
	}

	if n <= 3 {
		return true, 0, 0, nil
	}

	for i := int64(n / 2); i > 1; i-- {
		divisor2 := getDivisor(n, i)
		remainder := getRemainder(n, i)
		if remainder != 0 && i <= 2 {
			return true, 0, 0, nil
		}
		if remainder == 0 {
			return false, i, divisor2, nil
		}
	}
	// This should never happen.
	err = errors.New("an error occurred finding the prime. No condition met")
	return false, 0, 0, err
}

// Operate entry point for the calculate package
func Operate() {
	for _, a := range os.Args[1:] {
		aAsInt, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			checkifPositive(aAsInt)
			continue
		}
		_, err = checkIfNumber(a)
		if err != nil {
			fmt.Errorf("%s is not a number", a)
			continue
		}
		_, err = checkifPositive(aAsInt)
		if err != nil {
			fmt.Errorf("%v is not a positive integer", aAsInt)
			continue
		}
		ifPrime, divisor1, divisor2, err := Calculate(aAsInt)
		for ifPrime != true {
			if ifPrime != true {
				isNotPrime(a, divisor1, divisor2)
				a = strconv.Itoa(int(divisor1))
				ifPrime, divisor1, divisor2, err = Calculate(divisor1)
			}
		}
		isPrimePrint(a)
		fmt.Println("------------------------------------------------")
	}
}
