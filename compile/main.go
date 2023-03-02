package main

import (
	"compile/lexical"
	"fmt"
	"go/scanner"
	"go/token"
	"time"
)

func main() {

	src := []byte(`"a"`)
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)
	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}

	var c lexical.Scanner
	c.Init("name", src)
	for {
		_, tok, lit := c.Scan()
		time.Sleep(1 * time.Second)
		fmt.Printf("%-8s%q\n", tok, lit)

		if tok == lexical.EOF {
			break
		}

	}

}
