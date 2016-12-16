// Generated from ArgParse.g4 by ANTLR 4.6.

package parser // ArgParse

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
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 8, 47, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23, 11, 2, 3,
	3, 3, 3, 3, 3, 5, 3, 28, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2,
	2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 3, 3, 2, 7, 8, 42, 2, 16, 3, 2, 2, 2,
	4, 27, 3, 2, 2, 2, 6, 29, 3, 2, 2, 2, 8, 32, 3, 2, 2, 2, 10, 37, 3, 2,
	2, 2, 12, 42, 3, 2, 2, 2, 14, 44, 3, 2, 2, 2, 16, 21, 5, 4, 3, 2, 17, 18,
	7, 5, 2, 2, 18, 20, 5, 4, 3, 2, 19, 17, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2,
	21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 3, 3, 2, 2, 2, 23, 21, 3, 2,
	2, 2, 24, 28, 5, 10, 6, 2, 25, 28, 5, 8, 5, 2, 26, 28, 5, 6, 4, 2, 27,
	24, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 5, 3, 2, 2,
	2, 29, 30, 7, 4, 2, 2, 30, 31, 5, 12, 7, 2, 31, 7, 3, 2, 2, 2, 32, 33,
	7, 4, 2, 2, 33, 34, 5, 12, 7, 2, 34, 35, 7, 5, 2, 2, 35, 36, 5, 14, 8,
	2, 36, 9, 3, 2, 2, 2, 37, 38, 7, 4, 2, 2, 38, 39, 5, 12, 7, 2, 39, 40,
	7, 3, 2, 2, 40, 41, 5, 14, 8, 2, 41, 11, 3, 2, 2, 2, 42, 43, 7, 7, 2, 2,
	43, 13, 3, 2, 2, 2, 44, 45, 9, 2, 2, 2, 45, 15, 3, 2, 2, 2, 4, 21, 27,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='",
}

var symbolicNames = []string{
	"", "", "DASH", "SPACE", "NEWLINE", "TEXT", "STRING",
}

var ruleNames = []string{
	"args", "arg", "dash_opt_k_bool", "dash_opt_kv", "dash_opt_kv_eq", "key",
	"val",
}

type ArgParseParser struct {
	*antlr.BaseParser
}

func NewArgParseParser(input antlr.TokenStream) *ArgParseParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(ArgParseParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ArgParse.g4"

	return this
}

// ArgParseParser tokens.
const (
	ArgParseParserEOF     = antlr.TokenEOF
	ArgParseParserT__0    = 1
	ArgParseParserDASH    = 2
	ArgParseParserSPACE   = 3
	ArgParseParserNEWLINE = 4
	ArgParseParserTEXT    = 5
	ArgParseParserSTRING  = 6
)

// ArgParseParser rules.
const (
	ArgParseParserRULE_args            = 0
	ArgParseParserRULE_arg             = 1
	ArgParseParserRULE_dash_opt_k_bool = 2
	ArgParseParserRULE_dash_opt_kv     = 3
	ArgParseParserRULE_dash_opt_kv_eq  = 4
	ArgParseParserRULE_key             = 5
	ArgParseParserRULE_val             = 6
)

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgParseParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgParseParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *ArgsContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *ArgsContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(ArgParseParserSPACE)
}

func (s *ArgsContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(ArgParseParserSPACE, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *ArgParseParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ArgParseParserRULE_args)
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
		p.SetState(14)
		p.Arg()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ArgParseParserSPACE {
		{
			p.SetState(15)
			p.Match(ArgParseParserSPACE)
		}
		{
			p.SetState(16)
			p.Arg()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgParseParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgParseParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Dash_opt_kv_eq() IDash_opt_kv_eqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDash_opt_kv_eqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDash_opt_kv_eqContext)
}

func (s *ArgContext) Dash_opt_kv() IDash_opt_kvContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDash_opt_kvContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDash_opt_kvContext)
}

func (s *ArgContext) Dash_opt_k_bool() IDash_opt_k_boolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDash_opt_k_boolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDash_opt_k_boolContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *ArgParseParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ArgParseParserRULE_arg)

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

	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Dash_opt_kv_eq()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Dash_opt_kv()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Dash_opt_k_bool()
		}

	}

	return localctx
}

// IDash_opt_k_boolContext is an interface to support dynamic dispatch.
type IDash_opt_k_boolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDash_opt_k_boolContext differentiates from other interfaces.
	IsDash_opt_k_boolContext()
}

type Dash_opt_k_boolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDash_opt_k_boolContext() *Dash_opt_k_boolContext {
	var p = new(Dash_opt_k_boolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgParseParserRULE_dash_opt_k_bool
	return p
}

func (*Dash_opt_k_boolContext) IsDash_opt_k_boolContext() {}

func NewDash_opt_k_boolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dash_opt_k_boolContext {
	var p = new(Dash_opt_k_boolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgParseParserRULE_dash_opt_k_bool

	return p
}

func (s *Dash_opt_k_boolContext) GetParser() antlr.Parser { return s.parser }

func (s *Dash_opt_k_boolContext) DASH() antlr.TerminalNode {
	return s.GetToken(ArgParseParserDASH, 0)
}

func (s *Dash_opt_k_boolContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Dash_opt_k_boolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dash_opt_k_boolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dash_opt_k_boolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.EnterDash_opt_k_bool(s)
	}
}

func (s *Dash_opt_k_boolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.ExitDash_opt_k_bool(s)
	}
}

