from ..abstracts.instruccion import *
from ..abstracts.returns import *


class For(Instruction):
    def __init__(self, variable, value1, instructions, line, column, value2=None):
        Instruction.__init__(self, line, column)
        self.variable = variable
        self.value1 = value1
        self.value2 = value2
        self.instructions = instructions

    def execute(self, env: Environment):
        val_left = self.value1.execute(env)
        if self.value2 is not None:
            val_right = self.value2.execute(env)
            env.save_var(self.variable, val_left.value, val_left.type)
            var = env.get_var(self.variable)
            if var is not None:
                while var.value <= val_right.value:
                    execution = self.instructions.execute(env)
                    if execution is not None:
                        return execution
                    var.value += 1
                    env.save_var(self.variable, var.value, var.type)
        else:
            num = 0
            env.save_var(self.variable, val_left.value[num], val_left.type)
            var = env.get_var(self.variable)
            tamano = len(val_left.value)
            if var is not None:
                while num < tamano:
                    element = self.instructions.execute(env)
                    if element is not None:
                        if element.type == Type.BREAK_ST:
                            break
                        elif element.type == Type.CONTINUE_ST:
                            continue
                        else:
                            return element
                    if len(val_left.value) > num + 1:
                        var.value = val_left.value[num]
                        env.save_var(self.variable, val_left.value[num + 1], val_left.type)
                    num += 1
