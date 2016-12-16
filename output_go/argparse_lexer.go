// Generated from ArgParse.g4 by ANTLR 4.6.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 8, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	5, 3, 27, 10, 3, 3, 4, 6, 4, 30, 10, 4, 13, 4, 14, 4, 31, 3, 5, 6, 5, 35,
	10, 5, 13, 5, 14, 5, 36, 3, 5, 3, 5, 3, 6, 6, 6, 42, 10, 6, 13, 6, 14,
	6, 43, 3, 7, 3, 7, 3, 7, 7, 7, 49, 10, 7, 12, 7, 14, 7, 52, 11, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 59, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	2, 17, 2, 19, 2, 3, 2, 7, 4, 2, 12, 12, 15, 15, 6, 2, 50, 59, 67, 92, 97,
	97, 99, 124, 4, 2, 36, 36, 94, 94, 10, 2, 36, 36, 49, 49, 94, 94, 100,
	100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99,
	104, 71, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 26,
	3, 2, 2, 2, 7, 29, 3, 2, 2, 2, 9, 34, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13,
	45, 3, 2, 2, 2, 15, 55, 3, 2, 2, 2, 17, 60, 3, 2, 2, 2, 19, 66, 3, 2, 2,
	2, 21, 22, 7, 63, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 47, 2, 2, 24, 27,
	7, 47, 2, 2, 25, 27, 7, 47, 2, 2, 26, 23, 3, 2, 2, 2, 26, 25, 3, 2, 2,
	2, 27, 6, 3, 2, 2, 2, 28, 30, 7, 34, 2, 2, 29, 28, 3, 2, 2, 2, 30, 31,
	3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 8, 3, 2, 2, 2,
	33, 35, 9, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 34, 3,
	2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 8, 5, 2, 2, 39,
	10, 3, 2, 2, 2, 40, 42, 9, 3, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2,
	2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 12, 3, 2, 2, 2, 45, 50,
	7, 36, 2, 2, 46, 49, 5, 15, 8, 2, 47, 49, 10, 4, 2, 2, 48, 46, 3, 2, 2,
	2, 48, 47, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51,
	3, 2, 2, 2, 51, 53, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7, 36, 2, 2,
	54, 14, 3, 2, 2, 2, 55, 58, 7, 94, 2, 2, 56, 59, 9, 5, 2, 2, 57, 59, 5,
	17, 9, 2, 58, 56, 3, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 16, 3, 2, 2, 2, 60,
	61, 7, 119, 2, 2, 61, 62, 5, 19, 10, 2, 62, 63, 5, 19, 10, 2, 63, 64, 5,
	19, 10, 2, 64, 65, 5, 19, 10, 2, 65, 18, 3, 2, 2, 2, 66, 67, 9, 6, 2, 2,
	67, 20, 3, 2, 2, 2, 10, 2, 26, 31, 36, 43, 48, 50, 58, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "DASH", "SPACE", "NEWLINE", "TEXT", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "DASH", "SPACE", "NEWLINE", "TEXT", "STRING", "ESC", "UNICODE",
	"HEX",
}

type ArgParseLexer struct {
	*antlr.BaseLexer
	modeNames []string
	// TODO: EOF string
}

func NewArgParseLexer(input antlr.CharStream) *ArgParseLexer {
	var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	l := new(ArgParseLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ArgParse.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ArgParseLexer tokens.
const (
	ArgParseLexerT__0    = 1
	ArgParseLexerDASH    = 2
	ArgParseLexerSPACE   = 3
	ArgParseLexerNEWLINE = 4
	ArgParseLexerTEXT    = 5
	ArgParseLexerSTRING  = 6
)
