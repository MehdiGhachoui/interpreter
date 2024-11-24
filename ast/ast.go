package ast

import (
	"go-interpreter/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Statement
type LetStatement struct {
	Token token.Token //token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode()
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) ExpressionNode()
func (ls *Identifier) TokenLiteral() string {
	return ls.Token.Literal
}
