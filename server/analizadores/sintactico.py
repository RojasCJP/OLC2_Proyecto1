from ..interprete.comandos.expressions.access import *
from ..interprete.comandos.expressions.access_expression import *
from ..interprete.comandos.expressions.arithmetic import *
from ..interprete.comandos.expressions.call_func import *
from ..interprete.comandos.expressions.literal import *
from ..interprete.comandos.expressions.relational import *

from ..interprete.comandos.statement import *
from ..interprete.comandos.nativas.print import *
from ..interprete.comandos.variables.asignacion import *
from ..interprete.comandos.variables.declaracion import *

reservadas = {
    "numero": "NUMERO",
    "println": "PRINTLN",
    "print": "PRINT",
    "while": "WHILE",
    "if": "IF",
    "else": "ELSE",
    "nothing": "NOTHING",
    "int64": "INT",
    "float64": "FLOAT",
    "bool": "BOOL",
    "char": "CHAR",
    "string": "STRING",
    "uppercase": "UPERCASE",
    "lowercase": "LOWERCASE",
    "log10": "LOG10",
    "log": "LOG",
    "sin": "SIN",
    "cos": "COS",
    "tan": "TAN",
    "sqrt": "SQRT",
    "true": "TRUE",
    "false": "FALSE",
    "local": "LOCAL",
    "global": "GLOBAL"
}

tokens = [
             "PTCOMA",
             "DOSP",
             "PUNTO",
             "COMA",
             "LLAVIZQ",
             "LLAVDER",
             "CORCHETEIZQ",
             "CORCHETEDER",
             "PARIZQ",
             "PARDER",
             "IGUAL",
             "MAS",
             "MENOS",
             "POR",
             "DIVIDIDO",
             "POTENCIA",
             "MODULO",
             "CONCAT",
             "MENQUE",
             "MENIGUALQUE",
             "MAYQUE",
             "MAYIGUALQUE",
             "OR",
             "AND",
             "NOT",
             "IGUALQUE",
             "NIGUALQUE",
             "DECIMAL",
             "ENTERO",
             "CADENA",
             "ID"
         ] + list(reservadas.values())

# Tokens
t_PTCOMA = r';'
t_DOSP = r':'
t_PUNTO = r'.'
t_COMA = r','
t_LLAVIZQ = r'{'
t_LLAVDER = r'}'
t_CORCHETEIZQ = r'\['
t_CORCHETEDER = r']'
t_PARIZQ = r'\('
t_PARDER = r'\)'
t_IGUAL = r'='
t_MAS = r'\+'
t_MENOS = r'-'
t_POR = r'\*'
t_DIVIDIDO = r'/'
t_POTENCIA = r'\^'
t_MODULO = r'%'
t_CONCAT = r'&'
t_MENQUE = r'<'
t_MENIGUALQUE = r'<='
t_MAYQUE = r'>'
t_MAYIGUALQUE = r'>='
t_OR = r'\|\|'
t_AND = r'&&'
t_NOT = r'!'
t_IGUALQUE = r'=='
t_NIGUALQUE = r'!='


def t_DECIMAL(t):
    r'\d+\.\d+'
    try:
        t.value = float(t.value)
    except ValueError:
        print("Float value too large %d", t.value)
        t.value = 0
    return t


def t_ENTERO(t):
    r'\d+'
    try:
        t.value = int(t.value)
    except ValueError:
        print("Integer value too large %d", t.value)
        t.value = 0
    return t


def t_ID(t):
    r'[a-zA-Z_][a-zA-Z_0-9]*'
    # Check for reserved words
    t.type = reservadas.get(t.value.lower(), 'ID')
    return t


def t_CADENA(t):
    r'\".*?\"'
    t.value = t.value[1:-1]  # remuevo las comillas
    return t


# Comentario de múltiples líneas /* .. */


def t_COMENTARIO_MULTILINEA(t):
    r'\#=(.|\n)*?=\#'
    t.lexer.lineno += t.value.count('\n')


# Comentario simple // ...


def t_COMENTARIO_SIMPLE(t):
    r'\#.*\n'
    t.lexer.lineno += 1


# Caracteres ignorados
t_ignore = " \t"


def t_newline(t):
    r'\n+'
    t.lexer.lineno += t.value.count("\n")


def t_error(t):
    print("Illegal character '%s'" % t.value[0])
    t.lexer.skip(1)


# Construyendo el analizador léxico
import ply.lex as lex

lexer = lex.lex()

