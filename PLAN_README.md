# Plan para el README

## 1. Título: Procesador L5 a HTML

## 2. Descripción:
Explicar que el proyecto parsea archivos con la sintaxis definida en `parser/L5.g4` y genera archivos HTML y listados de tokens. Mencionar que está construido con Go y ANTLR.

## 3. Estructura del Proyecto:
*   `main.go`: Punto de entrada principal.
*   `parser/`: Contiene los archivos generados por ANTLR y la gramática `L5.g4`.
*   `html_listener.go`: Implementación del listener de ANTLR para generar HTML.
*   `input.l5`: Archivo de entrada de ejemplo.
*   `output.html`: Archivo HTML generado (salida).
*   `listado.txt`: Archivo con el listado de tokens (salida).
*   `go.mod` / `go.sum`: Archivos de módulos de Go para dependencias.

## 4. Instalación:
*   Requisitos: Go (versión mínima si aplica).
*   Clonar el repositorio.
*   Instalar dependencias: `go mod download` o `go mod tidy`.
*   (Opcional) Si se necesita regenerar el parser: Instalar ANTLR y ejecutar el comando apropiado. 
    Codigo de compilacion de la herramienta ANTLR  'antlr4 -Dlanguage=Go L5.g4'
## 5. Uso:
*   Crear o modificar el archivo `input.l5` con la sintaxis L5.
*   Ejecutar el programa: `go run main.go html_listener.go `. 
*   Verificar los archivos generados: `output.html` y `listado.txt`.

## 6. Sintaxis L5:
Proporcionar un breve ejemplo de la sintaxis L5 o un enlace/referencia al archivo `parser/L5.g4`.

## 7. Contribución:
(Se puede añadir si el usuario lo desea).

## 8. Licencia:
(Se puede añadir si el usuario lo desea).