func (p *ArgParseParser) Dash_opt_k_bool() (localctx IDash_opt_k_boolContext) {
	localctx = NewDash_opt_k_boolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ArgParseParserRULE_dash_opt_k_bool)

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
		p.SetState(27)
		p.Match(ArgParseParserDASH)
	}
	{
		p.SetState(28)
		p.Key()
	}

	return localctx
}

// IDash_opt_kvContext is an interface to support dynamic dispatch.
type IDash_opt_kvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDash_opt_kvContext differentiates from other interfaces.
	IsDash_opt_kvContext()
}

type Dash_opt_kvContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDash_opt_kvContext() *Dash_opt_kvContext {
	var p = new(Dash_opt_kvContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgParseParserRULE_dash_opt_kv
	return p
}

func (*Dash_opt_kvContext) IsDash_opt_kvContext() {}

func NewDash_opt_kvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dash_opt_kvContext {
	var p = new(Dash_opt_kvContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgParseParserRULE_dash_opt_kv

	return p
}

func (s *Dash_opt_kvContext) GetParser() antlr.Parser { return s.parser }

func (s *Dash_opt_kvContext) DASH() antlr.TerminalNode {
	return s.GetToken(ArgParseParserDASH, 0)
}

func (s *Dash_opt_kvContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Dash_opt_kvContext) SPACE() antlr.TerminalNode {
	return s.GetToken(ArgParseParserSPACE, 0)
}

func (s *Dash_opt_kvContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *Dash_opt_kvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dash_opt_kvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dash_opt_kvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.EnterDash_opt_kv(s)
	}
}

func (s *Dash_opt_kvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.ExitDash_opt_kv(s)
	}
}

func (p *ArgParseParser) Dash_opt_kv() (localctx IDash_opt_kvContext) {
	localctx = NewDash_opt_kvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ArgParseParserRULE_dash_opt_kv)

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
		p.SetState(30)
		p.Match(ArgParseParserDASH)
	}
	{
		p.SetState(31)
		p.Key()
	}
	{
		p.SetState(32)
		p.Match(ArgParseParserSPACE)
	}
	{
		p.SetState(33)
		p.Val()
	}

	return localctx
}

// IDash_opt_kv_eqContext is an interface to support dynamic dispatch.
type IDash_opt_kv_eqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDash_opt_kv_eqContext differentiates from other interfaces.
	IsDash_opt_kv_eqContext()
}

type Dash_opt_kv_eqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDash_opt_kv_eqContext() *Dash_opt_kv_eqContext {
	var p = new(Dash_opt_kv_eqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgParseParserRULE_dash_opt_kv_eq
	return p
}

func (*Dash_opt_kv_eqContext) IsDash_opt_kv_eqContext() {}

func NewDash_opt_kv_eqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dash_opt_kv_eqContext {
	var p = new(Dash_opt_kv_eqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgParseParserRULE_dash_opt_kv_eq

	return p
}

func (s *Dash_opt_kv_eqContext) GetParser() antlr.Parser { return s.parser }

func (s *Dash_opt_kv_eqContext) DASH() antlr.TerminalNode {
	return s.GetToken(ArgParseParserDASH, 0)
}

func (s *Dash_opt_kv_eqContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Dash_opt_kv_eqContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *Dash_opt_kv_eqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dash_opt_kv_eqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dash_opt_kv_eqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.EnterDash_opt_kv_eq(s)
	}
}

func (s *Dash_opt_kv_eqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.ExitDash_opt_kv_eq(s)
	}
}

func (p *ArgParseParser) Dash_opt_kv_eq() (localctx IDash_opt_kv_eqContext) {
	localctx = NewDash_opt_kv_eqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ArgParseParserRULE_dash_opt_kv_eq)

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
		p.SetState(35)
		p.Match(ArgParseParserDASH)
	}
	{
		p.SetState(36)
		p.Key()
	}
	{
		p.SetState(37)
		p.Match(ArgParseParserT__0)
	}
	{
		p.SetState(38)
		p.Val()
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgParseParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgParseParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) TEXT() antlr.TerminalNode {
	return s.GetToken(ArgParseParserTEXT, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *ArgParseParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ArgParseParserRULE_key)

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
		p.SetState(40)
		p.Match(ArgParseParserTEXT)
	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgParseParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgParseParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) TEXT() antlr.TerminalNode {
	return s.GetToken(ArgParseParserTEXT, 0)
}

func (s *ValContext) STRING() antlr.TerminalNode {
	return s.GetToken(ArgParseParserSTRING, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgParseListener); ok {
		listenerT.ExitVal(s)
	}
}

func (p *ArgParseParser) Val() (localctx IValContext) {
	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ArgParseParserRULE_val)
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
	p.SetState(42)
	_la = p.GetTokenStream().LA(1)

	if !(_la == ArgParseParserTEXT || _la == ArgParseParserSTRING) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
