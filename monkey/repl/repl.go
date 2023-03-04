package repl

import (
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	line := `!-a
	`
	l := lexer.New(line)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 { // 查看p.Error()是否是空的
		printParserErrors(out, p.Errors())
	}
	io.WriteString(out, program.String())
	io.WriteString(out, "\n")
}

// func Start(in io.Reader, out io.Writer) {
// 	scanner := bufio.NewScanner(in)
// 	// 先是输出>> 然后不断的读取一行数据让lexer进行解析
// 	// 直到解析到EOF进行停止
// 	for {
// 		fmt.Fprintf(out, PROMPT)
// 		scanned := scanner.Scan()
// 		if !scanned {
// 			return
// 		}
// 		line := scanner.Text()
// 		l := lexer.New(line)
// 		p := parser.New(l)
// 		program := p.ParseProgram()
// 		if len(p.Errors()) != 0 { // 查看p.Error()是否是空的
// 			printParserErrors(out, p.Errors())
// 			continue
// 		}
// 		io.WriteString(out, program.String())
// 		io.WriteString(out, "\n")
// 	}
// }

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`
