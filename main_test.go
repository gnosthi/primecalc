package main

import (
    "os"
    "os/exec"
    "testing"
)

func TestMain(m *testing.M) {
    switch os.Getenv("GO_TEST_METHOD") {
    case "":
        // default testcase
        os.Exit(m.Run())
    }
}

func TestEcho(t *testing.T) {
    cmd := exec.Command(os.Args[0])
    cmd.Env = []string{"GO_TEST_METHOD=echo"}
    output, err := cmd.Output()
    if err != nil {
        t.Errorf("echo: %v\n", err)
    }
    if r, e := string(output), ""; r != e {
        t.Errorf("echo: No output from command\n")
    }
}
