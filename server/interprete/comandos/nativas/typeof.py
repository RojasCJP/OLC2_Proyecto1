from ..abstracts.instruccion import *
from ..expressions.literal import *


class TypeOf(Instruction):

    def __init__(self, value, line, column):
        Instruction.__init__(self, line, column)
        self.value = value

    def execute(self, env: Environment):
        value = self.value.execute(env)
        while not isinstance(value, Return):
            value = value.exec(env)
        result = Return(Type.UNDEFINED, Type.UNDEFINED)
        if value is not None:
            result.value = value.type
        return result
