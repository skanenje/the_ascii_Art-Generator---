package pkg

import (
	"fmt"
	"strings"
)

const asciiHeight = 8
func PrintArt(s string, p map[rune]string, justify string, width int) {
	for i := 0; i < asciiHeight; i++ {
		var b strings.Builder
		for _, v := range s {
			art, ok := p[v]
			if !ok {
				fmt.Println("char not found")
				return
			}
			artLines := strings.Split(art, "\n")
			if len(artLines) <= i {
				continue
			}
			b.WriteString(artLines[i])
		}
		lineOutput := b.String()
		padding := getsize(lineOutput, justify, width)
		f := strings.Repeat(" ", padding) + lineOutput
		fmt.Println(f)
	}
}