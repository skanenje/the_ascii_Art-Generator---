package pkg

import (
	"strings"
)

func ParseFlag(sl []string) (string, string, []string, error) {
	var color string
	var substring string
	var args []string

	for i, arg := range sl {
		if strings.HasPrefix(arg, "--color=") {
			color = strings.TrimPrefix(arg, "--color=")
			args = append(sl[:i], sl[i+1:]...)
			
			// Check if there's a substring specified
			if len(args) > 1 {
				substring = args[0]
				args = args[1:]
			}
			
			return color, substring, args, nil
		}
	}

	// If we get here, it means we didn't find a valid --color flag
	return "", "", nil, nil
}