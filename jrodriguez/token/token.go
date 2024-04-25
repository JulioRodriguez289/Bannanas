package token 


// token type defined as string such that we are able to distinguish btwn diff types of tokens 

type TokenType string 

type Token struct{
	Type TokenType
	Literal string 
}

//  token types defined as of chapter 1 

const (
	ILLEGAL = "ILLEGAL" // a token that we dont know about 
	EOF     = "EOF" // signifies end of file, tells parser to stop 

	//Identifiers + literal
	IDENT = "IDENT" //add, foobar, x,y,...
	INT   = "INT"   //1234...

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "()"
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

