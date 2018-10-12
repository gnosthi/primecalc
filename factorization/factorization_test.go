package factorization

import (
	"fmt"
    "testing"
)
var testfactors = []struct{
	divisor string
	factors []string
	sum string
}{
	{"23",[]string{"2","2","31"},"2852"},
	{"31",[]string{"3","3","31"}, "8649"},
	{"11", []string{"2","11"}, "242"},
	{"1607",[]string{"3","2","2"}, "19284"},
	{"2", []string{"2","2","2","2","2","2","2","2","2"}, "1024"},

}
func TestPrimeFactors(t *testing.T) {
	if fmt.Sprintf("%v", PrimeFactors(23)) != `[23]` {
		t.Error(23)
	}
	if fmt.Sprintf("%v", PrimeFactors(12)) != `[2 2 3]` {
		t.Error(12)
	}
	if fmt.Sprintf("%v", PrimeFactors(360)) != `[2 2 2 3 3 5]` {
		t.Error(360)
	}
	if fmt.Sprintf("%v", PrimeFactors(97)) != `[97]` {
		t.Error(97)
	}
}

func TestFactorMultiplication(t *testing.T) {
	// Test string construction
	for _, tt := range testfactors {
		t.Run(tt.divisor, func(t *testing.T) {
			value := FactorMultiplication(tt.divisor,tt.factors,tt.sum)
			if value == "" {
				t.Errorf("error stringing factors, no string returned")
			}
			expected := tt.divisor + " * "
			for _,i := range tt.factors {
				if i == tt.factors[len(tt.factors)-1] {
					expected = expected + i + " = " + tt.sum
					break
				}
				expected = expected + i + " * "
			}
			if value != expected {
				t.Errorf("wanted %s, got %s", expected, value)
			}
		})
	}
	// Test empty string conditions
	for i := 0; i < len(testfactors); i++ {
		value := FactorMultiplication("2",[]string{""},"2")
		if value != "" {
			t.Errorf("expected empty string, got %s", value)
		}
	}
}
