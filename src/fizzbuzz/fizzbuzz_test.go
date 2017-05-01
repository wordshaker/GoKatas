package fizzbuzz

import (
	"testing"
)

func TestFizz(t *testing.T){
	var expected = "12Fizz"
	var actual = FizzBuzz(3)
	if actual != expected{
		t.Error("FB Fail. Expected:" + expected + " Actual :" + actual)
	}
}

func TestBuzz(t *testing.T){
	var expected = "12Fizz4Buzz"
	var actual = FizzBuzz(5)
	if actual != expected{
		t.Error("FB Fail. Expected:" + expected + " Actual :" + actual)
	}
}