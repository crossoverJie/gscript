// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 228,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 3, 7, 3, 54, 10, 3, 12, 3, 14,
	3, 57, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 5, 6, 5, 66,
	10, 5, 13, 5, 14, 5, 67, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 80, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 97, 10, 6, 12, 6,
	14, 6, 100, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 5, 8, 112, 10, 8, 3, 8, 3, 8, 5, 8, 116, 10, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 5, 8, 122, 10, 8, 3, 9, 5, 9, 125, 10, 9, 3, 9, 3, 9, 5, 9, 129,
	10, 9, 3, 9, 3, 9, 5, 9, 133, 10, 9, 3, 10, 3, 10, 5, 10, 137, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 143, 10, 11, 12, 11, 14, 11, 146, 11,
	11, 3, 12, 3, 12, 3, 12, 5, 12, 151, 10, 12, 3, 13, 3, 13, 3, 13, 7, 13,
	156, 10, 13, 12, 13, 14, 13, 159, 11, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 5, 15, 168, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 18, 5, 18, 177, 10, 18, 3, 18, 3, 18, 7, 18, 181, 10, 18, 12,
	18, 14, 18, 184, 11, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	192, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 7, 21, 199, 10, 21, 12,
	21, 14, 21, 202, 11, 21, 3, 22, 3, 22, 5, 22, 206, 10, 22, 3, 23, 3, 23,
	3, 23, 7, 23, 211, 10, 23, 12, 23, 14, 23, 214, 11, 23, 3, 24, 3, 24, 3,
	24, 7, 24, 219, 10, 24, 12, 24, 14, 24, 222, 11, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 2, 3, 10, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 7, 3, 2, 12, 13, 3,
	2, 14, 15, 4, 2, 18, 19, 21, 22, 4, 2, 20, 20, 23, 23, 3, 2, 36, 37, 2,
	237, 2, 50, 3, 2, 2, 2, 4, 55, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 65, 3,
	2, 2, 2, 10, 79, 3, 2, 2, 2, 12, 101, 3, 2, 2, 2, 14, 121, 3, 2, 2, 2,
	16, 124, 3, 2, 2, 2, 18, 136, 3, 2, 2, 2, 20, 138, 3, 2, 2, 2, 22, 147,
	3, 2, 2, 2, 24, 152, 3, 2, 2, 2, 26, 160, 3, 2, 2, 2, 28, 167, 3, 2, 2,
	2, 30, 169, 3, 2, 2, 2, 32, 171, 3, 2, 2, 2, 34, 176, 3, 2, 2, 2, 36, 185,
	3, 2, 2, 2, 38, 187, 3, 2, 2, 2, 40, 195, 3, 2, 2, 2, 42, 205, 3, 2, 2,
	2, 44, 207, 3, 2, 2, 2, 46, 215, 3, 2, 2, 2, 48, 223, 3, 2, 2, 2, 50, 51,
	5, 4, 3, 2, 51, 3, 3, 2, 2, 2, 52, 54, 5, 6, 4, 2, 53, 52, 3, 2, 2, 2,
	54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 5, 3, 2,
	2, 2, 57, 55, 3, 2, 2, 2, 58, 59, 5, 20, 11, 2, 59, 60, 7, 3, 2, 2, 60,
	63, 3, 2, 2, 2, 61, 63, 5, 14, 8, 2, 62, 58, 3, 2, 2, 2, 62, 61, 3, 2,
	2, 2, 63, 7, 3, 2, 2, 2, 64, 66, 5, 10, 6, 2, 65, 64, 3, 2, 2, 2, 66, 67,
	3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 70, 7, 2, 2, 3, 70, 9, 3, 2, 2, 2, 71, 72, 8, 6, 1, 2, 72, 73, 7, 4,
	2, 2, 73, 74, 5, 10, 6, 2, 74, 75, 7, 5, 2, 2, 75, 80, 3, 2, 2, 2, 76,
	80, 5, 28, 15, 2, 77, 78, 7, 15, 2, 2, 78, 80, 5, 10, 6, 8, 79, 71, 3,
	2, 2, 2, 79, 76, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 98, 3, 2, 2, 2, 81,
	82, 12, 7, 2, 2, 82, 83, 9, 2, 2, 2, 83, 97, 5, 10, 6, 8, 84, 85, 12, 6,
	2, 2, 85, 86, 7, 16, 2, 2, 86, 97, 5, 10, 6, 7, 87, 88, 12, 5, 2, 2, 88,
	89, 9, 3, 2, 2, 89, 97, 5, 10, 6, 6, 90, 91, 12, 4, 2, 2, 91, 92, 9, 4,
	2, 2, 92, 97, 5, 10, 6, 5, 93, 94, 12, 3, 2, 2, 94, 95, 9, 5, 2, 2, 95,
	97, 5, 10, 6, 4, 96, 81, 3, 2, 2, 2, 96, 84, 3, 2, 2, 2, 96, 87, 3, 2,
	2, 2, 96, 90, 3, 2, 2, 2, 96, 93, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98,
	96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 11, 3, 2, 2, 2, 100, 98, 3, 2,
	2, 2, 101, 102, 7, 6, 2, 2, 102, 103, 5, 4, 3, 2, 103, 104, 7, 7, 2, 2,
	104, 13, 3, 2, 2, 2, 105, 122, 5, 12, 7, 2, 106, 107, 7, 27, 2, 2, 107,
	108, 5, 48, 25, 2, 108, 111, 5, 14, 8, 2, 109, 110, 7, 28, 2, 2, 110, 112,
	5, 14, 8, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 122, 3, 2,
	2, 2, 113, 115, 7, 29, 2, 2, 114, 116, 5, 10, 6, 2, 115, 114, 3, 2, 2,
	2, 115, 116, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 122, 7, 3, 2, 2, 118,
	119, 5, 10, 6, 2, 119, 120, 7, 3, 2, 2, 120, 122, 3, 2, 2, 2, 121, 105,
	3, 2, 2, 2, 121, 106, 3, 2, 2, 2, 121, 113, 3, 2, 2, 2, 121, 118, 3, 2,
	2, 2, 122, 15, 3, 2, 2, 2, 123, 125, 5, 18, 10, 2, 124, 123, 3, 2, 2, 2,
	124, 125, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 128, 7, 3, 2, 2, 127,
	129, 5, 10, 6, 2, 128, 127, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130,
	3, 2, 2, 2, 130, 132, 7, 3, 2, 2, 131, 133, 5, 46, 24, 2, 132, 131, 3,
	2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 17, 3, 2, 2, 2, 134, 137, 5, 20, 11,
	2, 135, 137, 5, 46, 24, 2, 136, 134, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2,
	137, 19, 3, 2, 2, 2, 138, 139, 5, 34, 18, 2, 139, 144, 5, 22, 12, 2, 140,
	141, 7, 8, 2, 2, 141, 143, 5, 22, 12, 2, 142, 140, 3, 2, 2, 2, 143, 146,
	3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 21, 3, 2,
	2, 2, 146, 144, 3, 2, 2, 2, 147, 150, 5, 24, 13, 2, 148, 149, 7, 17, 2,
	2, 149, 151, 5, 26, 14, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2,
	151, 23, 3, 2, 2, 2, 152, 157, 7, 35, 2, 2, 153, 154, 7, 9, 2, 2, 154,
	156, 7, 10, 2, 2, 155, 153, 3, 2, 2, 2, 156, 159, 3, 2, 2, 2, 157, 155,
	3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 25, 3, 2, 2, 2, 159, 157, 3, 2,
	2, 2, 160, 161, 5, 10, 6, 2, 161, 27, 3, 2, 2, 2, 162, 168, 7, 33, 2, 2,
	163, 168, 7, 34, 2, 2, 164, 168, 7, 31, 2, 2, 165, 168, 7, 30, 2, 2, 166,
	168, 7, 32, 2, 2, 167, 162, 3, 2, 2, 2, 167, 163, 3, 2, 2, 2, 167, 164,
	3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 166, 3, 2, 2, 2, 168, 29, 3, 2,
	2, 2, 169, 170, 7, 33, 2, 2, 170, 31, 3, 2, 2, 2, 171, 172, 7, 34, 2, 2,
	172, 33, 3, 2, 2, 2, 173, 177, 5, 44, 23, 2, 174, 177, 5, 38, 20, 2, 175,
	177, 5, 36, 19, 2, 176, 173, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 175,
	3, 2, 2, 2, 177, 182, 3, 2, 2, 2, 178, 179, 7, 9, 2, 2, 179, 181, 7, 10,
	2, 2, 180, 178, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2,
	182, 183, 3, 2, 2, 2, 183, 35, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 186,
	9, 6, 2, 2, 186, 37, 3, 2, 2, 2, 187, 188, 7, 24, 2, 2, 188, 189, 5, 42,
	22, 2, 189, 191, 7, 4, 2, 2, 190, 192, 5, 40, 21, 2, 191, 190, 3, 2, 2,
	2, 191, 192, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 7, 5, 2, 2, 194,
	39, 3, 2, 2, 2, 195, 200, 5, 34, 18, 2, 196, 197, 7, 8, 2, 2, 197, 199,
	5, 34, 18, 2, 198, 196, 3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3,
	2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 41, 3, 2, 2, 2, 202, 200, 3, 2, 2,
	2, 203, 206, 5, 34, 18, 2, 204, 206, 7, 25, 2, 2, 205, 203, 3, 2, 2, 2,
	205, 204, 3, 2, 2, 2, 206, 43, 3, 2, 2, 2, 207, 212, 7, 35, 2, 2, 208,
	209, 7, 11, 2, 2, 209, 211, 7, 35, 2, 2, 210, 208, 3, 2, 2, 2, 211, 214,
	3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 45, 3, 2,
	2, 2, 214, 212, 3, 2, 2, 2, 215, 220, 5, 10, 6, 2, 216, 217, 7, 8, 2, 2,
	217, 219, 5, 10, 6, 2, 218, 216, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220,
	218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 47, 3, 2, 2, 2, 222, 220, 3,
	2, 2, 2, 223, 224, 7, 4, 2, 2, 224, 225, 5, 10, 6, 2, 225, 226, 7, 5, 2,
	2, 226, 49, 3, 2, 2, 2, 26, 55, 62, 67, 79, 96, 98, 111, 115, 121, 124,
	128, 132, 136, 144, 150, 157, 167, 176, 182, 191, 200, 205, 212, 220,
}
var literalNames = []string{
	"", "';'", "'('", "')'", "'{'", "'}'", "','", "'['", "']'", "'.'", "'*'",
	"'/'", "'+'", "'-'", "'%'", "'='", "'>'", "'<'", "'=='", "'<='", "'>='",
	"'!='", "'func'", "'void'", "'for'", "'if'", "'else'", "'return'", "",
	"", "'null'", "", "", "", "", "'string'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "MULT", "DIV", "PLUS", "SUB", "MOD",
	"ASSIGN", "GT", "LT", "EQUAL", "LE", "GE", "NOTEQUAL", "FUNCTION", "VOID",
	"FOR", "IF", "ELSE", "RETURN", "BOOL_LITERAL", "STRING_LITERAL", "NULL_LITERAL",
	"DECIMAL_LITERAL", "FLOAT_LITERAL", "IDENTIFIER", "NUMBER", "STRING", "SPACES",
}

