// Code generated from GrammarLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarlexerLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "LPAREN", "RPAREN", "ADD", "SUB", "MUL", "DIV", "ID", "NUMBER",
		"WS",
	}
	staticData.RuleNames = []string{
		"LPAREN", "RPAREN", "ADD", "SUB", "MUL", "DIV", "ID", "NUMBER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 59, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 5,
		6, 34, 8, 6, 10, 6, 12, 6, 37, 9, 6, 1, 7, 4, 7, 40, 8, 7, 11, 7, 12, 7,
		41, 1, 7, 1, 7, 5, 7, 46, 8, 7, 10, 7, 12, 7, 49, 9, 7, 3, 7, 51, 8, 7,
		1, 8, 4, 8, 54, 8, 8, 11, 8, 12, 8, 55, 1, 8, 1, 8, 0, 0, 9, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0, 4, 2, 0, 65, 90, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10,
		12, 13, 32, 32, 63, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3, 21, 1,
		0, 0, 0, 5, 23, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9, 27, 1, 0, 0, 0, 11, 29,
		1, 0, 0, 0, 13, 31, 1, 0, 0, 0, 15, 39, 1, 0, 0, 0, 17, 53, 1, 0, 0, 0,
		19, 20, 5, 40, 0, 0, 20, 2, 1, 0, 0, 0, 21, 22, 5, 41, 0, 0, 22, 4, 1,
		0, 0, 0, 23, 24, 5, 43, 0, 0, 24, 6, 1, 0, 0, 0, 25, 26, 5, 45, 0, 0, 26,
		8, 1, 0, 0, 0, 27, 28, 5, 42, 0, 0, 28, 10, 1, 0, 0, 0, 29, 30, 5, 47,
		0, 0, 30, 12, 1, 0, 0, 0, 31, 35, 7, 0, 0, 0, 32, 34, 7, 1, 0, 0, 33, 32,
		1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0,
		36, 14, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 40, 7, 2, 0, 0, 39, 38, 1,
		0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42,
		50, 1, 0, 0, 0, 43, 47, 5, 46, 0, 0, 44, 46, 7, 2, 0, 0, 45, 44, 1, 0,
		0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 51,
		1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 43, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0,
		51, 16, 1, 0, 0, 0, 52, 54, 7, 3, 0, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1,
		0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57,
		58, 6, 8, 0, 0, 58, 18, 1, 0, 0, 0, 6, 0, 35, 41, 47, 50, 55, 1, 6, 0,
		0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GrammarLexerInit initializes any static state used to implement GrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.once.Do(grammarlexerLexerInit)
}

// NewGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGrammarLexer(input antlr.CharStream) *GrammarLexer {
	GrammarLexerInit()
	l := new(GrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "GrammarLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrammarLexer tokens.
const (
	GrammarLexerLPAREN = 1
	GrammarLexerRPAREN = 2
	GrammarLexerADD    = 3
	GrammarLexerSUB    = 4
	GrammarLexerMUL    = 5
	GrammarLexerDIV    = 6
	GrammarLexerID     = 7
	GrammarLexerNUMBER = 8
	GrammarLexerWS     = 9
)
