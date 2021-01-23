package ast

import "../token"

type Node interface {
	TokenLiteral() string
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

type LetStatement struct {
	Token      token.Token
	Identifier *Identifier
	Expression *Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) expressionNode() {}
func (ls *LetStatement) statementNode() {}

type ReturnStatement struct {
	Token      token.Token
	ReturnValue *Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) expressionNode() {}
func (rs *ReturnStatement) statementNode() {}

type ExpressionStatement struct {
	Token      token.Token
	Expression *Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) expressionNode() {}
func (es *ExpressionStatement) statementNode() {}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}
