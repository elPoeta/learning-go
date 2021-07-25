package main

import "fmt"

type card struct {
	value string
	suit  string
}

func main() {
	spade := card{
		value: "Ace",
		suit:  "Spades",
	}
	spade.print()
	spade.updateValue("Four")
	fmt.Println()
	spade.print()
}
func (c *card) updateValue(value string) {
	(*c).value = value
}

func (c card) print() {
	fmt.Printf("%+v", c)
}
