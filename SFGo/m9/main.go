package main

import (
	"fmt"
	"sort"
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
	suspects := []string{"Alex", "Ronda", "Stipe", "Den", "Cain"}
	badlist := []Man{}

	Ppl := make(People)
	Ppl["Johny"] = Man{"Johny", "Doe", 33, "male", 4}
	Ppl["Markus"] = Man{"Markus", "Miller", 78, "male", 5}
	Ppl["Sahsa"] = Man{"Sahsa", "Smith", 17, "female", 11135}
	Ppl["Maxwell"] = Man{"Maxwell", "Donza", 29, "male", 8}
	Ppl["Killian"] = Man{"Killian", "Rowdy", 51, "male", 100}

	for _, value := range suspects {

		if _, ok := Ppl[value]; ok {
			badlist = append(badlist, Ppl[value])
		}

	}
	if len(badlist) == 0 {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
	if len(badlist) >= 1 {
		sort.SliceStable(badlist, func(i, j int) bool {
			return badlist[i].Crimes < badlist[j].Crimes
		})
		fmt.Println(badlist[len(badlist)-1].Name)
	}
}
