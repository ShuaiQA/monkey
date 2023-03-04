package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// 初始化解析器结构体,该结构体包含词法分析,以及当前的token和可以查看的下一个token
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 读取两个词法单元，以设置curToken和peekToken
	p.nextToken()
	p.nextToken()
	return p
}

// 将后面的赋值给前面的,peekToken进行更新
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
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
		// 向后遍历
		p.nextToken()
	}
	return program
}

// 根据curToken进行解析语句
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
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
		return false
	}
}
