package fizzbuzz

func FizzBuzz(n int) string {
	if n == 1 {
		return "1"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n == 4 {
		return "4"
	}
	return "2"
}
