from .Environment import *


class Generator:
    generator = None

    def __init__(self):
        self.count_temp = 0
        self.count_label = 0
        self.code = ''
        self.funcs = ''
        self.natives = ''
        self.in_func = False
        self.in_natives = False
        self.temps = []
        self.print_string = False

    def clean_all(self):
        self.count_temp = 0
        self.count_label = 0
        self.code = ''
        self.funcs = ''
        self.natives = ''
        self.in_func = False
        self.in_natives = False
        self.temps = []
        self.print_string = False
        Generator.generator = Generator()

    def get_header(self):
        ret = '/*----HEADER----*/\npackage main;\n\nimport (\n\t"fmt"\n)\n\n'
        if len(self.temps) > 0:
            ret += 'var '
            for temp in range(len(self.temps)):
                ret += self.temps[temp]
                if temp != (len(self.temps) - 1):
                    ret += ", "
            ret += " float64\n"
        ret += "var P, H float64;\nvar stack [30101999]float64;\nvar heap [30101999]float64;\n\n"
        return ret

    def get_code(self):
        return f'{self.getHeader()}{self.natives}\n{self.funcs}\nfunc main(){{\n{self.code}\n}}'

    def code_in(self, code, tab="\t"):
        if(self.in_natives):
            if(self.natives == ''):
                self.natives = self.natives + '/*-----NATIVES-----*/\n'
            self.natives = self.natives + tab + code
        elif(self.in_func):
            if(self.funcs == ''):
                self.funcs = self.funcs + '/*-----FUNCS-----*/\n'
            self.funcs = self.funcs + tab + code
        else:
            self.code = self.code + '\t' + code

    def add_comment(self, comment):
        self.codeIn(f'/* {comment} */\n')

    def get_instance(self):
        if Generator.generator == None:
            Generator.generator = Generator()
        return Generator.generator

    def add_space(self):
        self.codeIn("\n")

    # Manejo de Temporales
    def add_temp(self):
        temp = f't{self.countTemp}'
        self.countTemp += 1
        self.temps.append(temp)
        return temp

    # Manejo de Labels
    def new_label(self):
        label = f'L{self.countLabel}'
        self.countLabel += 1
        return label

    def put_label(self, label):
        self.codeIn(f'{label}:\n')

    # GOTO
    def add_goto(self, label):
        self.codeIn(f'goto {label};\n')

    # IF
    def add_if(self, left, right, op, label):
        self.codeIn(f'if {left} {op} {right} {{goto {label};}}\n')

    # EXPRESIONES
    def add_expression(self, result, left, right, op):
        self.codeIn(f'{result}={left}{op}{right};\n')

    # FUNCS
    def add_begin_func(self, id):
        if(not self.in_natives):
            self.in_func = True
        self.codeIn(f'func {id}(){{\n', '')

    def add_end_func(self):
        self.codeIn('return;\n}\n')
        if(not self.in_natives):
            self.in_func = False

    # STACK
    def set_stack(self, pos, value):
        self.codeIn(f'stack[int({pos})]={value};\n')

    def get_stack(self, place, pos):
        self.codeIn(f'{place}=stack[int({pos})];\n')

    # ENVS
    def new_env(self, size):
        self.codeIn(f'P=P+{size};\n')

    def call_fun(self, id):
        self.codeIn(f'{id}();\n')

    def ret_env(self, size):
        self.codeIn(f'P=P-{size};\n')

    # HEAP
    def set_heap(self, pos, value):
        self.codeIn(f'heap[int({pos})]={value};\n')

    def get_heap(self, place, pos):
        self.codeIn(f'{place}=heap[int({pos})];\n')

    def next_heap(self):
        self.codeIn('H=H+1;\n')

    # INSTRUCCIONES
    def add_print(self, type, value):
        self.codeIn(f'fmt.Printf("%{type}", int({value}));\n')

    def print_true(self):
        self.addPrint("c", 116)
        self.addPrint("c", 114)
        self.addPrint("c", 117)
        self.addPrint("c", 101)

    def print_false(self):
        self.addPrint("c", 102)
        self.addPrint("c", 97)
        self.addPrint("c", 108)
        self.addPrint("c", 115)
        self.addPrint("c", 101)

    # NATIVES
    def fprint_string(self):
        if(self.print_string):
            return
        self.print_string = True
        self.in_natives = True

        self.add_begin_func('print_string')
        # Label para salir de la funcion
        returnLbl = self.new_label()
        # Label para la comparacion para buscar fin de cadena
        compareLbl = self.new_label()

        # Temporal puntero a Stack
        tempP = self.add_temp()

        # Temporal puntero a Heap
        tempH = self.add_temp()

        self.add_expression(tempP, 'P', '1', '+')

        self.get_stack(tempH, tempP)

        # Temporal para comparark
        tempC = self.add_temp()

        self.put_label(compareLbl)

        self.get_heap(tempC, tempH)

        self.add_if(tempC, '-1', '==', returnLbl)

        self.add_print('c', tempC)

        self.add_expression(tempH, tempH, '1', '+')

        self.add_goto(compareLbl)

        self.put_label(returnLbl)
        self.add_end_func()
        self.in_natives = False
