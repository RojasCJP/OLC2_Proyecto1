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
        left_value = self.left.compile(env)
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

        if (self.type == ArithmethicOption.RAISED):
            generator.f_potencia()
            param_temp = generator.add_temp()
            generator.add_expression(param_temp, 'P', env.size, '+')
            generator.add_expression(param_temp, param_temp, '1', '+')
            generator.set_stack(param_temp, left_value.value)
            generator.add_expression(param_temp, param_temp, '1', '+')
            generator.set_stack(param_temp, right_value.value)
            generator.new_env(env.size)
            generator.call_fun('potencia')
            temp = generator.add_temp()
            generator.get_stack(temp, 'P')
            generator.ret_env(env.size)
            return Return(temp, Type.INT, True)
        else:
            generator.add_expression(
                temp, left_value.value, right_value.value, op)
            return Return(temp, Type.INT, True)
