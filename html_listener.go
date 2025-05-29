// Package main provides an ANTLR listener for converting a custom HTML-like syntax to HTML.
package main

import (
	"antlr-go-html/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type HTMLListener struct {
	*parser.BaseL5Listener
	html          strings.Builder
	currentList   strings.Builder
	currentLink   strings.Builder
	inLink        bool
	inImage       bool
	currentStyles []string
}

func NewHTMLListener() *HTMLListener {
	return &HTMLListener{}
}

func (l *HTMLListener) HTML() string {
	return l.html.String()
}

func (l *HTMLListener) EnterPagina(ctx *parser.PaginaContext) {
	l.html.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"UTF-8\">\n")
}

func (l *HTMLListener) ExitPagina(ctx *parser.PaginaContext) {
	l.html.WriteString("</body>\n</html>")
}

func (l *HTMLListener) EnterCabecera(ctx *parser.CabeceraContext) {
	l.html.WriteString("<title>")
}

func (l *HTMLListener) ExitCabecera(ctx *parser.CabeceraContext) {
	l.html.WriteString("</title>\n</head>\n<body>\n")
}

func (l *HTMLListener) EnterTitulo(ctx *parser.TituloContext) {
	if ctx.Contenido() != nil {
		l.html.WriteString(processContent(ctx.Contenido()))
	}
}

func (l *HTMLListener) EnterCuerpo(ctx *parser.CuerpoContext) {
	// Body content is handled by individual elements
}

func (l *HTMLListener) EnterPie(ctx *parser.PieContext) {
	l.html.WriteString("<footer>")
	if ctx.Contenido() != nil {
		l.html.WriteString(processContent(ctx.Contenido()))
	}
	l.html.WriteString("</footer>\n")
}

func (l *HTMLListener) EnterEstilo(ctx *parser.EstiloContext) {
	tagMapping := map[string]string{
		"Negrita":     "b",
		"Cursiva":     "i",
		"Subrayado":   "u",
		"Tachado":     "s",
		"Subindice":   "sub",
		"Superindice": "sup",
		"Tamaño":      "span",
	}

	startToken := ctx.GetStart().GetText()
	tag, ok := tagMapping[startToken]
	if !ok {
		tag = "span"
	}

	styleAttr := ""
	if startToken == "Tamaño" && ctx.ATTR_VAL() != nil {
		fontSize := strings.Trim(ctx.ATTR_VAL().GetText(), `"`)
		if _, err := strconv.Atoi(fontSize); err == nil {
			fontSize += "px"
		}
		styleAttr = " style=\"font-size: " + fontSize + ";\""
	}

	l.html.WriteString("<" + tag + styleAttr + ">")

	if ctx.Contenido() != nil {
		l.html.WriteString(processContent(ctx.Contenido()))
	}

	l.html.WriteString("</" + tag + ">")
}

func (l *HTMLListener) EnterFormato(ctx *parser.FormatoContext) {
	switch ctx.GetStart().GetTokenType() {
	case parser.L5ParserSANGRADO:
		l.html.WriteString("<div style=\"margin-left: 40px;\">")
		if ctx.Contenido() != nil {
			l.html.WriteString(processContent(ctx.Contenido()))
		}
		l.html.WriteString("</div>")

	case parser.L5ParserPARRAFO:
		align := "left"
		class := ""
		id := ""

		// Process attributes
		for _, attr := range ctx.AllAtributo() {
			if pos := attr.POSICION_KW(); pos != nil {
				switch attr.POSICION().GetText() {
				case "centrada":
					align = "center"
				case "derecha":
					align = "right"
				}
			} else if attrId := attr.ID(); attrId != nil {
				id = " id=\"" + strings.Trim(attr.ATTR_VAL().GetText(), `"`) + "\""
			} else if attrClass := attr.CLASE(); attrClass != nil {
				class = " class=\"" + strings.Trim(attr.ATTR_VAL().GetText(), `"`) + "\""
			}
		}

		l.html.WriteString("<p style=\"text-align: " + align + "\"" + id + class + ">")

		// Process content
		if ctx.Contenido() != nil {
			l.html.WriteString(processContent(ctx.Contenido()))
		}

		l.html.WriteString("</p>")
	}
}

func (l *HTMLListener) EnterLista(ctx *parser.ListaContext) {
	l.html.WriteString("<ul>")
}

func (l *HTMLListener) ExitLista(ctx *parser.ListaContext) {
	l.html.WriteString("</ul>")
}

