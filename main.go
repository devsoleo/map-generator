package main

import (
	"devsoleo/map-generator/lib"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: missing arguments")
		return
	}

	switch os.Args[1] {
	case "-f": // Take a file
		file, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid file")
			return
		}

		var config lib.Config
		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			fmt.Println("Error: invalid config")
			return
		}

		lib.Generate(config)
	case "-c": // Take stdin
		var config lib.Config
		err := json.Unmarshal([]byte(os.Args[2]), &config)
		if err != nil {
			fmt.Println("Error: invalid config")
			return
		}

		lib.Generate(config)
	default:
		fmt.Println("Error: unknown flag")
	}
}
