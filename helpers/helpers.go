package helpers

import (
    "fmt"
    "os"
)

func PrintHelp() {
    cmd := os.Args[0]
    fmt.Printf("%v is used to test a number or a set of numbers and check if it is a prime or not\nUsage; %v number1 number2 number3 ...\n", cmd, cmd)
    os.Exit(1)
}
