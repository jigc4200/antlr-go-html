// Generated from /home/jigc4200/Documents/antlr-go-html/parser/L5.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class L5Parser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, POSICION=43, TEXTO=44, WS=45;
	public static final int
		RULE_pagina = 0, RULE_cabecera = 1, RULE_titulo = 2, RULE_cuerpo = 3, 
		RULE_pie = 4, RULE_elemento = 5, RULE_estilo = 6, RULE_formato = 7, RULE_posicion = 8, 
		RULE_lista = 9, RULE_elementoLista = 10, RULE_enlace = 11, RULE_singleton = 12, 
		RULE_texto = 13;
	private static String[] makeRuleNames() {
		return new String[] {
			"pagina", "cabecera", "titulo", "cuerpo", "pie", "elemento", "estilo", 
			"formato", "posicion", "lista", "elementoLista", "enlace", "singleton", 
			"texto"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Pagina'", "'FinPagina'", "'Cabecera'", "'FinCabecera'", "'Titulo'", 
			"'FinTitulo'", "'Cuerpo'", "'FinCuerpo'", "'Ppagina'", "'FPpagina'", 
			"'Negrita'", "'FinNegrita'", "'Cursiva'", "'FinCursiva'", "'Subrayado'", 
			"'FinSubrayado'", "'Tachado'", "'FinTachado'", "'Subindice'", "'FinSubindice'", 
			"'Superindice'", "'FinSuperindice'", "'Tama\\u00F1o'", "'FinTama\\u00F1o'", 
			"'Sangrado'", "'FinSangrado'", "'Parrafo'", "'FinParrafo'", "'Posicion'", 
			"'FinPosicion'", "'Lista'", "'FinLista'", "'ElementoLista'", "'FinElementoLista'", 
			"'Enlace'", "'Con'", "'FinCon'", "'Mostrar'", "'FinMostrar'", "'FinEnlace'", 
			"'Linea'", "'Salto'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, "POSICION", "TEXTO", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
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
	public String getGrammarFileName() { return "L5.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public L5Parser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PaginaContext extends ParserRuleContext {
		public CabeceraContext cabecera() {
			return getRuleContext(CabeceraContext.class,0);
		}
		public CuerpoContext cuerpo() {
			return getRuleContext(CuerpoContext.class,0);
		}
		public PieContext pie() {
			return getRuleContext(PieContext.class,0);
		}
		public PaginaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pagina; }
	}

	public final PaginaContext pagina() throws RecognitionException {
		PaginaContext _localctx = new PaginaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_pagina);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(28);
			match(T__0);
			setState(29);
			cabecera();
			setState(30);
			cuerpo();
			setState(31);
			pie();
			setState(32);
			match(T__1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CabeceraContext extends ParserRuleContext {
		public TituloContext titulo() {
			return getRuleContext(TituloContext.class,0);
		}
		public CabeceraContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cabecera; }
	}

	public final CabeceraContext cabecera() throws RecognitionException {
		CabeceraContext _localctx = new CabeceraContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_cabecera);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(34);
			match(T__2);
			setState(35);
			titulo();
			setState(36);
			match(T__3);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TituloContext extends ParserRuleContext {
		public TextoContext texto() {
			return getRuleContext(TextoContext.class,0);
		}
		public TituloContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_titulo; }
	}

	public final TituloContext titulo() throws RecognitionException {
		TituloContext _localctx = new TituloContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_titulo);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			match(T__4);
			setState(39);
			texto();
			setState(40);
			match(T__5);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CuerpoContext extends ParserRuleContext {
		public List<ElementoContext> elemento() {
			return getRuleContexts(ElementoContext.class);
		}
		public ElementoContext elemento(int i) {
			return getRuleContext(ElementoContext.class,i);
		}
		public CuerpoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cuerpo; }
	}

	public final CuerpoContext cuerpo() throws RecognitionException {
		CuerpoContext _localctx = new CuerpoContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_cuerpo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42);
			match(T__6);
			setState(46);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 6633755944960L) != 0)) {
				{
				{
				setState(43);
				elemento();
				}
				}
				setState(48);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(49);
			match(T__7);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PieContext extends ParserRuleContext {
		public TextoContext texto() {
			return getRuleContext(TextoContext.class,0);
		}
		public PieContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pie; }
	}

	public final PieContext pie() throws RecognitionException {
		PieContext _localctx = new PieContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_pie);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51);
			match(T__8);
			setState(52);
			texto();
			setState(53);
			match(T__9);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ElementoContext extends ParserRuleContext {
		public EstiloContext estilo() {
			return getRuleContext(EstiloContext.class,0);
		}
		public FormatoContext formato() {
			return getRuleContext(FormatoContext.class,0);
		}
		public ListaContext lista() {
			return getRuleContext(ListaContext.class,0);
		}
		public EnlaceContext enlace() {
			return getRuleContext(EnlaceContext.class,0);
		}
		public SingletonContext singleton() {
			return getRuleContext(SingletonContext.class,0);
		}
		public ElementoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elemento; }
	}

	public final ElementoContext elemento() throws RecognitionException {
		ElementoContext _localctx = new ElementoContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_elemento);
		try {
			setState(60);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__10:
			case T__12:
			case T__14:
			case T__16:
			case T__18:
			case T__20:
			case T__22:
				enterOuterAlt(_localctx, 1);
				{
				setState(55);
				estilo();
				}
				break;
			case T__24:
			case T__26:
				enterOuterAlt(_localctx, 2);
				{
				setState(56);
				formato();
				}
				break;
			case T__30:
				enterOuterAlt(_localctx, 3);
				{
				setState(57);
				lista();
				}
				break;
			case T__34:
				enterOuterAlt(_localctx, 4);
				{
				setState(58);
				enlace();
				}
				break;
			case T__40:
			case T__41:
				enterOuterAlt(_localctx, 5);
				{
				setState(59);
				singleton();
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EstiloContext extends ParserRuleContext {
		public TextoContext texto() {
			return getRuleContext(TextoContext.class,0);
		}
		public EstiloContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_estilo; }
	}

	public final EstiloContext estilo() throws RecognitionException {
		EstiloContext _localctx = new EstiloContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_estilo);
		try {
			setState(90);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__10:
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				match(T__10);
				setState(63);
				texto();
				setState(64);
				match(T__11);
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				match(T__12);
				setState(67);
				texto();
				setState(68);
				match(T__13);
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 3);
				{
				setState(70);
				match(T__14);
				setState(71);
				texto();
				setState(72);
				match(T__15);
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 4);
				{
				setState(74);
				match(T__16);
				setState(75);
				texto();
				setState(76);
				match(T__17);
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 5);
				{
				setState(78);
				match(T__18);
				setState(79);
				texto();
				setState(80);
				match(T__19);
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 6);
				{
				setState(82);
				match(T__20);
				setState(83);
				texto();
				setState(84);
				match(T__21);
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 7);
				{
				setState(86);
				match(T__22);
				setState(87);
				texto();
				setState(88);
				match(T__23);
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FormatoContext extends ParserRuleContext {
		public TextoContext texto() {
			return getRuleContext(TextoContext.class,0);
		}
		public PosicionContext posicion() {
			return getRuleContext(PosicionContext.class,0);
		}
		public FormatoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formato; }
	}

	public final FormatoContext formato() throws RecognitionException {
		FormatoContext _localctx = new FormatoContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_formato);
		int _la;
		try {
			setState(103);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__24:
				enterOuterAlt(_localctx, 1);
				{
				setState(92);
				match(T__24);
				setState(93);
				texto();
				setState(94);
				match(T__25);
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 2);
				{
				setState(96);
				match(T__26);
				setState(98);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__28) {
					{
					setState(97);
					posicion();
					}
				}

				setState(100);
				texto();
				setState(101);
				match(T__27);
				}
				break;
			default:
				throw new NoViableAltException(this);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PosicionContext extends ParserRuleContext {
		public TerminalNode POSICION() { return getToken(L5Parser.POSICION, 0); }
		public PosicionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_posicion; }
	}

	public final PosicionContext posicion() throws RecognitionException {
		PosicionContext _localctx = new PosicionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_posicion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(T__28);
			setState(106);
			match(POSICION);
			setState(107);
			match(T__29);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ListaContext extends ParserRuleContext {
		public List<ElementoListaContext> elementoLista() {
			return getRuleContexts(ElementoListaContext.class);
		}
		public ElementoListaContext elementoLista(int i) {
			return getRuleContext(ElementoListaContext.class,i);
		}
		public ListaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista; }
	}

	public final ListaContext lista() throws RecognitionException {
		ListaContext _localctx = new ListaContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_lista);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(T__30);
			setState(111); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(110);
				elementoLista();
				}
				}
				setState(113); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__32 );
			setState(115);
			match(T__31);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ElementoListaContext extends ParserRuleContext {
		public TextoContext texto() {
			return getRuleContext(TextoContext.class,0);
		}
		public ElementoListaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elementoLista; }
	}

	public final ElementoListaContext elementoLista() throws RecognitionException {
		ElementoListaContext _localctx = new ElementoListaContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_elementoLista);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(T__32);
			setState(118);
			texto();
			setState(119);
			match(T__33);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnlaceContext extends ParserRuleContext {
		public List<TextoContext> texto() {
			return getRuleContexts(TextoContext.class);
		}
		public TextoContext texto(int i) {
			return getRuleContext(TextoContext.class,i);
		}
		public EnlaceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enlace; }
	}

	public final EnlaceContext enlace() throws RecognitionException {
		EnlaceContext _localctx = new EnlaceContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_enlace);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			match(T__34);
			setState(122);
			match(T__35);
			setState(123);
			texto();
			setState(124);
			match(T__36);
			setState(125);
			match(T__37);
			setState(126);
			texto();
			setState(127);
			match(T__38);
			setState(128);
			match(T__39);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SingletonContext extends ParserRuleContext {
		public SingletonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleton; }
	}

	public final SingletonContext singleton() throws RecognitionException {
		SingletonContext _localctx = new SingletonContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_singleton);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			_la = _input.LA(1);
			if ( !(_la==T__40 || _la==T__41) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class TextoContext extends ParserRuleContext {
		public List<TerminalNode> TEXTO() { return getTokens(L5Parser.TEXTO); }
		public TerminalNode TEXTO(int i) {
			return getToken(L5Parser.TEXTO, i);
		}
		public TextoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_texto; }
	}

	public final TextoContext texto() throws RecognitionException {
		TextoContext _localctx = new TextoContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_texto);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(132);
				match(TEXTO);
				}
				}
				setState(135); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==TEXTO );
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
		"\u0004\u0001-\u008a\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003"+
		"\u0005\u0003-\b\u0003\n\u0003\f\u00030\t\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005=\b\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006[\b\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007"+
		"c\b\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007h\b\u0007\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0004\tp\b\t\u000b\t\f\tq\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\f\u0001\f\u0001\r\u0004\r\u0086\b\r\u000b\r\f\r\u0087\u0001"+
		"\r\u0000\u0000\u000e\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014"+
		"\u0016\u0018\u001a\u0000\u0001\u0001\u0000)*\u008a\u0000\u001c\u0001\u0000"+
		"\u0000\u0000\u0002\"\u0001\u0000\u0000\u0000\u0004&\u0001\u0000\u0000"+
		"\u0000\u0006*\u0001\u0000\u0000\u0000\b3\u0001\u0000\u0000\u0000\n<\u0001"+
		"\u0000\u0000\u0000\fZ\u0001\u0000\u0000\u0000\u000eg\u0001\u0000\u0000"+
		"\u0000\u0010i\u0001\u0000\u0000\u0000\u0012m\u0001\u0000\u0000\u0000\u0014"+
		"u\u0001\u0000\u0000\u0000\u0016y\u0001\u0000\u0000\u0000\u0018\u0082\u0001"+
		"\u0000\u0000\u0000\u001a\u0085\u0001\u0000\u0000\u0000\u001c\u001d\u0005"+
		"\u0001\u0000\u0000\u001d\u001e\u0003\u0002\u0001\u0000\u001e\u001f\u0003"+
		"\u0006\u0003\u0000\u001f \u0003\b\u0004\u0000 !\u0005\u0002\u0000\u0000"+
		"!\u0001\u0001\u0000\u0000\u0000\"#\u0005\u0003\u0000\u0000#$\u0003\u0004"+
		"\u0002\u0000$%\u0005\u0004\u0000\u0000%\u0003\u0001\u0000\u0000\u0000"+
		"&\'\u0005\u0005\u0000\u0000\'(\u0003\u001a\r\u0000()\u0005\u0006\u0000"+
		"\u0000)\u0005\u0001\u0000\u0000\u0000*.\u0005\u0007\u0000\u0000+-\u0003"+
		"\n\u0005\u0000,+\u0001\u0000\u0000\u0000-0\u0001\u0000\u0000\u0000.,\u0001"+
		"\u0000\u0000\u0000./\u0001\u0000\u0000\u0000/1\u0001\u0000\u0000\u0000"+
		"0.\u0001\u0000\u0000\u000012\u0005\b\u0000\u00002\u0007\u0001\u0000\u0000"+
		"\u000034\u0005\t\u0000\u000045\u0003\u001a\r\u000056\u0005\n\u0000\u0000"+
		"6\t\u0001\u0000\u0000\u00007=\u0003\f\u0006\u00008=\u0003\u000e\u0007"+
		"\u00009=\u0003\u0012\t\u0000:=\u0003\u0016\u000b\u0000;=\u0003\u0018\f"+
		"\u0000<7\u0001\u0000\u0000\u0000<8\u0001\u0000\u0000\u0000<9\u0001\u0000"+
		"\u0000\u0000<:\u0001\u0000\u0000\u0000<;\u0001\u0000\u0000\u0000=\u000b"+
		"\u0001\u0000\u0000\u0000>?\u0005\u000b\u0000\u0000?@\u0003\u001a\r\u0000"+
		"@A\u0005\f\u0000\u0000A[\u0001\u0000\u0000\u0000BC\u0005\r\u0000\u0000"+
		"CD\u0003\u001a\r\u0000DE\u0005\u000e\u0000\u0000E[\u0001\u0000\u0000\u0000"+
		"FG\u0005\u000f\u0000\u0000GH\u0003\u001a\r\u0000HI\u0005\u0010\u0000\u0000"+
		"I[\u0001\u0000\u0000\u0000JK\u0005\u0011\u0000\u0000KL\u0003\u001a\r\u0000"+
		"LM\u0005\u0012\u0000\u0000M[\u0001\u0000\u0000\u0000NO\u0005\u0013\u0000"+
		"\u0000OP\u0003\u001a\r\u0000PQ\u0005\u0014\u0000\u0000Q[\u0001\u0000\u0000"+
		"\u0000RS\u0005\u0015\u0000\u0000ST\u0003\u001a\r\u0000TU\u0005\u0016\u0000"+
		"\u0000U[\u0001\u0000\u0000\u0000VW\u0005\u0017\u0000\u0000WX\u0003\u001a"+
		"\r\u0000XY\u0005\u0018\u0000\u0000Y[\u0001\u0000\u0000\u0000Z>\u0001\u0000"+
		"\u0000\u0000ZB\u0001\u0000\u0000\u0000ZF\u0001\u0000\u0000\u0000ZJ\u0001"+
		"\u0000\u0000\u0000ZN\u0001\u0000\u0000\u0000ZR\u0001\u0000\u0000\u0000"+
		"ZV\u0001\u0000\u0000\u0000[\r\u0001\u0000\u0000\u0000\\]\u0005\u0019\u0000"+
		"\u0000]^\u0003\u001a\r\u0000^_\u0005\u001a\u0000\u0000_h\u0001\u0000\u0000"+
		"\u0000`b\u0005\u001b\u0000\u0000ac\u0003\u0010\b\u0000ba\u0001\u0000\u0000"+
		"\u0000bc\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000de\u0003\u001a"+
		"\r\u0000ef\u0005\u001c\u0000\u0000fh\u0001\u0000\u0000\u0000g\\\u0001"+
		"\u0000\u0000\u0000g`\u0001\u0000\u0000\u0000h\u000f\u0001\u0000\u0000"+
		"\u0000ij\u0005\u001d\u0000\u0000jk\u0005+\u0000\u0000kl\u0005\u001e\u0000"+
		"\u0000l\u0011\u0001\u0000\u0000\u0000mo\u0005\u001f\u0000\u0000np\u0003"+
		"\u0014\n\u0000on\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qo\u0001"+
		"\u0000\u0000\u0000qr\u0001\u0000\u0000\u0000rs\u0001\u0000\u0000\u0000"+
		"st\u0005 \u0000\u0000t\u0013\u0001\u0000\u0000\u0000uv\u0005!\u0000\u0000"+
		"vw\u0003\u001a\r\u0000wx\u0005\"\u0000\u0000x\u0015\u0001\u0000\u0000"+
		"\u0000yz\u0005#\u0000\u0000z{\u0005$\u0000\u0000{|\u0003\u001a\r\u0000"+
		"|}\u0005%\u0000\u0000}~\u0005&\u0000\u0000~\u007f\u0003\u001a\r\u0000"+
		"\u007f\u0080\u0005\'\u0000\u0000\u0080\u0081\u0005(\u0000\u0000\u0081"+
		"\u0017\u0001\u0000\u0000\u0000\u0082\u0083\u0007\u0000\u0000\u0000\u0083"+
		"\u0019\u0001\u0000\u0000\u0000\u0084\u0086\u0005,\u0000\u0000\u0085\u0084"+
		"\u0001\u0000\u0000\u0000\u0086\u0087\u0001\u0000\u0000\u0000\u0087\u0085"+
		"\u0001\u0000\u0000\u0000\u0087\u0088\u0001\u0000\u0000\u0000\u0088\u001b"+
		"\u0001\u0000\u0000\u0000\u0007.<Zbgq\u0087";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}