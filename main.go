package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	flagSource  = flag.String("source", "", "go source file or folder")
	flagLeftpad = flag.String("leftpad", "", "left padding characters")
	flagVersion = flag.Bool("version", false, "print the version and exit")
	fset        *token.FileSet
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of struct2oas:\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("struct2oas: ")
	flag.Usage = Usage
	flag.Parse()

	if *flagVersion {
		fmt.Println("struct2oas v1.0.0")
		os.Exit(0)
	}

	process(*flagSource)
}

func process(name string) {
	// check if source is a file or folder
	fileInfo, err := os.Stat(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch mode := fileInfo.Mode(); {
	case mode.IsDir():
		processDir(name)
	case mode.IsRegular():
		processFile(name)
	}
}

func processDir(name string) {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < len(files); i++ {
		if path.Ext(files[i].Name()) == ".go" {
			process(path.Join(name, files[i].Name()))
		}
	}
}

func processFile(name string) {
	log.Printf("Read source %q\n", name)
	fset = token.NewFileSet()
	node, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
	if err != nil {
		log.Println("Error parsing source: " + name)
		os.Exit(255)
	}

	g := Generator{}
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			if t.Name.IsExported() {
				g.Generate(t)
			}
		}
		return true
	})
	g.WriteFile()
}

type Generator struct {
	buf  bytes.Buffer // Accumulated output.
	Name string
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, *flagLeftpad+format, args...)
}

func (g *Generator) Generate(t *ast.TypeSpec) {
	g.Name = t.Name.String()
	log.Println("Generate", t.Name)

	g.Printf("# Code generated by \"struct2oas -source %s\"; DO NOT EDIT.\n", *flagSource)
	g.Printf("%s:\n", t.Name.Name)
	g.Printf("  type: \"object\"\n")
	g.Printf("  properties:\n")

	switch s := t.Type.(type) {
	case *ast.StructType:
		for i := 0; i < len(s.Fields.List); i++ {
			log.Printf("  Prop Name: %s\tType: %s\tTag: %s\n", s.Fields.List[i].Names[0], s.Fields.List[i].Type, s.Fields.List[i].Tag.Value)
			g.Printf("    %s:\n", s.Fields.List[i].Names[0])

			switch ty := s.Fields.List[i].Type.(type) {
			case *ast.Ident:
				typ, format := TypeToSchema(s.Fields.List[i].Type)
				g.Printf("      type: %q\n", typ)
				if format != "" {
					g.Printf("      format: %q\n", format)
				}
				description := strings.Replace(s.Fields.List[i].Doc.Text(), "\n", "", -1)
				if description != "" {
					g.Printf("      description: %q\n", description)
				}

			case *ast.SelectorExpr:
				if ty.X.(*ast.Ident).Name == "time" && ty.Sel.Name == "Time" {
					g.Printf("      type: \"string\"\n")
					g.Printf("      format: \"date-time\"\n")
					description := strings.Replace(s.Fields.List[i].Doc.Text(), "\n", "", -1)
					if description != "" {
						g.Printf("      description: %q\n", description)
					}
				}

			case *ast.ArrayType:
				// log.Printf("----------- TODO %+v\n", ty.(*ast.ArrayType).Name)
				// ast.Print(fset, ty)
				// log.Println("TYYYYY", ty, TypeToSchema(s.Fields.List[i].Type))

				g.Printf("      type: \"array\"\n")
				g.Printf("      items:\n")
				typ, format := TypeToSchema(ty.Elt)
				g.Printf("        type: %q\n", typ)
				if format != "" {
					g.Printf("        format: %q\n", format)
				}
				// g.Printf("        type: %q\n", ty.Elt.(*ast.Ident).Name)
				// g.Printf("        properties:\n")

			case *ast.MapType:
				g.Printf("      type: \"object\"\n")
				description := strings.Replace(s.Fields.List[i].Doc.Text(), "\n", "", -1)
				if description != "" {
					g.Printf("      description: %q\n", description)
				}

				// case *ast.StructType:
				// 	g.Printf("      type: \"object\"\n")
				// 	description := strings.Replace(s.Fields.List[i].Doc.Text(), "\n", "", -1)
				// 	if description != "" {
				// 		g.Printf("      description: %q\n", description)
				// 	}
			}
			// ast.Print(fset, s.Fields.List[i].Type)
		}
	}
}

func (g *Generator) WriteFile() {
	err := ioutil.WriteFile(g.Name+".yml", g.buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

// https://swagger.io/docs/specification/data-models/data-types/#array
func TypeToSchema(e ast.Expr) (t string, f string) {
	switch e.(*ast.Ident).Name {
	case "string":
		t = "string"
		f = ""
		break
	case "bool":
		t = "boolean"
		f = ""
		break
	case "int", "int8", "int16", "uint", "uint8", "uint16", "byte", "rune":
		t = "integer"
		f = ""
		break
	case "int32", "uint32":
		t = "integer"
		f = "int32"
		break
	case "int64", "uint64":
		t = "integer"
		f = "int64"
		break
	case "float32":
		t = "number"
		f = "float"
		break
	case "float64", "complex64", "complex128":
		t = "number"
		f = "double"
		break
	}
	return
}
