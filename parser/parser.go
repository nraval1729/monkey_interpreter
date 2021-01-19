package parser

import (
	"../ast"
	"../lexer"
	"../token"
	"fmt"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}

	for p.currToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Literal {
	case "let":
		return p.parseLetStatement()
	default:
		panic(fmt.Errorf("invalid statement type %s\n", p.currToken.Literal))
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	ls := &ast.LetStatement{Token: p.currToken}

	if !p.eat(token.IDENT) {
		//	TODO: handle syntax error
		return nil
	}
	ls.Identifier = &ast.Identifier{Token: token.NewToken(token.IDENT, "ident"), Value: p.currToken.Literal}

	if !p.eat(token.ASSIGN) {
		//	TODO: handle syntax error
		return nil
	}

	for !p.currTokenIsOfType(token.SEMICOLON) {
		p.nextToken()
	}
	return ls
}

func (p *Parser) eat(tt token.TokenType) bool {
	if p.peekToken.Type == tt {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) currTokenIsOfType(tt token.TokenType) bool {
	return p.currToken.Type == tt
}