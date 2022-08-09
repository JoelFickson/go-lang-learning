package data_maps

func CreateMaps() map[string]int {

	mySecondMap := make(map[string]int)
	mySecondMap["James"] = 20

	print(mySecondMap)

	myMap := map[string]int{
		"Joel":   20,
		"Eunice": 20,
		"June":   20,
	}

	return myMap

}
