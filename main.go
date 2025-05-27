package main

import (
	"fmt"
	"os"

	"antlr-go-html/parser" // Cambia si tu módulo es diferente

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, err := os.ReadFile("input.l5")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al leer input.l5: %v\n", err)
		os.Exit(1)
	}

	// Lexer y parser
	is := antlr.NewInputStream(string(input))
	lexer := parser.NewL5Lexer(is)
	tokLexer := parser.NewL5Lexer(antlr.NewInputStream(string(input)))

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewL5Parser(stream)

	// crear txt de tokens
	tokens := tokLexer.GetAllTokens()
	for _, token := range tokens {
		tokenType := token.GetTokenType()
		var tokenTypeName string
		if tokenType >= 0 && tokenType < len(parser.L5LexerLexerStaticData.LiteralNames) {
			tokenTypeName = parser.L5LexerLexerStaticData.LiteralNames[tokenType]
		} else {
			tokenTypeName = "Desconocido" // Valor por defecto para tipos de token desconocidos
		}
		fmt.Printf("Token: %s, Tipo: %s, Posición: %d\n", token.GetText(), tokenTypeName, token.GetStart())
	}

	tree := p.Pagina()

	// Crear e invocar listener personalizado
	listener := NewHTMLListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	// Guardar en archivo
	err = os.WriteFile("output.html", []byte(listener.HTML()), 0644)
	if err != nil {
		fmt.Println("Error al guardar HTML:", err)
	} else {
		fmt.Println("HTML guardado en output.html")
	}

	// Guardar los tokens en un archivo .txt
	tokenFile, err := os.Create("listado.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al crear tokens.txt: %v\n", err)
		os.Exit(1)
	}
	defer tokenFile.Close()

	for _, token := range tokens {
		tokenType := token.GetTokenType()
		var tokenTypeName string
		if tokenType >= 0 && tokenType < len(parser.L5LexerLexerStaticData.LiteralNames) {
			tokenTypeName = parser.L5LexerLexerStaticData.LiteralNames[tokenType]
		} else {
			tokenTypeName = "Desconocido"
		}
		line := fmt.Sprintf("Token: %s, Tipo: %s, Posición: %d\n", token.GetText(), tokenTypeName, token.GetStart())
		tokenFile.WriteString(line)
	}
}
