// Generated from ArgParse.g4 by ANTLR 4.6
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ArgParseParser}.
 */
public interface ArgParseListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ArgParseParser#args}.
	 * @param ctx the parse tree
	 */
	void enterArgs(ArgParseParser.ArgsContext ctx);
	/**
	 * Exit a parse tree produced by {@link ArgParseParser#args}.
	 * @param ctx the parse tree
	 */
	void exitArgs(ArgParseParser.ArgsContext ctx);
	/**
	 * Enter a parse tree produced by {@link ArgParseParser#arg}.
	 * @param ctx the parse tree
	 */
	void enterArg(ArgParseParser.ArgContext ctx);
	/**
	 * Exit a parse tree produced by {@link ArgParseParser#arg}.
	 * @param ctx the parse tree
	 */
	void exitArg(ArgParseParser.ArgContext ctx);
	/**
	 * Enter a parse tree produced by {@link ArgParseParser#dash_opt_k_bool}.
	 * @param ctx the parse tree
	 */
	void enterDash_opt_k_bool(ArgParseParser.Dash_opt_k_boolContext ctx);
	/**
	 * Exit a parse tree produced by {@link ArgParseParser#dash_opt_k_bool}.
	 * @param ctx the parse tree
	 */
	void exitDash_opt_k_bool(ArgParseParser.Dash_opt_k_boolContext ctx);
	/**
	 * Enter a parse tree produced by {@link ArgParseParser#dash_opt_kv}.
	 * @param ctx the parse tree
	 */
	void enterDash_opt_kv(ArgParseParser.Dash_opt_kvContext ctx);
	/**
	 * Exit a parse tree produced by {@link ArgParseParser#dash_opt_kv}.
	 * @param ctx the parse tree
	 */
	void exitDash_opt_kv(ArgParseParser.Dash_opt_kvContext ctx);
	/**
	 * Enter a parse tree produced by {@link ArgParseParser#dash_opt_kv_eq}.
	 * @param ctx the parse tree
	 */
	void enterDash_opt_kv_eq(ArgParseParser.Dash_opt_kv_eqContext ctx);
	/**
	 * Exit a parse tree produced by {@link ArgParseParser#dash_opt_kv_eq}.
	 * @param ctx the parse tree
	 */
	void exitDash_opt_kv_eq(ArgParseParser.Dash_opt_kv_eqContext ctx);
	/**
	 * Enter a parse tree produced by {@link ArgParseParser#key}.
	 * @param ctx the parse tree
	 */
	void enterKey(ArgParseParser.KeyContext ctx);
	/**
	 * Exit a parse tree produced by {@link ArgParseParser#key}.
	 * @param ctx the parse tree
	 */
	void exitKey(ArgParseParser.KeyContext ctx);
	/**
	 * Enter a parse tree produced by {@link ArgParseParser#val}.
	 * @param ctx the parse tree
	 */
	void enterVal(ArgParseParser.ValContext ctx);
	/**
	 * Exit a parse tree produced by {@link ArgParseParser#val}.
	 * @param ctx the parse tree
	 */
	void exitVal(ArgParseParser.ValContext ctx);
}