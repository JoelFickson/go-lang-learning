package data_slices

func CreateSlice() []string {
	listOfMajorCities := []string{
		"Mzuzu",
		"Lilongwe",
		"Blantyre",
	}
	listOfMajorCities = append(listOfMajorCities, "Zomba")

	return listOfMajorCities

}
