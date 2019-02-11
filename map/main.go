package main

import "fmt"

func main() {
	// Declaring a map #1
	// var colors map[string]string
	// Declaring a map #2
	// colors := make(map[string]string)

	// Declaring a map #3
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"

	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for : ", color, " is :", hex)
	}
}
