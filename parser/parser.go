package parser

import (
	"../ast"
	"../lexer"
	"../token"
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(expression ast.Expression) ast.Expression
)

type Parser struct {
	lexer          *lexer.Lexer
	currToken      token.Token
	peekToken      token.Token
	errors         []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefixParseFn(token.IDENT, p.parseIdentifier)
	p.registerPrefixParseFn(token.INT, p.parseIntegerLiteral)
	return p
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

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Literal {
	case "let":
		return p.parseLetStatement()
	case "return":
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	rs := &ast.ReturnStatement{Token: p.currToken}

	for !p.currTokenIsOfType(token.SEMICOLON) {
		p.nextToken()
	}
	return rs
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	es := &ast.ExpressionStatement{Token: p.currToken}

	es.Expression = p.parseExpression(LOWEST)

	if !p.eat(token.SEMICOLON) {
		return nil
	}

	return es
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixFn := p.prefixParseFns[p.currToken.Type]
	if prefixFn == nil {
		return nil
	}
	leftExp := prefixFn()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	intLiteral, err := strconv.Atoi(p.currToken.Literal)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("error converting intLiteral from string to int %v\n", err))
		return nil
	}
	return &ast.IntegerLiteral{Token: p.currToken, Value: intLiteral}
}

func (p *Parser) eat(tt token.TokenType) bool {
	if p.peekToken.Type == tt {
		p.nextToken()
		return true
	} else {
		p.eatError(tt)
		return false
	}
}

func (p *Parser) eatError(tt token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("expected next token of type %s but got %s instead\n", tt, p.peekToken.Type))
}

func (p *Parser) currTokenIsOfType(tt token.TokenType) bool {
	return p.currToken.Type == tt
}

func (p *Parser) registerPrefixParseFn(tt token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}

func (p *Parser) registerInfixParseFn(tt token.TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}
