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
        elif self.type == ArithmethicOption.MODULE:
            op = '%'

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
            if(left_value.type == Type.FLOAT or right_value.type == Type.FLOAT or self.type == ArithmethicOption.DIV):
                generator.add_expression(
                    temp, left_value.value, right_value.value, op)
                return Return(temp, Type.FLOAT, True)
            elif (left_value.type == Type.STRING and right_value.type == Type.STRING):
                # TODO tengo que jalar los 2 valores.value y esos son los punteros de mis string
                left_temp = generator.add_temp()
                right_temp = generator.add_temp()
                ret_temp = generator.add_temp()
                generator.add_expression(ret_temp, 'H', '', '')

                generator.add_expression(left_temp, left_value.value, '', '')
                generator.add_expression(right_temp, right_value.value, '', '')

                left_label = generator.new_label()
                right_label = generator.new_label()
                left_swaper = generator.add_temp()
                right_swaper = generator.add_temp()
                generator.get_heap(left_swaper, left_temp)
                generator.get_heap(right_swaper, right_temp)

                generator.put_label(left_label)
                generator.set_heap('H', left_swaper)
                generator.next_heap()
                generator.add_expression(left_temp, left_temp, '1', '+')
                generator.get_heap(left_swaper, left_temp)
                generator.add_if(left_swaper, '-1', '!=', left_label)

                generator.put_label(right_label)
                generator.set_heap('H', right_swaper)
                generator.next_heap()
                generator.add_expression(right_temp, right_temp, '1', '+')
                generator.get_heap(right_swaper, right_temp)
                generator.add_if(right_swaper, '-1', '!=', right_label)
                generator.set_heap('H', '-1')
                generator.next_heap()

                return Return(ret_temp, Type.STRING, True)
            else:
                generator.add_expression(
                    temp, left_value.value, right_value.value, op)
                return Return(temp, Type.INT, True)
