from ..abstracts.expression import *
from ..abstracts.returns import *


class Literal(Expression):
    def __init__(self, value, type, line, column):
        Expression.__init__(self, line, column)
        self.value = value
        self.type = type

    def execute(self, environment: Environment):
        return Return(self.value, self.type)
