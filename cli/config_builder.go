package cli

import "sort"

//Struct that represents the parsed configuration
type Configuration struct {
	Keywords []string
	Span     bool
}

//Parses the passed os arguments to cmd parameter
func Parse(args []string) Configuration {
	return Configuration{
		Keywords: remove(args, "span"),
		Span:     contains(args, "span"),
	}
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
