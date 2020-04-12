package fizzbuzz_test

import (
	"testing"

	"github.com/goignat/tigger/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("1 given, '1' returned", func(t *testing.T) {
		got := fizzbuzz.FizzBuzz(1)
		assert(t, "1", got)
	})

	t.Run("2 given, '2' returned", func(t *testing.T) {
		got := fizzbuzz.FizzBuzz(2)
		assert(t, "2", got)
	})

	t.Run("3 given, 'Fizz' returned", func(t *testing.T) {
		got := fizzbuzz.FizzBuzz(3)
		assert(t, "Fizz", got)
	})

}

func assert(t *testing.T, want, got string) {
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}
