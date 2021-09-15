from ..abstracts.expression import *
from ..abstracts.returns import *
from ..expressions.access import *
from ..expressions.arithmetic import *


class AccessList(Expression):

    def __init__(self, idd, index, line, column):
        Expression.__init__(self, line, column)
        self.id = idd
        self.index = index

    def execute(self, environment: Environment):
        value = environment.get_var(self.id)
        valor_retorno = value.value
        if value is None:
            print("error la variable que busca no existe")
            return
        else:
            for element in self.index:
                elemento = element
                if isinstance(element, Access) or isinstance(element, Arithmetic):
                    elemento = element.execute(environment)
                    elemento = elemento.value
                else:
                    elemento = elemento.value
                valor_retorno = valor_retorno[elemento-1]
            return Return(valor_retorno, value.type)
