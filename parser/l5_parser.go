// Code generated from L5.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // L5

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type L5Parser struct {
	*antlr.BaseParser
}

var L5ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func l5ParserInit() {
	staticData := &L5ParserStaticData
	staticData.LiteralNames = []string{
		"", "'Pagina'", "'FinPagina'", "'Cabecera'", "'FinCabecera'", "'Titulo'",
		"'FinTitulo'", "'Cuerpo'", "'FinCuerpo'", "'Ppagina'", "'FPpagina'",
		"'Negrita'", "'FinNegrita'", "'Cursiva'", "'FinCursiva'", "'Subrayado'",
		"'FinSubrayado'", "'Tachado'", "'FinTachado'", "'Subindice'", "'FinSubindice'",
		"'Superindice'", "'FinSuperindice'", "'Tama\\u00F1o'", "'FinTama\\u00F1o'",
		"'Sangrado'", "'FinSangrado'", "'Parrafo'", "'FinParrafo'", "'Posicion'",
		"'FinPosicion'", "'Lista'", "'FinLista'", "'ElementoLista'", "'FinElementoLista'",
		"'Enlace'", "'Con'", "'FinCon'", "'Mostrar'", "'FinMostrar'", "'FinEnlace'",
		"'Linea'", "'Salto'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "POSICION", "TEXTO", "WS",
	}
	staticData.RuleNames = []string{
		"pagina", "cabecera", "titulo", "cuerpo", "pie", "elemento", "estilo",
		"formato", "posicion", "lista", "elementoLista", "enlace", "singleton",
		"texto",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 138, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 5,
		3, 45, 8, 3, 10, 3, 12, 3, 48, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 61, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 91, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 99, 8, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 104, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 4, 9,
		112, 8, 9, 11, 9, 12, 9, 113, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 4, 13, 134, 8, 13, 11, 13, 12, 13, 135, 1, 13, 0, 0, 14, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 1, 1, 0, 41, 42, 138, 0, 28,
		1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 42, 1, 0, 0, 0, 8,
		51, 1, 0, 0, 0, 10, 60, 1, 0, 0, 0, 12, 90, 1, 0, 0, 0, 14, 103, 1, 0,
		0, 0, 16, 105, 1, 0, 0, 0, 18, 109, 1, 0, 0, 0, 20, 117, 1, 0, 0, 0, 22,
		121, 1, 0, 0, 0, 24, 130, 1, 0, 0, 0, 26, 133, 1, 0, 0, 0, 28, 29, 5, 1,
		0, 0, 29, 30, 3, 2, 1, 0, 30, 31, 3, 6, 3, 0, 31, 32, 3, 8, 4, 0, 32, 33,
		5, 2, 0, 0, 33, 1, 1, 0, 0, 0, 34, 35, 5, 3, 0, 0, 35, 36, 3, 4, 2, 0,
		36, 37, 5, 4, 0, 0, 37, 3, 1, 0, 0, 0, 38, 39, 5, 5, 0, 0, 39, 40, 3, 26,
		13, 0, 40, 41, 5, 6, 0, 0, 41, 5, 1, 0, 0, 0, 42, 46, 5, 7, 0, 0, 43, 45,
		3, 10, 5, 0, 44, 43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 50, 5,
		8, 0, 0, 50, 7, 1, 0, 0, 0, 51, 52, 5, 9, 0, 0, 52, 53, 3, 26, 13, 0, 53,
		54, 5, 10, 0, 0, 54, 9, 1, 0, 0, 0, 55, 61, 3, 12, 6, 0, 56, 61, 3, 14,
		7, 0, 57, 61, 3, 18, 9, 0, 58, 61, 3, 22, 11, 0, 59, 61, 3, 24, 12, 0,
		60, 55, 1, 0, 0, 0, 60, 56, 1, 0, 0, 0, 60, 57, 1, 0, 0, 0, 60, 58, 1,
		0, 0, 0, 60, 59, 1, 0, 0, 0, 61, 11, 1, 0, 0, 0, 62, 63, 5, 11, 0, 0, 63,
		64, 3, 26, 13, 0, 64, 65, 5, 12, 0, 0, 65, 91, 1, 0, 0, 0, 66, 67, 5, 13,
		0, 0, 67, 68, 3, 26, 13, 0, 68, 69, 5, 14, 0, 0, 69, 91, 1, 0, 0, 0, 70,
		71, 5, 15, 0, 0, 71, 72, 3, 26, 13, 0, 72, 73, 5, 16, 0, 0, 73, 91, 1,
		0, 0, 0, 74, 75, 5, 17, 0, 0, 75, 76, 3, 26, 13, 0, 76, 77, 5, 18, 0, 0,
		77, 91, 1, 0, 0, 0, 78, 79, 5, 19, 0, 0, 79, 80, 3, 26, 13, 0, 80, 81,
		5, 20, 0, 0, 81, 91, 1, 0, 0, 0, 82, 83, 5, 21, 0, 0, 83, 84, 3, 26, 13,
		0, 84, 85, 5, 22, 0, 0, 85, 91, 1, 0, 0, 0, 86, 87, 5, 23, 0, 0, 87, 88,
		3, 26, 13, 0, 88, 89, 5, 24, 0, 0, 89, 91, 1, 0, 0, 0, 90, 62, 1, 0, 0,
		0, 90, 66, 1, 0, 0, 0, 90, 70, 1, 0, 0, 0, 90, 74, 1, 0, 0, 0, 90, 78,
		1, 0, 0, 0, 90, 82, 1, 0, 0, 0, 90, 86, 1, 0, 0, 0, 91, 13, 1, 0, 0, 0,
		92, 93, 5, 25, 0, 0, 93, 94, 3, 26, 13, 0, 94, 95, 5, 26, 0, 0, 95, 104,
		1, 0, 0, 0, 96, 98, 5, 27, 0, 0, 97, 99, 3, 16, 8, 0, 98, 97, 1, 0, 0,
		0, 98, 99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 3, 26, 13, 0, 101,
		102, 5, 28, 0, 0, 102, 104, 1, 0, 0, 0, 103, 92, 1, 0, 0, 0, 103, 96, 1,
		0, 0, 0, 104, 15, 1, 0, 0, 0, 105, 106, 5, 29, 0, 0, 106, 107, 5, 43, 0,
		0, 107, 108, 5, 30, 0, 0, 108, 17, 1, 0, 0, 0, 109, 111, 5, 31, 0, 0, 110,
		112, 3, 20, 10, 0, 111, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 111,
		1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 5, 32,
		0, 0, 116, 19, 1, 0, 0, 0, 117, 118, 5, 33, 0, 0, 118, 119, 3, 26, 13,
		0, 119, 120, 5, 34, 0, 0, 120, 21, 1, 0, 0, 0, 121, 122, 5, 35, 0, 0, 122,
		123, 5, 36, 0, 0, 123, 124, 3, 26, 13, 0, 124, 125, 5, 37, 0, 0, 125, 126,
		5, 38, 0, 0, 126, 127, 3, 26, 13, 0, 127, 128, 5, 39, 0, 0, 128, 129, 5,
		40, 0, 0, 129, 23, 1, 0, 0, 0, 130, 131, 7, 0, 0, 0, 131, 25, 1, 0, 0,
		0, 132, 134, 5, 44, 0, 0, 133, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 27, 1, 0, 0, 0, 7, 46, 60,
		90, 98, 103, 113, 135,
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

// L5ParserInit initializes any static state used to implement L5Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewL5Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func L5ParserInit() {
	staticData := &L5ParserStaticData
	staticData.once.Do(l5ParserInit)
}

// NewL5Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewL5Parser(input antlr.TokenStream) *L5Parser {
	L5ParserInit()
	this := new(L5Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &L5ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "L5.g4"

	return this
}

// L5Parser tokens.
const (
	L5ParserEOF      = antlr.TokenEOF
	L5ParserT__0     = 1
	L5ParserT__1     = 2
	L5ParserT__2     = 3
	L5ParserT__3     = 4
	L5ParserT__4     = 5
	L5ParserT__5     = 6
	L5ParserT__6     = 7
	L5ParserT__7     = 8
	L5ParserT__8     = 9
	L5ParserT__9     = 10
	L5ParserT__10    = 11
	L5ParserT__11    = 12
	L5ParserT__12    = 13
	L5ParserT__13    = 14
	L5ParserT__14    = 15
	L5ParserT__15    = 16
	L5ParserT__16    = 17
	L5ParserT__17    = 18
	L5ParserT__18    = 19
	L5ParserT__19    = 20
	L5ParserT__20    = 21
	L5ParserT__21    = 22
	L5ParserT__22    = 23
	L5ParserT__23    = 24
	L5ParserT__24    = 25
	L5ParserT__25    = 26
	L5ParserT__26    = 27
	L5ParserT__27    = 28
	L5ParserT__28    = 29
	L5ParserT__29    = 30
	L5ParserT__30    = 31
	L5ParserT__31    = 32
	L5ParserT__32    = 33
	L5ParserT__33    = 34
	L5ParserT__34    = 35
	L5ParserT__35    = 36
	L5ParserT__36    = 37
	L5ParserT__37    = 38
	L5ParserT__38    = 39
	L5ParserT__39    = 40
	L5ParserT__40    = 41
	L5ParserT__41    = 42
	L5ParserPOSICION = 43
	L5ParserTEXTO    = 44
	L5ParserWS       = 45
)

// L5Parser rules.
const (
	L5ParserRULE_pagina        = 0
	L5ParserRULE_cabecera      = 1
	L5ParserRULE_titulo        = 2
	L5ParserRULE_cuerpo        = 3
	L5ParserRULE_pie           = 4
	L5ParserRULE_elemento      = 5
	L5ParserRULE_estilo        = 6
	L5ParserRULE_formato       = 7
	L5ParserRULE_posicion      = 8
	L5ParserRULE_lista         = 9
	L5ParserRULE_elementoLista = 10
	L5ParserRULE_enlace        = 11
	L5ParserRULE_singleton     = 12
	L5ParserRULE_texto         = 13
)

// IPaginaContext is an interface to support dynamic dispatch.
type IPaginaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Cabecera() ICabeceraContext
	Cuerpo() ICuerpoContext
	Pie() IPieContext

	// IsPaginaContext differentiates from other interfaces.
	IsPaginaContext()
}

type PaginaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPaginaContext() *PaginaContext {
	var p = new(PaginaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_pagina
	return p
}

func InitEmptyPaginaContext(p *PaginaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_pagina
}

func (*PaginaContext) IsPaginaContext() {}

func NewPaginaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PaginaContext {
	var p = new(PaginaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_pagina

	return p
}

func (s *PaginaContext) GetParser() antlr.Parser { return s.parser }

func (s *PaginaContext) Cabecera() ICabeceraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICabeceraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICabeceraContext)
}

func (s *PaginaContext) Cuerpo() ICuerpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICuerpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICuerpoContext)
}

func (s *PaginaContext) Pie() IPieContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPieContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPieContext)
}

func (s *PaginaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PaginaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PaginaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterPagina(s)
	}
}

func (s *PaginaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitPagina(s)
	}
}

func (p *L5Parser) Pagina() (localctx IPaginaContext) {
	localctx = NewPaginaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, L5ParserRULE_pagina)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(L5ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Cabecera()
	}
	{
		p.SetState(30)
		p.Cuerpo()
	}
	{
		p.SetState(31)
		p.Pie()
	}
	{
		p.SetState(32)
		p.Match(L5ParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICabeceraContext is an interface to support dynamic dispatch.
type ICabeceraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Titulo() ITituloContext

	// IsCabeceraContext differentiates from other interfaces.
	IsCabeceraContext()
}

type CabeceraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCabeceraContext() *CabeceraContext {
	var p = new(CabeceraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_cabecera
	return p
}

func InitEmptyCabeceraContext(p *CabeceraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_cabecera
}

func (*CabeceraContext) IsCabeceraContext() {}

func NewCabeceraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CabeceraContext {
	var p = new(CabeceraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_cabecera

	return p
}

func (s *CabeceraContext) GetParser() antlr.Parser { return s.parser }

func (s *CabeceraContext) Titulo() ITituloContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITituloContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITituloContext)
}

func (s *CabeceraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CabeceraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CabeceraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterCabecera(s)
	}
}

