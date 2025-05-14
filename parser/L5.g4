grammar L5;

pagina       : 'Pagina' cabecera cuerpo pie 'FinPagina' ;
cabecera     : 'Cabecera' titulo 'FinCabecera' ;
titulo       : 'Titulo' texto 'FinTitulo' ;
cuerpo       : 'Cuerpo' elemento* 'FinCuerpo' ;
pie          : 'Ppagina' texto 'FPpagina' ;

elemento
    : estilo
    | formato
    | lista
    | enlace
    | singleton
    ;

estilo
    : 'Negrita' texto 'FinNegrita'
    | 'Cursiva' texto 'FinCursiva'
    | 'Subrayado' texto 'FinSubrayado'
    | 'Tachado' texto 'FinTachado'
    | 'Subindice' texto 'FinSubindice'
    | 'Superindice' texto 'FinSuperindice'
    | 'Tamaño' texto 'FinTamaño'
    ;

formato
    : 'Sangrado' texto 'FinSangrado'
    | 'Parrafo' posicion? texto 'FinParrafo'
    ;

posicion
    : 'Posicion' POSICION 'FinPosicion'
    ;

POSICION : 'izquierda' | 'centrada' | 'derecha' 
    ;

lista
    : 'Lista' elementoLista+ 'FinLista'
    ;

elementoLista
    : 'ElementoLista' texto 'FinElementoLista'
    ;

enlace
    : 'Enlace'
      'Con' texto 'FinCon'
      'Mostrar' texto 'FinMostrar'
      'FinEnlace'
    ;
singleton
    : 'Linea'
    | 'Salto'
    ;

texto : TEXTO+ ;

TEXTO : ~[<>\r\n \t]+ ;


WS : [ \t\r\n]+ -> skip ;