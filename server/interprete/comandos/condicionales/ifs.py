from ..abstracts.instruccion import *
from ..abstracts.returns import *


class If(Instruction):
    def __init__(self, condition, instructions, line, column, else_state=None):
        Instruction.__init__(self, line, column)
        self.condition = condition
        self.instructions = instructions
        self.else_state = else_state

    def execute(self, env: Environment):
        condition = self.condition.execute(env)
        if condition.type != Type.BOOL:
            print("la condicion no cumple con se boolean")
            return
        if condition.value:
            return self.instructions.execute(env)
        elif self.else_state is not None:
            return self.else_state.execute(env)
