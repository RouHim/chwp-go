package cli

//Struct that represents the parsed configuration
type Configuration struct {
	Keywords []string
	Span     bool
}

//Parses the passed os arguments to cmd parameter
func Parse(args []string) Configuration {
	spans := contains(args, "span")
	args = remove(args, "span")

	if len(args) <= 0 {
		args = append(args, "wallpaper")
	}

	return Configuration{
		Keywords: args,
		Span:     spans,
	}
}

func contains(slice []string, term string) bool {
	for _, item := range slice {
		if item == term {
			return true
		}
	}
	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
