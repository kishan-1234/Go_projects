package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() Deck {
	cards := Deck{}
	cardSuites := []string{"Club", "Spade", "Heart", "Diamond"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func deal(d Deck, Handsize int) (Deck, Deck) {
	return d[:Handsize], d[Handsize:]
}

func (d Deck) toString() string {

	return strings.Join([]string(d), ", ")
}

func (d Deck) savetofile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromfile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1) //len(slicevariable) to get the length of the slice
		d[index], d[newPosition] = d[newPosition], d[index]

	}
}
