package factorization

import "fmt"

func PrimeFactors(n int) (pfs []int) {
    for n%2 == 0 {
        pfs = append(pfs, 2)
        n = n / 2
    }

    // n must be odd at this point
    for i := 3; i*i <= n; i = i + 2 {
        for n%i == 0 {
            pfs = append(pfs, i)
            n = n / i
        }
    }

    // This condition is to handle the case when n is a prime number
    // greater than 2
    if n > 2 {
        pfs = append(pfs,n)
    }
    return
}

func FactorMultiplication(primeDivisor string, factors []string, sum string) string {
    var factorial string
    if len(factors) > 1 {
        factorial = primeDivisor + " * "
        for _,i := range factors {
            if i == factors[len(factors)-1] {
                factorial = factorial + i + " = " + sum
                break
            }
            factorial = factorial + i + " * "
        }
        fmt.Println(factorial)
        return factorial
    }
    return ""
}