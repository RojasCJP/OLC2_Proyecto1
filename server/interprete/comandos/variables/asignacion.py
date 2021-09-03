from ..abstracts.instruccion import *
from ..abstracts.returns import *
from .declaracion import *


class Asignation(Instruction):
    def __init__(self, idd, value, line, column):
        Instruction.__init__(self, line, column)
        self.value = value
        self.id = idd

    def execute(self, env: Environment):
        value = self.value
        if self.id in env.variables.keys():
            env.save_var(self.id, value, env.variables[self.id].types)
        else:
            Declaration(self.id, self.value, self.line, self.column)
# todo falta hacer cosas si es un struct
