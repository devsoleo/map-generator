package main

import (
	"fmt"
	"math/rand"
)

func GenerateMap() string {
	// return a 1000 random chars length string

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 1_000_000)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return "{\"data\": \"" + string(b) + "\" }"
}

func main() {
	/*for i, arg := range os.Args {
		fmt.Printf("Argument %d: %s\n", i, arg)
	}*/

	fmt.Println(GenerateMap())
}
