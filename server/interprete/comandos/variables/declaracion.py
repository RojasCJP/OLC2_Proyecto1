from ..abstracts.instruccion import *
from ..abstracts.returns import *
from ..expressions.literal import *


class Declaration(Instruction):
    def __init__(self, idd, value, line, column, globall, types=None):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.value = value
        self.globall = globall
        self.type = types

    def execute(self, environment):
        value = self.value
        if value is None:
            value = Literal(0, Type.NULL, self.line, self.column)
        while not isinstance(value, Return):
            value = value.execute(environment)
        # todo aqui tengo que ver que pedo con los tipos
        # todo esto lo tengo que ver despues
        if value.type == Type.STRUCT:
            environment.save_var_struct(self.id, value.value, value.auxType)
        else:
            if self.type is None:
                if self.globall:
                    environment.get_global_env().save_var(self.id, value.value, value.type)
                else:
                    environment.save_var(self.id, value.value, value.type)
            else:
                if value.type == self.type:
                    if self.globall:
                        environment.get_global_env().save_var(self.id, value.value, value.type)
                    else:
                        environment.save_var(self.id, value.value, value.type)
                else:
                    print("el tipo de datos no es el mismo que el dato ingresado no se puede asignar el valor")
