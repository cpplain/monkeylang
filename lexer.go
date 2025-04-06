package main

type Lexer struct {
	input     string
	tokens    []Token
	charIndex int
}

func newLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() byte {
	if l.charIndex >= len(l.input) {
		return 0
	}

	return l.input[l.charIndex]
}

func (l *Lexer) peekChar() byte {
	if l.charIndex+1 >= len(l.input) {
		return 0
	}

	return l.input[l.charIndex+1]
}

func (l *Lexer) skipWhitespace() {
	for l.readChar() == ' ' || l.readChar() == '\t' || l.readChar() == '\n' || l.readChar() == '\r' {
		l.charIndex += 1
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.charIndex

	for isLetter(l.peekChar()) {
		l.charIndex += 1
	}

	return l.input[start : l.charIndex+1]
}

func (l *Lexer) readNumber() {
	for isDigit(l.peekChar()) {
		l.charIndex += 1
	}
}

func (l *Lexer) tokenize() {
	for l.charIndex <= len(l.input) {
		l.skipWhitespace()

		tok := Token{start: l.charIndex}
		ch := l.readChar()

		switch ch {
		case '=':
			if l.peekChar() == '=' {
				l.charIndex += 1
				tok.tag = equalEqual
			} else {
				tok.tag = equal
			}
		case '+':
			tok.tag = plus
		case '-':
			tok.tag = minus
		case '!':
			if l.peekChar() == '=' {
				l.charIndex += 1
				tok.tag = bangEqual
			} else {
				tok.tag = bang
			}
		case '/':
			tok.tag = slash
		case '*':
			tok.tag = asterisk
		case '<':
			tok.tag = lAngleBracket
		case '>':
			tok.tag = rAngleBracket
		case ',':
			tok.tag = comma
		case ';':
			tok.tag = semicolon
		case '(':
			tok.tag = lParen
		case ')':
			tok.tag = rParen
		case '{':
			tok.tag = lBrace
		case '}':
			tok.tag = rBrace
		case 0:
			tok.tag = eof
		default:
			if isLetter(ch) {
				identifier := l.readIdentifier()
				tok.tag = lookupIdent(identifier)
			} else if isDigit(ch) {
				l.readNumber()
				tok.tag = numberLiteral
			} else {
				tok.tag = illegal
			}
		}

		tok.end = l.charIndex
		l.tokens = append(l.tokens, tok)
		l.charIndex += 1
	}
}
