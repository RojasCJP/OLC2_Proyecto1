from ..abstracts.instruccion import *
from ..abstracts.returns import *


class CreateStruct(Instruction):

    def __init__(self, idd, attributes, line, column):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.attributes = attributes

    def execute(self, env: Environment):
        env.save_struct(self.id, self.attributes)
