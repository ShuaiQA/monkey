package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

// 定义表达式的优先级
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string // 对于每一个语句进行记录相关的错误
	// 主要是记录相应的token对应的相关的函数
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// 初始化解析器结构体,该结构体包含词法分析,以及当前的token和可以查看的下一个token
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// 读取两个词法单元，以设置curToken和peekToken
	p.nextToken()
	p.nextToken()
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	// 针对于IDENT,创建一个前缀解析函数
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	// 注册中缀表达式
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	return p
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	// 注意此处进行解析右面的部分,返回后的值赋值给当前的右面的表达式
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // 生成一个程序的根节点
	program.Statements = []ast.Statement{}
	// 遍历当前的token的类型,到达EOF进行结束
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement() // 解析每一条语句
		if stmt != nil {           // 如果解析出来的语句没有出现错误,将该语句添加到根节点program中
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// 根据curToken进行解析语句
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// 解析标识符值
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// 解析整型值,将字符串进行转换
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

// 定义前缀解析函数和中缀解析函数
// 在前缀位置遇到相关的词法单元,调用prefixParseFn
// 在中缀位置遇到词法单元类型时会调用infixParseFn
// 根据函数的参数和返回值也可以看出前缀和中缀的区别
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// 将相关的tokenType注册到对应的函数上面
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// 将后面的赋值给前面的,peekToken进行更新
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token is %s but get token %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// 在该函数中优先级最低的会传递给parseExpression
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST) // 表达式解析LOWEST代表什么
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// 检查前缀位置是否有与tokenType相关的解析函数,如果有相应的函数调用解析函数
func (p *Parser) parseExpression(precedence int) ast.Expression {
	// 首先是前缀解析获取前缀表达式
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	// 之后进行中缀解析,只有当优先级较大的时候才会执行下面的循环
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	// TODO: 跳过对表达式的处理，直到遇见分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// 解析Let语句
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// 对curToken进行赋值,便于后续使用TokenLiteral()方法进行调试
	stmt := &ast.LetStatement{Token: p.curToken}

	// 当前期望下一个token是标识符
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	// 将标识符添加到语句中
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	// 期望下一个操作符是赋值
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: 跳过对表达式的处理，直到遇见分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// 如果下一个token符合我们的预测那么消耗一个token
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// 建立操作符和优先级之间的关系
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

// 如果没有找到返回最低的优先级
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
