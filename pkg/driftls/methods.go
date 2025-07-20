package driftls

import (
	"encoding/json"

	"github.com/driftsl/driftc/pkg/driftc"
)

func (s *Server) initialize(id any) error {
	var initializeResult struct {
		Capabilities struct {
			TextDocumentSync int `json:"textDocumentSync"`

			SemanticTokensProvider struct {
				Legend struct {
					TokenTypes     []string `json:"tokenTypes"`
					TokenModifiers []string `json:"tokenModifiers"`
				} `json:"legend"`

				Full bool `json:"full"`
			} `json:"semanticTokensProvider"`
		} `json:"capabilities"`

		ServerInfo struct {
			Name string `json:"name"`
		} `json:"serverInfo"`
	}

	initializeResult.ServerInfo.Name = "driftls"

	initializeResult.Capabilities.TextDocumentSync = 1

	initializeResult.Capabilities.SemanticTokensProvider.Full = true
	initializeResult.Capabilities.SemanticTokensProvider.Legend.TokenTypes = tokensArray[:]
	initializeResult.Capabilities.SemanticTokensProvider.Legend.TokenModifiers = make([]string, 0)

	return s.sendResponse(id, initializeResult)
}

func (s *Server) sendTokens(id any, rawParams json.RawMessage) error {
	var params DocumentParams[TextDocumentIdentifier]

	if err := json.Unmarshal(rawParams, &params); err != nil {
		return err
	}

	lexer := driftc.Lexer{}

	tokens, err := lexer.Tokenize([]rune(s.documents.Get(params.TextDocument.Uri)), true)
	if err != nil {
		return err
	}

	var result struct {
		Data []uint `json:"data"`
	}

	result.Data = make([]uint, 0, len(tokens)*5)

	prevLine := -1
	prevColumn := 0

	for _, tok := range tokens {
		tokenType := mapTokenType(tok.Type)
		if tokenType < 0 {
			continue
		}

		line := tok.Line - 1
		column := tok.Column - 1

		deltaLine := line
		if prevLine != -1 {
			deltaLine -= prevLine
		}

		deltaStart := column
		if prevLine == line {
			deltaStart -= prevColumn
		}

		length := len(tok.Value)

		result.Data = append(result.Data,
			uint(deltaLine),
			uint(deltaStart),
			uint(length),
			uint(tokenType),
			0,
		)

		prevLine, prevColumn = line, column
	}

	return s.sendResponse(id, result)
}
