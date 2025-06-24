package driftc

import (
	"fmt"
	"slices"
)

type Parser struct {
	tokens []Token
	pos    int
}

func (p *Parser) peek() Token {
	return p.tokens[p.pos]
}

func (p *Parser) expect(tokenTypes ...TokenType) (Token, error) {
	token := p.peek()

	p.pos += 1

	if slices.Contains(tokenTypes, token.Type) {
		return token, nil
	}

	var expectedString string

	if len(tokenTypes) == 1 {
		expectedString = "'" + string(tokenTypes[0]) + "'"
	} else {
		for _, t := range tokenTypes {
			expectedString += "'" + string(t) + "' or "
		}

		expectedString = expectedString[:len(expectedString)-4]
	}

	return token, fmt.Errorf("unexpected token '%s', want %s", token.Value, expectedString)
}

func (p *Parser) Parse(tokens []Token) (RootNode, error) {
	p.tokens = tokens
	p.pos = 0

	result := RootNode{
		Imports: make([]ImportNode, 0),
	}

	for {
		token, err := p.expect(TokenImport, TokenEOF)
		if err != nil {
			return result, err
		}

		switch token.Type {
		case TokenImport:
			importNode, err := p.parseImport()
			if err != nil {
				return result, err
			}
			if _, err := p.expect(TokenSemicolon); err != nil {
				return result, err
			}
			result.Imports = append(result.Imports, importNode)
		case TokenEOF:
			return result, nil
		}
	}
}

func (p *Parser) parseImport() (ImportNode, error) {
	var result ImportNode

	token, err := p.expect(TokenOpenBrace, TokenName)
	if err != nil {
		return result, err
	}

	switch token.Type {
	case TokenName:
		result.To = token
	case TokenOpenBrace:
		result.To, err = p.parseObjectDestructuring()
		if err != nil {
			return result, err
		}
	}

	if _, err := p.expect(TokenFrom); err != nil {
		return result, err
	}

	result.From, err = p.expect(TokenStringLiteral)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *Parser) parseObjectDestructuring() (DeconstructionNode, error) {
	result := DeconstructionNode{
		NameTokens: make([]Token, 0),
	}

	name := true

	for {
		var expectedType TokenType

		if name {
			expectedType = TokenName
		} else {
			expectedType = TokenComma
		}

		token, err := p.expect(TokenCloseBrace, expectedType)
		if err != nil {
			return result, err
		}

		switch token.Type {
		case TokenName:
			result.NameTokens = append(result.NameTokens, token)
		case TokenCloseBrace:
			return result, nil
		}

		name = !name
	}
}