var ruleNames = []string{
	"prog", "blockStatements", "blockStatement", "parse", "expr", "block",
	"statement", "forControl", "forInit", "variableDeclarators", "variableDeclarator",
	"variableDeclaratorId", "variableInitializer", "literal", "integerLiteral",
	"floatLiteral", "typeType", "primitiveType", "functionType", "typeList",
	"typeTypeOrVoid", "classOrInterfaceType", "expressionList", "parExpression",
}

type GScriptParser struct {
	*antlr.BaseParser
}

// NewGScriptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GScriptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGScriptParser(input antlr.TokenStream) *GScriptParser {
	this := new(GScriptParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GScript.g4"

	return this
}

// GScriptParser tokens.
const (
	GScriptParserEOF             = antlr.TokenEOF
	GScriptParserT__0            = 1
	GScriptParserT__1            = 2
	GScriptParserT__2            = 3
	GScriptParserT__3            = 4
	GScriptParserT__4            = 5
	GScriptParserT__5            = 6
	GScriptParserT__6            = 7
	GScriptParserT__7            = 8
	GScriptParserT__8            = 9
	GScriptParserMULT            = 10
	GScriptParserDIV             = 11
	GScriptParserPLUS            = 12
	GScriptParserSUB             = 13
	GScriptParserMOD             = 14
	GScriptParserASSIGN          = 15
	GScriptParserGT              = 16
	GScriptParserLT              = 17
	GScriptParserEQUAL           = 18
	GScriptParserLE              = 19
	GScriptParserGE              = 20
	GScriptParserNOTEQUAL        = 21
	GScriptParserFUNCTION        = 22
	GScriptParserVOID            = 23
	GScriptParserFOR             = 24
	GScriptParserIF              = 25
	GScriptParserELSE            = 26
	GScriptParserRETURN          = 27
	GScriptParserBOOL_LITERAL    = 28
	GScriptParserSTRING_LITERAL  = 29
	GScriptParserNULL_LITERAL    = 30
	GScriptParserDECIMAL_LITERAL = 31
	GScriptParserFLOAT_LITERAL   = 32
	GScriptParserIDENTIFIER      = 33
	GScriptParserNUMBER          = 34
	GScriptParserSTRING          = 35
	GScriptParserSPACES          = 36
)

