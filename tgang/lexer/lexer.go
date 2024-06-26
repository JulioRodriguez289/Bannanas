// lexer/lexer.go
package lexer

type Lexer struct {
	input        string
	position     int  // curr position in input
	readPosition int  // curr reading position
	char         byte //
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
