package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special Characters
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers & Literals
	IDENT = "IDENT" // add, foobar, x, y, ..
	INT   = "INT"   // 123254

	// Operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	ASTERIX = "*"
	BANG    = "!"
	SLASH   = "/"

	LT = "<"
	RT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keyboards
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