func (s *CabeceraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitCabecera(s)
	}
}

func (p *L5Parser) Cabecera() (localctx ICabeceraContext) {
	localctx = NewCabeceraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, L5ParserRULE_cabecera)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(L5ParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Titulo()
	}
	{
		p.SetState(36)
		p.Match(L5ParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITituloContext is an interface to support dynamic dispatch.
type ITituloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Texto() ITextoContext

	// IsTituloContext differentiates from other interfaces.
	IsTituloContext()
}

type TituloContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTituloContext() *TituloContext {
	var p = new(TituloContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_titulo
	return p
}

func InitEmptyTituloContext(p *TituloContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_titulo
}

func (*TituloContext) IsTituloContext() {}

func NewTituloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TituloContext {
	var p = new(TituloContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_titulo

	return p
}

func (s *TituloContext) GetParser() antlr.Parser { return s.parser }

func (s *TituloContext) Texto() ITextoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextoContext)
}

func (s *TituloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TituloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TituloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterTitulo(s)
	}
}

func (s *TituloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitTitulo(s)
	}
}

func (p *L5Parser) Titulo() (localctx ITituloContext) {
	localctx = NewTituloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, L5ParserRULE_titulo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(L5ParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Texto()
	}
	{
		p.SetState(40)
		p.Match(L5ParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICuerpoContext is an interface to support dynamic dispatch.
type ICuerpoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllElemento() []IElementoContext
	Elemento(i int) IElementoContext

	// IsCuerpoContext differentiates from other interfaces.
	IsCuerpoContext()
}

type CuerpoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCuerpoContext() *CuerpoContext {
	var p = new(CuerpoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_cuerpo
	return p
}

func InitEmptyCuerpoContext(p *CuerpoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_cuerpo
}

func (*CuerpoContext) IsCuerpoContext() {}

func NewCuerpoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CuerpoContext {
	var p = new(CuerpoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_cuerpo

	return p
}

func (s *CuerpoContext) GetParser() antlr.Parser { return s.parser }

func (s *CuerpoContext) AllElemento() []IElementoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementoContext); ok {
			len++
		}
	}

	tst := make([]IElementoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementoContext); ok {
			tst[i] = t.(IElementoContext)
			i++
		}
	}

	return tst
}

