package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("<Type: %v, Literal: %v>", t.Type, t.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	BANG      = "BANG"
	BACKSLASH = "BACKSLASH"
	ASTERISK  = "ASTERISK"
	LT        = "LT"
	GT        = "GT"
	EQ        = "EQ"
	NEQ       = "NEQ"

	// Special characters
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"

	// Keywords
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywordToTokenType = map[string]TokenType{
	"fn":     FUNCTION,
	"return": RETURN,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func NewToken(tt TokenType, l string) Token {
	return Token{tt, l}
}

func GetTokenType(kw string) TokenType {
	if tt, ok := keywordToTokenType[kw]; ok {
		return tt
	} else {
		return IDENT
	}
}
