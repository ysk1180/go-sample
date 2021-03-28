package culc_test

import (
	"go-sample/test-sample/culc"
	"testing"
)

var flagtests = []struct {
	in  int
	out bool
}{
	{4, true},
	{98, true},
	{99, false},
}

func TestIsEven(t *testing.T) {
	for _, test := range flagtests {
		result := culc.IsEven(test.in)
		if result != test.out {
			t.Errorf("expect: %t but actual: %t, test in: %d", test.out, result, test.in)
		}
	}
}