func (s *CuerpoContext) Elemento(i int) IElementoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementoContext)
}

func (s *CuerpoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CuerpoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CuerpoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterCuerpo(s)
	}
}

func (s *CuerpoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitCuerpo(s)
	}
}

func (p *L5Parser) Cuerpo() (localctx ICuerpoContext) {
	localctx = NewCuerpoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, L5ParserRULE_cuerpo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(L5ParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6633755944960) != 0 {
		{
			p.SetState(43)
			p.Elemento()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(L5ParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPieContext is an interface to support dynamic dispatch.
type IPieContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Texto() ITextoContext

	// IsPieContext differentiates from other interfaces.
	IsPieContext()
}

type PieContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPieContext() *PieContext {
	var p = new(PieContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_pie
	return p
}

func InitEmptyPieContext(p *PieContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_pie
}

func (*PieContext) IsPieContext() {}

func NewPieContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PieContext {
	var p = new(PieContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_pie

	return p
}

func (s *PieContext) GetParser() antlr.Parser { return s.parser }

func (s *PieContext) Texto() ITextoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextoContext)
}

func (s *PieContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PieContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PieContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterPie(s)
	}
}

func (s *PieContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitPie(s)
	}
}

func (p *L5Parser) Pie() (localctx IPieContext) {
	localctx = NewPieContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, L5ParserRULE_pie)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(L5ParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Texto()
	}
	{
		p.SetState(53)
		p.Match(L5ParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementoContext is an interface to support dynamic dispatch.
type IElementoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Estilo() IEstiloContext
	Formato() IFormatoContext
	Lista() IListaContext
	Enlace() IEnlaceContext
	Singleton() ISingletonContext

	// IsElementoContext differentiates from other interfaces.
	IsElementoContext()
}

type ElementoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementoContext() *ElementoContext {
	var p = new(ElementoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_elemento
	return p
}

func InitEmptyElementoContext(p *ElementoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_elemento
}

func (*ElementoContext) IsElementoContext() {}

func NewElementoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementoContext {
	var p = new(ElementoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_elemento

	return p
}

func (s *ElementoContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementoContext) Estilo() IEstiloContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstiloContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEstiloContext)
}

func (s *ElementoContext) Formato() IFormatoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormatoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormatoContext)
}

