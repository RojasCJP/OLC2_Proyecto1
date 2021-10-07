from abstract.Expression import *
from abstract.Return import *
from symbol.Generator import Generator
import uuid


class Literal(Expression):
    def __init__(self, value, type, line, column):
        Expression.__init__(self, line, column)
        self.value = value
        self.type = type

    def compile(self, env):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        if (self.type == Type.INT or self.type == Type.FLOAT):
            return Return(str(self.value), self.type, False)
        elif self.type == Type.BOOL:
            self.check_labels()
            if self.value:
                generator.add_goto(self.true_lbl)
                generator.add_comment("goto para evitar errores")
                generator.add_goto(self.false_lbl)
            else:
                generator.add_goto(self.false_lbl)
                generator.add_comment("goto para evitar errores")
                generator.add_goto(self.true_lbl)
            ret = Return(self.value, self.type, False)
            ret.true_lbl = self.true_lbl
            ret.false_lbl = self.false_lbl
            return ret
        elif self.type == Type.STRING:
            ret_temp = generator.add_temp()
            generator.add_expression(ret_temp, 'H', '', '')
            for char in str(self.value):
                generator.set_heap('H', ord(char))
                generator.next_heap()

            generator.set_heap('H', '-1')
            generator.next_heap()
            return Return(ret_temp, Type.STRING, True)
        else:
            print("falta hacer")
            # TODO estas son las de arreglos y structs supongo

    def check_labels(self):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        if self.true_lbl == '':
            self.true_lbl = generator.new_label()
        if self.false_lbl == '':
            self.false_lbl = generator.new_label()
