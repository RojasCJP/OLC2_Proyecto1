from ..abstracts.instruccion import *


class Print(Instruction):

    def __init__(self, value, line, column, newLine=False):
        Instruction.__init__(self, line, column)
        self.value = value
        self.newLine = newLine

    def execute(self, env: Environment):
        if self.newLine:
            for value in self.value:
                val = value.execute(env)
                print(val.value)
        else:
            for value in self.value:
                val = value.execute(env)
                print(val.value, end="")
        # todo tengo que mandarlo a la cola de impresion porque no se va a imprimir en consola
