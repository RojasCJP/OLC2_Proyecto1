import ply.yacc as yacc
import ply.lex as lex

from ..interprete.comandos.expressions.access import *
from ..interprete.comandos.expressions.access_expression import *
from ..interprete.comandos.expressions.arithmetic import *
from ..interprete.comandos.expressions.logic import *
from ..interprete.comandos.expressions.call_func import *
from ..interprete.comandos.expressions.literal import *
from ..interprete.comandos.expressions.relational import *

from ..interprete.comandos.funciones.param import *
from ..interprete.comandos.funciones.function import *
from ..interprete.comandos.funciones.return_st import *

from ..interprete.comandos.condicionales.ifs import *

from ..interprete.comandos.ciclos.whiles import *

from ..interprete.comandos.ciclos.breaks import *
from ..interprete.comandos.ciclos.continues import *

from ..interprete.comandos.statement import *

from ..interprete.comandos.nativas.print import *
from ..interprete.comandos.nativas.log import *
from ..interprete.comandos.nativas.cos import *
from ..interprete.comandos.nativas.sin import *
from ..interprete.comandos.nativas.tan import *
from ..interprete.comandos.nativas.sqrt import *
from ..interprete.comandos.nativas.uppercase import *
from ..interprete.comandos.nativas.lowercase import *
from ..interprete.comandos.nativas.string import *
from ..interprete.comandos.nativas.float import *
from ..interprete.comandos.nativas.parse import *
from ..interprete.comandos.nativas.trunc import *
from ..interprete.comandos.nativas.typeof import *
from ..interprete.comandos.variables.asignacion import *
from ..interprete.comandos.variables.declaracion import *

