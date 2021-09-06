from ..abstracts.instruccion import *
from ..abstracts.returns import *


class Declaration(Instruction):
    def __init__(self, idd, value, line, column, globall):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.value = value
        self.globall = globall

    def execute(self, environment):
        value = self.value
        # todo aqui tengo que ver que pedo con los tipos
        if value.type == Type.STRUCT:
            environment.save_var_struct(self.id, value.value, value.auxType)
        else:
            if self.globall:
                environment.get_global_env().save_var(self.id, value.value, value.type)
            else:
                environment.save_var(self.id, value.value, value.type)
