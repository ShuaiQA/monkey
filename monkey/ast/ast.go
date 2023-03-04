// ast/ast.go
package ast

import "monkey/token"

// 无论是语句还是标识符都会实现该方法
type Node interface {
	TokenLiteral() string
}

type Statement interface { // 语句
	Node
	statementNode()
}

type Expression interface { // 表达式
	Node
	expressionNode()
}

// 程序是由多个语句构成的
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET词法单元
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// 返回当前的token的标识,用于调试当前是否是对应的目标值
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT词法单元
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
