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
		PAGINA=1, CABECERA=2, TITULO=3, FINTITULO=4, FINCABECERA=5, CUERPO=6, 
		FINCUERPO=7, PPAGINA=8, FPPAGINA=9, FINPAGINA=10, NEGRITA=11, FINNEGRITA=12, 
		CURSIVA=13, FINCURSIVA=14, SUBRAYADO=15, FINSUBRAYADO=16, TACHADO=17, 
		FINTACHADO=18, SUBINDICE=19, FINSUBINDICE=20, SUPERINDICE=21, FINSUPERINDICE=22, 
		TAMANO=23, FINTAMANO=24, SANGRADO=25, FINSANGRADO=26, PARRAFO=27, FINPARRAFO=28, 
		POSICION_KW=29, FINPOSICION=30, ID=31, FINID=32, CLASE=33, FINCLASE=34, 
		LISTA=35, FINLISTA=36, ELEMENTOLISTA=37, FINELEMENTOLISTA=38, ENLACE=39, 
		FINENLACE=40, CON=41, FINCON=42, MOSTRAR=43, FINMOSTRAR=44, IMAGEN=45, 
		FINIMAGEN=46, SRC=47, FINSRC=48, ALT=49, FINALT=50, ENCABEZADO=51, FINENCABEZADO=52, 
		LINEA=53, SALTO=54, ESTILO=55, FINESTILO=56, POSICION=57, ATTR_VAL=58, 
		WS=59;
	public static final int
		RULE_pagina = 0, RULE_cabecera = 1, RULE_titulo = 2, RULE_cuerpo = 3, 
		RULE_pie = 4, RULE_elemento = 5, RULE_estilo = 6, RULE_formato = 7, RULE_atributo = 8, 
		RULE_lista = 9, RULE_elementoLista = 10, RULE_enlace = 11, RULE_imagen = 12, 
		RULE_encabezado = 13, RULE_singleton = 14, RULE_contenido = 15;
	private static String[] makeRuleNames() {
		return new String[] {
			"pagina", "cabecera", "titulo", "cuerpo", "pie", "elemento", "estilo", 
			"formato", "atributo", "lista", "elementoLista", "enlace", "imagen", 
			"encabezado", "singleton", "contenido"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Pagina'", "'Cabecera'", "'Titulo'", "'FinTitulo'", "'FinCabecera'", 
			"'Cuerpo'", "'FinCuerpo'", "'Ppagina'", "'FPpagina'", "'FinPagina'", 
			"'Negrita'", "'FinNegrita'", "'Cursiva'", "'FinCursiva'", "'Subrayado'", 
			"'FinSubrayado'", "'Tachado'", "'FinTachado'", "'Subindice'", "'FinSubindice'", 
			"'Superindice'", "'FinSuperindice'", "'Tama\\u00F1o'", "'FinTama\\u00F1o'", 
			"'Sangrado'", "'FinSangrado'", "'Parrafo'", "'FinParrafo'", "'Posicion'", 
			"'FinPosicion'", "'Id'", "'FinId'", "'Clase'", "'FinClase'", "'Lista'", 
			"'FinLista'", "'ElementoLista'", "'FinElementoLista'", "'Enlace'", "'FinEnlace'", 
			"'Con'", "'FinCon'", "'Mostrar'", "'FinMostrar'", "'Imagen'", "'FinImagen'", 
			"'Src'", "'FinSrc'", "'Alt'", "'FinAlt'", "'Encabezado'", "'FinEncabezado'", 
			"'Linea'", "'Salto'", "'Estilo'", "'FinEstilo'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PAGINA", "CABECERA", "TITULO", "FINTITULO", "FINCABECERA", "CUERPO", 
			"FINCUERPO", "PPAGINA", "FPPAGINA", "FINPAGINA", "NEGRITA", "FINNEGRITA", 
			"CURSIVA", "FINCURSIVA", "SUBRAYADO", "FINSUBRAYADO", "TACHADO", "FINTACHADO", 
			"SUBINDICE", "FINSUBINDICE", "SUPERINDICE", "FINSUPERINDICE", "TAMANO", 
			"FINTAMANO", "SANGRADO", "FINSANGRADO", "PARRAFO", "FINPARRAFO", "POSICION_KW", 
			"FINPOSICION", "ID", "FINID", "CLASE", "FINCLASE", "LISTA", "FINLISTA", 
			"ELEMENTOLISTA", "FINELEMENTOLISTA", "ENLACE", "FINENLACE", "CON", "FINCON", 
			"MOSTRAR", "FINMOSTRAR", "IMAGEN", "FINIMAGEN", "SRC", "FINSRC", "ALT", 
			"FINALT", "ENCABEZADO", "FINENCABEZADO", "LINEA", "SALTO", "ESTILO", 
			"FINESTILO", "POSICION", "ATTR_VAL", "WS"
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
		public TerminalNode PAGINA() { return getToken(L5Parser.PAGINA, 0); }
		public CabeceraContext cabecera() {
			return getRuleContext(CabeceraContext.class,0);
		}
		public CuerpoContext cuerpo() {
			return getRuleContext(CuerpoContext.class,0);
		}
		public PieContext pie() {
			return getRuleContext(PieContext.class,0);
		}
		public TerminalNode FINPAGINA() { return getToken(L5Parser.FINPAGINA, 0); }
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
			setState(32);
			match(PAGINA);
			setState(33);
			cabecera();
			setState(34);
			cuerpo();
			setState(35);
			pie();
			setState(36);
			match(FINPAGINA);
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
		public TerminalNode CABECERA() { return getToken(L5Parser.CABECERA, 0); }
		public TituloContext titulo() {
			return getRuleContext(TituloContext.class,0);
		}
		public TerminalNode FINCABECERA() { return getToken(L5Parser.FINCABECERA, 0); }
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
			setState(38);
			match(CABECERA);
			setState(39);
			titulo();
			setState(40);
			match(FINCABECERA);
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
		public TerminalNode TITULO() { return getToken(L5Parser.TITULO, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FINTITULO() { return getToken(L5Parser.FINTITULO, 0); }
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
			setState(42);
			match(TITULO);
			setState(43);
			contenido();
			setState(44);
			match(FINTITULO);
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
		public TerminalNode CUERPO() { return getToken(L5Parser.CUERPO, 0); }
		public TerminalNode FINCUERPO() { return getToken(L5Parser.FINCUERPO, 0); }
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
			setState(46);
			match(CUERPO);
			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 29309166244505600L) != 0)) {
				{
				{
				setState(47);
				elemento();
				}
				}
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(53);
			match(FINCUERPO);
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
		public TerminalNode PPAGINA() { return getToken(L5Parser.PPAGINA, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FPPAGINA() { return getToken(L5Parser.FPPAGINA, 0); }
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
			setState(55);
			match(PPAGINA);
			setState(56);
			contenido();
			setState(57);
			match(FPPAGINA);
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
		public ImagenContext imagen() {
			return getRuleContext(ImagenContext.class,0);
		}
		public EncabezadoContext encabezado() {
			return getRuleContext(EncabezadoContext.class,0);
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
			setState(66);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NEGRITA:
			case CURSIVA:
			case SUBRAYADO:
			case TACHADO:
			case SUBINDICE:
			case SUPERINDICE:
			case TAMANO:
				enterOuterAlt(_localctx, 1);
				{
				setState(59);
				estilo();
				}
				break;
			case SANGRADO:
			case PARRAFO:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				formato();
				}
				break;
			case LISTA:
				enterOuterAlt(_localctx, 3);
				{
				setState(61);
				lista();
				}
				break;
			case ENLACE:
				enterOuterAlt(_localctx, 4);
				{
				setState(62);
				enlace();
				}
				break;
			case LINEA:
			case SALTO:
				enterOuterAlt(_localctx, 5);
				{
				setState(63);
				singleton();
				}
				break;
			case IMAGEN:
				enterOuterAlt(_localctx, 6);
				{
				setState(64);
				imagen();
				}
				break;
			case ENCABEZADO:
				enterOuterAlt(_localctx, 7);
				{
				setState(65);
				encabezado();
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
		public TerminalNode NEGRITA() { return getToken(L5Parser.NEGRITA, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FINNEGRITA() { return getToken(L5Parser.FINNEGRITA, 0); }
		public TerminalNode CURSIVA() { return getToken(L5Parser.CURSIVA, 0); }
		public TerminalNode FINCURSIVA() { return getToken(L5Parser.FINCURSIVA, 0); }
		public TerminalNode SUBRAYADO() { return getToken(L5Parser.SUBRAYADO, 0); }
		public TerminalNode FINSUBRAYADO() { return getToken(L5Parser.FINSUBRAYADO, 0); }
		public TerminalNode TACHADO() { return getToken(L5Parser.TACHADO, 0); }
		public TerminalNode FINTACHADO() { return getToken(L5Parser.FINTACHADO, 0); }
		public TerminalNode SUBINDICE() { return getToken(L5Parser.SUBINDICE, 0); }
		public TerminalNode FINSUBINDICE() { return getToken(L5Parser.FINSUBINDICE, 0); }
		public TerminalNode SUPERINDICE() { return getToken(L5Parser.SUPERINDICE, 0); }
		public TerminalNode FINSUPERINDICE() { return getToken(L5Parser.FINSUPERINDICE, 0); }
		public TerminalNode TAMANO() { return getToken(L5Parser.TAMANO, 0); }
		public TerminalNode ATTR_VAL() { return getToken(L5Parser.ATTR_VAL, 0); }
		public TerminalNode FINTAMANO() { return getToken(L5Parser.FINTAMANO, 0); }
		public EstiloContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_estilo; }
	}

	public final EstiloContext estilo() throws RecognitionException {
		EstiloContext _localctx = new EstiloContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_estilo);
		try {
			setState(97);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NEGRITA:
				enterOuterAlt(_localctx, 1);
				{
				setState(68);
				match(NEGRITA);
				setState(69);
				contenido();
				setState(70);
				match(FINNEGRITA);
				}
				break;
			case CURSIVA:
				enterOuterAlt(_localctx, 2);
				{
				setState(72);
				match(CURSIVA);
				setState(73);
				contenido();
				setState(74);
				match(FINCURSIVA);
				}
				break;
			case SUBRAYADO:
				enterOuterAlt(_localctx, 3);
				{
				setState(76);
				match(SUBRAYADO);
				setState(77);
				contenido();
				setState(78);
				match(FINSUBRAYADO);
				}
				break;
			case TACHADO:
				enterOuterAlt(_localctx, 4);
				{
				setState(80);
				match(TACHADO);
				setState(81);
				contenido();
				setState(82);
				match(FINTACHADO);
				}
				break;
			case SUBINDICE:
				enterOuterAlt(_localctx, 5);
				{
				setState(84);
				match(SUBINDICE);
				setState(85);
				contenido();
				setState(86);
				match(FINSUBINDICE);
				}
				break;
			case SUPERINDICE:
				enterOuterAlt(_localctx, 6);
				{
				setState(88);
				match(SUPERINDICE);
				setState(89);
				contenido();
				setState(90);
				match(FINSUPERINDICE);
				}
				break;
			case TAMANO:
				enterOuterAlt(_localctx, 7);
				{
				setState(92);
				match(TAMANO);
				setState(93);
				match(ATTR_VAL);
				setState(94);
				contenido();
				setState(95);
				match(FINTAMANO);
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
		public TerminalNode SANGRADO() { return getToken(L5Parser.SANGRADO, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FINSANGRADO() { return getToken(L5Parser.FINSANGRADO, 0); }
		public TerminalNode PARRAFO() { return getToken(L5Parser.PARRAFO, 0); }
		public TerminalNode FINPARRAFO() { return getToken(L5Parser.FINPARRAFO, 0); }
		public EstiloContext estilo() {
			return getRuleContext(EstiloContext.class,0);
		}
		public List<AtributoContext> atributo() {
			return getRuleContexts(AtributoContext.class);
		}
		public AtributoContext atributo(int i) {
			return getRuleContext(AtributoContext.class,i);
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
			setState(116);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SANGRADO:
				enterOuterAlt(_localctx, 1);
				{
				setState(99);
				match(SANGRADO);
				setState(100);
				contenido();
				setState(101);
				match(FINSANGRADO);
				}
				break;
			case PARRAFO:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				match(PARRAFO);
				setState(105);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
				case 1:
					{
					setState(104);
					estilo();
					}
					break;
				}
				setState(110);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 11274289152L) != 0)) {
					{
					{
					setState(107);
					atributo();
					}
					}
					setState(112);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(113);
				contenido();
				setState(114);
				match(FINPARRAFO);
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
	public static class AtributoContext extends ParserRuleContext {
		public TerminalNode POSICION_KW() { return getToken(L5Parser.POSICION_KW, 0); }
		public TerminalNode POSICION() { return getToken(L5Parser.POSICION, 0); }
		public TerminalNode FINPOSICION() { return getToken(L5Parser.FINPOSICION, 0); }
		public TerminalNode ID() { return getToken(L5Parser.ID, 0); }
		public TerminalNode ATTR_VAL() { return getToken(L5Parser.ATTR_VAL, 0); }
		public TerminalNode FINID() { return getToken(L5Parser.FINID, 0); }
		public TerminalNode CLASE() { return getToken(L5Parser.CLASE, 0); }
		public TerminalNode FINCLASE() { return getToken(L5Parser.FINCLASE, 0); }
		public AtributoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atributo; }
	}

	public final AtributoContext atributo() throws RecognitionException {
		AtributoContext _localctx = new AtributoContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_atributo);
		try {
			setState(127);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case POSICION_KW:
				enterOuterAlt(_localctx, 1);
				{
				setState(118);
				match(POSICION_KW);
				setState(119);
				match(POSICION);
				setState(120);
				match(FINPOSICION);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(121);
				match(ID);
				setState(122);
				match(ATTR_VAL);
				setState(123);
				match(FINID);
				}
				break;
			case CLASE:
				enterOuterAlt(_localctx, 3);
				{
				setState(124);
				match(CLASE);
				setState(125);
				match(ATTR_VAL);
				setState(126);
				match(FINCLASE);
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
	public static class ListaContext extends ParserRuleContext {
		public TerminalNode LISTA() { return getToken(L5Parser.LISTA, 0); }
		public TerminalNode FINLISTA() { return getToken(L5Parser.FINLISTA, 0); }
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
			setState(129);
			match(LISTA);
			setState(131); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(130);
				elementoLista();
				}
				}
				setState(133); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ELEMENTOLISTA );
			setState(135);
			match(FINLISTA);
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
		public TerminalNode ELEMENTOLISTA() { return getToken(L5Parser.ELEMENTOLISTA, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FINELEMENTOLISTA() { return getToken(L5Parser.FINELEMENTOLISTA, 0); }
		public List<ElementoContext> elemento() {
			return getRuleContexts(ElementoContext.class);
		}
		public ElementoContext elemento(int i) {
			return getRuleContext(ElementoContext.class,i);
		}
		public ElementoListaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elementoLista; }
	}

	public final ElementoListaContext elementoLista() throws RecognitionException {
		ElementoListaContext _localctx = new ElementoListaContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_elementoLista);
		int _la;
		try {
			setState(149);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(137);
				match(ELEMENTOLISTA);
				setState(138);
				contenido();
				setState(139);
				match(FINELEMENTOLISTA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(141);
				match(ELEMENTOLISTA);
				setState(143); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(142);
					elemento();
					}
					}
					setState(145); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 29309166244505600L) != 0) );
				setState(147);
				match(FINELEMENTOLISTA);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EnlaceContext extends ParserRuleContext {
		public TerminalNode ENLACE() { return getToken(L5Parser.ENLACE, 0); }
		public TerminalNode CON() { return getToken(L5Parser.CON, 0); }
		public TerminalNode ATTR_VAL() { return getToken(L5Parser.ATTR_VAL, 0); }
		public TerminalNode FINCON() { return getToken(L5Parser.FINCON, 0); }
		public TerminalNode MOSTRAR() { return getToken(L5Parser.MOSTRAR, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FINMOSTRAR() { return getToken(L5Parser.FINMOSTRAR, 0); }
		public TerminalNode FINENLACE() { return getToken(L5Parser.FINENLACE, 0); }
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
			setState(151);
			match(ENLACE);
			setState(152);
			match(CON);
			setState(153);
			match(ATTR_VAL);
			setState(154);
			match(FINCON);
			setState(155);
			match(MOSTRAR);
			setState(156);
			contenido();
			setState(157);
			match(FINMOSTRAR);
			setState(158);
			match(FINENLACE);
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
	public static class ImagenContext extends ParserRuleContext {
		public TerminalNode IMAGEN() { return getToken(L5Parser.IMAGEN, 0); }
		public TerminalNode SRC() { return getToken(L5Parser.SRC, 0); }
		public TerminalNode ATTR_VAL() { return getToken(L5Parser.ATTR_VAL, 0); }
		public TerminalNode FINSRC() { return getToken(L5Parser.FINSRC, 0); }
		public TerminalNode FINIMAGEN() { return getToken(L5Parser.FINIMAGEN, 0); }
		public TerminalNode ALT() { return getToken(L5Parser.ALT, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FINALT() { return getToken(L5Parser.FINALT, 0); }
		public ImagenContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imagen; }
	}

	public final ImagenContext imagen() throws RecognitionException {
		ImagenContext _localctx = new ImagenContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_imagen);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(IMAGEN);
			setState(161);
			match(SRC);
			setState(162);
			match(ATTR_VAL);
			setState(163);
			match(FINSRC);
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ALT) {
				{
				setState(164);
				match(ALT);
				setState(165);
				contenido();
				setState(166);
				match(FINALT);
				}
			}

			setState(170);
			match(FINIMAGEN);
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
	public static class EncabezadoContext extends ParserRuleContext {
		public TerminalNode ENCABEZADO() { return getToken(L5Parser.ENCABEZADO, 0); }
		public TerminalNode ATTR_VAL() { return getToken(L5Parser.ATTR_VAL, 0); }
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public TerminalNode FINENCABEZADO() { return getToken(L5Parser.FINENCABEZADO, 0); }
		public EncabezadoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_encabezado; }
	}

	public final EncabezadoContext encabezado() throws RecognitionException {
		EncabezadoContext _localctx = new EncabezadoContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_encabezado);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(172);
			match(ENCABEZADO);
			setState(173);
			match(ATTR_VAL);
			setState(174);
			contenido();
			setState(175);
			match(FINENCABEZADO);
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
		public TerminalNode LINEA() { return getToken(L5Parser.LINEA, 0); }
		public TerminalNode SALTO() { return getToken(L5Parser.SALTO, 0); }
		public SingletonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleton; }
	}

	public final SingletonContext singleton() throws RecognitionException {
		SingletonContext _localctx = new SingletonContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_singleton);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			_la = _input.LA(1);
			if ( !(_la==LINEA || _la==SALTO) ) {
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
	public static class ContenidoContext extends ParserRuleContext {
		public List<TerminalNode> ATTR_VAL() { return getTokens(L5Parser.ATTR_VAL); }
		public TerminalNode ATTR_VAL(int i) {
			return getToken(L5Parser.ATTR_VAL, i);
		}
		public List<EstiloContext> estilo() {
			return getRuleContexts(EstiloContext.class);
		}
		public EstiloContext estilo(int i) {
			return getRuleContext(EstiloContext.class,i);
		}
		public ContenidoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_contenido; }
	}

	public final ContenidoContext contenido() throws RecognitionException {
		ContenidoContext _localctx = new ContenidoContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_contenido);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 288230376162895872L) != 0)) {
				{
				setState(181);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case ATTR_VAL:
					{
					setState(179);
					match(ATTR_VAL);
					}
					break;
				case NEGRITA:
				case CURSIVA:
				case SUBRAYADO:
				case TACHADO:
				case SUBINDICE:
				case SUPERINDICE:
				case TAMANO:
					{
					setState(180);
					estilo();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(185);
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

	public static final String _serializedATN =
		"\u0004\u0001;\u00bb\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0005\u00031\b\u0003"+
		"\n\u0003\f\u00034\t\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005C\b\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006b\b\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0003\u0007j\b\u0007\u0001\u0007\u0005\u0007m\b\u0007\n\u0007\f\u0007"+
		"p\t\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007u\b\u0007\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003"+
		"\b\u0080\b\b\u0001\t\u0001\t\u0004\t\u0084\b\t\u000b\t\f\t\u0085\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0004\n\u0090"+
		"\b\n\u000b\n\f\n\u0091\u0001\n\u0001\n\u0003\n\u0096\b\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0003\f\u00a9\b\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0005\u000f"+
		"\u00b6\b\u000f\n\u000f\f\u000f\u00b9\t\u000f\u0001\u000f\u0000\u0000\u0010"+
		"\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e\u0000\u0001\u0001\u000056\u00c2\u0000 \u0001\u0000\u0000"+
		"\u0000\u0002&\u0001\u0000\u0000\u0000\u0004*\u0001\u0000\u0000\u0000\u0006"+
		".\u0001\u0000\u0000\u0000\b7\u0001\u0000\u0000\u0000\nB\u0001\u0000\u0000"+
		"\u0000\fa\u0001\u0000\u0000\u0000\u000et\u0001\u0000\u0000\u0000\u0010"+
		"\u007f\u0001\u0000\u0000\u0000\u0012\u0081\u0001\u0000\u0000\u0000\u0014"+
		"\u0095\u0001\u0000\u0000\u0000\u0016\u0097\u0001\u0000\u0000\u0000\u0018"+
		"\u00a0\u0001\u0000\u0000\u0000\u001a\u00ac\u0001\u0000\u0000\u0000\u001c"+
		"\u00b1\u0001\u0000\u0000\u0000\u001e\u00b7\u0001\u0000\u0000\u0000 !\u0005"+
		"\u0001\u0000\u0000!\"\u0003\u0002\u0001\u0000\"#\u0003\u0006\u0003\u0000"+
		"#$\u0003\b\u0004\u0000$%\u0005\n\u0000\u0000%\u0001\u0001\u0000\u0000"+
		"\u0000&\'\u0005\u0002\u0000\u0000\'(\u0003\u0004\u0002\u0000()\u0005\u0005"+
		"\u0000\u0000)\u0003\u0001\u0000\u0000\u0000*+\u0005\u0003\u0000\u0000"+
		"+,\u0003\u001e\u000f\u0000,-\u0005\u0004\u0000\u0000-\u0005\u0001\u0000"+
		"\u0000\u0000.2\u0005\u0006\u0000\u0000/1\u0003\n\u0005\u00000/\u0001\u0000"+
		"\u0000\u000014\u0001\u0000\u0000\u000020\u0001\u0000\u0000\u000023\u0001"+
		"\u0000\u0000\u000035\u0001\u0000\u0000\u000042\u0001\u0000\u0000\u0000"+
		"56\u0005\u0007\u0000\u00006\u0007\u0001\u0000\u0000\u000078\u0005\b\u0000"+
		"\u000089\u0003\u001e\u000f\u00009:\u0005\t\u0000\u0000:\t\u0001\u0000"+
		"\u0000\u0000;C\u0003\f\u0006\u0000<C\u0003\u000e\u0007\u0000=C\u0003\u0012"+
		"\t\u0000>C\u0003\u0016\u000b\u0000?C\u0003\u001c\u000e\u0000@C\u0003\u0018"+
		"\f\u0000AC\u0003\u001a\r\u0000B;\u0001\u0000\u0000\u0000B<\u0001\u0000"+
		"\u0000\u0000B=\u0001\u0000\u0000\u0000B>\u0001\u0000\u0000\u0000B?\u0001"+
		"\u0000\u0000\u0000B@\u0001\u0000\u0000\u0000BA\u0001\u0000\u0000\u0000"+
		"C\u000b\u0001\u0000\u0000\u0000DE\u0005\u000b\u0000\u0000EF\u0003\u001e"+
		"\u000f\u0000FG\u0005\f\u0000\u0000Gb\u0001\u0000\u0000\u0000HI\u0005\r"+
		"\u0000\u0000IJ\u0003\u001e\u000f\u0000JK\u0005\u000e\u0000\u0000Kb\u0001"+
		"\u0000\u0000\u0000LM\u0005\u000f\u0000\u0000MN\u0003\u001e\u000f\u0000"+
		"NO\u0005\u0010\u0000\u0000Ob\u0001\u0000\u0000\u0000PQ\u0005\u0011\u0000"+
		"\u0000QR\u0003\u001e\u000f\u0000RS\u0005\u0012\u0000\u0000Sb\u0001\u0000"+
		"\u0000\u0000TU\u0005\u0013\u0000\u0000UV\u0003\u001e\u000f\u0000VW\u0005"+
		"\u0014\u0000\u0000Wb\u0001\u0000\u0000\u0000XY\u0005\u0015\u0000\u0000"+
		"YZ\u0003\u001e\u000f\u0000Z[\u0005\u0016\u0000\u0000[b\u0001\u0000\u0000"+
		"\u0000\\]\u0005\u0017\u0000\u0000]^\u0005:\u0000\u0000^_\u0003\u001e\u000f"+
		"\u0000_`\u0005\u0018\u0000\u0000`b\u0001\u0000\u0000\u0000aD\u0001\u0000"+
		"\u0000\u0000aH\u0001\u0000\u0000\u0000aL\u0001\u0000\u0000\u0000aP\u0001"+
		"\u0000\u0000\u0000aT\u0001\u0000\u0000\u0000aX\u0001\u0000\u0000\u0000"+
		"a\\\u0001\u0000\u0000\u0000b\r\u0001\u0000\u0000\u0000cd\u0005\u0019\u0000"+
		"\u0000de\u0003\u001e\u000f\u0000ef\u0005\u001a\u0000\u0000fu\u0001\u0000"+
		"\u0000\u0000gi\u0005\u001b\u0000\u0000hj\u0003\f\u0006\u0000ih\u0001\u0000"+
		"\u0000\u0000ij\u0001\u0000\u0000\u0000jn\u0001\u0000\u0000\u0000km\u0003"+
		"\u0010\b\u0000lk\u0001\u0000\u0000\u0000mp\u0001\u0000\u0000\u0000nl\u0001"+
		"\u0000\u0000\u0000no\u0001\u0000\u0000\u0000oq\u0001\u0000\u0000\u0000"+
		"pn\u0001\u0000\u0000\u0000qr\u0003\u001e\u000f\u0000rs\u0005\u001c\u0000"+
		"\u0000su\u0001\u0000\u0000\u0000tc\u0001\u0000\u0000\u0000tg\u0001\u0000"+
		"\u0000\u0000u\u000f\u0001\u0000\u0000\u0000vw\u0005\u001d\u0000\u0000"+
		"wx\u00059\u0000\u0000x\u0080\u0005\u001e\u0000\u0000yz\u0005\u001f\u0000"+
		"\u0000z{\u0005:\u0000\u0000{\u0080\u0005 \u0000\u0000|}\u0005!\u0000\u0000"+
		"}~\u0005:\u0000\u0000~\u0080\u0005\"\u0000\u0000\u007fv\u0001\u0000\u0000"+
		"\u0000\u007fy\u0001\u0000\u0000\u0000\u007f|\u0001\u0000\u0000\u0000\u0080"+
		"\u0011\u0001\u0000\u0000\u0000\u0081\u0083\u0005#\u0000\u0000\u0082\u0084"+
		"\u0003\u0014\n\u0000\u0083\u0082\u0001\u0000\u0000\u0000\u0084\u0085\u0001"+
		"\u0000\u0000\u0000\u0085\u0083\u0001\u0000\u0000\u0000\u0085\u0086\u0001"+
		"\u0000\u0000\u0000\u0086\u0087\u0001\u0000\u0000\u0000\u0087\u0088\u0005"+
		"$\u0000\u0000\u0088\u0013\u0001\u0000\u0000\u0000\u0089\u008a\u0005%\u0000"+
		"\u0000\u008a\u008b\u0003\u001e\u000f\u0000\u008b\u008c\u0005&\u0000\u0000"+
		"\u008c\u0096\u0001\u0000\u0000\u0000\u008d\u008f\u0005%\u0000\u0000\u008e"+
		"\u0090\u0003\n\u0005\u0000\u008f\u008e\u0001\u0000\u0000\u0000\u0090\u0091"+
		"\u0001\u0000\u0000\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0091\u0092"+
		"\u0001\u0000\u0000\u0000\u0092\u0093\u0001\u0000\u0000\u0000\u0093\u0094"+
		"\u0005&\u0000\u0000\u0094\u0096\u0001\u0000\u0000\u0000\u0095\u0089\u0001"+
		"\u0000\u0000\u0000\u0095\u008d\u0001\u0000\u0000\u0000\u0096\u0015\u0001"+
		"\u0000\u0000\u0000\u0097\u0098\u0005\'\u0000\u0000\u0098\u0099\u0005)"+
		"\u0000\u0000\u0099\u009a\u0005:\u0000\u0000\u009a\u009b\u0005*\u0000\u0000"+
		"\u009b\u009c\u0005+\u0000\u0000\u009c\u009d\u0003\u001e\u000f\u0000\u009d"+
		"\u009e\u0005,\u0000\u0000\u009e\u009f\u0005(\u0000\u0000\u009f\u0017\u0001"+
		"\u0000\u0000\u0000\u00a0\u00a1\u0005-\u0000\u0000\u00a1\u00a2\u0005/\u0000"+
		"\u0000\u00a2\u00a3\u0005:\u0000\u0000\u00a3\u00a8\u00050\u0000\u0000\u00a4"+
		"\u00a5\u00051\u0000\u0000\u00a5\u00a6\u0003\u001e\u000f\u0000\u00a6\u00a7"+
		"\u00052\u0000\u0000\u00a7\u00a9\u0001\u0000\u0000\u0000\u00a8\u00a4\u0001"+
		"\u0000\u0000\u0000\u00a8\u00a9\u0001\u0000\u0000\u0000\u00a9\u00aa\u0001"+
		"\u0000\u0000\u0000\u00aa\u00ab\u0005.\u0000\u0000\u00ab\u0019\u0001\u0000"+
		"\u0000\u0000\u00ac\u00ad\u00053\u0000\u0000\u00ad\u00ae\u0005:\u0000\u0000"+
		"\u00ae\u00af\u0003\u001e\u000f\u0000\u00af\u00b0\u00054\u0000\u0000\u00b0"+
		"\u001b\u0001\u0000\u0000\u0000\u00b1\u00b2\u0007\u0000\u0000\u0000\u00b2"+
		"\u001d\u0001\u0000\u0000\u0000\u00b3\u00b6\u0005:\u0000\u0000\u00b4\u00b6"+
		"\u0003\f\u0006\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b5\u00b4\u0001"+
		"\u0000\u0000\u0000\u00b6\u00b9\u0001\u0000\u0000\u0000\u00b7\u00b5\u0001"+
		"\u0000\u0000\u0000\u00b7\u00b8\u0001\u0000\u0000\u0000\u00b8\u001f\u0001"+
		"\u0000\u0000\u0000\u00b9\u00b7\u0001\u0000\u0000\u0000\r2Baint\u007f\u0085"+
		"\u0091\u0095\u00a8\u00b5\u00b7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}