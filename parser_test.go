package parser_test

import (
	"fmt"
	"testing"

	parser "github.com/Sotaneum/go-args-parser"
)

func TestRunner(t *testing.T) {
	fmt.Println(parser.Args(map[string]string{}))
	fmt.Println(parser.EnvAll(map[string]string{}))
	fmt.Println(parser.Env((map[string]string{"TMPDIR":"test"})))
	fmt.Println(parser.ArgsJoinEnv((map[string]string{"TMPDIR":"test"})))
}

