package pkg

import (
    "fmt"
    "strings"
)

func PrintAsciiArt(text string, p map[rune]string, color string, substring string) {
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
        return
    }

    var lines []string
    if strings.Contains(text, "\\n") {
        lines = strings.Split(text, "\\n")
    } else {
        lines = strings.Split(text, "\n")
    }

    for _, line := range lines {
        if line == "" {
            fmt.Println()
        } else {
            PrintArt(line, p, color, substring)
        }
    }
}