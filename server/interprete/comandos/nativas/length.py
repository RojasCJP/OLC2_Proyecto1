from ..abstracts.instruccion import *
from ..expressions.literal import *


class Length(Instruction):

    def __init__(self, value, line, column):
        Instruction.__init__(self, line, column)
        self.value = value

    def execute(self, env: Environment):
        value = self.value.execute(env)
        result = Literal(0, Type.INT, self.line, self.column)
        if value is not None and (value.type == Type.ARRAY or value.type == Type.STRING):
            result.value = len(value.value)
        return result
