from ..abstracts.instruccion import *
from ..abstracts.returns import *
from ..expressions.literal import *
from .declaracion import *


class Asignation(Instruction):
    def __init__(self, idd, value, line, column, globall):
        Instruction.__init__(self, line, column)
        self.value = value
        self.id = idd
        self.globall = globall

    def execute(self, env: Environment):
        value = self.value
        while not isinstance(value, Return):
            value = value.execute(env)
        if self.id in env.variables.keys():
            if self.globall:
                env.get_global_env().save_var(self.id, value, env.variables[self.id].type)
            else:
                env.save_var(self.id, value, env.variables[self.id].type)
        else:
            declaration = Declaration(self.id, self.value, self.line, self.column, self.globall)
            declaration.execute(env)
# todo falta hacer cosas si es un struct
