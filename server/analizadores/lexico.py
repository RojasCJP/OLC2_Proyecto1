import ply.lex as lex
reservadas = {
    'numero': 'NUMERO',
    'println': 'PRINTLN',
    'print': 'PRINT',
    'while': 'WHILE',
    'if': 'IF',
    'else': 'ELSE',
    'nothing': 'NOTHING',
    'int64': 'INT',
    'float64': 'FLOAT',
    'bool': 'BOOL',
    'char': 'CHAR',
    'string': 'STRING',
    'uppercase': 'UPERCASE',
    'lowercase': 'LOWERCASE',
    'log10': 'LOG10',
    'log': 'LOG',
    'sin': 'SIN',
    'cos': 'COS',
    'tan': 'TAN',
    'sqrt': 'SQRT',
    'true': 'TRUE',
    'false': 'FALSE',
    'local': 'LOCAL',
    'global': 'GLOBAL'
}

tokens = [
    'PTCOMA',
    'DOSP',
    'PUNTO',
    'COMA',
    'LLAVIZQ',
    'LLAVDER',
    'CORCHETEIZQ',
    'CORCHETEDER',
    'PARIZQ',
    'PARDER',
    'IGUAL',
    'MAS',
    'MENOS',
    'POR',
    'DIVIDIDO',
    'POTENCIA',
    'MODULO',
    'CONCAT',
    'MENQUE',
    'MENIGUALQUE',
    'MAYQUE',
    'MAYIGUALQUE',
    'OR',
    'AND',
    'NOT',
    'IGUALQUE',
    'NIGUALQUE',
    'DECIMAL',
    'ENTERO',
    'CADENA',
    'ID'
] + list(reservadas.values())

# Tokens
t_PTCOMA = r';'
t_DOSP = r':'
t_PUNTO = r'.'
t_COMA = r','
t_LLAVIZQ = r'{'
t_LLAVDER = r'}'
t_CORCHETEIZQ = r'['
t_CORCHETEDER = r']'
t_PARIZQ = r'\('
t_PARDER = r'\)'
t_IGUAL = r'='
t_MAS = r'\+'
t_MENOS = r'-'
t_POR = r'\*'
t_DIVIDIDO = r'/'
t_POTENCIA = r'^'
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
    r'#=(.|\n)*?=#'
    t.lexer.lineno += t.value.count('\n')

# Comentario simple // ...


def t_COMENTARIO_SIMPLE(t):
    r'#.*\n'
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
