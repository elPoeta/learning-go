package main

import "fmt"

func main() {
	//Array
	colors := [3]string{"green", "blue", "red"}
	fmt.Println(colors)

	//sclice

	languges := []string{"go", "javascript"}
	languges = append(languges, "c")
	languges = append(languges, "java")
	fmt.Println(languges)

	fmt.Println(len(languges))

	for index, lang := range languges {
		fmt.Println(index, lang)
	}

}
