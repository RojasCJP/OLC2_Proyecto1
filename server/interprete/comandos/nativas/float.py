from ..abstracts.instruccion import *
from ..expressions.literal import *


class ToFloat(Instruction):

    def __init__(self, value, line, column):
        Instruction.__init__(self, line, column)
        self.value = value

    def execute(self, env: Environment):
        value = self.value.execute(env)
        result = Literal(0, Type.FLOAT, self.line, self.column)
        if value is not None:
            result.value = float(value.value)
        return result