func (s *ElementoContext) Lista() IListaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *ElementoContext) Enlace() IEnlaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnlaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnlaceContext)
}

func (s *ElementoContext) Singleton() ISingletonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingletonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingletonContext)
}

func (s *ElementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterElemento(s)
	}
}

func (s *ElementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitElemento(s)
	}
}

func (p *L5Parser) Elemento() (localctx IElementoContext) {
	localctx = NewElementoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, L5ParserRULE_elemento)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case L5ParserT__10, L5ParserT__12, L5ParserT__14, L5ParserT__16, L5ParserT__18, L5ParserT__20, L5ParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Estilo()
		}

	case L5ParserT__24, L5ParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Formato()
		}

	case L5ParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Lista()
		}

	case L5ParserT__34:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Enlace()
		}

	case L5ParserT__40, L5ParserT__41:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.Singleton()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEstiloContext is an interface to support dynamic dispatch.
type IEstiloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Texto() ITextoContext

	// IsEstiloContext differentiates from other interfaces.
	IsEstiloContext()
}

type EstiloContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEstiloContext() *EstiloContext {
	var p = new(EstiloContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_estilo
	return p
}

func InitEmptyEstiloContext(p *EstiloContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_estilo
}

func (*EstiloContext) IsEstiloContext() {}

func NewEstiloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EstiloContext {
	var p = new(EstiloContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_estilo

	return p
}

func (s *EstiloContext) GetParser() antlr.Parser { return s.parser }

func (s *EstiloContext) Texto() ITextoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextoContext)
}

func (s *EstiloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EstiloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EstiloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterEstilo(s)
	}
}

func (s *EstiloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitEstilo(s)
	}
}

