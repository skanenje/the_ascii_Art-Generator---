package pkg

import "flag"

func CheckFlag(args []string) (string, []string){
	flagSet := flag.NewFlagSet("flags", flag.ContinueOnError)
	var justify string
	flagSet.StringVar(&justify, "align", "left", "string to represent printing area")
	err := flagSet.Parse(args)
	if err != nil {
		return "", nil
	}
	args = flagSet.Args()
	return justify, args

	
}
func getsize(s string, justify string, width int) int{
	padding := 0
	
	switch justify {

	case "left":
		padding = 0
	case "right":
		padding = (width - len(s))
	case "center":
		padding = (width - len(s)) / 2
	case "justify":
		padding =(width - len(s))/len(s)
	default:
		padding = 0
	}

	return padding
}