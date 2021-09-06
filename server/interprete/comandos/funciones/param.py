from ..abstracts.instruccion import *


class Param(Instruction):
    def __init__(self, idd, line, column):
        Instruction.__init__(self, line, column)
        self.id = idd

    def execute(self, env: Environment):
        return self