// GScriptParser rules.
const (
	GScriptParserRULE_prog                 = 0
	GScriptParserRULE_blockStatements      = 1
	GScriptParserRULE_blockStatement       = 2
	GScriptParserRULE_parse                = 3
	GScriptParserRULE_expr                 = 4
	GScriptParserRULE_block                = 5
	GScriptParserRULE_statement            = 6
	GScriptParserRULE_forControl           = 7
	GScriptParserRULE_forInit              = 8
	GScriptParserRULE_variableDeclarators  = 9
	GScriptParserRULE_variableDeclarator   = 10
	GScriptParserRULE_variableDeclaratorId = 11
	GScriptParserRULE_variableInitializer  = 12
	GScriptParserRULE_literal              = 13
	GScriptParserRULE_integerLiteral       = 14
	GScriptParserRULE_floatLiteral         = 15
	GScriptParserRULE_typeType             = 16
	GScriptParserRULE_primitiveType        = 17
	GScriptParserRULE_functionType         = 18
	GScriptParserRULE_typeList             = 19
	GScriptParserRULE_typeTypeOrVoid       = 20
	GScriptParserRULE_classOrInterfaceType = 21
	GScriptParserRULE_expressionList       = 22
	GScriptParserRULE_parExpression        = 23
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) BlockStatements() IBlockStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GScriptParserRULE_prog)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.BlockStatements()
	}

	return localctx
}

