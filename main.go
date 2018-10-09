package main

import (
    "os"
    "github.com/gnosthi/primecalc/helpers"
)

func main() {
    if len(os.Args[1:]) == 0 {
        helpers.PrintHelp()
    }
}
