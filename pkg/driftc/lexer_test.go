package driftc_test

import (
	"fmt"
	"testing"

	"github.com/driftsl/driftc/pkg/driftc"
)

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		want          []driftc.Token
		wantErr       bool
		parseComments bool
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
			name:  "divide operator and comments (parse comments)",
			input: "/= / // this is a comment\n//// another comment\n/",
			want: []driftc.Token{
				{Type: driftc.TokenDivideAssign, Value: "/="},
				{Type: driftc.TokenDivide, Value: "/"},
				{Type: driftc.TokenComment, Value: "// this is a comment"},
				{Type: driftc.TokenComment, Value: "//// another comment"},
				{Type: driftc.TokenDivide, Value: "/"},
			},
			parseComments: true,
		},
		{
			name:  "divide operator and comments (ignore comments)",
			input: "/= / // this is a comment\n//// another comment\n/",
			want: []driftc.Token{
				{Type: driftc.TokenDivideAssign, Value: "/="},
				{Type: driftc.TokenDivide, Value: "/"},
				{Type: driftc.TokenDivide, Value: "/"},
			},
		},
		{
			name:  "bit and logical operators",
			input: "| | |= || ||= & & &= && &&= == = ! !=",
			want: []driftc.Token{
				{Type: driftc.TokenBitOr, Value: "|"},
				{Type: driftc.TokenBitOr, Value: "|"},
				{Type: driftc.TokenBitOrAssign, Value: "|="},
				{Type: driftc.TokenLogicalOr, Value: "||"},
				{Type: driftc.TokenLogicalOrAssign, Value: "||="},
				{Type: driftc.TokenBitAnd, Value: "&"},
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
			got, gotErr := l.Tokenize([]rune(tt.input), tt.parseComments)
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
				gotString := ""
				for i, token := range got {
					gotString += fmt.Sprintf("\t%d: \t%+v\n", i, token)
				}

				wantString := ""
				for i, token := range tt.want {
					wantString += fmt.Sprintf("\n\t%d: \t%+v", i, token)
				}

				t.Errorf("\nTokenize():\n%swant:%s", gotString, wantString)
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