// IBlockStatementsContext is an interface to support dynamic dispatch.
type IBlockStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementsContext differentiates from other interfaces.
	IsBlockStatementsContext()
}

type BlockStatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementsContext() *BlockStatementsContext {
	var p = new(BlockStatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_blockStatements
	return p
}

func (*BlockStatementsContext) IsBlockStatementsContext() {}

func NewBlockStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementsContext {
	var p = new(BlockStatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_blockStatements

	return p
}

func (s *BlockStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementsContext) CopyFrom(ctx *BlockStatementsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BlockStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockStmsContext struct {
	*BlockStatementsContext
}

func NewBlockStmsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmsContext {
	var p = new(BlockStmsContext)

	p.BlockStatementsContext = NewEmptyBlockStatementsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementsContext))

	return p
}

func (s *BlockStmsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmsContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *BlockStmsContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockStmsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockStms(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) BlockStatements() (localctx IBlockStatementsContext) {
	localctx = NewBlockStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GScriptParserRULE_blockStatements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewBlockStmsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserT__1)|(1<<GScriptParserT__3)|(1<<GScriptParserSUB)|(1<<GScriptParserFUNCTION)|(1<<GScriptParserIF)|(1<<GScriptParserRETURN)|(1<<GScriptParserBOOL_LITERAL)|(1<<GScriptParserSTRING_LITERAL)|(1<<GScriptParserNULL_LITERAL)|(1<<GScriptParserDECIMAL_LITERAL))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GScriptParserFLOAT_LITERAL-32))|(1<<(GScriptParserIDENTIFIER-32))|(1<<(GScriptParserNUMBER-32))|(1<<(GScriptParserSTRING-32)))) != 0) {
		{
			p.SetState(50)
			p.BlockStatement()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) CopyFrom(ctx *BlockStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockStmContext struct {
	*BlockStatementContext
}

func NewBlockStmContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmContext {
	var p = new(BlockStmContext)

	p.BlockStatementContext = NewEmptyBlockStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementContext))

	return p
}

func (s *BlockStmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockStm(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockVarDeclarContext struct {
	*BlockStatementContext
}

func NewBlockVarDeclarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockVarDeclarContext {
	var p = new(BlockVarDeclarContext)

	p.BlockStatementContext = NewEmptyBlockStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementContext))

	return p
}

func (s *BlockVarDeclarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockVarDeclarContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *BlockVarDeclarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockVarDeclar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GScriptParserRULE_blockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserFUNCTION, GScriptParserIDENTIFIER, GScriptParserNUMBER, GScriptParserSTRING:
		localctx = NewBlockVarDeclarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.VariableDeclarators()
		}
		{
			p.SetState(57)
			p.Match(GScriptParserT__0)
		}

	case GScriptParserT__1, GScriptParserT__3, GScriptParserSUB, GScriptParserIF, GScriptParserRETURN, GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL:
		localctx = NewBlockStmContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetExpr_list returns the expr_list rule context list.
	GetExpr_list() []IExprContext

	// SetExpr_list sets the expr_list rule context list.
	SetExpr_list([]IExprContext)

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_expr     IExprContext
	expr_list []IExprContext
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Get_expr() IExprContext { return s._expr }

func (s *ParseContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ParseContext) GetExpr_list() []IExprContext { return s.expr_list }

func (s *ParseContext) SetExpr_list(v []IExprContext) { s.expr_list = v }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(GScriptParserEOF, 0)
}

