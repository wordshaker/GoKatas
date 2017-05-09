package fizzbuzz

import (
	"strconv"
)
func FizzBuzz(num int) string{
	var name = ""

	for i :=1; i < num+1; i ++ {
		if i%3 == 0 && i%5 == 0{
			name += "FizzBuzz"
		} else if i%3 == 0 {
			name += "Fizz"
		} else if i%5 == 0 {
			name += "Buzz"
		} else {
			name += strconv.Itoa(i)
		}
	}
	return name
}