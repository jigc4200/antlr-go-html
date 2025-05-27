// Code generated from L5.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // L5

import "github.com/antlr4-go/antlr/v4"

// BaseL5Listener is a complete listener for a parse tree produced by L5Parser.
type BaseL5Listener struct{}

var _ L5Listener = &BaseL5Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseL5Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseL5Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseL5Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseL5Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPagina is called when production pagina is entered.
func (s *BaseL5Listener) EnterPagina(ctx *PaginaContext) {}

// ExitPagina is called when production pagina is exited.
func (s *BaseL5Listener) ExitPagina(ctx *PaginaContext) {}

// EnterCabecera is called when production cabecera is entered.
func (s *BaseL5Listener) EnterCabecera(ctx *CabeceraContext) {}

// ExitCabecera is called when production cabecera is exited.
func (s *BaseL5Listener) ExitCabecera(ctx *CabeceraContext) {}

// EnterTitulo is called when production titulo is entered.
func (s *BaseL5Listener) EnterTitulo(ctx *TituloContext) {}

// ExitTitulo is called when production titulo is exited.
func (s *BaseL5Listener) ExitTitulo(ctx *TituloContext) {}

// EnterCuerpo is called when production cuerpo is entered.
func (s *BaseL5Listener) EnterCuerpo(ctx *CuerpoContext) {}

// ExitCuerpo is called when production cuerpo is exited.
func (s *BaseL5Listener) ExitCuerpo(ctx *CuerpoContext) {}

// EnterPie is called when production pie is entered.
func (s *BaseL5Listener) EnterPie(ctx *PieContext) {}

// ExitPie is called when production pie is exited.
func (s *BaseL5Listener) ExitPie(ctx *PieContext) {}

// EnterElemento is called when production elemento is entered.
func (s *BaseL5Listener) EnterElemento(ctx *ElementoContext) {}

// ExitElemento is called when production elemento is exited.
func (s *BaseL5Listener) ExitElemento(ctx *ElementoContext) {}

// EnterEstilo is called when production estilo is entered.
func (s *BaseL5Listener) EnterEstilo(ctx *EstiloContext) {}

// ExitEstilo is called when production estilo is exited.
func (s *BaseL5Listener) ExitEstilo(ctx *EstiloContext) {}

// EnterFormato is called when production formato is entered.
func (s *BaseL5Listener) EnterFormato(ctx *FormatoContext) {}

// ExitFormato is called when production formato is exited.
func (s *BaseL5Listener) ExitFormato(ctx *FormatoContext) {}

// EnterAtributo is called when production atributo is entered.
func (s *BaseL5Listener) EnterAtributo(ctx *AtributoContext) {}

// ExitAtributo is called when production atributo is exited.
func (s *BaseL5Listener) ExitAtributo(ctx *AtributoContext) {}

// EnterLista is called when production lista is entered.
func (s *BaseL5Listener) EnterLista(ctx *ListaContext) {}

// ExitLista is called when production lista is exited.
func (s *BaseL5Listener) ExitLista(ctx *ListaContext) {}

// EnterElementoLista is called when production elementoLista is entered.
func (s *BaseL5Listener) EnterElementoLista(ctx *ElementoListaContext) {}

// ExitElementoLista is called when production elementoLista is exited.
func (s *BaseL5Listener) ExitElementoLista(ctx *ElementoListaContext) {}

// EnterEnlace is called when production enlace is entered.
func (s *BaseL5Listener) EnterEnlace(ctx *EnlaceContext) {}

// ExitEnlace is called when production enlace is exited.
func (s *BaseL5Listener) ExitEnlace(ctx *EnlaceContext) {}

// EnterImagen is called when production imagen is entered.
func (s *BaseL5Listener) EnterImagen(ctx *ImagenContext) {}

// ExitImagen is called when production imagen is exited.
func (s *BaseL5Listener) ExitImagen(ctx *ImagenContext) {}

// EnterEncabezado is called when production encabezado is entered.
func (s *BaseL5Listener) EnterEncabezado(ctx *EncabezadoContext) {}

// ExitEncabezado is called when production encabezado is exited.
func (s *BaseL5Listener) ExitEncabezado(ctx *EncabezadoContext) {}

// EnterSingleton is called when production singleton is entered.
func (s *BaseL5Listener) EnterSingleton(ctx *SingletonContext) {}

// ExitSingleton is called when production singleton is exited.
func (s *BaseL5Listener) ExitSingleton(ctx *SingletonContext) {}

// EnterContenido is called when production contenido is entered.
func (s *BaseL5Listener) EnterContenido(ctx *ContenidoContext) {}

// ExitContenido is called when production contenido is exited.
func (s *BaseL5Listener) ExitContenido(ctx *ContenidoContext) {}
