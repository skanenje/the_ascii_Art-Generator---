package main

import (
    "fmt"
    "os"
	"strings"
    "practice1/pkg"
)

func main() {
	args := os.Args[1:]
	color, substring, argz, err := pkg.ParseFlag(args)
	if err != nil || len(argz) < 1 || color == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	fileName := "standard.txt"
	text := argz[0]

	if len(argz) > 1 {
		// Check if the last argument is a valid banner file
		switch argz[len(argz)-1] {
		case "shadow", "standard", "thinkertoy":
			fileName = argz[len(argz)-1] + ".txt"
			text = strings.Join(argz[:len(argz)-1], " ")
		default:
			// If it's not a valid banner file, assume it's part of the text
			text = strings.Join(argz, " ")
		}
	}

	err = pkg.ValidateFileChecksum(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	theMap, err := pkg.ReadFile(fileName)
	if err != nil {
		fmt.Println("error")
		return
	}

	// If substring is empty, color the whole string
	if substring == "" {
		substring = text
	}

	// Call your printing function here, for example:
	pkg.PrintArt(text, theMap, color, substring)
}