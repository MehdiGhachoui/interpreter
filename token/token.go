package token

// Tokens are representation of our source code

type TokenType string

const (
	ILLEGAL = "ILLEGAL" // Token/character we don't know about
	EOF     = "EOF"     // END OF FILE

	// Identifiers + literal
	IDENT = "IDENT" // var names,func names ...
	INT   = "INT"   // the type currently only INT

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string // holds the value of the token
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Check if it's keyword or Identifier
func LookupIdent(ident string) TokenType {
	if tk, ok := keywords[ident]; ok {
		return tk
	}
	return IDENT
}
