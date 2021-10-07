from abstract.Expression import *
from abstract.Return import *
from symbol.Generator import *
from enum import Enum


class RelationalOption(Enum):
    GREATER = 0
    LESS = 1
    GREATEREQUAL = 2
    LESSEQUAL = 3
    EQUAL = 4
    DISTINCT = 5


class Relational(Expression):

    def __init__(self, left, right, type, line, column):
        Expression.__init__(line, column)
        self.left = left
        self.right = right
        self.type = type

    def compile(self, env):
        gen_aux = Generator()
        generator = gen_aux.get_instance()

        generator.add_comment("inicio de expression realcional")
        left = self.left.compile(env)
        right = None
        result = Return(None, Type.BOOL, False)

        if left.type != Type.BOOL:
            right = self.right.compile(env)
            if (left.type == Type.INT or left.type == Type.FLOAT) and (right.type == Type.INT or right.type == Type.FLOAT):
                self.check_labels()
                generator.add_if(left.value, right.value,
                                 self.get_operation(), self.true_lbl)
                generator.add_goto(self.false_lbl)
            elif left.type == Type.STRING and right.type == Type.STRING:
                print('comparacion de cadenas')
                # TODO tengo que hacer la comparacion de cadenas
        else:
            goto_right = generator.new_label()
            left_temp = generator.add_temp()
            generator.put_label(left.true_lbl)
            generator.add_expression(left_temp, '1', '', '')
            generator.add_goto(goto_right)
            generator.put_label(left.false_lbl)
            generator.add_expression(left_temp, '0', '', '')
            generator.put_label(goto_right)
            right = self.right.compile(env)
            if right.type != Type.BOOL:
                print("error, no se pueden comparar")
                return
                # TODO tengo que pasarlo a la printlist
            goto_end = generator.new_label()
            right_temp = generator.add_temp()
            generator.put_label(right.true_lbl)
            generator.add_expression(right_temp, '1', '', '')
            generator.add_goto(goto_end)
            generator.put_label(right.false_lbl)
            generator.add_expression(right_temp, '0', '', '')
            generator.put_label(goto_end)
            self.check_labels()
            generator.add_if(left_temp, right_temp,
                             self.get_operation(), self.true_lbl)
            generator.add_goto(self.false_lbl)
        generator.add_comment("fin de la expression relacional")
        generator.add_space()
        result.true_lbl = self.true_lbl
        result.false_lbl = self.false_lbl
        return result

    def check_labels(self):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        if self.true_lbl == '':
            self.true_lbl = generator.new_label()
        if self.false_lbl == '':
            self.false_lbl = generator.new_label()

    def get_operation(self):
        if self.type == RelationalOption.GREATER:
            return '>'
        elif self.type == RelationalOption.LESS:
            return '<'
        elif self.type == RelationalOption.GREATEREQUAL:
            return '>='
        elif self.type == RelationalOption.LESSEQUAL:
            return '<='
        elif self.type == RelationalOption.EQUAL:
            return '=='
        elif self.type == RelationalOption.DISTINCT:
            return '!='
