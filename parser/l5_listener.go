// Code generated from L5.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // L5

import "github.com/antlr4-go/antlr/v4"

// L5Listener is a complete listener for a parse tree produced by L5Parser.
type L5Listener interface {
	antlr.ParseTreeListener

	// EnterPagina is called when entering the pagina production.
	EnterPagina(c *PaginaContext)

	// EnterCabecera is called when entering the cabecera production.
	EnterCabecera(c *CabeceraContext)

	// EnterTitulo is called when entering the titulo production.
	EnterTitulo(c *TituloContext)

	// EnterCuerpo is called when entering the cuerpo production.
	EnterCuerpo(c *CuerpoContext)

	// EnterPie is called when entering the pie production.
	EnterPie(c *PieContext)

	// EnterElemento is called when entering the elemento production.
	EnterElemento(c *ElementoContext)

	// EnterEstilo is called when entering the estilo production.
	EnterEstilo(c *EstiloContext)

	// EnterFormato is called when entering the formato production.
	EnterFormato(c *FormatoContext)

	// EnterAtributo is called when entering the atributo production.
	EnterAtributo(c *AtributoContext)

	// EnterLista is called when entering the lista production.
	EnterLista(c *ListaContext)

	// EnterElementoLista is called when entering the elementoLista production.
	EnterElementoLista(c *ElementoListaContext)

	// EnterEnlace is called when entering the enlace production.
	EnterEnlace(c *EnlaceContext)

	// EnterImagen is called when entering the imagen production.
	EnterImagen(c *ImagenContext)

	// EnterEncabezado is called when entering the encabezado production.
	EnterEncabezado(c *EncabezadoContext)

	// EnterSingleton is called when entering the singleton production.
	EnterSingleton(c *SingletonContext)

	// EnterContenido is called when entering the contenido production.
	EnterContenido(c *ContenidoContext)

	// ExitPagina is called when exiting the pagina production.
	ExitPagina(c *PaginaContext)

	// ExitCabecera is called when exiting the cabecera production.
	ExitCabecera(c *CabeceraContext)

	// ExitTitulo is called when exiting the titulo production.
	ExitTitulo(c *TituloContext)

	// ExitCuerpo is called when exiting the cuerpo production.
	ExitCuerpo(c *CuerpoContext)

	// ExitPie is called when exiting the pie production.
	ExitPie(c *PieContext)

	// ExitElemento is called when exiting the elemento production.
	ExitElemento(c *ElementoContext)

	// ExitEstilo is called when exiting the estilo production.
	ExitEstilo(c *EstiloContext)

	// ExitFormato is called when exiting the formato production.
	ExitFormato(c *FormatoContext)

	// ExitAtributo is called when exiting the atributo production.
	ExitAtributo(c *AtributoContext)

	// ExitLista is called when exiting the lista production.
	ExitLista(c *ListaContext)

	// ExitElementoLista is called when exiting the elementoLista production.
	ExitElementoLista(c *ElementoListaContext)

	// ExitEnlace is called when exiting the enlace production.
	ExitEnlace(c *EnlaceContext)

	// ExitImagen is called when exiting the imagen production.
	ExitImagen(c *ImagenContext)

	// ExitEncabezado is called when exiting the encabezado production.
	ExitEncabezado(c *EncabezadoContext)

	// ExitSingleton is called when exiting the singleton production.
	ExitSingleton(c *SingletonContext)

	// ExitContenido is called when exiting the contenido production.
	ExitContenido(c *ContenidoContext)
}
