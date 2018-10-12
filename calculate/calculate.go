package calculate

import (
    "fmt"
    "errors"
    "strconv"
    "os"
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
    // This should never happen.
    err =  errors.New("an error occurred finding the prime. No condition met")
    return false, 0, 0, err
}

// TODO Issue running factorization in Operate(), No output seems to occur. - 10/11/2018 (Commented out relevant code for now)

func Operate() {
    // var factor []string
    for _, a := range os.Args[1:] {
        //copyOfA := a
        aAsInt, err := strconv.Atoi(a)
        if err != nil {
            checkIfNumber(a)
            continue
        }
        ifPrime, divisor1, divisor2, err := Calculate(aAsInt)
        for ifPrime != true {
            /* d2ToA := strconv.Itoa(divisor2)
        	factor = append(factor, d2ToA) */
            if ifPrime != true {
                isNotPrime(a, divisor1, divisor2)
                a = strconv.Itoa(divisor1)
                ifPrime, divisor1, divisor2, err = Calculate(divisor1)
            }
        }
        // text := factorization.FactorMultiplication(a,factor,copyOfA)
        // fmt.Printf("%s\n", text)
        isPrimePrint(a)
        fmt.Println("------------------------------------------------")
    }
}
