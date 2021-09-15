from ..abstracts.instruccion import *
from ..expressions.literal import *
from ..expressions.access import *


class Push(Instruction):

    def __init__(self, value, idd, line, column):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.value = value

    def execute(self, env: Environment):
        value = self.value.execute(env)
        result = Literal(False, Type.BOOL, self.line, self.column)
        if value is not None:
            retorno_arreglo = self.id.execute(env)
            retorno_arreglo.value.append(value.value)
            result.value = True
        return None
