package main

import (
	"encoding/json"
	"fmt"
)

// Sampe program to how marshal & unmarshal

type Person struct {
	First string
	Last  string
}

func main() {
	/* This will marshal the JSON into []bytes */

	p1 := Person{"alice", "bob"}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

	/* This will unmarshal the JSON from []bytes */

	var p2 Person
	bs = []byte(`{"First":"alice","Last":"bob"}`)
	json.Unmarshal(bs, &p2)
	fmt.Println(p2)

}
