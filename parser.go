package parser

import (
	"os"
	"strings"
)

func toKeyValue(data, pattern string) (string, string) {
	items := strings.Split(data, pattern)
	if len(items) != 2 {
		if pattern == ":" {
			return data, "true"
		}
		return toKeyValue(data, ":")
	}
	return items[0], items[1]
}

func Args(options map[string]string) map[string]string {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		println(arg)
		key, value := toKeyValue(arg, "=")
		if key == "" || value == "" {
			continue
		}
		options[key] = value
	}
	return options
}
