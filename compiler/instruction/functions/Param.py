from abstract.Instruction import *


class Param(Instruction):

    def __init__(self, id, line, column):
        Instruction.__init__(self, line, column)
        self.id = id
