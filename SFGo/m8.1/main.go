package main

import (
	"fmt"
	electronic "m81/pkg"
)

func main() {
	apple := electronic.AppleCreate("3G", "IOS")                     //
	android := electronic.AndroidCreate("Samsung", "S20", "Android") //
	radiophone := electronic.RadioCreate("Siemens", "CX65", 9)       //"

	printCharacteristics(apple)
	printCharacteristics(android)
	printCharacteristics(radiophone)
}

func printCharacteristics(a electronic.Phone) {
	switch a.(type) {
	case electronic.Smartphone:
		casted := a.(electronic.Smartphone)
		fmt.Printf("Brand: %s\nModel: %s\nType: %s\nOS: %s\n", casted.Brand(), casted.Model(), casted.Type(), casted.OS())
	case electronic.StationPhone:
		busted := a.(electronic.StationPhone)
		fmt.Printf("Brand: %s\nModel: %s\nType: %s\nButtonsCount: %d\n", busted.Brand(), busted.Model(), busted.Type(), busted.ButtonsCount())
	default:
		fmt.Println("нихуя не выводица, капса")
	}

}