func (s *ParseContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ParseContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GScriptParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(GScriptParserT__1-2))|(1<<(GScriptParserSUB-2))|(1<<(GScriptParserBOOL_LITERAL-2))|(1<<(GScriptParserSTRING_LITERAL-2))|(1<<(GScriptParserNULL_LITERAL-2))|(1<<(GScriptParserDECIMAL_LITERAL-2))|(1<<(GScriptParserFLOAT_LITERAL-2)))) != 0) {
		{
			p.SetState(62)

			var _x = p.expr(0)

			localctx.(*ParseContext)._expr = _x
		}
		localctx.(*ParseContext).expr_list = append(localctx.(*ParseContext).expr_list, localctx.(*ParseContext)._expr)

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(GScriptParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultDivExprContext struct {
	*ExprContext
	lhs IExprContext
	bop antlr.Token
	rhs IExprContext
}

func NewMultDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultDivExprContext {
	var p = new(MultDivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultDivExprContext) GetBop() antlr.Token { return s.bop }

func (s *MultDivExprContext) SetBop(v antlr.Token) { s.bop = v }

func (s *MultDivExprContext) GetLhs() IExprContext { return s.lhs }

func (s *MultDivExprContext) GetRhs() IExprContext { return s.rhs }

func (s *MultDivExprContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *MultDivExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *MultDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultDivExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultDivExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultDivExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(GScriptParserMULT, 0)
}

func (s *MultDivExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(GScriptParserDIV, 0)
}

func (s *MultDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitMultDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GLeContext struct {
	*ExprContext
	bop antlr.Token
}

func NewGLeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GLeContext {
	var p = new(GLeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GLeContext) GetBop() antlr.Token { return s.bop }

func (s *GLeContext) SetBop(v antlr.Token) { s.bop = v }

func (s *GLeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GLeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GLeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GLeContext) LE() antlr.TerminalNode {
	return s.GetToken(GScriptParserLE, 0)
}

func (s *GLeContext) GE() antlr.TerminalNode {
	return s.GetToken(GScriptParserGE, 0)
}

func (s *GLeContext) GT() antlr.TerminalNode {
	return s.GetToken(GScriptParserGT, 0)
}

func (s *GLeContext) LT() antlr.TerminalNode {
	return s.GetToken(GScriptParserLT, 0)
}

func (s *GLeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitGLe(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiterContext struct {
	*ExprContext
	liter ILiteralContext
}

func NewLiterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiterContext {
	var p = new(LiterContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiterContext) GetLiter() ILiteralContext { return s.liter }

func (s *LiterContext) SetLiter(v ILiteralContext) { s.liter = v }

func (s *LiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiterContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitLiter(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusSubExprContext struct {
	*ExprContext
	lhs IExprContext
	bop antlr.Token
	rhs IExprContext
}

func NewPlusSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusSubExprContext {
	var p = new(PlusSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PlusSubExprContext) GetBop() antlr.Token { return s.bop }

func (s *PlusSubExprContext) SetBop(v antlr.Token) { s.bop = v }

func (s *PlusSubExprContext) GetLhs() IExprContext { return s.lhs }

func (s *PlusSubExprContext) GetRhs() IExprContext { return s.rhs }

func (s *PlusSubExprContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *PlusSubExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *PlusSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusSubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PlusSubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PlusSubExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GScriptParserPLUS, 0)
}

func (s *PlusSubExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(GScriptParserSUB, 0)
}

func (s *PlusSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitPlusSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedExprContext struct {
	*ExprContext
}

func NewNestedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExprContext {
	var p = new(NestedExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NestedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NestedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitNestedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModExprContext struct {
	*ExprContext
	lhs IExprContext
	bop antlr.Token
	rhs IExprContext
}

func NewModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModExprContext {
	var p = new(ModExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ModExprContext) GetBop() antlr.Token { return s.bop }

func (s *ModExprContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ModExprContext) GetLhs() IExprContext { return s.lhs }

func (s *ModExprContext) GetRhs() IExprContext { return s.rhs }

func (s *ModExprContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ModExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ModExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ModExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(GScriptParserMOD, 0)
}

func (s *ModExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitModExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualOrNotContext struct {
	*ExprContext
	bop antlr.Token
}

func NewEqualOrNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualOrNotContext {
	var p = new(EqualOrNotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualOrNotContext) GetBop() antlr.Token { return s.bop }

func (s *EqualOrNotContext) SetBop(v antlr.Token) { s.bop = v }

func (s *EqualOrNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualOrNotContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualOrNotContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualOrNotContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserEQUAL, 0)
}

func (s *EqualOrNotContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserNOTEQUAL, 0)
}

func (s *EqualOrNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitEqualOrNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	*ExprContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(GScriptParserSUB, 0)
}

func (s *UnaryExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GScriptParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, GScriptParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserT__1:
		localctx = NewNestedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(70)
			p.Match(GScriptParserT__1)
		}
		{
			p.SetState(71)
			p.expr(0)
		}
		{
			p.SetState(72)
			p.Match(GScriptParserT__2)
		}

	case GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL:
		localctx = NewLiterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)

			var _x = p.Literal()

			localctx.(*LiterContext).liter = _x
		}

	case GScriptParserSUB:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(75)
			p.Match(GScriptParserSUB)
		}
		{
			p.SetState(76)
			p.expr(6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultDivExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(80)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultDivExprContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GScriptParserMULT || _la == GScriptParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultDivExprContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(81)

					var _x = p.expr(6)

					localctx.(*MultDivExprContext).rhs = _x
				}

			case 2:
				localctx = NewModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ModExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(83)

					var _m = p.Match(GScriptParserMOD)

					localctx.(*ModExprContext).bop = _m
				}
				{
					p.SetState(84)

					var _x = p.expr(5)

					localctx.(*ModExprContext).rhs = _x
				}

			case 3:
				localctx = NewPlusSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*PlusSubExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(85)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(86)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*PlusSubExprContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GScriptParserPLUS || _la == GScriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*PlusSubExprContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(87)

					var _x = p.expr(4)

					localctx.(*PlusSubExprContext).rhs = _x
				}

			case 4:
				localctx = NewGLeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(89)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*GLeContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserGT)|(1<<GScriptParserLT)|(1<<GScriptParserLE)|(1<<GScriptParserGE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*GLeContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(90)
					p.expr(3)
				}

			case 5:
				localctx = NewEqualOrNotContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(92)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualOrNotContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GScriptParserEQUAL || _la == GScriptParserNOTEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualOrNotContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(93)
					p.expr(2)
				}

			}

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) BlockStatements() IBlockStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GScriptParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(GScriptParserT__3)
	}
	{
		p.SetState(100)
		p.BlockStatements()
	}
	{
		p.SetState(101)
		p.Match(GScriptParserT__4)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnContext struct {
	*StatementContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GScriptParserRETURN, 0)
}

func (s *ReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseContext struct {
	*StatementContext
}

func NewIfElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseContext {
	var p = new(IfElseContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseContext) IF() antlr.TerminalNode {
	return s.GetToken(GScriptParserIF, 0)
}

func (s *IfElseContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *IfElseContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfElseContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GScriptParserELSE, 0)
}

func (s *IfElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitIfElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmExprContext struct {
	*StatementContext
	statementExpression IExprContext
}

func NewStmExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmExprContext {
	var p = new(StmExprContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *StmExprContext) GetStatementExpression() IExprContext { return s.statementExpression }

func (s *StmExprContext) SetStatementExpression(v IExprContext) { s.statementExpression = v }

func (s *StmExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitStmExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockLabelContext struct {
	*StatementContext
	blockLabel IBlockContext
}

func NewBlockLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockLabelContext {
	var p = new(BlockLabelContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BlockLabelContext) GetBlockLabel() IBlockContext { return s.blockLabel }

func (s *BlockLabelContext) SetBlockLabel(v IBlockContext) { s.blockLabel = v }

func (s *BlockLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLabelContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GScriptParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserT__3:
		localctx = NewBlockLabelContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)

			var _x = p.Block()

			localctx.(*BlockLabelContext).blockLabel = _x
		}

	case GScriptParserIF:
		localctx = NewIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(GScriptParserIF)
		}
		{
			p.SetState(105)
			p.ParExpression()
		}
		{
			p.SetState(106)
			p.Statement()
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(107)
				p.Match(GScriptParserELSE)
			}
			{
				p.SetState(108)
				p.Statement()
			}

		}

	case GScriptParserRETURN:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.Match(GScriptParserRETURN)
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(GScriptParserT__1-2))|(1<<(GScriptParserSUB-2))|(1<<(GScriptParserBOOL_LITERAL-2))|(1<<(GScriptParserSTRING_LITERAL-2))|(1<<(GScriptParserNULL_LITERAL-2))|(1<<(GScriptParserDECIMAL_LITERAL-2))|(1<<(GScriptParserFLOAT_LITERAL-2)))) != 0 {
			{
				p.SetState(112)
				p.expr(0)
			}

		}
		{
			p.SetState(115)
			p.Match(GScriptParserT__0)
		}

	case GScriptParserT__1, GScriptParserSUB, GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL:
		localctx = NewStmExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)

			var _x = p.expr(0)

			localctx.(*StmExprContext).statementExpression = _x
		}
		{
			p.SetState(117)
			p.Match(GScriptParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetForUpdate returns the forUpdate rule contexts.
	GetForUpdate() IExpressionListContext

	// SetForUpdate sets the forUpdate rule contexts.
	SetForUpdate(IExpressionListContext)

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	forUpdate IExpressionListContext
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) GetForUpdate() IExpressionListContext { return s.forUpdate }

func (s *ForControlContext) SetForUpdate(v IExpressionListContext) { s.forUpdate = v }

func (s *ForControlContext) ForInit() IForInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForControlContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForControlContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GScriptParserRULE_forControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserT__1)|(1<<GScriptParserSUB)|(1<<GScriptParserFUNCTION)|(1<<GScriptParserBOOL_LITERAL)|(1<<GScriptParserSTRING_LITERAL)|(1<<GScriptParserNULL_LITERAL)|(1<<GScriptParserDECIMAL_LITERAL))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GScriptParserFLOAT_LITERAL-32))|(1<<(GScriptParserIDENTIFIER-32))|(1<<(GScriptParserNUMBER-32))|(1<<(GScriptParserSTRING-32)))) != 0) {
		{
			p.SetState(121)
			p.ForInit()
		}

	}
	{
		p.SetState(124)
		p.Match(GScriptParserT__0)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(GScriptParserT__1-2))|(1<<(GScriptParserSUB-2))|(1<<(GScriptParserBOOL_LITERAL-2))|(1<<(GScriptParserSTRING_LITERAL-2))|(1<<(GScriptParserNULL_LITERAL-2))|(1<<(GScriptParserDECIMAL_LITERAL-2))|(1<<(GScriptParserFLOAT_LITERAL-2)))) != 0 {
		{
			p.SetState(125)
			p.expr(0)
		}

	}
	{
		p.SetState(128)
		p.Match(GScriptParserT__0)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(GScriptParserT__1-2))|(1<<(GScriptParserSUB-2))|(1<<(GScriptParserBOOL_LITERAL-2))|(1<<(GScriptParserSTRING_LITERAL-2))|(1<<(GScriptParserNULL_LITERAL-2))|(1<<(GScriptParserDECIMAL_LITERAL-2))|(1<<(GScriptParserFLOAT_LITERAL-2)))) != 0 {
		{
			p.SetState(129)

			var _x = p.ExpressionList()

			localctx.(*ForControlContext).forUpdate = _x
		}

	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GScriptParserRULE_forInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserFUNCTION, GScriptParserIDENTIFIER, GScriptParserNUMBER, GScriptParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.VariableDeclarators()
		}

	case GScriptParserT__1, GScriptParserSUB, GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.ExpressionList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableDeclarators
	return p
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem())
	var tst = make([]IVariableDeclaratorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclaratorContext)
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GScriptParserRULE_variableDeclarators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.TypeType()
	}
	{
		p.SetState(137)
		p.VariableDeclarator()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__5 {
		{
			p.SetState(138)
			p.Match(GScriptParserT__5)
		}
		{
			p.SetState(139)
			p.VariableDeclarator()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableDeclarator
	return p
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *VariableDeclaratorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GScriptParserASSIGN, 0)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GScriptParserRULE_variableDeclarator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.VariableDeclaratorId()
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GScriptParserASSIGN {
		{
			p.SetState(146)
			p.Match(GScriptParserASSIGN)
		}
		{
			p.SetState(147)
			p.VariableInitializer()
		}

	}

	return localctx
}

// IVariableDeclaratorIdContext is an interface to support dynamic dispatch.
type IVariableDeclaratorIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorIdContext differentiates from other interfaces.
	IsVariableDeclaratorIdContext()
}

type VariableDeclaratorIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorIdContext() *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableDeclaratorId
	return p
}

func (*VariableDeclaratorIdContext) IsVariableDeclaratorIdContext() {}

func NewVariableDeclaratorIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableDeclaratorId

	return p
}

func (s *VariableDeclaratorIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorIdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GScriptParserIDENTIFIER, 0)
}

func (s *VariableDeclaratorIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableDeclaratorId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableDeclaratorId() (localctx IVariableDeclaratorIdContext) {
	localctx = NewVariableDeclaratorIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GScriptParserRULE_variableDeclaratorId)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(GScriptParserIDENTIFIER)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__6 {
		{
			p.SetState(151)
			p.Match(GScriptParserT__6)
		}
		{
			p.SetState(152)
			p.Match(GScriptParserT__7)
		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableInitializer
	return p
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GScriptParserRULE_variableInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.expr(0)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FloatContext struct {
	*LiteralContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserFLOAT_LITERAL, 0)
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	*LiteralContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserNULL_LITERAL, 0)
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	*LiteralContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserBOOL_LITERAL, 0)
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*LiteralContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserSTRING_LITERAL, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	*LiteralContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserDECIMAL_LITERAL, 0)
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GScriptParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserDECIMAL_LITERAL:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Match(GScriptParserDECIMAL_LITERAL)
		}

	case GScriptParserFLOAT_LITERAL:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Match(GScriptParserFLOAT_LITERAL)
		}

	case GScriptParserSTRING_LITERAL:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.Match(GScriptParserSTRING_LITERAL)
		}

	case GScriptParserBOOL_LITERAL:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(163)
			p.Match(GScriptParserBOOL_LITERAL)
		}

	case GScriptParserNULL_LITERAL:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(164)
			p.Match(GScriptParserNULL_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserDECIMAL_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GScriptParserRULE_integerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(GScriptParserDECIMAL_LITERAL)
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GScriptParserRULE_floatLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(GScriptParserFLOAT_LITERAL)
	}

	return localctx
}

