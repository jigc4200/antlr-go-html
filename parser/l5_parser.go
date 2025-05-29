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
		"", "'Pagina'", "'Cabecera'", "'Titulo'", "'FinTitulo'", "'FinCabecera'",
		"'Cuerpo'", "'FinCuerpo'", "'Ppagina'", "'FPpagina'", "'FinPagina'",
		"'Negrita'", "'FinNegrita'", "'Cursiva'", "'FinCursiva'", "'Subrayado'",
		"'FinSubrayado'", "'Tachado'", "'FinTachado'", "'Subindice'", "'FinSubindice'",
		"'Superindice'", "'FinSuperindice'", "'Tama\\u00F1o'", "'FinTama\\u00F1o'",
		"'Sangrado'", "'FinSangrado'", "'Parrafo'", "'FinParrafo'", "'Posicion'",
		"'FinPosicion'", "'Id'", "'FinId'", "'Clase'", "'FinClase'", "'Lista'",
		"'FinLista'", "'ElementoLista'", "'FinElementoLista'", "'Enlace'", "'FinEnlace'",
		"'Con'", "'FinCon'", "'Mostrar'", "'FinMostrar'", "'Imagen'", "'FinImagen'",
		"'Src'", "'FinSrc'", "'Alt'", "'FinAlt'", "'Encabezado'", "'FinEncabezado'",
		"'Linea'", "'Salto'", "'Estilo'", "'FinEstilo'",
	}
	staticData.SymbolicNames = []string{
		"", "PAGINA", "CABECERA", "TITULO", "FINTITULO", "FINCABECERA", "CUERPO",
		"FINCUERPO", "PPAGINA", "FPPAGINA", "FINPAGINA", "NEGRITA", "FINNEGRITA",
		"CURSIVA", "FINCURSIVA", "SUBRAYADO", "FINSUBRAYADO", "TACHADO", "FINTACHADO",
		"SUBINDICE", "FINSUBINDICE", "SUPERINDICE", "FINSUPERINDICE", "TAMANO",
		"FINTAMANO", "SANGRADO", "FINSANGRADO", "PARRAFO", "FINPARRAFO", "POSICION_KW",
		"FINPOSICION", "ID", "FINID", "CLASE", "FINCLASE", "LISTA", "FINLISTA",
		"ELEMENTOLISTA", "FINELEMENTOLISTA", "ENLACE", "FINENLACE", "CON", "FINCON",
		"MOSTRAR", "FINMOSTRAR", "IMAGEN", "FINIMAGEN", "SRC", "FINSRC", "ALT",
		"FINALT", "ENCABEZADO", "FINENCABEZADO", "LINEA", "SALTO", "ESTILO",
		"FINESTILO", "POSICION", "ATTR_VAL", "WS",
	}
	staticData.RuleNames = []string{
		"pagina", "cabecera", "titulo", "cuerpo", "pie", "elemento", "estilo",
		"formato", "atributo", "lista", "elementoLista", "enlace", "imagen",
		"encabezado", "singleton", "contenido",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 187, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 49, 8, 3, 10, 3, 12, 3, 52, 9, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 67, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 98, 8, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 3, 7, 106, 8, 7, 1, 7, 5, 7, 109, 8, 7, 10, 7, 12,
		7, 112, 9, 7, 1, 7, 1, 7, 1, 7, 3, 7, 117, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 128, 8, 8, 1, 9, 1, 9, 4, 9, 132, 8,
		9, 11, 9, 12, 9, 133, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 4, 10, 144, 8, 10, 11, 10, 12, 10, 145, 1, 10, 1, 10, 3, 10, 150, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 169, 8, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		5, 15, 182, 8, 15, 10, 15, 12, 15, 185, 9, 15, 1, 15, 0, 0, 16, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 1, 1, 0, 53, 54, 194,
		0, 32, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0, 6, 46, 1, 0, 0,
		0, 8, 55, 1, 0, 0, 0, 10, 66, 1, 0, 0, 0, 12, 97, 1, 0, 0, 0, 14, 116,
		1, 0, 0, 0, 16, 127, 1, 0, 0, 0, 18, 129, 1, 0, 0, 0, 20, 149, 1, 0, 0,
		0, 22, 151, 1, 0, 0, 0, 24, 160, 1, 0, 0, 0, 26, 172, 1, 0, 0, 0, 28, 177,
		1, 0, 0, 0, 30, 183, 1, 0, 0, 0, 32, 33, 5, 1, 0, 0, 33, 34, 3, 2, 1, 0,
		34, 35, 3, 6, 3, 0, 35, 36, 3, 8, 4, 0, 36, 37, 5, 10, 0, 0, 37, 1, 1,
		0, 0, 0, 38, 39, 5, 2, 0, 0, 39, 40, 3, 4, 2, 0, 40, 41, 5, 5, 0, 0, 41,
		3, 1, 0, 0, 0, 42, 43, 5, 3, 0, 0, 43, 44, 3, 30, 15, 0, 44, 45, 5, 4,
		0, 0, 45, 5, 1, 0, 0, 0, 46, 50, 5, 6, 0, 0, 47, 49, 3, 10, 5, 0, 48, 47,
		1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0,
		51, 53, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 7, 0, 0, 54, 7, 1, 0,
		0, 0, 55, 56, 5, 8, 0, 0, 56, 57, 3, 30, 15, 0, 57, 58, 5, 9, 0, 0, 58,
		9, 1, 0, 0, 0, 59, 67, 3, 12, 6, 0, 60, 67, 3, 14, 7, 0, 61, 67, 3, 18,
		9, 0, 62, 67, 3, 22, 11, 0, 63, 67, 3, 28, 14, 0, 64, 67, 3, 24, 12, 0,
		65, 67, 3, 26, 13, 0, 66, 59, 1, 0, 0, 0, 66, 60, 1, 0, 0, 0, 66, 61, 1,
		0, 0, 0, 66, 62, 1, 0, 0, 0, 66, 63, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66,
		65, 1, 0, 0, 0, 67, 11, 1, 0, 0, 0, 68, 69, 5, 11, 0, 0, 69, 70, 3, 30,
		15, 0, 70, 71, 5, 12, 0, 0, 71, 98, 1, 0, 0, 0, 72, 73, 5, 13, 0, 0, 73,
		74, 3, 30, 15, 0, 74, 75, 5, 14, 0, 0, 75, 98, 1, 0, 0, 0, 76, 77, 5, 15,
		0, 0, 77, 78, 3, 30, 15, 0, 78, 79, 5, 16, 0, 0, 79, 98, 1, 0, 0, 0, 80,
		81, 5, 17, 0, 0, 81, 82, 3, 30, 15, 0, 82, 83, 5, 18, 0, 0, 83, 98, 1,
		0, 0, 0, 84, 85, 5, 19, 0, 0, 85, 86, 3, 30, 15, 0, 86, 87, 5, 20, 0, 0,
		87, 98, 1, 0, 0, 0, 88, 89, 5, 21, 0, 0, 89, 90, 3, 30, 15, 0, 90, 91,
		5, 22, 0, 0, 91, 98, 1, 0, 0, 0, 92, 93, 5, 23, 0, 0, 93, 94, 5, 58, 0,
		0, 94, 95, 3, 30, 15, 0, 95, 96, 5, 24, 0, 0, 96, 98, 1, 0, 0, 0, 97, 68,
		1, 0, 0, 0, 97, 72, 1, 0, 0, 0, 97, 76, 1, 0, 0, 0, 97, 80, 1, 0, 0, 0,
		97, 84, 1, 0, 0, 0, 97, 88, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 98, 13, 1,
		0, 0, 0, 99, 100, 5, 25, 0, 0, 100, 101, 3, 30, 15, 0, 101, 102, 5, 26,
		0, 0, 102, 117, 1, 0, 0, 0, 103, 105, 5, 27, 0, 0, 104, 106, 3, 12, 6,
		0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 110, 1, 0, 0, 0, 107,
		109, 3, 16, 8, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108,
		1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 110, 1, 0,
		0, 0, 113, 114, 3, 30, 15, 0, 114, 115, 5, 28, 0, 0, 115, 117, 1, 0, 0,
		0, 116, 99, 1, 0, 0, 0, 116, 103, 1, 0, 0, 0, 117, 15, 1, 0, 0, 0, 118,
		119, 5, 29, 0, 0, 119, 120, 5, 57, 0, 0, 120, 128, 5, 30, 0, 0, 121, 122,
		5, 31, 0, 0, 122, 123, 5, 58, 0, 0, 123, 128, 5, 32, 0, 0, 124, 125, 5,
		33, 0, 0, 125, 126, 5, 58, 0, 0, 126, 128, 5, 34, 0, 0, 127, 118, 1, 0,
		0, 0, 127, 121, 1, 0, 0, 0, 127, 124, 1, 0, 0, 0, 128, 17, 1, 0, 0, 0,
		129, 131, 5, 35, 0, 0, 130, 132, 3, 20, 10, 0, 131, 130, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135,
		1, 0, 0, 0, 135, 136, 5, 36, 0, 0, 136, 19, 1, 0, 0, 0, 137, 138, 5, 37,
		0, 0, 138, 139, 3, 30, 15, 0, 139, 140, 5, 38, 0, 0, 140, 150, 1, 0, 0,
		0, 141, 143, 5, 37, 0, 0, 142, 144, 3, 10, 5, 0, 143, 142, 1, 0, 0, 0,
		144, 145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		147, 1, 0, 0, 0, 147, 148, 5, 38, 0, 0, 148, 150, 1, 0, 0, 0, 149, 137,
		1, 0, 0, 0, 149, 141, 1, 0, 0, 0, 150, 21, 1, 0, 0, 0, 151, 152, 5, 39,
		0, 0, 152, 153, 5, 41, 0, 0, 153, 154, 5, 58, 0, 0, 154, 155, 5, 42, 0,
		0, 155, 156, 5, 43, 0, 0, 156, 157, 3, 30, 15, 0, 157, 158, 5, 44, 0, 0,
		158, 159, 5, 40, 0, 0, 159, 23, 1, 0, 0, 0, 160, 161, 5, 45, 0, 0, 161,
		162, 5, 47, 0, 0, 162, 163, 5, 58, 0, 0, 163, 168, 5, 48, 0, 0, 164, 165,
		5, 49, 0, 0, 165, 166, 3, 30, 15, 0, 166, 167, 5, 50, 0, 0, 167, 169, 1,
		0, 0, 0, 168, 164, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0,
		0, 170, 171, 5, 46, 0, 0, 171, 25, 1, 0, 0, 0, 172, 173, 5, 51, 0, 0, 173,
		174, 5, 58, 0, 0, 174, 175, 3, 30, 15, 0, 175, 176, 5, 52, 0, 0, 176, 27,
		1, 0, 0, 0, 177, 178, 7, 0, 0, 0, 178, 29, 1, 0, 0, 0, 179, 182, 5, 58,
		0, 0, 180, 182, 3, 12, 6, 0, 181, 179, 1, 0, 0, 0, 181, 180, 1, 0, 0, 0,
		182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		31, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 13, 50, 66, 97, 105, 110, 116, 127,
		133, 145, 149, 168, 181, 183,
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
	L5ParserEOF              = antlr.TokenEOF
	L5ParserPAGINA           = 1
	L5ParserCABECERA         = 2
	L5ParserTITULO           = 3
	L5ParserFINTITULO        = 4
	L5ParserFINCABECERA      = 5
	L5ParserCUERPO           = 6
	L5ParserFINCUERPO        = 7
	L5ParserPPAGINA          = 8
	L5ParserFPPAGINA         = 9
	L5ParserFINPAGINA        = 10
	L5ParserNEGRITA          = 11
	L5ParserFINNEGRITA       = 12
	L5ParserCURSIVA          = 13
	L5ParserFINCURSIVA       = 14
	L5ParserSUBRAYADO        = 15
	L5ParserFINSUBRAYADO     = 16
	L5ParserTACHADO          = 17
	L5ParserFINTACHADO       = 18
	L5ParserSUBINDICE        = 19
	L5ParserFINSUBINDICE     = 20
	L5ParserSUPERINDICE      = 21
	L5ParserFINSUPERINDICE   = 22
	L5ParserTAMANO           = 23
	L5ParserFINTAMANO        = 24
	L5ParserSANGRADO         = 25
	L5ParserFINSANGRADO      = 26
	L5ParserPARRAFO          = 27
	L5ParserFINPARRAFO       = 28
	L5ParserPOSICION_KW      = 29
	L5ParserFINPOSICION      = 30
	L5ParserID               = 31
	L5ParserFINID            = 32
	L5ParserCLASE            = 33
	L5ParserFINCLASE         = 34
	L5ParserLISTA            = 35
	L5ParserFINLISTA         = 36
	L5ParserELEMENTOLISTA    = 37
	L5ParserFINELEMENTOLISTA = 38
	L5ParserENLACE           = 39
	L5ParserFINENLACE        = 40
	L5ParserCON              = 41
	L5ParserFINCON           = 42
	L5ParserMOSTRAR          = 43
	L5ParserFINMOSTRAR       = 44
	L5ParserIMAGEN           = 45
	L5ParserFINIMAGEN        = 46
	L5ParserSRC              = 47
	L5ParserFINSRC           = 48
	L5ParserALT              = 49
	L5ParserFINALT           = 50
	L5ParserENCABEZADO       = 51
	L5ParserFINENCABEZADO    = 52
	L5ParserLINEA            = 53
	L5ParserSALTO            = 54
	L5ParserESTILO           = 55
	L5ParserFINESTILO        = 56
	L5ParserPOSICION         = 57
	L5ParserATTR_VAL         = 58
	L5ParserWS               = 59
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
	L5ParserRULE_atributo      = 8
	L5ParserRULE_lista         = 9
	L5ParserRULE_elementoLista = 10
	L5ParserRULE_enlace        = 11
	L5ParserRULE_imagen        = 12
	L5ParserRULE_encabezado    = 13
	L5ParserRULE_singleton     = 14
	L5ParserRULE_contenido     = 15
)

