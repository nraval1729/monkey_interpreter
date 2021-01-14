package lexer

import (
	"../token"
	"strconv"
)

type Lexer struct {
	input   string
	currIdx int
	nextIdx int
	ch      byte
}

func New(input string) *Lexer {
	l := Lexer{input, 0, 0, 0}
	l.readChar()

	return &l
}

func (l *Lexer) NextToken() token.Token {
	var currTok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		currTok = token.NewToken(token.LPAREN, "(")
	case ')':
		currTok = token.NewToken(token.RPAREN, ")")
	case '{':
		currTok = token.NewToken(token.LBRACE, "{")
	case '}':
		currTok = token.NewToken(token.RBRACE, "}")
	case ';':
		currTok = token.NewToken(token.SEMICOLON, ";")
	case ',':
		currTok = token.NewToken(token.COMMA, ",")
	case '=':
		nextChar := l.peekNextChar()
		if nextChar == '=' {
			currTok = token.NewToken(token.EQ, "==")
			l.readChar()
		} else {
			currTok = token.NewToken(token.ASSIGN, "=")
		}
	case '+':
		currTok = token.NewToken(token.PLUS, "+")
	case '-':
		currTok = token.NewToken(token.MINUS, "-")
	case '*':
		currTok = token.NewToken(token.ASTERISK, "*")
	case '/':
		currTok = token.NewToken(token.BACKSLASH, "/")
	case '!':
		nextChar := l.peekNextChar()
		if nextChar == '=' {
			currTok = token.NewToken(token.NEQ, "!=")
			l.readChar()
		} else {
			currTok = token.NewToken(token.BANG, "!")
		}
	case '>':
		currTok = token.NewToken(token.GT, ">")
	case '<':
		currTok = token.NewToken(token.LT, "<")
	case 0:
		currTok = token.NewToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			currTok = token.NewToken(token.GetTokenType(literal), literal)
		} else if isInt(l.ch) {
			currTok = token.NewToken(token.INT, l.readInt())
		} else {
			currTok = token.NewToken(token.ILLEGAL, "")
		}
		return currTok
	}
	l.readChar()

	return currTok
}

func (l *Lexer) readChar() {
	if l.nextIdx >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextIdx]
	}
	l.currIdx = l.nextIdx
	l.nextIdx++
}

func (l *Lexer) peekNextChar() byte {
	if l.nextIdx < len(l.input) {
		return l.input[l.nextIdx]
	}
	return 0
}

func (l *Lexer) readIdentifier() string {
	startIdx := l.currIdx
	for isLetter(l.input[l.currIdx]) {
		l.readChar()
	}
	return l.input[startIdx:l.currIdx]
}

func (l *Lexer) readInt() string {
	startIdx := l.currIdx

	for isInt(l.input[l.currIdx]) {
		l.readChar()
	}

	return l.input[startIdx:l.currIdx]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isInt(c byte) bool {
	_, err := strconv.Atoi(string(c))

	return err == nil
}
