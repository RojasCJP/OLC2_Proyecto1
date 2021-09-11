from ..abstracts.instruccion import *
from ..abstracts.returns import *


class While(Instruction):
    def __init__(self, condition, instructions, line, column):
        Instruction.__init__(self, line, column)
        self.condition = condition
        self.instruction = instructions

    def execute(self, env: Environment):
        condition = self.condition.execute(env)
        if condition.type != Type.BOOL:
            print("error en la condicion, no es un valor booleano")
            return
        while condition.value:
            element = self.instruction.execute(env)
            if element is not None:
                if element.type == Type.BREAK_ST:
                    break
                elif element.type == Type.CONTINUE_ST:
                    continue
                else:
                    return element
            condition = self.condition.execute(env)
            if condition.type != Type.BOOL:
                print("error en la condicion, no es un valor booleano")
                return