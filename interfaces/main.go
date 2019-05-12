package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishbot struct {
}

type spanishbot struct {
}

func main() {
	eb := englishbot{}
	sb := spanishbot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(sb spanishbot) {
// 	fmt.Println(sb.getGreeting())
// }
func (eb englishbot) getGreeting() string {
	return "HI THERE !"
}

func (sb spanishbot) getGreeting() string {
	return "HOLA!"
}
