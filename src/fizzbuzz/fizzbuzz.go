package fizzbuzz

import (
	"strconv"
)
func FizzBuzz(num int) string{
	var name = ""

	for i :=1; i < num+1; i ++ {
		if ModOfThree(i) && ModOfFive(i){
			name += "FizzBuzz"
		} else if ModOfThree(i) {
			name += "Fizz"
		} else if ModOfFive(i) {
			name += "Buzz"
		} else {
			name += strconv.Itoa(i)
		}
	}
	return name
}

func ModOfThree(x int) bool{
	if x%3 == 0{
		return true
	}
	return false
}

func ModOfFive(x int) bool{
	if x%5 == 0{
		return true
	}
	return false
}