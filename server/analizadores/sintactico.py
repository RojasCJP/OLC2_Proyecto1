import ply.yacc as yacc
from lexico import *

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


precedence = (
    ('left', 'CONCAT'),
    ('left', 'MAS', 'MENOS'),
    ('left', 'POR', 'DIVIDIDO'),
    ('left', 'POTENCIA', 'MODULO'),
    ('right', 'UMENOS'),
)

# Definición de la gramática


def p_init(t):
    'init            : instrucciones'
    t[0] = t[1]


def p_instrucciones_lista(t):
    'instrucciones    : instrucciones instruccion'
    t[1].append(t[2])
    t[0] = t[1]


def p_instrucciones_instruccion(t):
    'instrucciones    : instruccion '
    t[0] = [t[1]]


def p_instruccion(t):
    '''instruccion      : print_instr
                        | println_instr
                        | definicion_instr
                        | asignacion_instr
                        | definicion_asignacion_instr
                        | while_instr
                        | if_instr
                        | if_else_instr'''
    t[0] = t[1]


def p_expression(t):
    '''expression       : MENOS expression %prec UMENOS
                        | NOT expression %prec UMENOS
                        | expression MAS expression
                        | expression MENOS expression
                        | expression POR expression
                        | expression DIVIDIDO expression
                        | expression MAYQUE expression
                        | expression MENQUE expression
                        | expression MENIGUALQUE expression
                        | expression MAYIGUALQUE expression
                        | expression IGUALQUE expression
                        | expression NIGUALQUE expression
                        | expression OR expression
                        | expression AND expression
                        | final_expression'''

#todo falta potencia y modulo
def p_final_expression(t):
    '''final_expression     : PARIZQ expression PARDER
                            | DECIMAL
                            | ENTERO
                            | CADENA
                            | ID
                            | TRUE
                            | FALSE
                            | call_func'''

def p_call_func(t):
    '''call_func        : ID PARIZQ PARDER
                        | ID PARIZQ exp_list PARDER'''

def p_exp_list(t):
    '''exp_list         : exp_list COMA expression
                        | expression'''

def p_print_instr(t):
    'print_instr     : PRINT PARIZQ expression PARDER PTCOMA'


def p_println_instr(t):
    'println_instr  : PRINTLN PARIZQ expression PARDER PTCOMA'


def p_tipo(t):
    '''tipo: INT
            |FLOAT
            |BOOL
            |CHAR
            |STRING'''


def p_definicion_instr(t):
    '''definicion_instr   :  LOCAL ID PTCOMA
                            |GLOBAL ID PTCOMA'''


def p_asignacion_instr(t):
    '''asignacion_instr   : ID IGUAL expression PTCOMA
                            |LOCAL ID IGUAL expression PTCOMA
                            |GLOBAL ID IGUAL expression PTCOMA'''


def p_definicion_asignacion_instr(t):
    '''definicion_asignacion_instr  : ID IGUAL expression DOSP DOSP tipo PTCOMA
                                    |LOCAL ID IGUAL expression DOSP DOSP tipo PTCOMA
                                    |GLOBAL ID IGUAL expression DOSP DOSP tipo PTCOMA'''


def p_mientras_instr(t):
    'mientras_instr     : MIENTRAS PARIZQ expresion_logica PARDER LLAVIZQ instrucciones LLAVDER'


def p_if_instr(t):
    'if_instr           : IF PARIZQ expresion_logica PARDER LLAVIZQ instrucciones LLAVDER'


def p_if_else_instr(t):
    'if_else_instr      : IF PARIZQ expresion_logica PARDER LLAVIZQ instrucciones LLAVDER ELSE LLAVIZQ instrucciones LLAVDER'


def p_expresion_binaria(t):
    '''expresion_numerica : expresion_numerica MAS expresion_numerica
                        | expresion_numerica MENOS expresion_numerica
                        | expresion_numerica POR expresion_numerica
                        | expresion_numerica DIVIDIDO expresion_numerica'''


def p_expresion_unaria(t):
    'expresion_numerica : MENOS expresion_numerica %prec UMENOS'


def p_expresion_agrupacion(t):
    'expresion_numerica : PARIZQ expresion_numerica PARDER'


def p_expresion_number(t):
    '''expresion_numerica : ENTERO
                        | DECIMAL'''


def p_expresion_id(t):
    'expresion_numerica   : ID'


def p_expresion_concatenacion(t):
    'expresion_cadena     : expresion_cadena CONCAT expresion_cadena'


def p_expresion_cadena(t):
    'expresion_cadena     : CADENA'


def p_expresion_cadena_numerico(t):
    'expresion_cadena     : expresion_numerica'


def p_expresion_logica(t):
    '''expresion_logica : expresion_numerica MAYQUE expresion_numerica
                        | expresion_numerica MENQUE expresion_numerica
                        | expresion_numerica IGUALQUE expresion_numerica
                        | expresion_numerica NIGUALQUE expresion_numerica'''


def p_error(t):
    print(t)
    print("Error sintáctico en '%s'" % t.value)


parser = yacc.yacc()


def parse(input):
    return parser.parse(input)