// IPaginaContext is an interface to support dynamic dispatch.
type IPaginaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAGINA() antlr.TerminalNode
	Cabecera() ICabeceraContext
	Cuerpo() ICuerpoContext
	Pie() IPieContext
	FINPAGINA() antlr.TerminalNode

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

func (s *PaginaContext) PAGINA() antlr.TerminalNode {
	return s.GetToken(L5ParserPAGINA, 0)
}

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

func (s *PaginaContext) FINPAGINA() antlr.TerminalNode {
	return s.GetToken(L5ParserFINPAGINA, 0)
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
		p.SetState(32)
		p.Match(L5ParserPAGINA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Cabecera()
	}
	{
		p.SetState(34)
		p.Cuerpo()
	}
	{
		p.SetState(35)
		p.Pie()
	}
	{
		p.SetState(36)
		p.Match(L5ParserFINPAGINA)
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
	CABECERA() antlr.TerminalNode
	Titulo() ITituloContext
	FINCABECERA() antlr.TerminalNode

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

func (s *CabeceraContext) CABECERA() antlr.TerminalNode {
	return s.GetToken(L5ParserCABECERA, 0)
}

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

func (s *CabeceraContext) FINCABECERA() antlr.TerminalNode {
	return s.GetToken(L5ParserFINCABECERA, 0)
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
		p.SetState(38)
		p.Match(L5ParserCABECERA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Titulo()
	}
	{
		p.SetState(40)
		p.Match(L5ParserFINCABECERA)
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
	TITULO() antlr.TerminalNode
	Contenido() IContenidoContext
	FINTITULO() antlr.TerminalNode

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

func (s *TituloContext) TITULO() antlr.TerminalNode {
	return s.GetToken(L5ParserTITULO, 0)
}

func (s *TituloContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *TituloContext) FINTITULO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINTITULO, 0)
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
		p.SetState(42)
		p.Match(L5ParserTITULO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Contenido()
	}
	{
		p.SetState(44)
		p.Match(L5ParserFINTITULO)
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
	CUERPO() antlr.TerminalNode
	FINCUERPO() antlr.TerminalNode
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

func (s *CuerpoContext) CUERPO() antlr.TerminalNode {
	return s.GetToken(L5ParserCUERPO, 0)
}

func (s *CuerpoContext) FINCUERPO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINCUERPO, 0)
}

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
		p.SetState(46)
		p.Match(L5ParserCUERPO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29309166244505600) != 0 {
		{
			p.SetState(47)
			p.Elemento()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(L5ParserFINCUERPO)
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
	PPAGINA() antlr.TerminalNode
	Contenido() IContenidoContext
	FPPAGINA() antlr.TerminalNode

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

func (s *PieContext) PPAGINA() antlr.TerminalNode {
	return s.GetToken(L5ParserPPAGINA, 0)
}

func (s *PieContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *PieContext) FPPAGINA() antlr.TerminalNode {
	return s.GetToken(L5ParserFPPAGINA, 0)
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
		p.SetState(55)
		p.Match(L5ParserPPAGINA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Contenido()
	}
	{
		p.SetState(57)
		p.Match(L5ParserFPPAGINA)
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
	Imagen() IImagenContext
	Encabezado() IEncabezadoContext

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

func (s *ElementoContext) Imagen() IImagenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImagenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImagenContext)
}

func (s *ElementoContext) Encabezado() IEncabezadoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEncabezadoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEncabezadoContext)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case L5ParserNEGRITA, L5ParserCURSIVA, L5ParserSUBRAYADO, L5ParserTACHADO, L5ParserSUBINDICE, L5ParserSUPERINDICE, L5ParserTAMANO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Estilo()
		}

	case L5ParserSANGRADO, L5ParserPARRAFO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Formato()
		}

	case L5ParserLISTA:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Lista()
		}

	case L5ParserENLACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.Enlace()
		}

	case L5ParserLINEA, L5ParserSALTO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Singleton()
		}

	case L5ParserIMAGEN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.Imagen()
		}

	case L5ParserENCABEZADO:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(65)
			p.Encabezado()
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
	NEGRITA() antlr.TerminalNode
	Contenido() IContenidoContext
	FINNEGRITA() antlr.TerminalNode
	CURSIVA() antlr.TerminalNode
	FINCURSIVA() antlr.TerminalNode
	SUBRAYADO() antlr.TerminalNode
	FINSUBRAYADO() antlr.TerminalNode
	TACHADO() antlr.TerminalNode
	FINTACHADO() antlr.TerminalNode
	SUBINDICE() antlr.TerminalNode
	FINSUBINDICE() antlr.TerminalNode
	SUPERINDICE() antlr.TerminalNode
	FINSUPERINDICE() antlr.TerminalNode
	TAMANO() antlr.TerminalNode
	ATTR_VAL() antlr.TerminalNode
	FINTAMANO() antlr.TerminalNode

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

