from ..abstracts.expression import *
from ..abstracts.returns import *


class Access(Expression):

    def __init__(self, idd, line, column):
        Expression.__init__(self, line, column)
        self.id = idd

    def execute(self, environment):
        value = environment.get_var(self.id)
        if value is None:
            print("error la variable que busca no existe")
            return
        else:
            if isinstance(value, dict):
                return Return(value, type(value))
            return Return(value.value, value.type)
