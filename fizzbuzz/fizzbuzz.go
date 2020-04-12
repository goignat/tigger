package fizzbuzz

import "fmt"

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func DisplayResult() {
	for n := 1; n <= 100; n++ {
		fmt.Println(FizzBuzz(n))
	}
}