func (s *EstiloContext) NEGRITA() antlr.TerminalNode {
	return s.GetToken(L5ParserNEGRITA, 0)
}

func (s *EstiloContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *EstiloContext) FINNEGRITA() antlr.TerminalNode {
	return s.GetToken(L5ParserFINNEGRITA, 0)
}

func (s *EstiloContext) CURSIVA() antlr.TerminalNode {
	return s.GetToken(L5ParserCURSIVA, 0)
}

func (s *EstiloContext) FINCURSIVA() antlr.TerminalNode {
	return s.GetToken(L5ParserFINCURSIVA, 0)
}

func (s *EstiloContext) SUBRAYADO() antlr.TerminalNode {
	return s.GetToken(L5ParserSUBRAYADO, 0)
}

func (s *EstiloContext) FINSUBRAYADO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINSUBRAYADO, 0)
}

func (s *EstiloContext) TACHADO() antlr.TerminalNode {
	return s.GetToken(L5ParserTACHADO, 0)
}

func (s *EstiloContext) FINTACHADO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINTACHADO, 0)
}

func (s *EstiloContext) SUBINDICE() antlr.TerminalNode {
	return s.GetToken(L5ParserSUBINDICE, 0)
}

func (s *EstiloContext) FINSUBINDICE() antlr.TerminalNode {
	return s.GetToken(L5ParserFINSUBINDICE, 0)
}

