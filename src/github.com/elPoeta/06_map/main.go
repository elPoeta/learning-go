package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}

	fmt.Println(colors)

	for key, value := range colors {
		fmt.Println("color ", key, "=>", value)
	}

	fmt.Println(colors["red"])

	delete(colors, "red")

	fmt.Println(colors)

	var countryCodes = make(map[string]string)

	countryCodes["ARG"] = "Argentina"
	countryCodes["ITA"] = "Italy"
	countryCodes["USA"] = "United States"
	countryCodes["JPN"] = "Japan"

	fmt.Println(countryCodes)

}
