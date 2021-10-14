from abstract.Expression import *
from abstract.Return import *
from sym.Generator import *
from enum import Enum
import uuid


class ArithmethicOption(Enum):
    PLUS = 0
    MINUS = 1
    TIMES = 2
    DIV = 3
    RAISED = 4
    MODULE = 5


class Arithmetic(Expression):
    def __init__(self, left, right, type, line, column):
        Expression.__init__(self, line, column)
        self.left = left
        self.right = right
        self.type = type

    def compile(self, env):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        left_value = self.left.complie(env)
        right_value = self.right.compile(env)

        temp = generator.add_temp()
        op = ''
        if self.type == ArithmethicOption.PLUS:
            op = '+'
        elif self.type == ArithmethicOption.MINUS:
            op = '-'
        elif self.type == ArithmethicOption.TIMES:
            op = '*'
        elif self.type == ArithmethicOption.DIV:
            op = '/'
        generator.add_expression(temp, left_value.value, right_value.value, op)
        return Return(temp, Type.INT, True)
