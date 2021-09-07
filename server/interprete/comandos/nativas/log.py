from ..abstracts.instruccion import *
from ..expressions.literal import *
from math import log


class Log(Instruction):

    def __init__(self, value, base, line, column):
        Instruction.__init__(self, line, column)
        self.value = value
        self.base = base

    def execute(self, env: Environment):
        value = self.value.execute(env)
        result = Literal(0, Type.FLOAT, self.line, self.column)
        if value is not None:
            result.value = log(value.value, self.base)
        return result
