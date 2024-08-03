package pkg

import (
    "fmt"
    "strings"
)

func PrintAsciiArt(text string, p map[rune]string, color string, substring string, justify string, width int) {
    if text == "" {
        return
    }

    specialChars := map[string]string{
        "\\t ": "Tab",
        "\\b ": "Backspace",
        "\\v ": "Vertical Tab",
        "\\0 ": "Null",
        "\\f ": "Form Feed",
        "\\r ": "Carriage Return",
    }
    for spChar, description := range specialChars {
        if strings.Contains(text, spChar) {
            fmt.Printf("Print Error: Special ASCII character (%s) or (%s) detected \n", spChar, description)
            return
        }
    }

    text = strings.ReplaceAll(text, "\\n", "\n")

    if text == "\n" {
        fmt.Println()
    }

    var lines []string
    if strings.Contains(text, "\\n") {
        lines = strings.Split(text, "\\n")
    } else {
        lines = strings.Split(text, "\n")
    }
    count := 0

    for _, line := range lines {
        if line == "" {
            count++
            if count < len(lines) {
                fmt.Println()
            }
        } else if len(line) > 0 {
            PrintArt(line, p, color, substring)
        }
    }
}