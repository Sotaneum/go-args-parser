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
		key, value := toKeyValue(arg, "=")
		if key == "" || value == "" {
			continue
		}
		options[key] = value
	}
	return options
}

func Env(options map[string]string) map[string]string{
	for key, _ := range options {
		value := os.Getenv(key)
		if value == ""{
			continue
		}
		options[key] = value
	}
	return options
}

func EnvAll(options map[string]string) map[string]string{
	for _, item := range os.Environ() {
		pair := strings.Split(item, "=")
		options[pair[0]] = pair[1]
	}
	return options
}

func ArgsJoinEnv(options map[string]string) map[string]string{
	return Args(Env(options))
}

func ArgsJoinEnvAll(options map[string]string) map[string]string{
	return Args(EnvAll(options))
}