func (s *EstiloContext) SUPERINDICE() antlr.TerminalNode {
	return s.GetToken(L5ParserSUPERINDICE, 0)
}

func (s *EstiloContext) FINSUPERINDICE() antlr.TerminalNode {
	return s.GetToken(L5ParserFINSUPERINDICE, 0)
}

func (s *EstiloContext) TAMANO() antlr.TerminalNode {
	return s.GetToken(L5ParserTAMANO, 0)
}

func (s *EstiloContext) ATTR_VAL() antlr.TerminalNode {
	return s.GetToken(L5ParserATTR_VAL, 0)
}

func (s *EstiloContext) FINTAMANO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINTAMANO, 0)
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
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case L5ParserNEGRITA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(L5ParserNEGRITA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Contenido()
		}
		{
			p.SetState(70)
			p.Match(L5ParserFINNEGRITA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserCURSIVA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(L5ParserCURSIVA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Contenido()
		}
		{
			p.SetState(74)
			p.Match(L5ParserFINCURSIVA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserSUBRAYADO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.Match(L5ParserSUBRAYADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Contenido()
		}
		{
			p.SetState(78)
			p.Match(L5ParserFINSUBRAYADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserTACHADO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)
			p.Match(L5ParserTACHADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Contenido()
		}
		{
			p.SetState(82)
			p.Match(L5ParserFINTACHADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserSUBINDICE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.Match(L5ParserSUBINDICE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Contenido()
		}
		{
			p.SetState(86)
			p.Match(L5ParserFINSUBINDICE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserSUPERINDICE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)
			p.Match(L5ParserSUPERINDICE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Contenido()
		}
		{
			p.SetState(90)
			p.Match(L5ParserFINSUPERINDICE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserTAMANO:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.Match(L5ParserTAMANO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(L5ParserATTR_VAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Contenido()
		}
		{
			p.SetState(95)
			p.Match(L5ParserFINTAMANO)
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
	SANGRADO() antlr.TerminalNode
	Contenido() IContenidoContext
	FINSANGRADO() antlr.TerminalNode
	PARRAFO() antlr.TerminalNode
	FINPARRAFO() antlr.TerminalNode
	Estilo() IEstiloContext
	AllAtributo() []IAtributoContext
	Atributo(i int) IAtributoContext

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

func (s *FormatoContext) SANGRADO() antlr.TerminalNode {
	return s.GetToken(L5ParserSANGRADO, 0)
}

func (s *FormatoContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *FormatoContext) FINSANGRADO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINSANGRADO, 0)
}

func (s *FormatoContext) PARRAFO() antlr.TerminalNode {
	return s.GetToken(L5ParserPARRAFO, 0)
}

func (s *FormatoContext) FINPARRAFO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINPARRAFO, 0)
}

func (s *FormatoContext) Estilo() IEstiloContext {
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

func (s *FormatoContext) AllAtributo() []IAtributoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtributoContext); ok {
			len++
		}
	}

	tst := make([]IAtributoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtributoContext); ok {
			tst[i] = t.(IAtributoContext)
			i++
		}
	}

	return tst
}

func (s *FormatoContext) Atributo(i int) IAtributoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributoContext); ok {
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

	return t.(IAtributoContext)
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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case L5ParserSANGRADO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(L5ParserSANGRADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Contenido()
		}
		{
			p.SetState(101)
			p.Match(L5ParserFINSANGRADO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserPARRAFO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(L5ParserPARRAFO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(104)
				p.Estilo()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&11274289152) != 0 {
			{
				p.SetState(107)
				p.Atributo()
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(113)
			p.Contenido()
		}
		{
			p.SetState(114)
			p.Match(L5ParserFINPARRAFO)
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

// IAtributoContext is an interface to support dynamic dispatch.
type IAtributoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POSICION_KW() antlr.TerminalNode
	POSICION() antlr.TerminalNode
	FINPOSICION() antlr.TerminalNode
	ID() antlr.TerminalNode
	ATTR_VAL() antlr.TerminalNode
	FINID() antlr.TerminalNode
	CLASE() antlr.TerminalNode
	FINCLASE() antlr.TerminalNode

	// IsAtributoContext differentiates from other interfaces.
	IsAtributoContext()
}

type AtributoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributoContext() *AtributoContext {
	var p = new(AtributoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_atributo
	return p
}

func InitEmptyAtributoContext(p *AtributoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_atributo
}

func (*AtributoContext) IsAtributoContext() {}

func NewAtributoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributoContext {
	var p = new(AtributoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_atributo

	return p
}

func (s *AtributoContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributoContext) POSICION_KW() antlr.TerminalNode {
	return s.GetToken(L5ParserPOSICION_KW, 0)
}

func (s *AtributoContext) POSICION() antlr.TerminalNode {
	return s.GetToken(L5ParserPOSICION, 0)
}

func (s *AtributoContext) FINPOSICION() antlr.TerminalNode {
	return s.GetToken(L5ParserFINPOSICION, 0)
}

func (s *AtributoContext) ID() antlr.TerminalNode {
	return s.GetToken(L5ParserID, 0)
}

func (s *AtributoContext) ATTR_VAL() antlr.TerminalNode {
	return s.GetToken(L5ParserATTR_VAL, 0)
}

func (s *AtributoContext) FINID() antlr.TerminalNode {
	return s.GetToken(L5ParserFINID, 0)
}

func (s *AtributoContext) CLASE() antlr.TerminalNode {
	return s.GetToken(L5ParserCLASE, 0)
}

func (s *AtributoContext) FINCLASE() antlr.TerminalNode {
	return s.GetToken(L5ParserFINCLASE, 0)
}

func (s *AtributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtributoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterAtributo(s)
	}
}

func (s *AtributoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitAtributo(s)
	}
}

func (p *L5Parser) Atributo() (localctx IAtributoContext) {
	localctx = NewAtributoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, L5ParserRULE_atributo)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case L5ParserPOSICION_KW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Match(L5ParserPOSICION_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(L5ParserPOSICION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(L5ParserFINPOSICION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Match(L5ParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(L5ParserATTR_VAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(L5ParserFINID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case L5ParserCLASE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(L5ParserCLASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(L5ParserATTR_VAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(L5ParserFINCLASE)
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

// IListaContext is an interface to support dynamic dispatch.
type IListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LISTA() antlr.TerminalNode
	FINLISTA() antlr.TerminalNode
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

func (s *ListaContext) LISTA() antlr.TerminalNode {
	return s.GetToken(L5ParserLISTA, 0)
}

func (s *ListaContext) FINLISTA() antlr.TerminalNode {
	return s.GetToken(L5ParserFINLISTA, 0)
}

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
		p.SetState(129)
		p.Match(L5ParserLISTA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == L5ParserELEMENTOLISTA {
		{
			p.SetState(130)
			p.ElementoLista()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
		p.Match(L5ParserFINLISTA)
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
	ELEMENTOLISTA() antlr.TerminalNode
	Contenido() IContenidoContext
	FINELEMENTOLISTA() antlr.TerminalNode
	AllElemento() []IElementoContext
	Elemento(i int) IElementoContext

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

func (s *ElementoListaContext) ELEMENTOLISTA() antlr.TerminalNode {
	return s.GetToken(L5ParserELEMENTOLISTA, 0)
}

func (s *ElementoListaContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *ElementoListaContext) FINELEMENTOLISTA() antlr.TerminalNode {
	return s.GetToken(L5ParserFINELEMENTOLISTA, 0)
}

func (s *ElementoListaContext) AllElemento() []IElementoContext {
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

func (s *ElementoListaContext) Elemento(i int) IElementoContext {
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
	var _la int

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(L5ParserELEMENTOLISTA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Contenido()
		}
		{
			p.SetState(139)
			p.Match(L5ParserFINELEMENTOLISTA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(L5ParserELEMENTOLISTA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29309166244505600) != 0) {
			{
				p.SetState(142)
				p.Elemento()
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(147)
			p.Match(L5ParserFINELEMENTOLISTA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IEnlaceContext is an interface to support dynamic dispatch.
type IEnlaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENLACE() antlr.TerminalNode
	CON() antlr.TerminalNode
	ATTR_VAL() antlr.TerminalNode
	FINCON() antlr.TerminalNode
	MOSTRAR() antlr.TerminalNode
	Contenido() IContenidoContext
	FINMOSTRAR() antlr.TerminalNode
	FINENLACE() antlr.TerminalNode

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

func (s *EnlaceContext) ENLACE() antlr.TerminalNode {
	return s.GetToken(L5ParserENLACE, 0)
}

func (s *EnlaceContext) CON() antlr.TerminalNode {
	return s.GetToken(L5ParserCON, 0)
}

func (s *EnlaceContext) ATTR_VAL() antlr.TerminalNode {
	return s.GetToken(L5ParserATTR_VAL, 0)
}

func (s *EnlaceContext) FINCON() antlr.TerminalNode {
	return s.GetToken(L5ParserFINCON, 0)
}

func (s *EnlaceContext) MOSTRAR() antlr.TerminalNode {
	return s.GetToken(L5ParserMOSTRAR, 0)
}

func (s *EnlaceContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *EnlaceContext) FINMOSTRAR() antlr.TerminalNode {
	return s.GetToken(L5ParserFINMOSTRAR, 0)
}

func (s *EnlaceContext) FINENLACE() antlr.TerminalNode {
	return s.GetToken(L5ParserFINENLACE, 0)
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
		p.SetState(151)
		p.Match(L5ParserENLACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(L5ParserCON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(L5ParserATTR_VAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(L5ParserFINCON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(L5ParserMOSTRAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Contenido()
	}
	{
		p.SetState(157)
		p.Match(L5ParserFINMOSTRAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(L5ParserFINENLACE)
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

// IImagenContext is an interface to support dynamic dispatch.
type IImagenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMAGEN() antlr.TerminalNode
	SRC() antlr.TerminalNode
	ATTR_VAL() antlr.TerminalNode
	FINSRC() antlr.TerminalNode
	FINIMAGEN() antlr.TerminalNode
	ALT() antlr.TerminalNode
	Contenido() IContenidoContext
	FINALT() antlr.TerminalNode

	// IsImagenContext differentiates from other interfaces.
	IsImagenContext()
}

type ImagenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImagenContext() *ImagenContext {
	var p = new(ImagenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_imagen
	return p
}

func InitEmptyImagenContext(p *ImagenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_imagen
}

func (*ImagenContext) IsImagenContext() {}

func NewImagenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImagenContext {
	var p = new(ImagenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_imagen

	return p
}

func (s *ImagenContext) GetParser() antlr.Parser { return s.parser }

func (s *ImagenContext) IMAGEN() antlr.TerminalNode {
	return s.GetToken(L5ParserIMAGEN, 0)
}

func (s *ImagenContext) SRC() antlr.TerminalNode {
	return s.GetToken(L5ParserSRC, 0)
}

func (s *ImagenContext) ATTR_VAL() antlr.TerminalNode {
	return s.GetToken(L5ParserATTR_VAL, 0)
}

func (s *ImagenContext) FINSRC() antlr.TerminalNode {
	return s.GetToken(L5ParserFINSRC, 0)
}

func (s *ImagenContext) FINIMAGEN() antlr.TerminalNode {
	return s.GetToken(L5ParserFINIMAGEN, 0)
}

func (s *ImagenContext) ALT() antlr.TerminalNode {
	return s.GetToken(L5ParserALT, 0)
}

func (s *ImagenContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *ImagenContext) FINALT() antlr.TerminalNode {
	return s.GetToken(L5ParserFINALT, 0)
}

func (s *ImagenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImagenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImagenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterImagen(s)
	}
}

func (s *ImagenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitImagen(s)
	}
}

func (p *L5Parser) Imagen() (localctx IImagenContext) {
	localctx = NewImagenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, L5ParserRULE_imagen)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(L5ParserIMAGEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(L5ParserSRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(L5ParserATTR_VAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(L5ParserFINSRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == L5ParserALT {
		{
			p.SetState(164)
			p.Match(L5ParserALT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Contenido()
		}
		{
			p.SetState(166)
			p.Match(L5ParserFINALT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(170)
		p.Match(L5ParserFINIMAGEN)
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

// IEncabezadoContext is an interface to support dynamic dispatch.
type IEncabezadoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENCABEZADO() antlr.TerminalNode
	ATTR_VAL() antlr.TerminalNode
	Contenido() IContenidoContext
	FINENCABEZADO() antlr.TerminalNode

	// IsEncabezadoContext differentiates from other interfaces.
	IsEncabezadoContext()
}

type EncabezadoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncabezadoContext() *EncabezadoContext {
	var p = new(EncabezadoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_encabezado
	return p
}

func InitEmptyEncabezadoContext(p *EncabezadoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_encabezado
}

func (*EncabezadoContext) IsEncabezadoContext() {}

func NewEncabezadoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncabezadoContext {
	var p = new(EncabezadoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_encabezado

	return p
}

func (s *EncabezadoContext) GetParser() antlr.Parser { return s.parser }

func (s *EncabezadoContext) ENCABEZADO() antlr.TerminalNode {
	return s.GetToken(L5ParserENCABEZADO, 0)
}

func (s *EncabezadoContext) ATTR_VAL() antlr.TerminalNode {
	return s.GetToken(L5ParserATTR_VAL, 0)
}

func (s *EncabezadoContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *EncabezadoContext) FINENCABEZADO() antlr.TerminalNode {
	return s.GetToken(L5ParserFINENCABEZADO, 0)
}

func (s *EncabezadoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncabezadoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncabezadoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterEncabezado(s)
	}
}

func (s *EncabezadoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitEncabezado(s)
	}
}

func (p *L5Parser) Encabezado() (localctx IEncabezadoContext) {
	localctx = NewEncabezadoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, L5ParserRULE_encabezado)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(L5ParserENCABEZADO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(L5ParserATTR_VAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Contenido()
	}
	{
		p.SetState(175)
		p.Match(L5ParserFINENCABEZADO)
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

	// Getter signatures
	LINEA() antlr.TerminalNode
	SALTO() antlr.TerminalNode

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

func (s *SingletonContext) LINEA() antlr.TerminalNode {
	return s.GetToken(L5ParserLINEA, 0)
}

func (s *SingletonContext) SALTO() antlr.TerminalNode {
	return s.GetToken(L5ParserSALTO, 0)
}

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
	p.EnterRule(localctx, 28, L5ParserRULE_singleton)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		_la = p.GetTokenStream().LA(1)

		if !(_la == L5ParserLINEA || _la == L5ParserSALTO) {
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

// IContenidoContext is an interface to support dynamic dispatch.
type IContenidoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllATTR_VAL() []antlr.TerminalNode
	ATTR_VAL(i int) antlr.TerminalNode
	AllEstilo() []IEstiloContext
	Estilo(i int) IEstiloContext

	// IsContenidoContext differentiates from other interfaces.
	IsContenidoContext()
}

type ContenidoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContenidoContext() *ContenidoContext {
	var p = new(ContenidoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_contenido
	return p
}

func InitEmptyContenidoContext(p *ContenidoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = L5ParserRULE_contenido
}

func (*ContenidoContext) IsContenidoContext() {}

func NewContenidoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContenidoContext {
	var p = new(ContenidoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = L5ParserRULE_contenido

	return p
}

func (s *ContenidoContext) GetParser() antlr.Parser { return s.parser }

func (s *ContenidoContext) AllATTR_VAL() []antlr.TerminalNode {
	return s.GetTokens(L5ParserATTR_VAL)
}

func (s *ContenidoContext) ATTR_VAL(i int) antlr.TerminalNode {
	return s.GetToken(L5ParserATTR_VAL, i)
}

func (s *ContenidoContext) AllEstilo() []IEstiloContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEstiloContext); ok {
			len++
		}
	}

	tst := make([]IEstiloContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEstiloContext); ok {
			tst[i] = t.(IEstiloContext)
			i++
		}
	}

	return tst
}

func (s *ContenidoContext) Estilo(i int) IEstiloContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstiloContext); ok {
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

	return t.(IEstiloContext)
}

func (s *ContenidoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContenidoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContenidoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.EnterContenido(s)
	}
}

func (s *ContenidoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(L5Listener); ok {
		listenerT.ExitContenido(s)
	}
}

func (p *L5Parser) Contenido() (localctx IContenidoContext) {
	localctx = NewContenidoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, L5ParserRULE_contenido)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&288230376162895872) != 0 {
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case L5ParserATTR_VAL:
			{
				p.SetState(179)
				p.Match(L5ParserATTR_VAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case L5ParserNEGRITA, L5ParserCURSIVA, L5ParserSUBRAYADO, L5ParserTACHADO, L5ParserSUBINDICE, L5ParserSUPERINDICE, L5ParserTAMANO:
			{
				p.SetState(180)
				p.Estilo()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(185)
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
