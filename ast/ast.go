package ast

import (
	"../token"
	"fmt"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}
	return p.Statements[0].TokenLiteral()
}

func (p *Program) String() string {
	var sb strings.Builder

	for _, stmt := range p.Statements {
		sb.WriteString(stmt.String())
	}

	return sb.String()
}

type LetStatement struct {
	Token      token.Token
	Identifier *Identifier
	Value      Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) expressionNode() {}
func (ls *LetStatement) statementNode()  {}

func (ls *LetStatement) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("let %s = ", ls.Identifier.String()))

	if ls.Value != nil {
		sb.WriteString(ls.Value.String())
	}
	sb.WriteString(";")

	return sb.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) expressionNode() {}
func (rs *ReturnStatement) statementNode()  {}

func (rs *ReturnStatement) String() string {
	var sb strings.Builder

	sb.WriteString("return")

	if rs.ReturnValue != nil {
		sb.WriteString(rs.ReturnValue.String())
	}
	sb.WriteString(";")

	return sb.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) expressionNode() {}
func (es *ExpressionStatement) statementNode()  {}

func (es *ExpressionStatement) String() string {
	var sb strings.Builder

	if es.Expression != nil {
		sb.WriteString(es.Expression.String())
		sb.WriteString(";")
	}

	return sb.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Value
}
