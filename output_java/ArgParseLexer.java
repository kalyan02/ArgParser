// Generated from ArgParse.g4 by ANTLR 4.6
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ArgParseLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.6", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, DASH=2, SPACE=3, NEWLINE=4, TEXT=5, STRING=6;
	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "DASH", "SPACE", "NEWLINE", "TEXT", "STRING", "ESC", "UNICODE", 
		"HEX"
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


	public ArgParseLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ArgParse.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u0430\ud6d1\u8206\uad2d\u4417\uaef1\u8d80\uaadd\2\bD\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2"+
		"\3\3\3\3\3\3\5\3\33\n\3\3\4\6\4\36\n\4\r\4\16\4\37\3\5\6\5#\n\5\r\5\16"+
		"\5$\3\5\3\5\3\6\6\6*\n\6\r\6\16\6+\3\7\3\7\3\7\7\7\61\n\7\f\7\16\7\64"+
		"\13\7\3\7\3\7\3\b\3\b\3\b\5\b;\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\2\2"+
		"\13\3\3\5\4\7\5\t\6\13\7\r\b\17\2\21\2\23\2\3\2\7\4\2\f\f\17\17\6\2\62"+
		";C\\aac|\4\2$$^^\n\2$$\61\61^^ddhhppttvv\5\2\62;CHchG\2\3\3\2\2\2\2\5"+
		"\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\3\25\3\2\2\2"+
		"\5\32\3\2\2\2\7\35\3\2\2\2\t\"\3\2\2\2\13)\3\2\2\2\r-\3\2\2\2\17\67\3"+
		"\2\2\2\21<\3\2\2\2\23B\3\2\2\2\25\26\7?\2\2\26\4\3\2\2\2\27\30\7/\2\2"+
		"\30\33\7/\2\2\31\33\7/\2\2\32\27\3\2\2\2\32\31\3\2\2\2\33\6\3\2\2\2\34"+
		"\36\7\"\2\2\35\34\3\2\2\2\36\37\3\2\2\2\37\35\3\2\2\2\37 \3\2\2\2 \b\3"+
		"\2\2\2!#\t\2\2\2\"!\3\2\2\2#$\3\2\2\2$\"\3\2\2\2$%\3\2\2\2%&\3\2\2\2&"+
		"\'\b\5\2\2\'\n\3\2\2\2(*\t\3\2\2)(\3\2\2\2*+\3\2\2\2+)\3\2\2\2+,\3\2\2"+
		"\2,\f\3\2\2\2-\62\7$\2\2.\61\5\17\b\2/\61\n\4\2\2\60.\3\2\2\2\60/\3\2"+
		"\2\2\61\64\3\2\2\2\62\60\3\2\2\2\62\63\3\2\2\2\63\65\3\2\2\2\64\62\3\2"+
		"\2\2\65\66\7$\2\2\66\16\3\2\2\2\67:\7^\2\28;\t\5\2\29;\5\21\t\2:8\3\2"+
		"\2\2:9\3\2\2\2;\20\3\2\2\2<=\7w\2\2=>\5\23\n\2>?\5\23\n\2?@\5\23\n\2@"+
		"A\5\23\n\2A\22\3\2\2\2BC\t\6\2\2C\24\3\2\2\2\n\2\32\37$+\60\62:\3\b\2"+
		"\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}