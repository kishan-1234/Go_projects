package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4f0000",
	}
	colors2 := make(map[string]string)
	colors2["brown"] = "#3f0010"
	delete(colors, "red") // delete keys in map
	fmt.Println(colors, colors2)
	printmap(colors)
}

func printmap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
