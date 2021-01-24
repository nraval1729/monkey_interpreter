package parser

import (
	"../ast"
	"../lexer"
	"../token"
	"testing"
)

func TestParseLetStatements(t *testing.T) {
	input := `
let five = 5;
let ten = 10;
let foobar = 17290022;
`

	parser := New(lexer.New(input))
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf("parser.ParseProgram() returned nil\n")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn't contain 3 statements. Got = %d\n", len(program.Statements))
	}

	tests := []struct{ expectedIdentifier string }{
		{"five"},
		{"ten"},
		{"foobar"},
	}

	for i, test := range tests {
		if !testLetStatement(t, program.Statements[i], test.expectedIdentifier) {
			return
		}
	}

}

func TestParseReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 17290022;
`

	parser := New(lexer.New(input))
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf("parser.ParseProgram() returned nil\n")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn't contain 3 statements. Got = %d\n", len(program.Statements))
	}

	tests := []struct{ expectedIdentifier string }{
		{"5"},
		{"10"},
		{"17290022"},
	}

	for i, test := range tests {
		if !testReturnStatement(t, program.Statements[i], test.expectedIdentifier) {
			return
		}
	}
}

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.NewToken(token.LET, "let"),
				Identifier: &ast.Identifier{
					Token: token.NewToken(token.IDENT, "ident"),
					Value: "foo",
				},
				Value: &ast.Identifier{
					Token: token.NewToken(token.IDENT, "ident"),
					Value: "bar",
				},
			},
		},
	}

	if program.String() != "let foo = bar;" {
		t.Errorf("program.String() wrong. Expected let foo = bar; but got %s\n", program.String())
	}
}

/*
	1. Test that stmt.TokenLiteral() == "let"
	2. Test that stmt is of type LetStatement
	3. Test that stmt.Identifier.TokenLiteral() == "let"
	4. Test that stmt.Identifier.Value == expectedIdentifierValue
*/
func testLetStatement(t *testing.T, stmt ast.Statement, expectedIdentifierValue string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral() is not let for stmt: %v\n", stmt)
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt %v is not an ast.LetStatement\n", stmt)
		return false
	}

	if letStmt.Identifier.TokenLiteral() != "ident" {
		t.Errorf("letStmt.Identifier.TokenLiteral() is not let. Got = %s\n", letStmt.Identifier.TokenLiteral())
		return false
	}

	if letStmt.Identifier.Value != expectedIdentifierValue {
		t.Errorf("letStmt.Identifier.Value doesn't match expectation. Expected = %s, got = %s\n", expectedIdentifierValue, letStmt.Identifier.Value)
		return false
	}

	return true
}

/*
	1. Test that stmt.TokenLiteral() == "return"
	2. Test that stmt is of type ReturnStatement
*/
func testReturnStatement(t *testing.T, stmt ast.Statement, expectedIdentifierValue string) bool {
	if stmt.TokenLiteral() != "return" {
		t.Errorf("stmt.TokenLiteral() is not return for stmt: %v\n", stmt)
	}

	returnStatement, ok := stmt.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("stmt %v is not an ast.ReturnStatement\n", stmt)
		return false
	}

	if returnStatement.TokenLiteral() != "return" {
		t.Errorf("returnStatement.TokenLiteral() is not return. Got = %s\n", returnStatement.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errs := p.Errors()

	if len(errs) == 0 {
		return
	}

	t.Errorf("parser has %d errors\n", len(errs))

	for _, err := range errs {
		t.Errorf("parser error: %s\n", err)
	}
	t.FailNow()
}