precedence = (
    ('left', 'CONCAT'),
    ('left', 'MAS', 'MENOS'),
    ('left', 'POR', 'DIVIDIDO'),
    ('left', 'POTENCIA', 'MODULO'),
    ('right', 'UMENOS'),
)


# Definición de la gramática


def p_init(t):
    '''init            : instrucciones'''
    t[0] = t[1]


def p_instrucciones_lista(t):
    '''instrucciones    : instrucciones instruccion
                        | instruccion'''
    if len(t) != 2:
        t[1].append(t[2])
        t[0] = t[1]
    else:
        t[0] = [t[1]]


def p_instruccion(t):
    '''instruccion      : print_instr
                        | println_instr
                        | definicion_instr
                        | asignacion_instr
                        | definicion_asignacion_instr'''
    t[0] = t[1]


def p_expression(t):
    '''expression       : MENOS expression %prec UMENOS
                        | NOT expression %prec UMENOS
                        | expression MAS expression
                        | expression MENOS expression
                        | expression POR expression
                        | expression DIVIDIDO expression
                        | expression POTENCIA expression
                        | expression MODULO expression
                        | expression MAYQUE expression
                        | expression MENQUE expression
                        | expression MENIGUALQUE expression
                        | expression MAYIGUALQUE expression
                        | expression IGUALQUE expression
                        | expression NIGUALQUE expression
                        | expression OR expression
                        | expression AND expression
                        | final_expression'''
    if len(t) == 2:
        t[0] = t[1]
    elif len(t) == 3:
        if t[1] == "-":
            # todo esto tiene que cambiar dependiendo del tipo del que me manden
            value = Literal(0, Type.INT, t.lineno(1), t.lexpos(0))
            t[0] = Arithmetic(value, t[2], ArithmeticEnum.MINUS, value.line, value.column)
        else:
            # todo tengo que hacer las logicas
            pass
    else:
        if t[2] == "+":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.PLUS, t.lineno(1), t.lexpos(0))
        elif t[2] == "-":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.MINUS, t.lineno(1), t.lexpos(0))
        elif t[2] == "*":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.TIMES, t.lineno(1), t.lexpos(0))
        elif t[2] == "/":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.DIV, t.lineno(1), t.lexpos(0))


# todo falta potencia y modulo
def p_final_expression(t):
    '''final_expression     : PARIZQ expression PARDER
                            | DECIMAL
                            | ENTERO
                            | CADENA
                            | ID
                            | TRUE
                            | FALSE
                            | call_func'''
    if len(t) == 2:
        # todo si no jala poner types
        if t.slice[1].type == "ENTERO":
            t[0] = Literal(int(t[1]), Type.INT, t.lineno(1), t.lexpos(0))
    else:
        t[0] = t[2]


def p_call_func(t):
    '''call_func        : ID PARIZQ PARDER
                        | ID PARIZQ exp_list PARDER'''


def p_exp_list(t):
    '''exp_list         : exp_list COMA expression
                        | expression'''


def p_print_instr(t):
    'print_instr    : PRINT PARIZQ expression PARDER PTCOMA'
    t[0] = Print(t[3], t.lineno(1), t.lexpos(0))


def p_println_instr(t):
    'println_instr  : PRINTLN PARIZQ expression PARDER PTCOMA'
    t[0] = Print(t[3], t.lineno(1), t.lexpos(0), 1)


def p_tipo(t):
    '''tipo     : INT
                | FLOAT
                | BOOL
                | CHAR
                | STRING'''


def p_definicion_instr(t):
    '''definicion_instr   :  LOCAL ID PTCOMA
                            | GLOBAL ID PTCOMA'''


def p_asignacion_instr(t):
    '''asignacion_instr   : ID IGUAL expression PTCOMA
                            | LOCAL ID IGUAL expression PTCOMA
                            | GLOBAL ID IGUAL expression PTCOMA'''


def p_definicion_asignacion_instr(t):
    '''definicion_asignacion_instr  : ID IGUAL expression DOSP DOSP tipo PTCOMA
                                    | LOCAL ID IGUAL expression DOSP DOSP tipo PTCOMA
                                    | GLOBAL ID IGUAL expression DOSP DOSP tipo PTCOMA'''


def p_error(t):
    print(t)
    print("Error sintáctico en '%s'" % t.value)


import ply.yacc as yacc

parser = yacc.yacc()


def parse():
    f = open("D:\\usac\\Compi2\\OLC2_Proyecto1\\server\\analizadores\\pruebas.jl", "r")
    input = f.read()
    print(input)
    # todo esto lo tengo que cambiar para jalarlo en el endpoint
    parserrr = parser.parse(input)
    return parser.parse(input)
