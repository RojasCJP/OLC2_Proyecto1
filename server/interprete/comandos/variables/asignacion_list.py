from ..abstracts.expression import *
from ..abstracts.returns import *
from ..expressions.literal import *
from ..expressions.access import *
from ..expressions.arithmetic import *


class AssignationList(Expression):

    def __init__(self, idd, index, value, line, column):
        Expression.__init__(self, line, column)
        self.id = idd
        self.index = index
        self.value = value

    def execute(self, environment):
        value = environment.get_var(self.id)
        result = Literal(0, Type.UNDEFINED, self.line, self.column)
        valor_retorno = value.value
        if value is None:
            print("error la variable que busca no existe")
            return
        else:
            if len(self.index) == 1:
                valor0 = self.return_value(self.index[0], environment)-1
                valor_retorno[valor0] = self.get_value(self.value, environment)
            elif len(self.index) == 2:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor_retorno[valor0][valor1] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 3:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor_retorno[valor0][valor1][valor2] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 4:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor3 = self.return_value(self.index[3], environment)-1
                valor_retorno[valor0][valor1][valor2][valor3] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 5:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor3 = self.return_value(self.index[3], environment)-1
                valor4 = self.return_value(self.index[4], environment)-1
                valor_retorno[valor0][valor1][valor2][valor3][valor4] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 6:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor3 = self.return_value(self.index[3], environment)-1
                valor4 = self.return_value(self.index[4], environment)-1
                valor5 = self.return_value(self.index[5], environment)-1
                valor_retorno[valor0][valor1][valor2][valor3][valor4][valor5] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 7:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor3 = self.return_value(self.index[3], environment)-1
                valor4 = self.return_value(self.index[4], environment)-1
                valor5 = self.return_value(self.index[5], environment)-1
                valor6 = self.return_value(self.index[6], environment)-1
                valor_retorno[valor0][valor1][valor2][valor3][valor4][valor5][valor6] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 8:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor3 = self.return_value(self.index[3], environment)-1
                valor4 = self.return_value(self.index[4], environment)-1
                valor5 = self.return_value(self.index[5], environment)-1
                valor6 = self.return_value(self.index[6], environment)-1
                valor7 = self.return_value(self.index[7], environment)-1
                valor_retorno[valor0][valor1][valor2][valor3][valor4][valor5][valor6][valor7] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 9:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor3 = self.return_value(self.index[3], environment)-1
                valor4 = self.return_value(self.index[4], environment)-1
                valor5 = self.return_value(self.index[5], environment)-1
                valor6 = self.return_value(self.index[6], environment)-1
                valor7 = self.return_value(self.index[7], environment)-1
                valor8 = self.return_value(self.index[8], environment)-1
                valor_retorno[valor0][valor1][valor2][valor3][valor4][valor5][valor6][valor7][valor8] = self.get_value(
                    self.value, environment)
            elif len(self.index) == 10:
                valor0 = self.return_value(self.index[0], environment)-1
                valor1 = self.return_value(self.index[1], environment)-1
                valor2 = self.return_value(self.index[2], environment)-1
                valor3 = self.return_value(self.index[3], environment)-1
                valor4 = self.return_value(self.index[4], environment)-1
                valor5 = self.return_value(self.index[5], environment)-1
                valor6 = self.return_value(self.index[6], environment)-1
                valor7 = self.return_value(self.index[7], environment)-1
                valor8 = self.return_value(self.index[8], environment)-1
                valor9 = self.return_value(self.index[9], environment)-1
                valor_retorno[valor0][valor1][valor2][valor3][valor4][valor5][valor6][valor7][valor8][valor9] = self.get_value(
                    self.value, environment)
        result.value = valor_retorno
        return None

    def return_value(self, element, environment):
        elemento = element
        if isinstance(element, Access) or isinstance(element, Arithmetic):
            elemento = element.execute(environment)
            elemento = elemento.value
        else:
            elemento = elemento.value
        return elemento

    def get_value(self, valor, env):
        array_return = []
        if isinstance(valor, list):
            for element in valor:
                if isinstance(element, int) or isinstance(element, bool) or isinstance(element, str) or isinstance(element, float):
                    array_return.append(element)
                elif isinstance(element, Return):
                    array_return.append(element.value)
                elif isinstance(element, list):
                    array_return.append(self.get_value(element))
                else:
                    element_value = element.execute(env)
                    array_return.append(element_value.value)
        else:
            try:
                if valor.value is not None:
                    if isinstance(valor.value, list):
                        array_return = self.get_value(valor.value, env)
                        return array_return
            except AttributeError:
                pass
            if isinstance(valor, int) or isinstance(valor, bool) or isinstance(valor, str) or isinstance(valor, float):
                array_return = valor
            elif isinstance(valor, Return):
                array_return = valor.value
            else:
                element_value = valor.execute(env)
                array_return = element_value.value
        return array_return
