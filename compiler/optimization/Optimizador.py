from optimization.Expressions.Access import *
from optimization.Expressions.Expression import *
from optimization.Expressions.Literal import *

from optimization.Gotos.Goto import *
from optimization.Gotos.If import *

from optimization.Instructions.Assignment import *
from optimization.Instructions.CallFun import *
from optimization.Instructions.Function import *
from optimization.Instructions.Label import *
from optimization.Instructions.Print import *
from optimization.Instructions.Return import *


class Optimizador():
    def __init__(self, packages, temps, code):
        self.packages = packages
        self.temps = temps
        self.code = code

    def get_code(self):
        ret = f'package main;\n\nimport(\n\t"{self.packages}"\n);\n'
        for temp in self.temps:
            ret = ret + f'var {temp}\n'
        ret = ret + '\n'

        for func in self.code:
            ret = ret + func.get_code()+'\n\n'
        return ret

    def Mirilla(self):
        for func in self.code:
            tamano = 10
            while tamano <= len(func.instr):
                flag_opt = False
                for i in range(5):
                    aux = 0
                    while (tamano+aux) <= len(func.instr):
                        flag_opt = flag_opt or self.Regla3(
                            func.instr[0+aux:tamano+aux])
                        flag_opt = flag_opt or self.Regla6(
                            func.instr[0+aux:tamano+aux])
                        aux = aux+1

    def Regla3(self, arreglo):
        ret = False
        for i in range(len(arreglo)):
            actual = arreglo[i]
            if type(actual) is If and not actual.deleted:
                next_inst = arreglo[i+1]
                if type(next_inst) is Goto and not next_inst.deleted:
                    actual.condition.get_contrary()
                    actual.label = next_inst.label
                    next_inst.deleted = True
                    arreglo[i+2].deleted = True
                    ret = True
        return ret

    def Regla6(self, arreglo):
        ret = False
        for i in range(len(arreglo)):
            actual = arreglo[i]
            if type(actual) is Assigment and not actual.deleted:
                if(actual.self_assigment()):
                    actual_opt = actual.exp.neutral_ops()
                    if actual_opt:
                        ret = True
                        actual.deleted = True
        return ret
