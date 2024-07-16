package main

import (
	"fmt"
	"os"

	"practice1/pkg"

)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	justify, args := pkg.CheckFlag(args)
	fileName := ""
	if len(args) > 1 {
		switch args[len(args)-1] {
		case "shadow":
			fileName = "shadow.txt"
		case "standard":
			fileName = "standard.txt"
		case "thinkertoy":
			fileName = "thinkertoy.txt"
		default:
			fileName = "standard.txt"
		}
	} 
		
		err := pkg.ValidateFileChecksum(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		
	}
	
	
	width, err := pkg.GetTerminalWidth()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	theMap, err := pkg.ReadFile(fileName)
	if err != nil {
		fmt.Println("error")
	}
	pkg.PrintAsciiArt(args[0], theMap, justify, width )

	// // fileName := pkg.TheFile(args)

	// // for _, v := range theMap {
	// // 	fmt.Printf("%s\n", v)
	// // }
	// pkg.PrintAsciiArt(args[0], theMap)
}
