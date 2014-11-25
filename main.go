package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

var dirNames []map[string]*ast.Package

func validate(t *doc.Func) {
	if strings.HasPrefix(t.Name, "Test") || strings.HasPrefix(t.Name, "Example") {
		return
	}
	if len(t.Doc) == 0 {
		fmt.Printf("	[W] %s() is missing godoc\n", t.Name)
		return
	}
	parsedDoc := strings.Split(t.Doc, " ")
	if parsedDoc[0] != t.Name {
		fmt.Printf("	[E] '%s' godoc name does not match func name %s()\n", parsedDoc[0], t.Name)
	}
}

func parseFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	if !!fi.IsDir() {
		fset := token.NewFileSet()
		d, err := parser.ParseDir(fset, fp, nil, parser.ParseComments)
		if err != nil {
			fmt.Println(err)
		} else {
			dirNames = append(dirNames, d)
		}
	}
	return nil
}

//func main() {
//    //specify directory below or walk through /
//    filepath.Walk("/", VisitFile)
//}

func main() {
	flag.Parse()

	pkgName := flag.Arg(0)
	goPath := os.Getenv("GOPATH")

	if len(goPath) == 0 {
		fmt.Printf("ERROR: The GOPATH must be set")
		os.Exit(1)
	}

	importPath := goPath + "/src/" + pkgName

	filepath.Walk(importPath, parseFile)

	for _, d := range dirNames {
		for k, f := range d {
			p := doc.New(f, "./", 1)

			if len(p.Funcs) > 0 {
				fmt.Printf("Analyzing %s:\n", k)
				for _, f := range p.Funcs {
					validate(f)
				}
			}

			for _, t := range p.Types {
				if len(t.Methods) == 0 {
					continue
				}
				fmt.Printf("Processing '%s' in package %s:\n", t.Name, k)
				for _, f := range t.Methods {
					validate(f)
				}
				for _, f := range t.Funcs {
					validate(f)
				}
			}
		}
	}
}
