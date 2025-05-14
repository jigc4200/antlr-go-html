package main

import (
	"antlr-go-html/parser" // Asegúrate que esta ruta sea correcta
	"log"
	"strings"

	"github.com/antlr4-go/antlr/v4" // Importación común para antlr.ParserRuleContext
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

// Helper function para reconstruir el texto con espacios
func (l *HTMLListener) reconstructTextFromContext(textoCtx parser.ITextoContext) string {
	if textoCtx == nil {
		return ""
	}

	var sb strings.Builder
	// ITextoContext debería tener AllTEXTO() que devuelve []antlr.TerminalNode
	textTokens := textoCtx.AllTEXTO()

	for i, tokenNode := range textTokens {
		sb.WriteString(tokenNode.GetText())
		if i < len(textTokens)-1 { // Si no es el último token
			sb.WriteString(" ") // Añade el espacio
		}
	}
	return sb.String()
}

func (l *HTMLListener) EnterPagina(ctx *parser.PaginaContext) {
	l.html.WriteString("<html><body>\n")
}

func (l *HTMLListener) ExitPagina(ctx *parser.PaginaContext) {
	l.html.WriteString("</body></html>\n")
}

func (l *HTMLListener) EnterTitulo(ctx *parser.TituloContext) {
	l.html.WriteString("<h1>")
}

func (l *HTMLListener) ExitTitulo(ctx *parser.TituloContext) {
	// El texto del título ya habrá sido escrito por EnterTexto
	l.html.WriteString("</h1>\n")
}

// EnterTexto modificado
func (l *HTMLListener) EnterTexto(ctx *parser.TextoContext) {
	parent := ctx.GetParent()
	// Si el padre es EnlaceContext, EnterEnlace se encargará de este texto.
	if _, ok := parent.(*parser.EnlaceContext); ok {
		return
	}

	reconstructedText := l.reconstructTextFromContext(ctx)
	l.html.WriteString(reconstructedText)
}

// ------ Estilos ------
func (l *HTMLListener) EnterEstilo(ctx *parser.EstiloContext) {
	// Usamos ctx.GetChild(0) para obtener el token específico del tipo de estilo (Negrita, Cursiva, etc.)
	// ya que GetStart() se refiere al primer token de toda la regla 'estilo', que es el tag de inicio.
	var estiloTagToken antlr.Token
	// El primer hijo de EstiloContext es el TerminalNode del tag (Ej: 'Negrita')
	if tn, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		estiloTagToken = tn.GetSymbol()
	} else {
		log.Println("Advertencia: No se pudo determinar el tipo de estilo.")
		l.html.WriteString("<span>") // Fallback
		return
	}

	tagMapping := map[string]string{
		"Negrita":     "b",
		"Cursiva":     "i",
		"Subrayado":   "u",
		"Tachado":     "s",
		"Subindice":   "sub",
		"Superindice": "sup",
		// Para Tamaño, el tag de apertura completo se define aquí
		"Tamaño": "span style='font-size: larger;'",
	}
	// Usamos el texto del token que identifica el tipo de estilo.
	tag, ok := tagMapping[estiloTagToken.GetText()]
	if !ok {
		log.Println("Advertencia: Tag de estilo desconocido -", estiloTagToken.GetText())
		tag = "span" // Fallback
	}

	// Si el tag ya incluye atributos (como en 'Tamaño'), no añadimos '<' y '>' dos veces.
	if strings.Contains(tag, " ") { // Heurística simple para tags con atributos
		l.html.WriteString("<" + tag + ">")
	} else {
		l.html.WriteString("<" + tag + ">")
	}
}

