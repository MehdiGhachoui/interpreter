package parser

import (
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParserError(t, parser)

	if program == nil {
		t.Fatalf("Parse Program -- returned NIL")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Program does not contain the 3 statements. got %d", len(program.Statements))
	}

	test := []struct {
		expectedIndentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for idx, tt := range test {
		stmt := program.Statements[idx]
		if !testLetStatement(t, stmt, tt.expectedIndentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("Token Literal not 'let'. got %q", stmt.TokenLiteral())
		return false
	}

	//golang type assertion :
	//asserts that the interface value stmt holds the concrete type (...) and assigns the underlying (...) value
	//to the variable 'letStmt'.
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not *ast.LetStatement. got=%T", stmt)
		return false
	}

	//Name is the Indentifier
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value != '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral != '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserError(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}
