package calculate

import (
    "fmt"
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

func isNotPrime(a string, divisor1 int, divisor2 int) {
    fmt.Printf("%s : not prime : %v * %v = %s\n", a, divisor1, divisor2, a)
}


func getDivisor(n int, i int) int {
    return int(n/i)
}

func getRemainder(n int, i int) int {
    return int(n%i)
}

func Calculate(n int) (isPrime bool, divisor1, divisor2 int, err error) {
    if n <= 0 {
        return false, 0, 0, nil
    }

    if n <= 3 {
        return true, 0, 0, nil
    }

    for i := int(n/2); i > 1; i-- {
        divisor2 := getDivisor(n, i)
        remainder := getRemainder(n, i)
        if remainder != 0 && i <= 2 {
            return true, 0, 0, nil
        }
        if remainder == 0 {
            return false, i, divisor2, nil
        }
    }
    err =  errors.New("An error occurred finding the prime. No condition met.")
    return false, 0, 0, err
}

func Operate() {
    for _, a := range os.Args[1:] {
        aAsInt, err := strconv.Atoi(a)
        if err != nil {
            checkIfNumber(a)
            continue
        }
        ifPrime, divisor1, divisor2, err :=  Calculate(aAsInt)
        for ifPrime != true {
            if ifPrime != true {
                isNotPrime(a, divisor1, divisor2)
                a = strconv.Atoi(divisor1)
                ifPrime, divisor1, divisor2, err = Calculate(a)
            } else {
                isPrimePrint(a)
            }
        }
}
