package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mailru/easyjson/bootstrap"
	"github.com/mailru/easyjson/parser"
)

var buildTags = flag.String("build_tags", "", "build tags to add to generated file")
var snakeCase = flag.Bool("snake_case", false, "use snake_case names instead of CamelCase by default")
var omitEmpty = flag.Bool("omit_empty", false, "omit empty fields by default")
var allStructs = flag.Bool("all", false, "generate un-/marshallers for all structs in a file")
var leaveTemps = flag.Bool("leave_temps", false, "do not delete temporary files")
var stubs = flag.Bool("stubs", false, "only generate stubs for marshallers/unmarshallers methods")

func generate(fname string) (err error) {
	p := parser.Parser{AllStructs: *allStructs}
	if err := p.Parse(fname); err != nil {
		return fmt.Errorf("Error parsing %v: %v", fname, err)
	}

	var outName string
	if s := strings.TrimSuffix(fname, ".go"); s == fname {
		return fmt.Errorf("Filename must end in '.go'")
	} else {
		outName = s + "_easyjson.go"
	}

	g := bootstrap.Generator{
		BuildTags:  *buildTags,
		PkgPath:    p.PkgPath,
		PkgName:    p.PkgName,
		Types:      p.StructNames,
		SnakeCase:  *snakeCase,
		OmitEmpty:  *omitEmpty,
		LeaveTemps: *leaveTemps,
		OutName:    outName,
		StubsOnly:  *stubs,
	}

	if err := g.Run(); err != nil {
		return fmt.Errorf("Bootstrap failed: %v", err)
	}
	return nil
}

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		files = []string{os.Getenv("GOFILE")}
	}

	for _, fname := range files {
		if err := generate(fname); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}