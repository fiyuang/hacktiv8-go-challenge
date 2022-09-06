package firstChallange

import "fmt"

func fizzbuzzSecond() {
	n := 30
	i := 1

	for {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
		i++
		if i > n {
			break
		}
	}
}
