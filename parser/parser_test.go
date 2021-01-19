package parser

import (
	"../ast"
	"../lexer"
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