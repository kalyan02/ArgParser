// Generated from ArgParse.g4 by ANTLR 4.6.

package parser // ArgParse

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseArgParseListener is a complete listener for a parse tree produced by ArgParseParser.
type BaseArgParseListener struct{}

var _ ArgParseListener = &BaseArgParseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArgParseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArgParseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArgParseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArgParseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseArgParseListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseArgParseListener) ExitArgs(ctx *ArgsContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseArgParseListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseArgParseListener) ExitArg(ctx *ArgContext) {}

// EnterDash_opt_k_bool is called when production dash_opt_k_bool is entered.
func (s *BaseArgParseListener) EnterDash_opt_k_bool(ctx *Dash_opt_k_boolContext) {}

// ExitDash_opt_k_bool is called when production dash_opt_k_bool is exited.
func (s *BaseArgParseListener) ExitDash_opt_k_bool(ctx *Dash_opt_k_boolContext) {}

// EnterDash_opt_kv is called when production dash_opt_kv is entered.
func (s *BaseArgParseListener) EnterDash_opt_kv(ctx *Dash_opt_kvContext) {}

// ExitDash_opt_kv is called when production dash_opt_kv is exited.
func (s *BaseArgParseListener) ExitDash_opt_kv(ctx *Dash_opt_kvContext) {}

// EnterDash_opt_kv_eq is called when production dash_opt_kv_eq is entered.
func (s *BaseArgParseListener) EnterDash_opt_kv_eq(ctx *Dash_opt_kv_eqContext) {}

// ExitDash_opt_kv_eq is called when production dash_opt_kv_eq is exited.
func (s *BaseArgParseListener) ExitDash_opt_kv_eq(ctx *Dash_opt_kv_eqContext) {}

// EnterKey is called when production key is entered.
func (s *BaseArgParseListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseArgParseListener) ExitKey(ctx *KeyContext) {}

// EnterVal is called when production val is entered.
func (s *BaseArgParseListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BaseArgParseListener) ExitVal(ctx *ValContext) {}
