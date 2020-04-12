package fizzbuzz_test

import (
	"testing"

	"github.com/goignat/tigger/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		want string
		n    int
	}{
		{name: "1 given, '1' returned", want: "1", n: 1},
		{name: "2 given, '2' returned", want: "2", n: 2},
		{name: "3 given, 'Fizz' returned", want: "Fizz", n: 3},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			t.Parallel()

			got := fizzbuzz.FizzBuzz(tst.n)
			assert(t, tst.want, got)
		})
	}
}

func assert(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}
