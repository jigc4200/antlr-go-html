grammar L5;

// === Reglas principales ===

pagina       : PAGINA cabecera cuerpo pie FINPAGINA ;
cabecera     : CABECERA titulo FINCABECERA ;
titulo       : TITULO contenido FINTITULO ;
cuerpo       : CUERPO elemento* FINCUERPO ;
pie          : PPAGINA contenido FPPAGINA ;

elemento
    : estilo
    | formato
    | lista
    | enlace
    | singleton
    | imagen
    | encabezado
    ;

estilo
    : NEGRITA contenido FINNEGRITA
    | CURSIVA contenido FINCURSIVA
    | SUBRAYADO contenido FINSUBRAYADO
    | TACHADO contenido FINTACHADO
    | SUBINDICE contenido FINSUBINDICE
    | SUPERINDICE contenido FINSUPERINDICE
    | TAMANO ATTR_VAL contenido FINTAMANO  
    ;

formato
    : SANGRADO contenido FINSANGRADO
    | PARRAFO estilo? atributo* contenido FINPARRAFO
    ;

atributo
    : POSICION_KW POSICION FINPOSICION
    | ID ATTR_VAL FINID
    | CLASE ATTR_VAL FINCLASE
    ;

lista
    : LISTA elementoLista+ FINLISTA
    ;

elementoLista
    : ELEMENTOLISTA contenido FINELEMENTOLISTA
    ;

enlace
    : ENLACE
      CON ATTR_VAL FINCON
      MOSTRAR contenido FINMOSTRAR
      FINENLACE
    ;

imagen
    : IMAGEN
      SRC ATTR_VAL FINSRC
      (ALT contenido FINALT)?
      FINIMAGEN
    ;

encabezado
    : ENCABEZADO ATTR_VAL contenido FINENCABEZADO  
    ;

singleton
    : LINEA
    | SALTO
    ;

contenido
    : (ATTR_VAL| elemento | estilo )+
    ;

// === Tokens ===

// Palabras clave
PAGINA          : 'Pagina' ;
CABECERA        : 'Cabecera' ;
TITULO          : 'Titulo' ;
FINTITULO       : 'FinTitulo' ;
FINCABECERA     : 'FinCabecera' ;
CUERPO          : 'Cuerpo' ;
FINCUERPO       : 'FinCuerpo' ;
PPAGINA         : 'Ppagina' ;
FPPAGINA        : 'FPpagina' ;
FINPAGINA       : 'FinPagina' ; // Added missing token definition
NEGRITA         : 'Negrita' ;
FINNEGRITA      : 'FinNegrita' ;
CURSIVA         : 'Cursiva' ;
FINCURSIVA      : 'FinCursiva' ;
SUBRAYADO       : 'Subrayado' ;
FINSUBRAYADO    : 'FinSubrayado' ;
TACHADO         : 'Tachado' ;
FINTACHADO      : 'FinTachado' ;
SUBINDICE       : 'Subindice' ;
FINSUBINDICE    : 'FinSubindice' ;
SUPERINDICE     : 'Superindice' ;
FINSUPERINDICE  : 'FinSuperindice' ;
TAMANO          : 'Tamaño' ;
FINTAMANO       : 'FinTamaño' ;
SANGRADO        : 'Sangrado' ;
FINSANGRADO     : 'FinSangrado' ;
PARRAFO         : 'Parrafo' ;
FINPARRAFO      : 'FinParrafo' ;
POSICION_KW     : 'Posicion' ; // Renombrado para evitar conflicto con token POSICION
FINPOSICION     : 'FinPosicion' ;
ID              : 'Id' ;
FINID           : 'FinId' ;
CLASE           : 'Clase' ;
FINCLASE        : 'FinClase' ;
LISTA           : 'Lista' ;
FINLISTA        : 'FinLista' ;
ELEMENTOLISTA   : 'ElementoLista' ;
FINELEMENTOLISTA: 'FinElementoLista' ;
ENLACE          : 'Enlace' ;
FINENLACE       : 'FinEnlace' ;
CON             : 'Con' ;
FINCON          : 'FinCon' ;
MOSTRAR         : 'Mostrar' ;
FINMOSTRAR      : 'FinMostrar' ;
IMAGEN          : 'Imagen' ;
FINIMAGEN       : 'FinImagen' ;
SRC             : 'Src' ;
FINSRC          : 'FinSrc' ;
ALT             : 'Alt' ;
FINALT          : 'FinAlt' ;
ENCABEZADO      : 'Encabezado' ;
FINENCABEZADO   : 'FinEncabezado' ;
LINEA           : 'Linea' ;
SALTO           : 'Salto' ;
ESTILO          : 'Estilo' ;
FINESTILO       : 'FinEstilo' ;


// Otros tokens
POSICION   : 'izquierda' | 'centrada' | 'derecha' ;
ATTR_VAL   : '"' (~["\r\n])* '"' ; // atributo como "valor"
WS         : [ \t\r\n]+ -> skip ;
//TEXTO      : ~[\r\n<>"\t]+ ; // Cualquier cosa que no sea palabras clave, delimitadores, etc.