func (l *HTMLListener) ExitEstilo(ctx *parser.EstiloContext) {
	var estiloTagToken antlr.Token
	if tn, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		estiloTagToken = tn.GetSymbol()
	} else {
		log.Println("Advertencia: No se pudo determinar el tipo de estilo al salir.")
		l.html.WriteString("</span>") // Fallback
		return
	}
	// Para el cierre, solo necesitamos el nombre del tag
	tagMapping := map[string]string{
		"Negrita":     "b",
		"Cursiva":     "i",
		"Subrayado":   "u",
		"Tachado":     "s",
		"Subindice":   "sub",
		"Superindice": "sup",
		"Tamaño":      "span", // Tag de cierre simple
	}
	tag, ok := tagMapping[estiloTagToken.GetText()]
	if !ok {
		log.Println("Advertencia: Tag de estilo desconocido al salir -", estiloTagToken.GetText())
		tag = "span" // Fallback
	}
	l.html.WriteString("</" + tag + ">")
}

// ------ Formato ------
func (l *HTMLListener) EnterFormato(ctx *parser.FormatoContext) {
	// Usamos ctx.GetChild(0) para obtener el token específico del tipo de formato (Sangrado, Parrafo)
	var formatoTagToken antlr.Token
	if tn, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		formatoTagToken = tn.GetSymbol()
	} else {
		log.Println("Advertencia: No se pudo determinar el tipo de formato.")
		// Manejar el error o usar un fallback
		return
	}

	if formatoTagToken.GetText() == "Sangrado" {
		l.html.WriteString("<div style='margin-left: 40px;'>")
	} else { // Parrafo
		align := "left" // Valor por defecto
		// ctx.Posicion() puede ser nil si la posición es opcional (posicion?)
		if ctx.Posicion() != nil && ctx.Posicion().POSICION() != nil {
			switch ctx.Posicion().POSICION().GetText() {
			case "centrada":
				align = "center"
			case "derecha":
				align = "right"
			}
		}
		l.html.WriteString("<p style='text-align: " + align + ";'>")
	}
}

func (l *HTMLListener) ExitFormato(ctx *parser.FormatoContext) {
	var formatoTagToken antlr.Token
	if tn, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		formatoTagToken = tn.GetSymbol()
	} else {
		// Este caso es menos probable si EnterFormato funcionó, pero por completitud.
		log.Println("Advertencia: No se pudo determinar el tipo de formato al salir.")
		return
	}

	if formatoTagToken.GetText() == "Sangrado" {
		l.html.WriteString("</div>\n")
	} else { // Parrafo
		l.html.WriteString("</p>\n")
	}
}

// ------ Lista ------
func (l *HTMLListener) EnterLista(ctx *parser.ListaContext) {
	l.html.WriteString("<ul>\n")
}

func (l *HTMLListener) ExitLista(ctx *parser.ListaContext) {
	l.html.WriteString("</ul>\n")
}

func (l *HTMLListener) EnterElementoLista(ctx *parser.ElementoListaContext) {
	l.html.WriteString("<li>")
	// El texto del ElementoLista será manejado por EnterTexto
}

func (l *HTMLListener) ExitElementoLista(ctx *parser.ElementoListaContext) {
	l.html.WriteString("</li>\n")
}

// ------ Enlace ------
func (l *HTMLListener) EnterEnlace(ctx *parser.EnlaceContext) {
	textoContexts := ctx.AllTexto() // []parser.ITextoContext

	if len(textoContexts) < 2 {
		log.Println("Advertencia: Enlace mal formado, no hay suficientes 'texto' para URL y display.")
		l.html.WriteString("")
		return
	}

	url := l.reconstructTextFromContext(textoContexts[0])
	txt := l.reconstructTextFromContext(textoContexts[1])

	l.html.WriteString("<a href='" + url + "'>" + txt + "</a>")
}

// ------ Singleton ------
func (l *HTMLListener) EnterSingleton(ctx *parser.SingletonContext) {
	// GetStart() aquí es el token 'Linea' o 'Salto'
	switch ctx.GetStart().GetText() {
	case "Linea":
		l.html.WriteString("<hr/>\n")
	case "Salto":
		l.html.WriteString("<br/>\n")
	}
}
