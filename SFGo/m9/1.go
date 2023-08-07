package main

import (
	"fmt"
	"os"
)

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

type People map[string]Man

func main() {
	suspects := []string{"Den", "Alex", "Khabib", "Den", "asc"}

	Ppl := make(People)
	Ppl["Johny"] = Man{"Johny", "Doe", 33, "male", 34}
	Ppl["Markus"] = Man{"Markus", "Miller", 78, "male", 0}
	Ppl["Sahsa"] = Man{"Sahsa", "Smith", 17, "female", 0}
	Ppl["Maxwell"] = Man{"Maxwell", "Donza", 29, "male", 1000}
	Ppl["Killian"] = Man{"Killian", "Rowdy", 51, "male", 223}

	var maxCrimes int
	var maxKey string
	var foundPpl bool

	for _, value := range suspects {

		if _, ok := Ppl[value]; ok {
			foundPpl = true
			person := Ppl[value]
			if person.Crimes > maxCrimes {
				maxCrimes = person.Crimes
				maxKey = person.Name
			}
		}
	}
	if !foundPpl {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
		os.Exit(1)
	}
	fmt.Println(maxKey)
}

// if _, ok := Ppl[value]; ok {
// 	foundPpl = true
// 	person := Ppl[value]
// 	if person.Crimes > maxCrimes {
// 		maxCrimes = person.Crimes
// 		maxKey = person.Name
// 	}
// 	fmt.Println(maxKey, maxCrimes)
// }
