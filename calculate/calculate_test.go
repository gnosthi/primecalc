package calculate

import (
	"os"
	"strconv"
	"testing"
)

var calculatetest = []struct {
	number   string
	prime    bool
	divisor  int64
	divisor2 int64
	err      error
}{
	{"-1", false, 0, 0, nil},
	{"1", true, 0, 0, nil},
	{"2", true, 0, 0, nil},
	{"3", true, 0, 0, nil},
	{"4", false, 2, 2, nil},
	{"5", true, 0, 0, nil},
	{"6", false, 3, 2, nil},
	{"7", true, 0, 0, nil},
	{"8", false, 4, 2, nil},
	{"9", false, 3, 3, nil},
	{"10", false, 5, 2, nil},
	{"11", true, 0, 0, nil},
	{"12", false, 6, 2, nil},
	{"13", true, 0, 0, nil},
	{"14", false, 7, 2, nil},
	{"15", false, 5, 3, nil},
	{"16", false, 8, 2, nil},
	{"17", true, 0, 0, nil},
	{"18", false, 9, 2, nil},
	{"19", true, 0, 0, nil},
	{"31", true, 0, 0, nil},
	{"24", false, 12, 2, nil},
	{"111", false, 37, 3, nil},
}

var divisortest = []struct {
	in  string
	out int
}{
	{"1", 0},
	{"10", 5},
	{"20", 10},
	{"7", 0},
	{"21", 7},
	{"33", 11},
	{"93", 31},
	{"111", 37},
}

var nantest = []struct {
	in  string
	out bool
}{
	{"1", true},
	{"1z", false},
	{"432#", false},
	{"abcd", false},
	{"1029381", true},
	{"abc123", false},
	{"93", true},
}

func TestCalculate(t *testing.T) {
	for _, tt := range calculatetest {
		t.Run(tt.number, func(t *testing.T) {
			n, _ := strconv.Atoi(tt.number)
			isPrime, divisor, divisor2, err := Calculate(int64(n))
			if isPrime != tt.prime {
				t.Errorf("Error: Expected isPrime value of %v got %v", tt.prime, isPrime)
			}
			if divisor != tt.divisor {
				t.Errorf("Error: Expected divisor value of %v got %v", tt.divisor, divisor)
			}
			if divisor2 != tt.divisor2 {
				t.Errorf("Error: Expected divisor2 value of %v got %v", tt.divisor2, divisor2)
			}
			if err != tt.err {
				t.Errorf("Error: Expected err value of %v got %v", tt.err, err)
			}
		})
	}
}

func BenchmarkCalculate(b *testing.B) {
	for _, tt := range calculatetest {
		b.Run(tt.number, func(b *testing.B) {
			n, _ := strconv.Atoi(tt.number)
			isPrime, divisor, divisor2, err := Calculate(int64(n))
			if isPrime != tt.prime {
				b.Errorf("Error: Expected isPrime value of %v got %v", tt.prime, isPrime)
			}
			if divisor != tt.divisor {
				b.Errorf("Error: Expected divisor value of %v got %v", tt.divisor, divisor)
			}
			if divisor2 != tt.divisor2 {
				b.Errorf("Error: Expected divisor2 value of %v got %v", tt.divisor2, divisor2)
			}
			if err != tt.err {
				b.Errorf("Error: Expected err value of %v got %v", tt.err, err)
			}

		})
	}
}

func TestCheckIfNumber(t *testing.T) {
	for _, tt := range nantest {
		t.Run(tt.in, func(t *testing.T) {
			isNum, _ := checkIfNumber(tt.in)
			if isNum != tt.out {
				t.Errorf("Error: Expected checkIfNumber to report %v, got %v instead", tt.out, isNum)
			}
		})
	}
	for i := 0; i < 1000; i++ {
		isNum, err := checkIfNumber("a")
		if err != nil {
			continue
		}
		if isNum == true {
			t.Errorf("Error: Expected false, got %t", isNum)
		}
	}
	for i := 0; i < 1000; i++ {
		isNum, _ := checkIfNumber(strconv.Itoa(i))
		if isNum != true {
			t.Errorf("Error: Expected true, got %t", isNum)
		}
	}
}

func BenchmarkCheckIfNumber(b *testing.B) {
	for _, tt := range nantest {
		b.Run(tt.in, func(b *testing.B) {
			isNum, _ := checkIfNumber(tt.in)
			if isNum != tt.out {
				b.Errorf("Error: Expected checkIfNumber to report %v, got %v instead", tt.out, isNum)
			}
		})
	}
	for i := 0; i < b.N; i++ {
		isNum, err := checkIfNumber("a")
		if err != nil {
			continue
		}
		if isNum == true {
			b.Errorf("Error: Expected false, got %t", isNum)
		}
	}
	for i := 0; i < b.N; i++ {
		isNum, _ := checkIfNumber(strconv.Itoa(i))
		if isNum != true {
			b.Errorf("Error: Expected true, got %t", isNum)
		}
	}
}

func TestIsPrimePrint(t *testing.T) {
	for _, tt := range calculatetest {
		t.Run(tt.number, func(t *testing.T) {
			isPrimePrint(tt.number)
			isNotPrime(tt.number, tt.divisor, tt.divisor2)
		})
	}
	for i := 0; i < 1000; i++ {
		isPrime, divisor, divisor2, err := Calculate(int64(i))
		if err != nil {
			t.Errorf("Error Occurred")
		}
		if isPrime == true {
			isPrimePrint(strconv.Itoa(i))
		}
		if isPrime == false {
			isNotPrime(strconv.Itoa(i), divisor, divisor2)
		}
	}
}

func BenchMarkIsPrimePrint(b *testing.B) {
	for _, tt := range calculatetest {
		b.Run(tt.number, func(b *testing.B) {
			isPrimePrint(tt.number)
			isNotPrime(tt.number, tt.divisor, tt.divisor2)
		})
	}
	for i := 0; i < b.N; i++ {
		isPrime, divisor, divisor2, err := Calculate(int64(i))
		if err != nil {
			b.Errorf("Error Occured")
		}
		if isPrime == true {
			isPrimePrint(strconv.Itoa(i))
		}
		if isPrime != true {
			isNotPrime(strconv.Itoa(i), divisor, divisor2)
		}
	}
}

func TestOperate(t *testing.T) {
	os.Args = []string{"1", "111", "93", "abc", "31", "101", "891", "777", "91", "118", "20394", "19283", "1928737", "123", "31", "666"}
	Operate()
	for i := 0; i < 1000; i++ {
		os.Args = []string{strconv.Itoa(i)}
		Operate()
	}
	for i := 0; i < 1000; i++ {
		os.Args = []string{"a"}
		Operate()
	}
}

func BenchmarkOperate(b *testing.B) {
	os.Args = []string{"1", "111", "93", "abc", "31", "101", "891", "777", "91", "118", "20394", "19283", "1928737", "123", "31", "666"}
	Operate()
	for i := 0; i < b.N; i++ {
		os.Args = []string{strconv.Itoa(i)}
		Operate()
	}
	for i := 0; i < b.N; i++ {
		os.Args = []string{"a"}
		Operate()
	}
}
