package main

import (
	"fmt"
	"learninggo/lib/conditions"
	"learninggo/lib/data-arrays"
	"learninggo/lib/data-maps"
	"learninggo/lib/data-slices"
	"learninggo/lib/functions"
	"learninggo/lib/loops"
	"learninggo/lib/mystruct"
)

func main() {
	fmt.Println("Welcome to my new Go Program")

	//structs
	mystruct.CreateStructs()

	//Slices
	resp := data_slices.CreateSlice()
	fmt.Println(resp)

	//Basic Function
	functions.Sum(1, 3)

	//CBasedForLoop
	loops.CBasedForLoop()
	loops.ConditionalFor()
	loops.ForRange()

	//Create Maps
	data_maps.CreateMaps()

	//Data Arrays
	data_arrays.CreateArray()

	//Conditions
	conditions.CheckCondition()
	conditions.FizzBuzz(100)

}
