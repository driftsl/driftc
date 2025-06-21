package driftc

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Lexer struct {
	input []rune

	pos    int
	line   int
	column int
}

func (l *Lexer) advance() {
	if l.pos >= len(l.input) {
		return
	}
	ch := l.input[l.pos]

	if ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}

	l.pos++

	return
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) readWhile(f func(rune) bool) string {
	var builder strings.Builder

	for {
		ch := l.peek()
		if ch == 0 || !f(ch) {
			break
		}
		builder.WriteRune(ch)
		l.advance()
	}

	return builder.String()
}

func (l *Lexer) skipWhitespace() {
	l.readWhile(unicode.IsSpace)
}

func (l *Lexer) readNumber() (string, error) {
	var builder strings.Builder

	dotReached := false

	for {
		ch := l.peek()
		if ch == 0 {
			break
		}

		if ch == '.' {
			if dotReached {
				return builder.String(), errors.New("double dot")
			}
			dotReached = true
		} else if !unicode.IsDigit(ch) {
			break
		}

		builder.WriteRune(ch)
		l.advance()
	}

	return builder.String(), nil
}

func (l *Lexer) readString() (string, error) {
	var builder strings.Builder

	builder.WriteRune('"')

	escaping := false

	l.advance()

	for {
		ch := l.peek()
		if ch == 0 {
			return builder.String(), errors.New("unterminated string")
		}

		l.advance()

		builder.WriteRune(ch)

		if escaping {
			escaping = false
		} else if ch == '\\' {
			escaping = true
		} else if ch == '"' {
			return builder.String(), nil
		}
	}
}

func (l *Lexer) readOperator(ch rune, simpleType TokenType, assignType TokenType) (string, TokenType) {
	next := l.peek()

	if next == '=' {
		l.advance()
		return string(ch) + string(next), assignType
	}

	return string(ch), simpleType
}

func (l *Lexer) nextToken() (Token, error) {
	l.skipWhitespace()
	ch := l.peek()

	token := Token{Column: l.column, Line: l.line, Pos: l.pos}

	if ch == 0 {
		token.Type = TokenEOF
		return token, nil
	}

	// number literals

	if unicode.IsDigit(ch) {
		var err error
		token.Value, err = l.readNumber()

		if strings.Contains(token.Value, ".") {
			token.Type = TokenFloatLiteral
		} else {
			token.Type = TokenIntLiteral
		}

		return token, err
	}

	// string literals

	if ch == '"' {
		token.Type = TokenStringLiteral

		var err error
		token.Value, err = l.readString()
		return token, err
	}

	// keywords & names

	if unicode.IsLetter(ch) {
		token.Value = l.readWhile(func(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) })

		switch token.Value {
		case "let":
			token.Type = TokenLet
		case "return":
			token.Type = TokenReturn

		case "function":
			token.Type = TokenFunction
		case "fragment":
			token.Type = TokenFragment
		case "vertex":
			token.Type = TokenVertex

		case "export":
			token.Type = TokenExport
		case "import":
			token.Type = TokenImport
		case "from":
			token.Type = TokenFrom

		case "float":
			token.Type = TokenFloat
		case "int":
			token.Type = TokenInt
		case "bool":
			token.Type = TokenBoolean
		case "vec2":
			token.Type = TokenVec2
		case "vec3":
			token.Type = TokenVec3
		case "vec4":
			token.Type = TokenVec4
		case "ivec2":
			token.Type = TokenIntVec2
		case "ivec3":
			token.Type = TokenIntVec3
		case "ivec4":
			token.Type = TokenIntVec4
		case "bvec2":
			token.Type = TokenBooleanVec2
		case "bvec3":
			token.Type = TokenBooleanVec3
		case "bvec4":
			token.Type = TokenBooleanVec4

		case "true", "false":
			token.Type = TokenBooleanLiteral

		default:
			token.Type = TokenName
		}

		return token, nil
	}

	// single-character tokens, operators, comments

	l.advance()

	token.Value = string(ch)

	switch ch {
	case ';':
		token.Type = TokenSemicolon
	case ':':
		token.Type = TokenColon
	case ',':
		token.Type = TokenComma
	case '.':
		token.Type = TokenDot

	case '(':
		token.Type = TokenOpenParen
	case ')':
		token.Type = TokenCloseParen
	case '{':
		token.Type = TokenOpenBrace
	case '}':
		token.Type = TokenCloseBrace
	case '[':
		token.Type = TokenOpenBracket
	case ']':
		token.Type = TokenCloseBracket

	case '+':
		token.Value, token.Type = l.readOperator(ch, TokenPlus, TokenPlusAssign)
	case '-':
		token.Value, token.Type = l.readOperator(ch, TokenMinus, TokenMinusAssign)
	case '/':
		if l.peek() == '/' {
			l.advance()
			l.readWhile(func(r rune) bool { return r != '\n' })
			l.advance()
			return l.nextToken()
		}
		token.Value, token.Type = l.readOperator(ch, TokenDivide, TokenDivideAssign)
	case '*':
		token.Value, token.Type = l.readOperator(ch, TokenMultiply, TokenMultiplyAssign)

	case '=':
		token.Type = TokenAssign

	default:
		return token, fmt.Errorf("unknown token starting with %c", ch)
	}

	return token, nil
}

func (l *Lexer) Tokenize(input []rune) ([]Token, error) {
	l.input = input
	l.line = 1
	l.column = 1
	l.pos = 0

	result := make([]Token, 0)

	for {
		token, err := l.nextToken()
		if err != nil {
			return result, err
		}
		if token.Type == TokenEOF {
			break
		}
		result = append(result, token)
	}

	return result, nil
}
