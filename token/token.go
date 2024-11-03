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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

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
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

type Token struct {
	Type    TokenType
	Literal string // holds the value of the token
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Check if it's keyword or Identifier
func LookupIdent(ident string) TokenType {
	if tk, ok := keywords[ident]; ok {
		return tk
	}
	return IDENT
}
