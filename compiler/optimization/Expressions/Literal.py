from optimization.Instruction import *


class Literal(Instruction):

    def __init__(self, value, line, column, constant=False):
        Instruction.__init__(self, line, column)
        self.value = value
        self.constant = constant

    def getCode(self):
        return str(self.value)
