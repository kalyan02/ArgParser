// Generated from ArgParse.g4 by ANTLR 4.6
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ArgParseParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.6", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, DASH=2, SPACE=3, NEWLINE=4, TEXT=5, STRING=6;
	public static final int
		RULE_args = 0, RULE_arg = 1, RULE_dash_opt_k_bool = 2, RULE_dash_opt_kv = 3, 
		RULE_dash_opt_kv_eq = 4, RULE_key = 5, RULE_val = 6;
	public static final String[] ruleNames = {
		"args", "arg", "dash_opt_k_bool", "dash_opt_kv", "dash_opt_kv_eq", "key", 
		"val"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'='"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "DASH", "SPACE", "NEWLINE", "TEXT", "STRING"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "ArgParse.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ArgParseParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class ArgsContext extends ParserRuleContext {
		public List<ArgContext> arg() {
			return getRuleContexts(ArgContext.class);
		}
		public ArgContext arg(int i) {
			return getRuleContext(ArgContext.class,i);
		}
		public List<TerminalNode> SPACE() { return getTokens(ArgParseParser.SPACE); }
		public TerminalNode SPACE(int i) {
			return getToken(ArgParseParser.SPACE, i);
		}
		public ArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_args; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).enterArgs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).exitArgs(this);
		}
	}

	public final ArgsContext args() throws RecognitionException {
		ArgsContext _localctx = new ArgsContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_args);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(14);
			arg();
			setState(19);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SPACE) {
				{
				{
				setState(15);
				match(SPACE);
				setState(16);
				arg();
				}
				}
				setState(21);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgContext extends ParserRuleContext {
		public Dash_opt_kv_eqContext dash_opt_kv_eq() {
			return getRuleContext(Dash_opt_kv_eqContext.class,0);
		}
		public Dash_opt_kvContext dash_opt_kv() {
			return getRuleContext(Dash_opt_kvContext.class,0);
		}
		public Dash_opt_k_boolContext dash_opt_k_bool() {
			return getRuleContext(Dash_opt_k_boolContext.class,0);
		}
		public ArgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arg; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).enterArg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).exitArg(this);
		}
	}

	public final ArgContext arg() throws RecognitionException {
		ArgContext _localctx = new ArgContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_arg);
		try {
			setState(25);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(22);
				dash_opt_kv_eq();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(23);
				dash_opt_kv();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(24);
				dash_opt_k_bool();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dash_opt_k_boolContext extends ParserRuleContext {
		public TerminalNode DASH() { return getToken(ArgParseParser.DASH, 0); }
		public KeyContext key() {
			return getRuleContext(KeyContext.class,0);
		}
		public Dash_opt_k_boolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dash_opt_k_bool; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).enterDash_opt_k_bool(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).exitDash_opt_k_bool(this);
		}
	}

	public final Dash_opt_k_boolContext dash_opt_k_bool() throws RecognitionException {
		Dash_opt_k_boolContext _localctx = new Dash_opt_k_boolContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_dash_opt_k_bool);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(27);
			match(DASH);
			setState(28);
			key();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dash_opt_kvContext extends ParserRuleContext {
		public TerminalNode DASH() { return getToken(ArgParseParser.DASH, 0); }
		public KeyContext key() {
			return getRuleContext(KeyContext.class,0);
		}
		public TerminalNode SPACE() { return getToken(ArgParseParser.SPACE, 0); }
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public Dash_opt_kvContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dash_opt_kv; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).enterDash_opt_kv(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).exitDash_opt_kv(this);
		}
	}

	public final Dash_opt_kvContext dash_opt_kv() throws RecognitionException {
		Dash_opt_kvContext _localctx = new Dash_opt_kvContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_dash_opt_kv);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			match(DASH);
			setState(31);
			key();
			setState(32);
			match(SPACE);
			setState(33);
			val();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Dash_opt_kv_eqContext extends ParserRuleContext {
		public TerminalNode DASH() { return getToken(ArgParseParser.DASH, 0); }
		public KeyContext key() {
			return getRuleContext(KeyContext.class,0);
		}
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public Dash_opt_kv_eqContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dash_opt_kv_eq; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).enterDash_opt_kv_eq(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).exitDash_opt_kv_eq(this);
		}
	}

	public final Dash_opt_kv_eqContext dash_opt_kv_eq() throws RecognitionException {
		Dash_opt_kv_eqContext _localctx = new Dash_opt_kv_eqContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_dash_opt_kv_eq);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			match(DASH);
			setState(36);
			key();
			setState(37);
			match(T__0);
			setState(38);
			val();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyContext extends ParserRuleContext {
		public TerminalNode TEXT() { return getToken(ArgParseParser.TEXT, 0); }
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).enterKey(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).exitKey(this);
		}
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_key);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			match(TEXT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValContext extends ParserRuleContext {
		public TerminalNode TEXT() { return getToken(ArgParseParser.TEXT, 0); }
		public TerminalNode STRING() { return getToken(ArgParseParser.STRING, 0); }
		public ValContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_val; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).enterVal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ArgParseListener ) ((ArgParseListener)listener).exitVal(this);
		}
	}

	public final ValContext val() throws RecognitionException {
		ValContext _localctx = new ValContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_val);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42);
			_la = _input.LA(1);
			if ( !(_la==TEXT || _la==STRING) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u0430\ud6d1\u8206\uad2d\u4417\uaef1\u8d80\uaadd\3\b/\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\2\7\2\24\n\2\f\2"+
		"\16\2\27\13\2\3\3\3\3\3\3\5\3\34\n\3\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3"+
		"\6\3\6\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\b\2\2\t\2\4\6\b\n\f\16\2\3\3\2\7"+
		"\b*\2\20\3\2\2\2\4\33\3\2\2\2\6\35\3\2\2\2\b \3\2\2\2\n%\3\2\2\2\f*\3"+
		"\2\2\2\16,\3\2\2\2\20\25\5\4\3\2\21\22\7\5\2\2\22\24\5\4\3\2\23\21\3\2"+
		"\2\2\24\27\3\2\2\2\25\23\3\2\2\2\25\26\3\2\2\2\26\3\3\2\2\2\27\25\3\2"+
		"\2\2\30\34\5\n\6\2\31\34\5\b\5\2\32\34\5\6\4\2\33\30\3\2\2\2\33\31\3\2"+
		"\2\2\33\32\3\2\2\2\34\5\3\2\2\2\35\36\7\4\2\2\36\37\5\f\7\2\37\7\3\2\2"+
		"\2 !\7\4\2\2!\"\5\f\7\2\"#\7\5\2\2#$\5\16\b\2$\t\3\2\2\2%&\7\4\2\2&\'"+
		"\5\f\7\2\'(\7\3\2\2()\5\16\b\2)\13\3\2\2\2*+\7\7\2\2+\r\3\2\2\2,-\t\2"+
		"\2\2-\17\3\2\2\2\4\25\33";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}