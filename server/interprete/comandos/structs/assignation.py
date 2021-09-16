from ..abstracts.instruccion import *


class AssignAccess(Instruction):

    def __init__(self, idd, access, expression, line, column):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.access = access
        self.expression = expression

    def execute(self, env: Environment):
        value = self.expression.execute(env)
        symbol = env.get_var(self.id)
        symbol.value[self.access] = value.value
