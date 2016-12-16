// Generated from ArgParse.g4 by ANTLR 4.6.

package parser // ArgParse

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ArgParseListener is a complete listener for a parse tree produced by ArgParseParser.
type ArgParseListener interface {
	antlr.ParseTreeListener

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterDash_opt_k_bool is called when entering the dash_opt_k_bool production.
	EnterDash_opt_k_bool(c *Dash_opt_k_boolContext)

	// EnterDash_opt_kv is called when entering the dash_opt_kv production.
	EnterDash_opt_kv(c *Dash_opt_kvContext)

	// EnterDash_opt_kv_eq is called when entering the dash_opt_kv_eq production.
	EnterDash_opt_kv_eq(c *Dash_opt_kv_eqContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterVal is called when entering the val production.
	EnterVal(c *ValContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitDash_opt_k_bool is called when exiting the dash_opt_k_bool production.
	ExitDash_opt_k_bool(c *Dash_opt_k_boolContext)

	// ExitDash_opt_kv is called when exiting the dash_opt_kv production.
	ExitDash_opt_kv(c *Dash_opt_kvContext)

	// ExitDash_opt_kv_eq is called when exiting the dash_opt_kv_eq production.
	ExitDash_opt_kv_eq(c *Dash_opt_kv_eqContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitVal is called when exiting the val production.
	ExitVal(c *ValContext)
}
