package fizzbuzz

import (
	"strconv"
)
func FizzBuzz(num int) string{
	var name = ""

	for i :=1; i < num+1; i ++ {
		if i%3 == 0 {
			name += "Fizz"
		}else{
			name += strconv.Itoa(i)
		}
	}
	return name
}
