from ..abstracts.instruccion import *


class Print(Instruction):

    def __init__(self, value, line, column, newLine=False):
        Instruction.__init__(self, line, column)
        self.value = value
        self.newLine = newLine

    def execute(self, env: Environment):
        value = self.value.execute(env)
        if self.newLine:
            print(value.value)
        else:
            print(value.value, end="")
        # todo tengo que mandarlo a la cola de impresion porque no se va a imprimir en consola
