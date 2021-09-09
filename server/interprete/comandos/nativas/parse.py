from ..abstracts.instruccion import *
from ..expressions.literal import *


class Parse(Instruction):

    def __init__(self, value, types, line, column):
        Instruction.__init__(self, line, column)
        self.type = types
        self.value = value

    def execute(self, env: Environment):
        value = self.value.execute(env)
        result = Literal(0, Type.STRING, self.line, self.column)
        if value is not None:
            if self.type == Type.STRING:
                result.value = str(value.value)
                result.type = Type.STRING
            if self.type == Type.FLOAT:
                result.value = float(value.value)
                result.type = Type.FLOAT
            if self.type == Type.INT:
                result.value = int(value.value)
                result.type = Type.INT
        return result
