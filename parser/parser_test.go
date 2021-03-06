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

func TestParseIdentifierExpression(t *testing.T) {
	input := `foobar;`

	parser := New(lexer.New(input))
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf("parser.ParseProgram() returned nil\n")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements doesn't contain 1 statement. Got = %d\n", len(program.Statements))
	}

	expStmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt %v is not an ast.ExpressionStatement\n", program.Statements[0])
	}

	ident, ok := expStmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expStmt.Expression is not of type identifier. Got %T\n", expStmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident is not of value foobar. Got = %s\n", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() is not foobar. Got = %s\n", ident.TokenLiteral())
	}
}

func TestParseIntegerLiteral(t *testing.T) {
	input := `5;`

	parser := New(lexer.New(input))
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf("parser.ParseProgram() returned nil\n")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements doesn't contain 1 statement. Got = %d\n", len(program.Statements))
	}

	expStmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt %v is not an ast.ExpressionStatement\n", program.Statements[0])
	}

	intLiteral, ok := expStmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expStmt.Expression is not of type IntegerLiteral. Got %T\n", expStmt.Expression)
	}

	if intLiteral.Value != 5 {
		t.Fatalf("intLiteral is not of value 5. Got = %d\n", intLiteral.Value)
	}

	if intLiteral.TokenLiteral() != "5" {
		t.Fatalf("intLiteral.TokenLiteral() is not 5. Got = %s\n", intLiteral.TokenLiteral())
	}
}

func TestParsePrefixExpressions(t *testing.T) {
	tt := []struct {
		input        string
		operator     string
		integerValue int
	}{
		{input: "!5;", operator: "!", integerValue: 5},
		{input: "-15;", operator: "-", integerValue: 15},
	}

	for _, tc := range tt {
		parser := New(lexer.New(tc.input))
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if program == nil {
			t.Fatalf("parser.ParseProgram() returned nil\n")
		}

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements doesn't contain 1 statement. Got = %d\n", len(program.Statements))
		}

		prefixStmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("stmt %v is not an ast.ExpressionStatement\n", program.Statements[0])
		}

		exp, ok := prefixStmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("prefixStmt.Expression is not of type ast.PrefixExpression. Got %T\n", prefixStmt.Expression)
		}

		if exp.Operator != tc.operator {
			t.Fatalf("Expected exp.Operator to be %s. Got %s\n", exp.Operator, tc.operator)
		}

		if !testIntegerLiteral(t, exp.Right, tc.integerValue) {
			return
		}
	}
}

// helper functions

func testIntegerLiteral(t *testing.T, il ast.Expression, value int) bool {
	integer, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("Expected ast.Expression to be an ast.IntegerLiteral. Got = %T\n", il)
		return false
	}

	if integer.Value != value {
		t.Errorf("Expected integer.value to be %d. Got = %d\n", integer.Value, value)
		return false
	}

	if integer.Token.Type != token.INT {
		t.Errorf("Expected integer.Token.Literal to be %s. Got = %s\n", token.INT, integer.Token.Literal)
		return false
	}

	return true
}

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
