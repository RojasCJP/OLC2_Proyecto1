from .abstracts.instruccion import *
from .abstracts.returns import *
from ..enviroment.environment import *


class Statement(Instruction):
    def __init__(self, instructions, line, column):
        Instruction.__init__(self, line, column)
        self.instructions = instructions

    def execute(self, env: Environment):
        newEnv = Environment(env)
        for ins in self.instructions:
            res = ins.execute(newEnv)
            if res is not None:
                return res
