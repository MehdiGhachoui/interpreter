package parser

import (
	"fmt"
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"go-interpreter/token"
)

const (
	_ int = iota // 0
	LOWEST
	EQUALS
	LESSGRATER // < or >
	SUM
	PRODUCT
	PREFIX //-x or !x
	CALL   //func(x)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)
type Parser struct {
	l *lexer.Lexer

	currentToken  token.Token
	peekToken     token.Token
	errors        []string
	prefixParseFn map[token.TokenType]prefixParseFn
	infixParseFn  map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFn[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFn[tokenType] = fn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Current is Null at first call
	p.nextToken()
	p.nextToken()

	p.prefixParseFn = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	return p
}

// checking that the return type  implements the interface function
// it's still a pointer just making sure it implements the interface
// it's an interface as long a it implements it's functions regardless of the actual type
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Create the root node of the AST
// Iterate over every token until we hit EOF
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	//Creating an empty slice of **Interface Types;
	//**Implement the interface <Statement> methods
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// Return the type of the statement we are checking
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.currentToken,
	}

	//Expect next to be a var name
	//else error
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	//Expect the token after to be an asign
	//else error
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//TODO: handle<expression>/Value
	//....
	p.nextToken()

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: p.currentToken,
	}

	//TODO: Handle <expression>/Value
	p.nextToken()

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.currentToken,
	}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precendece int) ast.Expression {
	//check whether we have a parsing function associated with currentToken TYPE
	prefix := p.prefixParseFn[p.currentToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s. got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