// ITypeTypeContext is an interface to support dynamic dispatch.
type ITypeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeContext differentiates from other interfaces.
	IsTypeTypeContext()
}

type TypeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeContext() *TypeTypeContext {
	var p = new(TypeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_typeType
	return p
}

func (*TypeTypeContext) IsTypeTypeContext() {}

func NewTypeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeContext {
	var p = new(TypeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_typeType

	return p
}

func (s *TypeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeContext) ClassOrInterfaceType() IClassOrInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOrInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassOrInterfaceTypeContext)
}

func (s *TypeTypeContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *TypeTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitTypeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) TypeType() (localctx ITypeTypeContext) {
	localctx = NewTypeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GScriptParserRULE_typeType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserIDENTIFIER:
		{
			p.SetState(171)
			p.ClassOrInterfaceType()
		}

	case GScriptParserFUNCTION:
		{
			p.SetState(172)
			p.FunctionType()
		}

	case GScriptParserNUMBER, GScriptParserSTRING:
		{
			p.SetState(173)
			p.PrimitiveType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__6 {
		{
			p.SetState(176)
			p.Match(GScriptParserT__6)
		}
		{
			p.SetState(177)
			p.Match(GScriptParserT__7)
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GScriptParserNUMBER, 0)
}

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(GScriptParserSTRING, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GScriptParserRULE_primitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GScriptParserNUMBER || _la == GScriptParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(GScriptParserFUNCTION, 0)
}

func (s *FunctionTypeContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeOrVoidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionTypeContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GScriptParserRULE_functionType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(GScriptParserFUNCTION)
	}
	{
		p.SetState(186)
		p.TypeTypeOrVoid()
	}
	{
		p.SetState(187)
		p.Match(GScriptParserT__1)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(GScriptParserFUNCTION-22))|(1<<(GScriptParserIDENTIFIER-22))|(1<<(GScriptParserNUMBER-22))|(1<<(GScriptParserSTRING-22)))) != 0 {
		{
			p.SetState(188)
			p.TypeList()
		}

	}
	{
		p.SetState(191)
		p.Match(GScriptParserT__2)
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) AllTypeType() []ITypeTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem())
	var tst = make([]ITypeTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeTypeContext)
		}
	}

	return tst
}

