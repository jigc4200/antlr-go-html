package main

import (
	"fmt"
	"os"

	"antlr-go-html/parser" // Cambia si tu módulo se llama diferente
	// Cambia si tu módulo se llama diferente

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
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewL5Parser(stream)

	tree := p.Pagina()

	// Crear e invocar listener personalizado
	// Create an instance of HTMLListener with required arguments
	listener := NewHTMLListener() // Replace with the correct constructor and arguments if needed
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	// Guardar en archivo
	err = os.WriteFile("output.html", []byte(listener.HTML()), 0644)
	if err != nil {
		fmt.Println("❌ Error al guardar HTML:", err)
	} else {
		fmt.Println("✅ HTML guardado en output.html")
	}
}
