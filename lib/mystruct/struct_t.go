package mystruct

import "fmt"

//func CreateStruct() {
//	type DistrictT struct {
//		name     string
//		location string
//	}
//}

func CreateStructs() {

	type DistrictT struct {
		region    string
		name      string
		latitude  string
		longitude string
	}

	type MalawiT struct {
		name string
		dis  DistrictT
	}

	centralRegion := DistrictT{
		"central",
		"Lilongwe",
		"10",
		"1",
	}
	country := MalawiT{
		"Malawi",
		centralRegion,
	}

	fmt.Println(country)
	fmt.Println(centralRegion)

}
