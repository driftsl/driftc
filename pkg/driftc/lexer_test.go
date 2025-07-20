package driftc_test

import (
	"github.com/driftsl/driftc/pkg/driftc"
	"testing"
)

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []driftc.Token
		wantErr bool
	}{
		{
			name:  "literals",
			input: "2 2 4.0 \"hello \\\\ world\"",
			want: []driftc.Token{
				{Type: driftc.TokenIntLiteral, Value: "2"},
				{Type: driftc.TokenIntLiteral, Value: "2"},
				{Type: driftc.TokenFloatLiteral, Value: "4.0"},
				{Type: driftc.TokenStringLiteral, Value: "\"hello \\\\ world\""},
			},
		},
		{
			name:  "divide operator and comments",
			input: "/= / // this is a comment\n//// another comment",
			want: []driftc.Token{
				{Type: driftc.TokenDivideAssign, Value: "/="},
				{Type: driftc.TokenDivide, Value: "/"},
			},
		},
		{
			name:  "bit and logical operators",
			input: "| |= || ||= & &= && &&= == = ! !=",
			want: []driftc.Token{
				{Type: driftc.TokenBitOr, Value: "|"},
				{Type: driftc.TokenBitOrAssign, Value: "|="},
				{Type: driftc.TokenLogicalOr, Value: "||"},
				{Type: driftc.TokenLogicalOrAssign, Value: "||="},
				{Type: driftc.TokenBitAnd, Value: "&"},
				{Type: driftc.TokenBitAndAssign, Value: "&="},
				{Type: driftc.TokenLogicalAnd, Value: "&&"},
				{Type: driftc.TokenLogicalAndAssign, Value: "&&="},
				{Type: driftc.TokenEqual, Value: "=="},
				{Type: driftc.TokenAssign, Value: "="},
				{Type: driftc.TokenNot, Value: "!"},
				{Type: driftc.TokenNotEqual, Value: "!="},
			},
		},
		{
			name:    "unterminated string",
			input:   "\"hello world",
			wantErr: true,
		},
	}

	var l driftc.Lexer
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := l.Tokenize([]rune(tt.input))
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Tokenize() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Tokenize() succeeded unexpectedly")
			}

			if got[len(got)-1].Type != driftc.TokenEOF {
				t.Fatal("Tokenize() does not end with EOF token")
			}

			got = got[:len(got)-1]

			if !compareTokens(got, tt.want) {
				t.Errorf("Tokenize() = %+v\nwant %+v", got, tt.want)
			}
		})
	}
}

func compareTokens(a, b []driftc.Token) bool {
	if len(a) != len(b) {
		return false
	}

	for i, tokenA := range a {
		tokenB := b[i]
		if tokenA.Type != tokenB.Type || tokenA.Value != tokenB.Value {
			return false
		}
	}

	return true
}
