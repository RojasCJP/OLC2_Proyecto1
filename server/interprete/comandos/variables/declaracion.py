from ..abstracts.instruccion import *
from ..abstracts.returns import *


class Declaration(Instruction):
    def __init__(self, idd, value, line, column):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.value = value

    def execute(self, environment):
        value = self.value
        if value.type == Type.STRUCT:
            environment.save_var_struct(self.id, value.value, value.auxType)
        else:
            environment.save_var(self.id, value.value, value.type)
