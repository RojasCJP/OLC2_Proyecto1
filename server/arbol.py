from interprete.comandos.abstracts.instruccion import Instruction
import ply.yacc as yacc
import ply.lex as lex

from interprete.comandos.node.node import *
reservadas = {
    "println": "PRINTLN",
    "print": "PRINT",
    "while": "WHILE",
    "for": "FOR",
    "if": "IF",
    "in": "IN",
    "elseif": "ELSEIF",
    "else": "ELSE",
    "nothing": "NOTHING",
    "Int64": "INT",
    "Float64": "FLOAT",
    "Bool": "BOOL",
    "Char": "CHAR",
    "String": "STRING",
    "struct": "STRUCT",
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
    "length": "LENGTH",
    "push": "PUSH",
    "pop": "POP",
    "function": "FUNCTION",
    "break": "BREAK",
    "continue": "CONTINUE",
    "return": "RETURN",
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
    "CARACTER",
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


def t_CARACTER(t):
    r'\'.\''
    t.value = t.value[1:-1]
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
    initial_node = Node("incial")
    initial_node.insert_child(t[1])
    t[0] = initial_node.get_graph_ast()
    file = open("./salida.txt", "w")
    file.write(initial_node.get_graph_ast())


def p_instrucciones_lista(t):
    '''instrucciones    : instrucciones instruccion
                        | instruccion'''
    instructions_node = Node("instrucciones")
    if len(t) != 2:
        instructions_node.insert_child(t[2])
        instructions_node.insert_child(t[1])
        t[0] = instructions_node
    else:
        instructions_node.insert_child(t[1])
        t[0] = instructions_node


def p_instruccion(t):
    '''instruccion      : print_instr PTCOMA
                        | println_instr PTCOMA
                        | definicion_instr PTCOMA
                        | asignacion_instr PTCOMA
                        | asignacion_arreglo_instr PTCOMA
                        | definicion_asignacion_instr PTCOMA
                        | call_function PTCOMA
                        | declare_function PTCOMA
                        | return_state PTCOMA
                        | break_state PTCOMA
                        | continue_state PTCOMA
                        | if_state PTCOMA
                        | while_state PTCOMA
                        | for_state PTCOMA
                        | nativas PTCOMA
                        | create_struct PTCOMA
                        | declare_struct PTCOMA
                        | assign_access PTCOMA'''
    instruction_node = Node("instruccion")
    instruction_node.insert_child(t[1])
    t[0] = instruction_node


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
        instruction_node = Node("expresion_final")
        instruction_node.insert_child(t[1])
        t[0] = instruction_node
    elif len(t) == 3:
        if t[1] == "-":
            # todo esto tiene que cambiar dependiendo del tipo del que me manden
            instruction_node = Node("negativo")
            instruction_node.insert_child(Node("-"))
            instruction_node.insert_child(t[2])
            t[0] = instruction_node
        else:
            instruction_node = Node("negado")
            instruction_node.insert_child(Node("!"))
            instruction_node.insert_child(t[2])
            t[0] = instruction_node
    else:
        instruction_node = Node("operacion")
        instruction_node.insert_child(t[1])
        instruction_node.insert_child(Node(t[2]))
        instruction_node.insert_child(t[3])
        t[0] = instruction_node


# todo falta potencia y modulo
def p_final_expression(t):
    '''final_expression     : PARIZQ expression PARDER
                            | CORCHETEIZQ exp_list CORCHETEDER
                            | DECIMAL
                            | ENTERO
                            | CADENA
                            | CARACTER
                            | ID
                            | ID index_list
                            | TRUE
                            | FALSE
                            | call_function
                            | nativas'''
    if len(t) == 2:

        if t.slice[1].type == "ID":
            nodeExpVariable = Node("variable")
            nodeExpVariable.insert_child(Node(t[1]))
            t[0] = nodeExpVariable
        elif t.slice[1].type == "call_function":
            t[0] = t[1]
        elif t.slice[1].type == "nativas":
            t[0] = t[1]
        else:
            t[0] = Node(t[1])
    elif len(t) == 3:
        nodeExpVariable = Node("lista")
        nodeExpVariable.insert_child(Node(t[1]))
        nodeExpVariable.insert_child(t[2])
        t[0] = nodeExpVariable
    else:
        if t.slice[1].type == "PARIZQ":
            nodeExpAgrupacion = Node("parentesis")
            nodeExpAgrupacion.insert_child(Node("("))
            nodeExpAgrupacion.insert_child(t[2])
            nodeExpAgrupacion.insert_child(Node(")"))
            t[0] = nodeExpAgrupacion
        else:
            nodeExpArray = Node("lista")
            nodeExpArray.insert_child(Node("["))
            nodeExpArray.insert_child(t[2])
            nodeExpArray.insert_child(Node("]"))
            t[0] = nodeExpArray


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
                        | LENGTH PARIZQ expression PARDER
                        | PUSH PARIZQ expression COMA expression PARDER
                        | POP PARIZQ expression PARDER
                        '''
    nativa = Node("nativa")
    if t.slice[1].type == "LOG":
        nativa.insert_child(Node("log"))
        nativa.insert_child(Node("("))
        nativa.insert_child(Node(t[3]))
        nativa.insert_child(Node(","))
        nativa.insert_child(t[5])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "LOG10":
        nativa.insert_child(Node("log10"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "COS":
        nativa.insert_child(Node("cos"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "SIN":
        nativa.insert_child(Node("sin"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "TAN":
        nativa.insert_child(Node("tan"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "SQRT":
        nativa.insert_child(Node("sqrt"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "UPPERCASE":
        nativa.insert_child(Node("uppercase"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "LOWERCASE":
        nativa.insert_child(Node("lowercase"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "TOSTRING":
        nativa.insert_child(Node("tostring"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "TOFLOAT":
        nativa.insert_child(Node("tofloat"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "TRUNC":
        nativa.insert_child(Node("log"))
        nativa.insert_child(Node("("))
        nativa.insert_child(Node(t[3]))
        nativa.insert_child(Node(","))
        nativa.insert_child(t[5])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "TYPEOF":
        nativa.insert_child(Node("type"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "PARSE":
        nativa.insert_child(Node("parse"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(","))
        nativa.insert_child(t[5])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "LENGTH":
        nativa.insert_child(Node("length"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "PUSH":
        nativa.insert_child(Node("push"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(","))
        nativa.insert_child(t[5])
        nativa.insert_child(Node(")"))
    elif t.slice[1].type == "POP":
        nativa.insert_child(Node("pop"))
        nativa.insert_child(Node("("))
        nativa.insert_child(t[3])
        nativa.insert_child(Node(")"))
    t[0] = nativa


def p_print_instr(t):
    'print_instr    : PRINT PARIZQ exp_list PARDER'
    print_instr = Node("print")
    print_instr.insert_child(Node("("))
    print_instr.insert_child(t[3])
    print_instr.insert_child(Node(")"))
    t[0] = print_instr


def p_println_instr(t):
    'println_instr  : PRINTLN PARIZQ exp_list PARDER'
    print_instr = Node("println")
    print_instr.insert_child(Node("("))
    print_instr.insert_child(t[3])
    print_instr.insert_child(Node(")"))
    t[0] = print_instr


def p_tipo(t):
    '''tipo     : INT
                | FLOAT
                | BOOL
                | CHAR
                | STRING
    '''
    nodo_tipo = Node("tipo")
    if t.slice[1].type == "INT":
        nodo_tipo.insert_child(Node("int"))
    elif t.slice[1].type == "FLOAT":
        nodo_tipo.insert_child(Node("float"))
    elif t.slice[1].type == "BOOL":
        nodo_tipo.insert_child(Node("bool"))
    elif t.slice[1].type == "CHAR":
        nodo_tipo.insert_child(Node("char"))
    elif t.slice[1].type == "STRING":
        nodo_tipo.insert_child(Node("string"))
    t[0] = nodo_tipo


def p_definicion_instr(t):
    '''definicion_instr   :  LOCAL ID
                            | GLOBAL ID'''
    declaracion_nodo = Node("declaracion")
    if t.slice[1].type == "LOCAL":
        declaracion_nodo.insert_child(Node("local"))
        declaracion_nodo.insert_child(Node(t[2]))
    elif t.slice[1].type == "GLOBAL":
        declaracion_nodo.insert_child(Node("global"))
        declaracion_nodo.insert_child(Node(t[2]))
    t[0] = declaracion_nodo


def p_asignacion_instr(t):
    '''asignacion_instr     : ID IGUAL expression
                            | LOCAL ID IGUAL expression
                            | GLOBAL ID IGUAL expression'''
    asignacion_nodo = Node("asignacion")
    if len(t) == 4:
        asignacion_nodo.insert_child(Node(t[1]))
        asignacion_nodo.insert_child(Node("="))
        asignacion_nodo.insert_child(t[3])
    else:
        if t.slice[1].type == "LOCAL":
            asignacion_nodo.insert_child(Node("local"))
            asignacion_nodo.insert_child(Node(t[2]))
            asignacion_nodo.insert_child(Node("="))
            asignacion_nodo.insert_child(t[4])
        else:
            asignacion_nodo.insert_child(Node("global"))
            asignacion_nodo.insert_child(Node(t[2]))
            asignacion_nodo.insert_child(Node("="))
            asignacion_nodo.insert_child(t[4])
    t[0] = asignacion_nodo


def p_definicion_asignacion_instr(t):
    '''definicion_asignacion_instr  : ID IGUAL expression DOSP DOSP tipo
                                    | LOCAL ID IGUAL expression DOSP DOSP tipo
                                    | GLOBAL ID IGUAL expression DOSP DOSP tipo'''
    declaracion_nodo = Node("declaracion")
    if len(t) == 7:
        declaracion_nodo.insert_child(Node(t[1]))
        declaracion_nodo.insert_child(Node("="))
        declaracion_nodo.insert_child(t[3])
        declaracion_nodo.insert_child(Node("::"))
        declaracion_nodo.insert_child(t[6])
    else:
        if t.slice[1].type == "LOCAL":
            declaracion_nodo.insert_child(Node("local"))
            declaracion_nodo.insert_child(Node(t[2]))
            declaracion_nodo.insert_child(Node("="))
            declaracion_nodo.insert_child(t[4])
            declaracion_nodo.insert_child(Node("::"))
            declaracion_nodo.insert_child(t[7])
        elif t.slice[1].type == "GLOBAL":
            declaracion_nodo.insert_child(Node("global"))
            declaracion_nodo.insert_child(Node(t[2]))
            declaracion_nodo.insert_child(Node("="))
            declaracion_nodo.insert_child(t[4])
            declaracion_nodo.insert_child(Node("::"))
            declaracion_nodo.insert_child(t[7])
    t[0] = declaracion_nodo


def p_asignacion_arreglo_instr(t):
    '''asignacion_arreglo_instr     : ID index_list IGUAL expression'''
    asignacion = Node("asignacion_arreglo")
    asignacion.insert_child(Node(t[1]))
    asignacion.insert_child(t[2])
    asignacion.insert_child(Node("="))
    asignacion.insert_child(t[4])
    t[0] = asignacion


# TODO tengo que hacer que se asignar y declarar los arreglos con index

def p_call_function_instr(t):
    '''call_function    : ID PARIZQ PARDER
                        | ID PARIZQ exp_list PARDER'''
    call_func = Node("llamda_funcion")
    if len(t) == 4:
        call_func.insert_child(Node(t[1]))
        call_func.insert_child(Node("("))
        call_func.insert_child(Node(")"))
    else:
        call_func.insert_child(Node(t[1]))
        call_func.insert_child(Node("("))
        call_func.insert_child(t[3])
        call_func.insert_child(Node(")"))
    t[0] = call_func


def p_exp_list_instr(t):
    '''exp_list         : exp_list COMA expression
                        | expression'''
    exp_list = Node("exp_list")
    if len(t) == 2:
        exp_list.insert_child(t[1])
    else:
        exp_list.insert_child(t[1])
        exp_list.insert_child(Node(","))
        exp_list.insert_child(t[3])
    t[0] = exp_list


def p_index_list_instr(t):
    '''index_list       : index_list CORCHETEIZQ expression CORCHETEDER
                        | CORCHETEIZQ expression CORCHETEDER'''
    index_list = Node("index_list")
    if len(t) == 4:
        index_list.insert_child(Node("["))
        index_list.insert_child(t[2])
        index_list.insert_child(Node("]"))
    else:
        index_list.insert_child(t[1])
        index_list.insert_child(Node("["))
        index_list.insert_child(t[3])
        index_list.insert_child(Node("]"))
    t[0] = index_list


def p_statement(t):
    '''statement        : instrucciones'''
    statement = Node("instrucciones")
    statement.insert_child(t[1])
    t[0] = statement


def p_declare_function(t):
    '''declare_function     : FUNCTION ID PARIZQ PARDER statement END
                            | FUNCTION ID PARIZQ dec_params PARDER statement END'''
    func_declare = Node("declaracion_funcion")
    if len(t) == 7:
        func_declare.insert_child(Node("function"))
        func_declare.insert_child(Node(t[2]))
        func_declare.insert_child(Node("("))
        func_declare.insert_child(Node(")"))
        func_declare.insert_child(t[5])
        func_declare.insert_child(Node("end"))
    else:
        func_declare.insert_child(Node("function"))
        func_declare.insert_child(Node(t[2]))
        func_declare.insert_child(Node("("))
        func_declare.insert_child(t[4])
        func_declare.insert_child(Node(")"))
        func_declare.insert_child(t[6])
        func_declare.insert_child(Node("end"))
    t[0] = func_declare


def p_dec_params(t):
    '''dec_params :    dec_params COMA ID
                    | ID'''
    params = Node("parametros")
    if len(t) == 2:
        params.insert_child(Node(t[1]))
    else:
        params.insert_child(t[1])
        params.insert_child(Node(","))
        params.insert_child(Node(t[3]))
    t[0] = params


def p_if_state(t):
    '''if_state     : IF expression statement END
                    | IF expression statement ELSE statement END
                    | IF expression statement else_if_list END'''
    if_node = Node("if")
    if len(t) == 5:
        if_node.insert_child(Node("if"))
        if_node.insert_child(t[2])
        if_node.insert_child(t[3])
        if_node.insert_child(Node("end"))
    elif len(t) == 6:
        if_node.insert_child(Node("if"))
        if_node.insert_child(t[2])
        if_node.insert_child(t[3])
        if_node.insert_child(t[4])
        if_node.insert_child(Node("end"))
    elif len(t) == 7:
        if_node.insert_child(Node("if"))
        if_node.insert_child(t[2])
        if_node.insert_child(t[3])
        if_node.insert_child(Node("else"))
        if_node.insert_child(t[5])
        if_node.insert_child(Node("end"))
    t[0] = if_node


def p_else_if_list(t):
    '''else_if_list     : ELSEIF expression statement
                        | ELSEIF expression statement ELSE statement
                        | ELSEIF expression statement else_if_list'''
    if_list = Node("else_if")
    if len(t) == 4:
        if_list.insert_child(Node("elseif"))
        if_list.insert_child(t[2])
        if_list.insert_child(t[3])
    elif len(t) == 5:
        if_list.insert_child(Node("elseif"))
        if_list.insert_child(t[2])
        if_list.insert_child(t[3])
        if_list.insert_child(t[4])
    elif len(t) == 6:
        if_list.insert_child(Node("elseif"))
        if_list.insert_child(t[2])
        if_list.insert_child(t[3])
        if_list.insert_child(Node("else"))
        if_list.insert_child(t[5])
    t[0] = if_list


def p_while_state(t):
    '''while_state      : WHILE expression statement END'''
    while_nodo = Node("while")
    while_nodo.insert_child(Node("while"))
    while_nodo.insert_child(t[2])
    while_nodo.insert_child(t[3])
    while_nodo.insert_child(Node("end"))
    t[0] = while_nodo


def p_for_state(t):
    '''for_state        : FOR ID IN expression DOSP expression statement END
                        | FOR ID IN expression statement END'''
    for_nodo = Node("for")
    if len(t) == 9:
        for_nodo.insert_child(Node("for"))
        for_nodo.insert_child(Node(t[2]))
        for_nodo.insert_child(Node("in"))
        for_nodo.insert_child(t[4])
        for_nodo.insert_child(Node(":"))
        for_nodo.insert_child(t[6])
        for_nodo.insert_child(t[7])
        for_nodo.insert_child(Node("end"))
    else:
        for_nodo.insert_child(Node("for"))
        for_nodo.insert_child(Node(t[2]))
        for_nodo.insert_child(Node("in"))
        for_nodo.insert_child(t[4])
        for_nodo.insert_child(t[5])
        for_nodo.insert_child(Node("end"))
    t[0] = for_nodo


def p_break(t):
    '''break_state      : BREAK'''
    t[0] = Node("break")


def p_continue(t):
    '''continue_state      : CONTINUE'''
    t[0] = Node("continue")


def p_return(t):
    '''return_state     : RETURN
                        | RETURN expression'''
    if len(t) == 2:
        t[0] = Node("return")
    else:
        return_nodo = Node("return")
        return_nodo.insert_child(t[2])
        t[0] = return_nodo


def p_createStruct(t):
    'create_struct : STRUCT ID attList END'
    create_node = Node("create_struct")
    create_node.insert_child(Node("struct"))
    create_node.insert_child(Node(t[2]))
    create_node.insert_child(t[3])
    create_node.insert_child(Node("end"))
    t[0] = create_node


def p_attList(t):
    '''attList :  attList ID PTCOMA
                | ID PTCOMA'''
    att_node = Node("list_attributes")
    if len(t) == 3:
        att_node.insert_child(Node(t[1]))
        att_node.insert_child(Node(";"))
    else:
        att_node.insert_child(t[1])
        att_node.insert_child(Node(t[2]))
        att_node.insert_child(Node(";"))
    t[0] = att_node


def p_declareStruct(t):
    'declare_struct : ID DOSP DOSP ID'
    declare_struct_node = Node("declaracion_structs")
    declare_struct_node.insert_child(Node(t[1]))
    declare_struct_node.insert_child(Node(":"))
    declare_struct_node.insert_child(Node(":"))
    declare_struct_node.insert_child(Node(t[4]))
    t[0] = declare_struct_node

# ASSIGN ACCESS


def p_assignAccess(t):
    'assign_access : ID PUNTO ID IGUAL expression'
    asignacion_node = Node("asignacion_struct")
    asignacion_node.insert_child(Node(t[1]))
    asignacion_node.insert_child(Node("."))
    asignacion_node.insert_child(Node(t[3]))
    asignacion_node.insert_child(Node("="))
    asignacion_node.insert_child(t[5])
    t[0] = asignacion_node


def p_error(t):
    print(t)
    print("Error sintáctico en '%s'" % t.value)


parser = yacc.yacc()


def arbol(input):
    # f = open("D:\\usac\\Compi2\\OLC2_Proyecto1\\server\\analizadores\\pruebas.jl", "r")
    # input = f.read()
    # todo esto lo tengo que cambiar para jalarlo en el endpoint
    parser.parse(input)
    return parser.parse(input)
