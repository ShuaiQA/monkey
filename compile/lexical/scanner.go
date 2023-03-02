package lexical

type Scanner struct {
	src        []byte                    // source
	file       *File                     // 文件的信息
	ch         byte                      // 记录offset下标的字符
	offset     int                       // 上一次读取到什么位置
	lineOffset int                       // 代表当前的行的偏移
	rdOffset   int                       // 下一次预读取的位置
	insertSemi bool                      // 是否在\n的时候插入一个;声明变量以及表达式计算可以插入,if for {}之后不需要插入
	err        func(pos int, msg string) // error reporting; or nil
}

// Init 进行初始化的过程
func (s *Scanner) Init(name string, src []byte) {
	s.src = src
	s.ch = ' '
	s.offset = 0
	s.rdOffset = 0
	s.lineOffset = 0
	s.insertSemi = false
	s.file = &File{name: name, size: len(src)}
	// 执行下一次操作,主要的目的是真正的赋值ch = src[0] offset = 0 rdOffset = 1
	s.next()
}

func NewScan(src []byte) *Scanner {
	return &Scanner{src: src, offset: 0, rdOffset: 0, ch: ' '}
}

func (s *Scanner) error(offs int, msg string) {
	if s.err != nil {
		s.err(offs, msg)
	}
}

// scanByte 查找到' 那么就结束了
func (s *Scanner) scanByte() string {
	offs := s.offset - 1
	for {
		ch := s.ch
		s.next()
		if ch == '\'' {
			break
		}
	}
	return string(s.src[offs:s.offset])
}

func (s *Scanner) Scan() (pos int, tok token, lit string) {
	s.skipWhitespace()
	pos = s.offset
	insertSemi := false
	switch ch := s.ch; {
	case isLetter(ch):
		lit = s.scanIdentifier()
		if len(lit) > 1 {
			// keywords are longer than one letter - avoid lookup otherwise
			tok = Lookup(lit)
			switch tok {
			case IDENT, BREAK, CONTINUE, FALLTHROUGH, RETURN:
				insertSemi = true
			}
		} else {
			insertSemi = true
			tok = IDENT
		}
	case isDecimal(ch) || ch == '.' && isDecimal(s.peek()):
		insertSemi = true
		tok, lit = s.scanNumber()
	default:
		s.next() // always make progress
		switch ch {
		case '0' - '0':
			if s.insertSemi {
				s.insertSemi = false // EOF consumed
				return pos, SEMICOLON, "\n"
			}
			tok = EOF
		case '\n':
			// we only reach here if s.insertSemi was
			// set in the first place and exited early
			// from s.skipWhitespace()
			s.insertSemi = false // newline consumed
			return pos, SEMICOLON, "\n"
		case '"':
			insertSemi = true
			tok = STRING
			lit = s.scanString()
		case '\'':
			insertSemi = true
			tok = CHAR
			lit = s.scanByte()
		case ':':
			tok = s.switch2(COLON, DEFINE)
		case '.':
			tok = PERIOD
		case ',':
			tok = COMMA
		case ';':
			tok = SEMICOLON
			lit = ";"
		case '(':
			tok = LPAREN
		case ')':
			insertSemi = true
			tok = RPAREN
		case '[':
			tok = LBRACK
		case ']':
			insertSemi = true
			tok = RBRACK
		case '{':
			tok = LBRACE
		case '}':
			insertSemi = true
			tok = RBRACE
		case '+':
			tok = ADD
		case '-':
			tok = SUB
		case '*':
			tok = MUL
		case '/':
			tok = QUO
		case '%':
			tok = REM
		case '^':
			tok = XOR
		case '<':
			tok = s.switch3(LSS, LEQ, '<', SHL)
		case '>':
			tok = s.switch3(GTR, GEQ, '>', SHR)
		case '=':
			tok = s.switch2(ASSIGN, EQL)
		case '!':
			tok = s.switch2(NOT, NEQ)
		case '&':
			if s.ch == '^' {
				s.next()
				tok = AND_NOT
			} else {
				if s.ch == '&' {
					s.next()
					tok = LAND
				} else {
					tok = AND
				}
			}
		case '|':
			if s.ch == '|' {
				s.next()
				tok = LOR
			} else {
				tok = OR
			}
		default:
			// next reports unexpected BOMs - don't repeat
			insertSemi = s.insertSemi // preserve insertSemi info
			tok = ILLEGAL
			lit = string(ch)
		}
	}
	s.insertSemi = insertSemi
	return
}

// scanString 规定string的检测结尾到达 "就是结束了
func (s *Scanner) scanString() string {
	offs := s.offset - 1
	for {
		ch := s.ch
		s.next()
		if ch == '"' {
			break
		}
	}
	return string(s.src[offs:s.offset])
}

// scanIdentifier 继续向下查找标识符可能的字符
func (s *Scanner) scanIdentifier() string {
	offs := s.offset
	for rdOffset, b := range s.src[s.rdOffset:] {
		if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
			continue
		}
		s.rdOffset += rdOffset
		s.next()
	}
	return string(s.src[offs:s.offset])
}

// scanNumber 仅仅处理10进制的数, 只能出现1个.
func (s *Scanner) scanNumber() (token, string) {
	offs := s.offset
	tok := INT
	for rdOffset, b := range s.src[s.rdOffset:] {
		if b >= '0' && b <= '9' {
			continue
		}
		s.rdOffset += rdOffset
		s.next()
	}
	if s.ch == '.' {
		tok = FLOAT
		s.next()
	}
	for rdOffset, b := range s.src[s.rdOffset:] {
		if b >= '0' && b <= '9' {
			continue
		}
		s.rdOffset += rdOffset
		s.next()
	}
	lit := string(s.src[offs:s.offset])
	return tok, lit
}

// skipWhitespace 不断的跳过空格,或者是\n但是对于后者来说只有不需要插入;的时候才继续执行,需要插入的时候需要停下来插入;操作
// example  if a==10 {     后面有多个空格当前不需要插入;所以可以执行next
func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' && !s.insertSemi {
		s.next()
	}
}

// next 将rdOffset同步到offset然后rdOffset向后移动一个,将ch设置为当前的rdOffset的值
func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		// 如果当前的ch是'\n'将offset下标添加到file的AddLine函数中
		if s.ch == '\n' {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}
		r, w := byte(s.src[s.rdOffset]), 1
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		s.ch = '0' - '0'
	}
	// 由此函数结束可以得到ch = src[offset]
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || ch == '_' || ch >= 'A' && ch <= 'Z'
}
func isDecimal(ch byte) bool { return '0' <= ch && ch <= '9' }
func (s *Scanner) peek() byte {
	if s.rdOffset < len(s.src) {
		return s.src[s.rdOffset]
	}
	return 0
}

// switch2 当前是: 下一个是= 优先是定义语句, 否则则是 : 可能是goto跳转语句
// 如果下一个是 = 就是tok1 否则则是tok0
func (s *Scanner) switch2(tok0, tok1 token) token {
	if s.ch == '=' {
		s.next()
		return tok1
	}
	return tok0
}

func (s *Scanner) switch3(tok0, tok1 token, ch2 byte, tok2 token) token {
	if s.ch == '=' {
		s.next()
		return tok1
	}
	if s.ch == ch2 {
		s.next()
		return tok2
	}
	return tok0
}
