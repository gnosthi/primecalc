package main

import (
    "os"
    "testing"
)

func TestMain(m *testing.M) {
    switch os.Getenv("GO_TEST_METHOD") {
    case "":
        // default testcase
        os.Exit(m.Run())
    }
}

func TestMainFunc(t *testing.T)  {
    os.Args = []string{"1", "3", "31", "93", "101", "111", "120931"}
    main()
}

func BenchMarkMainFunc(b *testing.B) {
    os.Args = []string{"1", "3", "31", "93", "101", "111", "120931"}
    main()
}

func TestMainPrintHelp(t *testing.T) {
    main()
}

func BenchmarkMainPrintHelp(b *testing.B) {
    main()
}