func (s *TypeListContext) TypeType(i int) ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GScriptParserRULE_typeList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.TypeType()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__5 {
		{
			p.SetState(194)
			p.Match(GScriptParserT__5)
		}
		{
			p.SetState(195)
			p.TypeType()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeTypeOrVoidContext is an interface to support dynamic dispatch.
type ITypeTypeOrVoidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeOrVoidContext differentiates from other interfaces.
	IsTypeTypeOrVoidContext()
}

type TypeTypeOrVoidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeOrVoidContext() *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_typeTypeOrVoid
	return p
}

func (*TypeTypeOrVoidContext) IsTypeTypeOrVoidContext() {}

func NewTypeTypeOrVoidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_typeTypeOrVoid

	return p
}

func (s *TypeTypeOrVoidContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeOrVoidContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeTypeOrVoidContext) VOID() antlr.TerminalNode {
	return s.GetToken(GScriptParserVOID, 0)
}

func (s *TypeTypeOrVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeOrVoidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeOrVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitTypeTypeOrVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) TypeTypeOrVoid() (localctx ITypeTypeOrVoidContext) {
	localctx = NewTypeTypeOrVoidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GScriptParserRULE_typeTypeOrVoid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserFUNCTION, GScriptParserIDENTIFIER, GScriptParserNUMBER, GScriptParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.TypeType()
		}

	case GScriptParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(GScriptParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClassOrInterfaceTypeContext is an interface to support dynamic dispatch.
type IClassOrInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOrInterfaceTypeContext differentiates from other interfaces.
	IsClassOrInterfaceTypeContext()
}

type ClassOrInterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOrInterfaceTypeContext() *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_classOrInterfaceType
	return p
}

func (*ClassOrInterfaceTypeContext) IsClassOrInterfaceTypeContext() {}

func NewClassOrInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_classOrInterfaceType

	return p
}

func (s *ClassOrInterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOrInterfaceTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserIDENTIFIER)
}

func (s *ClassOrInterfaceTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserIDENTIFIER, i)
}

func (s *ClassOrInterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOrInterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOrInterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitClassOrInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ClassOrInterfaceType() (localctx IClassOrInterfaceTypeContext) {
	localctx = NewClassOrInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GScriptParserRULE_classOrInterfaceType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(GScriptParserIDENTIFIER)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__8 {
		{
			p.SetState(206)
			p.Match(GScriptParserT__8)
		}
		{
			p.SetState(207)
			p.Match(GScriptParserIDENTIFIER)
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GScriptParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.expr(0)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__5 {
		{
			p.SetState(214)
			p.Match(GScriptParserT__5)
		}
		{
			p.SetState(215)
			p.expr(0)
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitParExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ParExpression() (localctx IParExpressionContext) {
	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GScriptParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(GScriptParserT__1)
	}
	{
		p.SetState(222)
		p.expr(0)
	}
	{
		p.SetState(223)
		p.Match(GScriptParserT__2)
	}

	return localctx
}

func (p *GScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GScriptParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
