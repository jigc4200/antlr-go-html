// Package main provides an ANTLR listener for converting a custom HTML-like syntax to HTML.
package main

import (
	"antlr-go-html/parser"
	"log"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type HTMLListener struct {
	*parser.BaseL5Listener
	html strings.Builder
}

func NewHTMLListener() *HTMLListener {
	return &HTMLListener{}
}

func (l *HTMLListener) HTML() string {
	return l.html.String()
}

func (l *HTMLListener) EnterPagina(ctx *parser.PaginaContext) {
	l.html.WriteString("<html><body>\n")
}

func (l *HTMLListener) ExitPagina(ctx *parser.PaginaContext) {
	l.html.WriteString("</body></html>\n")
}

func (l *HTMLListener) EnterTitulo(ctx *parser.TituloContext) {
	// Get text from the 'contenido' rule
	if ctx.Contenido() != nil {
		l.html.WriteString("<h1>" + stripQuotes(ctx.Contenido().GetText()) + "</h1>\n")
	} else {
		log.Println("Advertencia: Título sin contenido")
		l.html.WriteString("<h1></h1>\n")
	}
}

func (l *HTMLListener) EnterEstilo(ctx *parser.EstiloContext) {
	tagMapping := map[string]string{
		"Negrita":     "b",
		"Cursiva":     "i",
		"Subrayado":   "u",
		"Tachado":     "s",
		"Subindice":   "sub",
		"Superindice": "sup",
		"Tamaño":      "span", // Handle style attribute separately
	}

	// Determine the tag based on the starting token
	startToken := ctx.GetStart().GetText()
	tag, ok := tagMapping[startToken]
	if !ok {
		log.Println("Estilo desconocido:", startToken)
		tag = "span" // Default to span for unknown styles
	}

	// Handle Tamaño style with font-size attribute
	styleAttr := ""
	if startToken == "Tamaño" {
		if ctx.ATTR_VAL() != nil {
			// Extract the value from the ATTR_VAL token (e.g., "14px")
			fontSize := strings.Trim(ctx.ATTR_VAL().GetText(), `"`)
			styleAttr = " style='font-size: " + fontSize + ";'"
		} else {
			log.Println("Advertencia: Estilo Tamaño sin valor de tamaño")
		}
	}

	l.html.WriteString("<" + tag + styleAttr + ">")

	// Get text from the 'contenido' rule
	if ctx.Contenido() != nil {
		l.html.WriteString(stripQuotes(ctx.Contenido().GetText()))

	} else {
		log.Println("Advertencia:", startToken, "sin contenido")
	}

	// Close the tag, handling the case of 'span style...'
	closeTag := strings.Split(tag, " ")[0]
	l.html.WriteString("</" + closeTag + ">")
}

func (l *HTMLListener) EnterFormato(ctx *parser.FormatoContext) {
	// Check which alternative of the 'formato' rule was matched
	// The first child will be the token 'Sangrado' or 'Parrafo'
	switch ctx.GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType() {
	case parser.L5ParserSANGRADO: // 'Sangrado'
		l.html.WriteString("<div style='margin-left: 40px;'>")
		if ctx.Contenido() != nil {
			l.html.WriteString(stripQuotes(ctx.Contenido().GetText()))

		} else {
			log.Println("Advertencia: Sangrado sin contenido")
		}
		l.html.WriteString("</div>\n")
	case parser.L5ParserPARRAFO: // 'Parrafo'
		align := "left"
		// Iterate through attributes to find Posicion
		for _, attrCtx := range ctx.AllAtributo() {
			// Check if the POSICION token exists in this attribute context
			if posToken := attrCtx.POSICION(); posToken != nil {
				switch posToken.GetText() {
				case "centrada":
					align = "center"
				case "derecha":
					align = "right"
				}
				// Assuming only one position attribute per paragraph
				break
			}
			// Could add checks for attrCtx.ATTR_VAL() here if needed for Id or Clase
		}
		l.html.WriteString("<p style='text-align: " + align + ";'>")
		// Get text from the 'contenido' rule
		if ctx.Contenido() != nil {
			l.html.WriteString(ctx.Contenido().GetText())
		} else {
			log.Println("Advertencia: Párrafo sin contenido")
		}
		l.html.WriteString("</p>\n")
	}
}

func (l *HTMLListener) EnterLista(ctx *parser.ListaContext) {
	l.html.WriteString("<ul>\n")
}

func (l *HTMLListener) ExitLista(ctx *parser.ListaContext) {
	l.html.WriteString("</ul>\n")
}

func (l *HTMLListener) EnterElementoLista(ctx *parser.ElementoListaContext) {
	// Get text from the 'contenido' rule
	if ctx.Contenido() != nil {
		l.html.WriteString("<li>" + stripQuotes(ctx.Contenido().GetText()) + "</li>\n")
	} else {
		log.Println("Advertencia: ElementoLista sin contenido")
		l.html.WriteString("<li></li>\n")
	}
}

func (l *HTMLListener) EnterEnlace(ctx *parser.EnlaceContext) {
	// The URL is the ATTR_VAL token after 'Con'
	urlToken := ctx.ATTR_VAL()
	if urlToken == nil {
		log.Println("Advertencia: Enlace sin URL")
		// Attempt to get content even without URL to avoid losing text
		if ctx.Contenido() != nil {
			l.html.WriteString(stripQuotes(ctx.Contenido().GetText()))
		}
		return
	}
	url := strings.Trim(urlToken.GetText(), `"`) // Remove quotes from ATTR_VAL

	// The link text is the 'contenido' rule after 'Mostrar'
	linkContent := ctx.Contenido()
	if linkContent == nil {
		log.Println("Advertencia: Enlace sin contenido")
		l.html.WriteString("<a href='" + url + "'></a>") // Create link with empty text
		return
	}

	l.html.WriteString("<a href='" + url + "'>")
	// Get text from the 'contenido' rule
	l.html.WriteString(linkContent.GetText())
	l.html.WriteString("</a>")
}

func (l *HTMLListener) EnterSingleton(ctx *parser.SingletonContext) {
	switch ctx.GetStart().GetText() {
	case "Linea":
		l.html.WriteString("<hr/>\n")
	case "Salto":
		l.html.WriteString("<br/>\n")
	}
}

func stripQuotes(s string) string {
	return strings.Trim(s, `"`)
}
