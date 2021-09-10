from ..abstracts.instruccion import *
from ..abstracts.returns import *

class Continue(Instruction):

    def __init__(self, line, column):
        Instruction.__init__(self, line,column)

    def execute(self, env: Environment):
        return Return(None, Type.CONTINUE_ST)