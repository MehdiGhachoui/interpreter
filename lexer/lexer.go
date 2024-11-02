package lexer

import "go-interpreter/token"

// The way to turn the source code into tokens
// our lexer only support ASCII characters

type Lexer struct {
	input        string
	position     int  // point to the character in the input that corresponds to the ch byte
	readPosition int  // points to the next character in the input
	ch           byte // current chat under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Give us next character and advance our position in the input string
// set ch to 0 once we read end of input ; 0 is ASCII code to "NUL"
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIndentifier()
			tk.Type = token.LookupIdent(tk.Literal)
			return tk
		} else if isNumber(l.ch) {
			tk.Literal = l.readNumber()
			tk.Type = token.INT
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tk
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Reads Indetifier and advances lexer position unti it encounters non-letter-characters
func (l *Lexer) readIndentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Checks whether a character is a letter
// in this case we are allowing "_" for naming non-letter-characters
// we can also add ! or ? , ...
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Note : we are only reading int in here
// float,hex, ...
func (l *Lexer) readNumber() string {
	position := l.position

	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
