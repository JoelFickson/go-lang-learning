package conditions

import (
	"fmt"
	"math/rand"
)

func CheckCondition() {
	/*
		Check basic conditions with if-else
	*/

	if dob := rand.Intn(10); dob > 1994 {
		fmt.Println("You are an oldie")
	} else {
		fmt.Println("Still Young")
	}
}

func FizzBuzz(limit int) {

	for i := 0; i < limit; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}

	}

}
