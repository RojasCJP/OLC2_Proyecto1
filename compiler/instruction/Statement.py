from abstract.Instruction import *


class Statement(Instruction):
    def __init__(self, instructions, line, column):
        Instruction.__init__(self, line, column)
        self.instructions = instructions
