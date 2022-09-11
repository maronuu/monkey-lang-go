package lexer

type Lexer struct {
	input string
	position int // current
	readPosition int // next 
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition // read done
	l.readPosition += 1	// go to next position
}