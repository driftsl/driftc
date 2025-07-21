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

	return s.sendServerResponse(id, initializeResult)
}

type Diagnostic struct {
	Range    Range  `json:"range"`
	Severity int    `json:"severity"`
	Code     *int   `json:"code,omitempty"`
	Source   string `json:"source"`
	Message  string `json:"message"`
}

func (s *Server) sendTokens(id any, rawParams json.RawMessage) error {
	var params DocumentParams[TextDocumentIdentifier]

	if err := json.Unmarshal(rawParams, &params); err != nil {
		return err
	}

	lexer := driftc.Lexer{ParseAllErrors: true, ParseComments: true}

	tokens, errors := lexer.Tokenize([]rune(s.documents.Get(params.TextDocument.Uri)))

	var notification struct {
		TextDocumentIdentifier
		Diagnostics []Diagnostic `json:"diagnostics"`
	}

	notification.Uri = params.TextDocument.Uri
	notification.Diagnostics = make([]Diagnostic, 0)

	for _, err := range errors {
		notification.Diagnostics = append(notification.Diagnostics, Diagnostic{
			Range: Range{
				Start: Position{
					Line:      err.Token.Line - 1,
					Character: err.Token.Column - 1,
				},
				End: Position{
					Line:      err.Token.Line - 1,
					Character: err.Token.Column - 1 + len(err.Token.Value),
				},
			},
			Severity: 1,
			Source:   "lexer",
			Message:  err.Err.Error(),
		})
	}

	s.sendNotification("textDocument/publishDiagnostics", notification)

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
		if prevLine != -1 { // absoulute, if first token
			deltaLine -= prevLine
		}

		deltaStart := column
		if prevLine == line { // delta, if on another line
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

	return s.sendServerResponse(id, result)
}
