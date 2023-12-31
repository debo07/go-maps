package main

import "fmt"

func main() {
	colors := map[string]string {
		"white": "#ffffff",
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

	colors["white"] = "white"
	fmt.Println(colors)

	// Delete a field from map
	delete(colors, "white")
	fmt.Println(colors)
}