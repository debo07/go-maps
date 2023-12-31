package main

import "fmt"

func main() {
	colors := map[string]string {
		"white": "#ffffff",
		"red": "#ff0000",
		"black": "#000000",
	}
	fmt.Println(colors)

	// Another way to declare a map
	var colors1 map[string]string
	fmt.Println(colors1)

	// Another way to declare a map using make
	colors2 := make(map[string]string)
	fmt.Println(colors2)

	// Update/Add value to map
	colors2["white"] = "white"
	fmt.Println(colors2)

	// Delete a field from map
	delete(colors2, "white")
	fmt.Println(colors2)

	// Iterate over map
	for color, hex := range colors {
		fmt.Println("Hex of", color, "is", hex)
	}
}