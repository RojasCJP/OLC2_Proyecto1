from ..abstracts.instruccion import *


class Print(Instruction):
    printlist = ""

    def __init__(self, value, line, column, newLine=False):
        Instruction.__init__(self, line, column)
        self.value = value
        self.newLine = newLine

    def execute(self, env: Environment):
        if self.newLine:
            for value in self.value:
                val = value.execute(env)
                Print.printlist += str(val.value)
                print(val.value)
            Print.printlist += '\n'
        else:
            for value in self.value:
                val = value.execute(env)
                print(val.value, end="")
                Print.printlist += str(val.value)
        # todo tengo que mandarlo a la cola de impresion porque no se va a imprimir en consola
