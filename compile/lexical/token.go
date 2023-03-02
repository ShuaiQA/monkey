package lexical

import "strconv"

type token int

// 主要是定义了token都有哪些,以及token对应的字符串和输出标准

const (
	// Special tokens
	ILLEGAL token = iota
	EOF
	// type
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	CHAR   // 'a'
	STRING // "abc"
	// 算术运算符
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %
	// 逻辑运算
	LOR  // ||
	LAND // &&
	NOT  // !

	// 其余运算
	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^
	// 关系运算符
	EQL   // ==
	LSS   // <
	GTR   // >
	NEQ   // !=
	LEQ   // <=
	GEQ   // >=
	ARROW // <-
	// 定义
	DEFINE // :=
	// 赋值运算
	ASSIGN    // =
	LPAREN    // (
	LBRACK    // [
	LBRACE    // {
	COMMA     // ,
	PERIOD    // .
	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	// 关键字
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE
	DEFAULT
	DEFER
	IF
	ELSE
	FALLTHROUGH
	FOR
	FUNC
	GO
	GOTO
	IMPORT
	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN
	SELECT

	STRUCT
	SWITCH
	TYPE
	VAR
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	LOR:  "||",
	LAND: "&&",
	NOT:  "!",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	NEQ:    "!=",
	LEQ:    "<=",
	GEQ:    ">=",
	ARROW:  "<-",
	DEFINE: ":=",

	ASSIGN:    "=",
	LPAREN:    "(",
	LBRACK:    "[",
	LBRACE:    "{",
	COMMA:     ",",
	PERIOD:    ".",
	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	BREAK:    "break",
	CASE:     "case",
	CHAN:     "chan",
	CONST:    "const",
	CONTINUE: "continue",

	DEFAULT:     "default",
	DEFER:       "defer",
	IF:          "if",
	ELSE:        "else",
	FALLTHROUGH: "fallthrough",
	FOR:         "for",

	FUNC:      "func",
	GO:        "go",
	GOTO:      "goto",
	IMPORT:    "import",
	INTERFACE: "interface",
	MAP:       "map",
	PACKAGE:   "package",
	RANGE:     "range",
	RETURN:    "return",
	SELECT:    "select",
	STRUCT:    "struct",
	SWITCH:    "switch",
	TYPE:      "type",
	VAR:       "var",
}

func (tok token) String() string {
	s := ""
	if 0 <= tok && tok < token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var keywords map[string]token

func init() {
	keywords = make(map[string]token)
	for i := BREAK; i < VAR; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
//
func Lookup(ident string) token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}
