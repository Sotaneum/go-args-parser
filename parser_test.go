package parser_test

import (
	"fmt"
	"testing"

	parser "github.com/Sotaneum/go-args-parser"
)

func TestRunner(t *testing.T) {
	fmt.Println(parser.Args(map[string]string{}))
}