func (l *HTMLListener) EnterElementoLista(ctx *parser.ElementoListaContext) {
	l.html.WriteString("<li>")
	if ctx.Contenido() != nil {
		l.html.WriteString(processContent(ctx.Contenido()))
	}
	l.html.WriteString("</li>")
}

func (l *HTMLListener) EnterEnlace(ctx *parser.EnlaceContext) {
	url := ""
	if conCtx := ctx.CON(); conCtx != nil {
		url = strings.Trim(ctx.ATTR_VAL().GetText(), `"`)
	}

	l.html.WriteString("<a href=\"" + url + "\">")

	if mostrarCtx := ctx.MOSTRAR(); mostrarCtx != nil {
		if ctx.Contenido() != nil {
			l.html.WriteString(processContent(ctx.Contenido()))
		}
	}

	l.html.WriteString("</a>")
}

func (l *HTMLListener) EnterImagen(ctx *parser.ImagenContext) {
	src := ""
	alt := ""

	if srcCtx := ctx.SRC(); srcCtx != nil {
		src = strings.Trim(ctx.ATTR_VAL().GetText(), `"`)
	}

	if altCtx := ctx.ALT(); altCtx != nil && ctx.Contenido() != nil {
		alt = processContent(ctx.Contenido())
	}

	l.html.WriteString("<img src=\"" + src + "\" alt=\"" + alt + "\">")
}

func (l *HTMLListener) EnterEncabezado(ctx *parser.EncabezadoContext) {
	level := "1"
	if ctx.ATTR_VAL() != nil {
		level = strings.Trim(ctx.ATTR_VAL().GetText(), `"`)
	}

	l.html.WriteString("<h" + level + ">")
	if ctx.Contenido() != nil {
		l.html.WriteString(processContent(ctx.Contenido()))
	}
	l.html.WriteString("</h" + level + ">")
}

func (l *HTMLListener) EnterSingleton(ctx *parser.SingletonContext) {
	switch ctx.GetStart().GetText() {
	case "Linea":
		l.html.WriteString("<hr>")
	case "Salto":
		l.html.WriteString("<br>")
	}
}

// Helper function to process nested content
func processContent(ctx parser.IContenidoContext) string {
	var sb strings.Builder
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		switch c := child.(type) {
		case antlr.TerminalNode:
			text := c.GetText()
			// Only unquote if it's a quoted string
			if strings.HasPrefix(text, `"`) && strings.HasSuffix(text, `"`) {
				text = text[1 : len(text)-1]
			}
			sb.WriteString(text)

		case *parser.EstiloContext:
			sb.WriteString(renderStyle(c))

		case *parser.ElementoContext:
			if c.Estilo() != nil {
				sb.WriteString(renderStyle(c.Estilo()))
			} else if c.Enlace() != nil {
				sb.WriteString(renderLink(c.Enlace()))
			}
			// Add other element types as needed
		}
	}
	return sb.String()
}

// Helper to render style elements
func renderStyle(ctx parser.IEstiloContext) string {
	var sb strings.Builder
	tagMapping := map[string]string{
		"Negrita":     "b",
		"Cursiva":     "i",
		"Subrayado":   "u",
		"Tachado":     "s",
		"Subindice":   "sub",
		"Superindice": "sup",
	}

	styleCtx := ctx.(*parser.EstiloContext)
	startToken := styleCtx.GetStart().GetText()
	tag, ok := tagMapping[startToken]
	if !ok {
		tag = "span"
	}

	sb.WriteString("<" + tag + ">")
	if styleCtx.Contenido() != nil {
		sb.WriteString(processContent(styleCtx.Contenido()))
	}
	sb.WriteString("</" + tag + ">")

	return sb.String()
}

// Helper to render links
func renderLink(ctx parser.IEnlaceContext) string {
	var sb strings.Builder
	linkCtx := ctx.(*parser.EnlaceContext)

	url := ""
	if conCtx := linkCtx.CON(); conCtx != nil {
		url = strings.Trim(linkCtx.ATTR_VAL().GetText(), `"`)
	}

	sb.WriteString("<a href=\"" + url + "\">")

	if mostrarCtx := linkCtx.MOSTRAR(); mostrarCtx != nil {
		if linkCtx.Contenido() != nil {
			sb.WriteString(processContent(linkCtx.Contenido()))
		}
	}

	sb.WriteString("</a>")
	return sb.String()
}

// Helper to remove quotes from attribute values
func stripQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
