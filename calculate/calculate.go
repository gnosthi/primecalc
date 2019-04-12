package calculate

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Ensure that the passed argument is an actual number.
func checkIfNumber(a string) bool {
	if _, err := strconv.Atoi(a); err != nil {
		fmt.Errorf("%s is not a number...skipping\n", a)
		return false
	}
	return true
}

func isPrimePrint(a string) {
	fmt.Printf("%s : prime\n", a)
}

func isNotPrime(a string, divisor1 uint64, divisor2 uint64) {
	fmt.Printf("%s : not prime : %v * %v = %s\n", a, divisor1, divisor2, a)
}

func getDivisor(n uint64, i uint64) uint64 {
	return uint64(n / i)
}

func getRemainder(n uint64, i uint64) uint64 {
	return uint64(n % i)
}

// Calculate takes n as unsigned int and returns whether n is prime
func Calculate(n uint64) (isPrime bool, divisor1, divisor2 uint64, err error) {
	if n <= 0 {
		return false, 0, 0, nil
	}

	if n <= 3 {
		return true, 0, 0, nil
	}

	for i := uint64(n / 2); i > 1; i-- {
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
		aAsInt, err := strconv.ParseUint(a, 10, 64)
		if err != nil {
			checkIfNumber(a)
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
