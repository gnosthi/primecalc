package main

import (
    "os"
    "github.com/gnosthi/primecalc/helpers"
    "github.com/gnosthi/primecalc/calculate"
)

func main() {
    if len(os.Args[1:]) == 0 {
        helpers.PrintHelp()
    }
    calculate.Operate()
}
