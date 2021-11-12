from abstract.Instruction import *
from abstract.Return import *
from sym.Generator import *


class Trunc(Instruction):
    def __init__(self, line, column, value):
        Instruction.__init__(self, line, column)
        self.value = value

    def compile(self, env):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        value = self.value.compile(env)
        temp = generator.add_temp()
        generator.add_trunc(temp, value.value)
        return Return(temp, Type.INT, True)