func (p *L5Parser) Estilo() (localctx IEstiloContext) {
	localctx = NewEstiloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, L5ParserRULE_estilo)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case L5ParserT__10:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(L5ParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Texto()
		}
		{
			p.SetState(64)
			p.Match(L5ParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(L5ParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Texto()
		}
		{
			p.SetState(68)
			p.Match(L5ParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserT__14:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Match(L5ParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Texto()
		}
		{
			p.SetState(72)
			p.Match(L5ParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserT__16:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			p.Match(L5ParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Texto()
		}
		{
			p.SetState(76)
			p.Match(L5ParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserT__18:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(78)
			p.Match(L5ParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Texto()
		}
		{
			p.SetState(80)
			p.Match(L5ParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserT__20:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(82)
			p.Match(L5ParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Texto()
		}
		{
			p.SetState(84)
			p.Match(L5ParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserT__22:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(86)
			p.Match(L5ParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Texto()
		}
		{
			p.SetState(88)
			p.Match(L5ParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormatoContext is an interface to support dynamic dispatch.
type IFormatoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Texto() ITextoContext
	Posicion() IPosicionContext

	// IsFormatoContext differentiates from other interfaces.
	IsFormatoContext()
}

type FormatoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormatoContext() *FormatoContext {
	var p = new(FormatoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_formato
	return p
}

func InitEmptyFormatoContext(p *FormatoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_formato
}

func (*FormatoContext) IsFormatoContext() {}

func NewFormatoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatoContext {
	var p = new(FormatoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_formato

	return p
}

func (s *FormatoContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatoContext) Texto() ITextoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextoContext)
}

func (s *FormatoContext) Posicion() IPosicionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPosicionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPosicionContext)
}

func (s *FormatoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterFormato(s)
	}
}

func (s *FormatoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitFormato(s)
	}
}

func (p *L5Parser) Formato() (localctx IFormatoContext) {
	localctx = NewFormatoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, L5ParserRULE_formato)
	var _la int

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case L5ParserT__24:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(L5ParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Texto()
		}
		{
			p.SetState(94)
			p.Match(L5ParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(L5ParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == L5ParserT__28 {
			{
				p.SetState(97)
				p.Posicion()
			}

		}
		{
			p.SetState(100)
			p.Texto()
		}
		{
			p.SetState(101)
			p.Match(L5ParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPosicionContext is an interface to support dynamic dispatch.
type IPosicionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POSICION() antlr.TerminalNode

	// IsPosicionContext differentiates from other interfaces.
	IsPosicionContext()
}

type PosicionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPosicionContext() *PosicionContext {
	var p = new(PosicionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_posicion
	return p
}

func InitEmptyPosicionContext(p *PosicionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_posicion
}

func (*PosicionContext) IsPosicionContext() {}

func NewPosicionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PosicionContext {
	var p = new(PosicionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_posicion

	return p
}

func (s *PosicionContext) GetParser() antlr.Parser { return s.parser }

func (s *PosicionContext) POSICION() antlr.TerminalNode {
	return s.GetToken(L5ParserPOSICION, 0)
}

func (s *PosicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PosicionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PosicionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterPosicion(s)
	}
}

func (s *PosicionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitPosicion(s)
	}
}

func (p *L5Parser) Posicion() (localctx IPosicionContext) {
	localctx = NewPosicionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, L5ParserRULE_posicion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(L5ParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(L5ParserPOSICION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(L5ParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaContext is an interface to support dynamic dispatch.
type IListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllElementoLista() []IElementoListaContext
	ElementoLista(i int) IElementoListaContext

	// IsListaContext differentiates from other interfaces.
	IsListaContext()
}

type ListaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaContext() *ListaContext {
	var p = new(ListaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_lista
	return p
}

func InitEmptyListaContext(p *ListaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_lista
}

func (*ListaContext) IsListaContext() {}

func NewListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaContext {
	var p = new(ListaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_lista

	return p
}

func (s *ListaContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaContext) AllElementoLista() []IElementoListaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementoListaContext); ok {
			len++
		}
	}

	tst := make([]IElementoListaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementoListaContext); ok {
			tst[i] = t.(IElementoListaContext)
			i++
		}
	}

	return tst
}

func (s *ListaContext) ElementoLista(i int) IElementoListaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementoListaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementoListaContext)
}

func (s *ListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterLista(s)
	}
}

func (s *ListaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitLista(s)
	}
}

func (p *L5Parser) Lista() (localctx IListaContext) {
	localctx = NewListaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, L5ParserRULE_lista)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(L5ParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == L5ParserT__32 {
		{
			p.SetState(110)
			p.ElementoLista()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(L5ParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementoListaContext is an interface to support dynamic dispatch.
type IElementoListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Texto() ITextoContext

	// IsElementoListaContext differentiates from other interfaces.
	IsElementoListaContext()
}

type ElementoListaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementoListaContext() *ElementoListaContext {
	var p = new(ElementoListaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_elementoLista
	return p
}

func InitEmptyElementoListaContext(p *ElementoListaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_elementoLista
}

func (*ElementoListaContext) IsElementoListaContext() {}

func NewElementoListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementoListaContext {
	var p = new(ElementoListaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_elementoLista

	return p
}

func (s *ElementoListaContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementoListaContext) Texto() ITextoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextoContext)
}

func (s *ElementoListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementoListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementoListaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterElementoLista(s)
	}
}

func (s *ElementoListaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitElementoLista(s)
	}
}

func (p *L5Parser) ElementoLista() (localctx IElementoListaContext) {
	localctx = NewElementoListaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, L5ParserRULE_elementoLista)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(L5ParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Texto()
	}
	{
		p.SetState(119)
		p.Match(L5ParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnlaceContext is an interface to support dynamic dispatch.
type IEnlaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTexto() []ITextoContext
	Texto(i int) ITextoContext

	// IsEnlaceContext differentiates from other interfaces.
	IsEnlaceContext()
}

type EnlaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnlaceContext() *EnlaceContext {
	var p = new(EnlaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_enlace
	return p
}

func InitEmptyEnlaceContext(p *EnlaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_enlace
}

func (*EnlaceContext) IsEnlaceContext() {}

func NewEnlaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnlaceContext {
	var p = new(EnlaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_enlace

	return p
}

func (s *EnlaceContext) GetParser() antlr.Parser { return s.parser }

func (s *EnlaceContext) AllTexto() []ITextoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITextoContext); ok {
			len++
		}
	}

	tst := make([]ITextoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITextoContext); ok {
			tst[i] = t.(ITextoContext)
			i++
		}
	}

	return tst
}

func (s *EnlaceContext) Texto(i int) ITextoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextoContext)
}

func (s *EnlaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnlaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnlaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterEnlace(s)
	}
}

func (s *EnlaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitEnlace(s)
	}
}

func (p *L5Parser) Enlace() (localctx IEnlaceContext) {
	localctx = NewEnlaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, L5ParserRULE_enlace)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(L5ParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(L5ParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Texto()
	}
	{
		p.SetState(124)
		p.Match(L5ParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(L5ParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Texto()
	}
	{
		p.SetState(127)
		p.Match(L5ParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(L5ParserT__39)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingletonContext is an interface to support dynamic dispatch.
type ISingletonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingletonContext differentiates from other interfaces.
	IsSingletonContext()
}

type SingletonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingletonContext() *SingletonContext {
	var p = new(SingletonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_singleton
	return p
}

func InitEmptySingletonContext(p *SingletonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_singleton
}

func (*SingletonContext) IsSingletonContext() {}

func NewSingletonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingletonContext {
	var p = new(SingletonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_singleton

	return p
}

func (s *SingletonContext) GetParser() antlr.Parser { return s.parser }
func (s *SingletonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingletonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingletonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterSingleton(s)
	}
}

func (s *SingletonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitSingleton(s)
	}
}

func (p *L5Parser) Singleton() (localctx ISingletonContext) {
	localctx = NewSingletonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, L5ParserRULE_singleton)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(_la == L5ParserT__40 || _la == L5ParserT__41) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITextoContext is an interface to support dynamic dispatch.
type ITextoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXTO() []antlr.TerminalNode
	TEXTO(i int) antlr.TerminalNode

	// IsTextoContext differentiates from other interfaces.
	IsTextoContext()
}

type TextoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextoContext() *TextoContext {
	var p = new(TextoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_texto
	return p
}

func InitEmptyTextoContext(p *TextoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_texto
}

func (*TextoContext) IsTextoContext() {}

func NewTextoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextoContext {
	var p = new(TextoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_texto

	return p
}

func (s *TextoContext) GetParser() antlr.Parser { return s.parser }

func (s *TextoContext) AllTEXTO() []antlr.TerminalNode {
	return s.GetTokens(L5ParserTEXTO)
}

func (s *TextoContext) TEXTO(i int) antlr.TerminalNode {
	return s.GetToken(L5ParserTEXTO, i)
}

func (s *TextoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterTexto(s)
	}
}

func (s *TextoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitTexto(s)
	}
}

func (p *L5Parser) Texto() (localctx ITextoContext) {
	localctx = NewTextoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, L5ParserRULE_texto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == L5ParserTEXTO {
		{
			p.SetState(132)
			p.Match(L5ParserTEXTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
