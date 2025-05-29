# Procesador L5 a HTML

Este proyecto es un procesador que convierte archivos escritos en un lenguaje de marcado personalizado, definido por la gramática ANTLR [`parser/L5.g4`](parser/L5.g4:1), a archivos HTML y listados de tokens. Está construido utilizando Go y la biblioteca ANTLR.

## Estructura del Proyecto

*   [`main.go`](main.go:1): El punto de entrada principal de la aplicación. Se encarga de leer el archivo de entrada, invocar el lexer y parser de ANTLR, usar el listener para generar HTML y guardar los archivos de salida.
*   `parser/`: Contiene los archivos Go generados automáticamente por ANTLR a partir de la gramática [`L5.g4`](parser/L5.g4:1), así como los archivos de tokens e intérprete.
*   [`html_listener.go`](html_listener.go): Implementa el listener de ANTLR para recorrer el árbol de parseo y generar la salida HTML correspondiente a la sintaxis L5.
*   `input.l5`: Un archivo de ejemplo que demuestra la sintaxis del lenguaje de marcado L5. Puedes modificar este archivo para probar diferentes estructuras L5.
*   `output.html`: El archivo de salida generado que contiene el HTML resultante del procesamiento de `input.l5`.
*   `listado.txt`: Un archivo de salida que contiene un listado de todos los tokens identificados por el lexer de ANTLR en el archivo `input.l5`.
*   [`go.mod`](go.mod) / [`go.sum`](go.sum): Archivos de módulos de Go que gestionan las dependencias del proyecto.

## Instalación

Asegúrate de tener Go instalado en tu sistema.

1.  Clona el repositorio:
    ```bash
    git clone <URL_DEL_REPOSITORIO>
    cd antlr-go-html
    ```
2.  Instala las dependencias de Go:
    ```bash
    go mod tidy
    ```
3.  (Opcional) Si modificas la gramática [`L5.g4`](parser/L5.g4:1) y necesitas regenerar los archivos del parser, asegúrate de tener ANTLR instalado y ejecuta el siguiente comando en el directorio raíz del proyecto:
    ```bash
    antlr4 -Dlanguage=Go parser/L5.g4
    ```

## Uso

1.  Edita el archivo `input.l5` para incluir el contenido que deseas procesar utilizando la sintaxis L5.
2.  Ejecuta el programa desde la terminal en el directorio raíz del proyecto:
    ```bash
    go run main.go
    ```
3.  El programa generará dos archivos en el directorio raíz:
    *   `output.html`: Contiene el HTML generado.
    *   `listado.txt`: Contiene el listado de tokens.

## Sintaxis L5

La sintaxis del lenguaje L5 está definida en el archivo de gramática [`parser/L5.g4`](parser/L5.g4:1). Aquí tienes un pequeño ejemplo:

```l5
Pagina
    Cabecera
        Titulo "Mi Página Web Personal" FinTitulo
    FinCabecera
    Cuerpo
        Encabezado "1" "Bienvenido a mi sitio web" FinEncabezado
        
        Parrafo
            Posicion centrada FinPosicion
            "Este es un " Negrita "ejemplo completo" FinNegrita " de lo que se puede hacer con esta gramática."
        FinParrafo
        
        Linea
        
        Parrafo
            "Texto con " Cursiva "énfasis" FinCursiva " y " Subrayado "formatos especiales" FinSubrayado "."
        FinParrafo
        
        Lista
            ElementoLista "Primer elemento de lista" FinElementoLista
            ElementoLista
                Negrita "Segundo elemento" FinNegrita
                " con formato"
            FinElementoLista
            ElementoLista
                Enlace
                    Con "https:www.google.com" FinCon
                    Mostrar "Enlace de ejemplo" FinMostrar
                FinEnlace
            FinElementoLista
        FinLista
        
        Imagen
            Src "/home/jigc4200/Pictures/imagen1.png" FinSrc
            Alt "Una imagen descriptiva" FinAlt
        FinImagen
        
        Salto
        
        Parrafo
            Tamaño "2" "Texto con tamaño modificado" FinTamaño
        FinParrafo
    FinCuerpo
    Ppagina "© 2023 - Todos los derechos reservados" Fppagina
FinPagina
```

Este ejemplo muestra una estructura básica de página con cabecera, cuerpo y pie, incluyendo elementos de estilo, un salto de línea y una imagen.
