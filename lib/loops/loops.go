package loops

import "fmt"

func CBasedForLoop() {

	var numbers []int

	for n := 0; n <= 30; n++ {
		numbers = append(numbers, n)

	}

	fmt.Printf("Array is %d", numbers)
}

func ConditionalFor() {

	for i := 1; i < 100; {
		fmt.Printf("Current Number Is : %d", i)
		i++
	}

}

func ForRange() {
	myDataStore := []string{
		"Malawi",
		"Zambia",
		"Zimbabwe",
		"South Africa",
		"Botswana",
		"Tanzania",
		"Mozambique",
		"Lesotho",
		"Swaziland",
	}

	for i, v := range myDataStore {
		fmt.Println(i, v)
	}

}
