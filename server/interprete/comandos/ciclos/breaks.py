from ..abstracts.instruccion import *
from ..abstracts.returns import *


class Break(Instruction):

    def __init__(self, line, column):
        Instruction.__init__(self, line, column)

    def execute(self, env: Environment):
        return Return(None, Type.BREAK_ST)