from ..abstracts.instruccion import *


class Function(Instruction):
    def __init__(self, idd, params, instructions, line, column):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.params = params
        self.instructions = instructions

    def execute(self, env: Environment):
        try:
            env.save_func(self.id, self)
        except Exception as e:
            print("no se pudo guardar la funcion")
            print(e)