reservadas = {
    "numero": "NUMERO",
    "println": "PRINTLN",
    "print": "PRINT",
    "while": "WHILE",
    "if": "IF",
    "elseif": "ELSEIF",
    "else": "ELSE",
    "nothing": "NOTHING",
    "Int64": "INT",
    "Float64": "FLOAT",
    "Bool": "BOOL",
    "Char": "CHAR",
    "String": "STRING",
    "uppercase": "UPPERCASE",
    "lowercase": "LOWERCASE",
    "log10": "LOG10",
    "log": "LOG",
    "sin": "SIN",
    "cos": "COS",
    "tan": "TAN",
    "trunc": "TRUNC",
    "float": "TOFLOAT",
    "string": "TOSTRING",
    "parse": "PARSE",
    "typeof": "TYPEOF",
    "sqrt": "SQRT",
    "function": "FUNCTION",
    "break": "BREAK",
    "continue": "CONTINUE",
    "return":"RETURN",
    "end": "END",
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
t_PUNTO = r'\.'
t_COMA = r','
t_LLAVIZQ = r'{'
t_LLAVDER = r'}'
t_CORCHETEIZQ = r'\['
t_CORCHETEDER = r']'
t_PARIZQ = r'\('
t_PARDER = r'\)'
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
t_IGUAL = r'='


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
    t.type = reservadas.get(t.value, 'ID')
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

lexer = lex.lex()

precedence = (
    ('left', 'OR'),
    ('left', 'AND'),
    ('left', 'IGUALQUE', 'NIGUALQUE'),
    ('left', 'MAYQUE', 'MENQUE', 'MAYIGUALQUE', 'MENIGUALQUE'),
    ('left', 'MAS', 'MENOS'),
    ('left', 'POR', 'DIVIDIDO', 'MODULO'),
    ('left', 'POTENCIA'),
    ('right', 'UMENOS')
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
    '''instruccion      : print_instr PTCOMA
                        | println_instr PTCOMA
                        | definicion_instr PTCOMA
                        | asignacion_instr PTCOMA
                        | definicion_asignacion_instr PTCOMA
                        | call_function PTCOMA
                        | declare_function PTCOMA
                        | return_state PTCOMA
                        | break_state PTCOMA
                        | continue_state PTCOMA
                        | if_state PTCOMA
                        | while_state PTCOMA'''
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
            t[0] = Arithmetic(value, t[2], ArithmeticEnum.MINUS,
                              value.line, value.column)
        else:
            t[0] = Logic(t[2], t[2], LogicEnum.NOT, t.lineno(1), t.lexpos(0))
            pass
    else:
        if t[2] == "+":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.PLUS,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == "-":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.MINUS,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == "*":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.TIMES,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == "/":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.DIV,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == "^":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.RAISED,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == "%":
            t[0] = Arithmetic(t[1], t[3], ArithmeticEnum.MODULE,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == "||":
            t[0] = Logic(t[1], t[3], LogicEnum.OR, t.lineno(1), t.lexpos(0))
        elif t[2] == "&&":
            t[0] = Logic(t[1], t[3], LogicEnum.AND, t.lineno(1), t.lexpos(0))
        elif t[2] == "<":
            t[0] = Relational(t[1], t[3], RelationalEnum.MINOR,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == ">":
            t[0] = Relational(t[1], t[3], RelationalEnum.GREATER,
                              t.lineno(1), t.lexpos(0))
        elif t[2] == "<=":
            t[0] = Relational(
                t[1], t[3], RelationalEnum.MINOR_EQUAL, t.lineno(1), t.lexpos(0))
        elif t[2] == ">=":
            t[0] = Relational(
                t[1], t[3], RelationalEnum.GREATER_EQUAL, t.lineno(1), t.lexpos(0))
        elif t[2] == "==":
            t[0] = Relational(
                t[1], t[3], RelationalEnum.EQUAL_EQUAL, t.lineno(1), t.lexpos(0))
        elif t[2] == "!=":
            t[0] = Relational(
                t[1], t[3], RelationalEnum.DIFFERENT, t.lineno(1), t.lexpos(0))


# todo falta potencia y modulo
def p_final_expression(t):
    '''final_expression     : PARIZQ expression PARDER
                            | DECIMAL
                            | ENTERO
                            | CADENA
                            | ID
                            | TRUE
                            | FALSE
                            | call_func
                            | nativas'''
    if len(t) == 2:
        if t.slice[1].type == "ENTERO":
            t[0] = Literal(int(t[1]), Type.INT, t.lineno(1), t.lexpos(0))
        if t.slice[1].type == "DECIMAL":
            t[0] = Literal(float(t[1]), Type.FLOAT, t.lineno(1), t.lexpos(0))
        elif t.slice[1].type == "FALSE":
            t[0] = Literal(False, Type.BOOL, t.lineno(1), t.lexpos(0))
        elif t.slice[1].type == "TRUE":
            t[0] = Literal(True, Type.BOOL, t.lineno(1), t.lexpos(0))
        elif t.slice[1].type == "CADENA":
            t[0] = Literal(str(t[1]), Type.STRING, t.lineno(1), t.lexpos(0))
        elif t.slice[1].type == "ID":
            t[0] = Access(t[1], t.lineno(1), t.lexpos(0))
        elif t.slice[1].type == "nativas":
            t[0] = t[1]

    else:
        t[0] = t[2]


def p_call_func(t):
    '''call_func        : ID PARIZQ PARDER
                        | ID PARIZQ exp_list PARDER'''
    if t.slice[1].type == "LOG":
        t[0] = Log(t[5], t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "LOG10":
        t[0] = Log(t[3], 10, t.lineno(1), t.lexpos(0))


def p_nativas(t):
    '''nativas          : LOG PARIZQ ENTERO COMA expression PARDER
                        | LOG10 PARIZQ expression PARDER
                        | COS PARIZQ expression PARDER
                        | SIN PARIZQ expression PARDER
                        | TAN PARIZQ expression PARDER
                        | SQRT PARIZQ expression PARDER
                        | UPPERCASE PARIZQ expression PARDER
                        | LOWERCASE PARIZQ expression PARDER
                        | TOSTRING PARIZQ expression PARDER
                        | TOFLOAT PARIZQ expression PARDER
                        | TRUNC PARIZQ INT COMA expression PARDER
                        | TYPEOF PARIZQ expression PARDER
                        | PARSE PARIZQ tipo COMA expression PARDER
                        '''
    if t.slice[1].type == "LOG":
        t[0] = Log(t[5], t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "LOG10":
        t[0] = Log(t[3], 10, t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "COS":
        t[0] = Cos(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "SIN":
        t[0] = Sin(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "TAN":
        t[0] = Tan(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "SQRT":
        t[0] = Sqrt(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "UPPERCASE":
        t[0] = Uppercase(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "LOWERCASE":
        t[0] = Lowercase(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "TOSTRING":
        t[0] = ToString(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "TOFLOAT":
        t[0] = ToFloat(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "TRUNC":
        t[0] = Trunc(t[5], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "TYPEOF":
        t[0] = TypeOf(t[3], t.lineno(1), t.lexpos(0))
    elif t.slice[1].type == "PARSE":
        t[0] = Parse(t[5], t[3], t.lineno(1), t.lexpos(0))


def p_print_instr(t):
    'print_instr    : PRINT PARIZQ expression PARDER'
    t[0] = Print(t[3], t.lineno(1), t.lexpos(0))


def p_println_instr(t):
    'println_instr  : PRINTLN PARIZQ expression PARDER'
    t[0] = Print(t[3], t.lineno(1), t.lexpos(0), 1)


def p_tipo(t):
    '''tipo     : INT
                | FLOAT
                | BOOL
                | CHAR
                | STRING
    '''
    if t.slice[1].type == "INT":
        t[0] = Type.INT
    elif t.slice[1].type == "FLOAT":
        t[0] = Type.FLOAT
    elif t.slice[1].type == "BOOL":
        t[0] = Type.BOOL
    elif t.slice[1].type == "CHAR":
        t[0] = Type.CHAR
    elif t.slice[1].type == "STRING":
        t[0] = Type.STRING


def p_definicion_instr(t):
    '''definicion_instr   :  LOCAL ID
                            | GLOBAL ID'''
    if t.slice[1].type == "LOCAL":
        t[0] = Declaration(t[2], None, t.lineno(1), t.lexpos(0), False)
    elif t.slice[1].type == "GLOBAL":
        t[0] = Declaration(t[2], None, t.lineno(1), t.lexpos(0), True)


def p_asignacion_instr(t):
    '''asignacion_instr   : ID IGUAL expression
                            | LOCAL ID IGUAL expression
                            | GLOBAL ID IGUAL expression'''
    if len(t) == 4:
        t[0] = Asignation(t[1], t[3], t.lineno(1), t.lexpos(0), False)
    else:
        if t.slice[1].type == "LOCAL":
            t[0] = Asignation(t[2], t[4], t.lineno(1), t.lexpos(0), False)
        else:
            t[0] = Asignation(t[2], t[4], t.lineno(1), t.lexpos(0), True)


def p_definicion_asignacion_instr(t):
    '''definicion_asignacion_instr  : ID IGUAL expression DOSP DOSP tipo
                                    | LOCAL ID IGUAL expression DOSP DOSP tipo
                                    | GLOBAL ID IGUAL expression DOSP DOSP tipo'''
    if len(t) == 8:
        t[0] = Declaration(t[2], t[3], t.lineno(1), t.lexpos(0), False)
    else:
        if t.slice[1].type == "LOCAL":
            t[0] = Declaration(t[2], t[4], t.lineno(1), t.lexpos(0), False)
        elif t.slice[1].type == "GLOBAL":
            t[0] = Declaration(t[2], t[4], t.lineno(1), t.lexpos(0), True)


def p_call_function_instr(t):
    '''call_function    : ID PARIZQ PARDER
                        | ID PARIZQ exp_list PARDER'''
    if len(t) == 4:
        t[0] = CallFunc(t[1], [], t.lineno(1), t.lexpos(1))
    else:
        t[0] = CallFunc(t[1], t[3], t.lineno(1), t.lexpos(1))


def p_exp_list_instr(t):
    '''exp_list         : exp_list COMA expression
                        | expression'''
    if len(t) == 2:
        t[0] = [t[1]]
    else:
        t[1].append(t[3])
        t[0] = t[1]


def p_statement(t):
    '''statement        : instrucciones'''
    t[0] = Statement(t[1], t.lineno(1), t.lexpos(0))


def p_declare_function(t):
    '''declare_function     : FUNCTION ID PARIZQ PARDER statement END
                            | FUNCTION ID PARIZQ dec_params PARDER statement END'''
    if len(t) == 7:
        t[0] = Function(t[2], [], t[5], t.lineno(1), t.lexpos(0))
    else:
        t[0] = Function(t[2], t[4], t[6], t.lineno(1), t.lexpos(0))


def p_dec_params(t):
    '''dec_params :    dec_params COMA ID
                    | ID'''
    if len(t) == 2:
        t[0] = [Param(t[1], t.lineno(1), t.lexpos(0))]
    else:
        t[1].append(Param(t[3], t.lineno(1), t.lexpos(0)))
        t[0] = t[1]


def p_if_state(t):
    '''if_state     : IF expression statement END
                    | IF expression statement ELSE statement END
                    | IF expression statement else_if_list END'''
    if len(t) == 5:
        t[0] = If(t[2], t[3], t.lineno(1), t.lexpos(0))
    elif len(t) == 6:
        t[0] = If(t[2], t[3], t.lineno(1), t.lexpos(0), t[4])
    elif len(t) == 7:
        t[0] = If(t[2], t[3], t.lineno(1), t.lexpos(0), t[5])


def p_else_if_list(t):
    '''else_if_list     : ELSEIF expression statement
                        | ELSEIF expression statement ELSE statement
                        | ELSEIF expression statement else_if_list'''
    if len(t) == 4:
        t[0] = If(t[2], t[3], t.lineno(1), t.lexpos(0))
    elif len(t) == 5:
        t[0] = If(t[2], t[3], t.lineno(1), t.lexpos(0), t[4])
    elif len(t) == 6:
        t[0] = If(t[2], t[3], t.lineno(1), t.lexpos(0), t[5])


def p_while_state(t):
    '''while_state      : WHILE expression statement END'''
    t[0] = While(t[2], t[3], t.lineno(1), t.lexpos(0))


def p_break(t):
    '''break_state      : BREAK'''
    t[0] = Break(t.lineno(1), t.lexpos(0))


def p_continue(t):
    '''continue_state      : CONTINUE'''
    t[0] = Continue(t.lineno(1), t.lexpos(0))


def p_return(t):
    '''return_state     : RETURN
                        | RETURN expression'''
    if len(t) == 2:
        t[0] = ReturnST(None,t.lineno(1), t.lexpos(0))
    else:
        t[0] = ReturnST(t[2],t.lineno(1), t.lexpos(0))

def p_error(t):
    print(t)
    print("Error sintáctico en '%s'" % t.value)


parser = yacc.yacc()


def parse(input):
    f = open("D:\\usac\\Compi2\\OLC2_Proyecto1\\server\\analizadores\\pruebas.jl", "r")
    input = f.read()
    # todo esto lo tengo que cambiar para jalarlo en el endpoint
    parser.parse(input)
    return parser.parse(input)
