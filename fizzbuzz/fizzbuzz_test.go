package fizzbuzz_test

import (
	"testing"

	"github.com/goignat/tigger/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("1 given, '1' returned", func(t *testing.T) {
		got := fizzbuzz.FizzBuzz(1)
		if "1" != got {
			t.Fatalf("want 1, got %q", got)
		}
	})

	t.Run("2 given, '1' returned", func(t *testing.T) {
		got := fizzbuzz.FizzBuzz(2)
		if "2" != got {
			t.Fatalf("want 2, got %q", got)
		}
	})
}
