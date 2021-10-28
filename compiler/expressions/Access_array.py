from os import truncate
from sym.Generator import *
from abstract.Expression import *
from abstract.Return import *


class AccessArray(Expression):
    def __init__(self, id, indexs, line, column):
        Expression.__init__(self, line, column)
        self.id = id
        self.indexs = indexs

    def compile(self, env):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        generator.add_comment('compilacion de acceso arreglos')
        array = env.get_var(self.id)
        if array == None:
            print("error no existe el arreglo")
            return
        temp = generator.add_temp()
        temp_pos = array.pos
        if not array.is_global:
            temp_pos = generator.add_temp()
            generator.add_expression(temp_pos, 'P', array.pos, '+')
        generator.get_stack(temp, temp_pos)
        tipo = Type.FLOAT
        for element in self.indexs:
            elemento = element.compile(env)
            sumado = generator.add_temp()
            generator.add_expression(sumado, elemento.value, temp, '+')
            generator.get_heap(temp, sumado)
            if(Generator.dict_temp[temp] % 1 != 0):
                if(Generator.heap[int(Generator.dict_temp[temp])] == 0):
                    tipo = Type.STRING
                else:
                    tipo = Type.ARRAY
        # print(0.123123 % 1)
        # print(Generator.heap)
        # print(Generator.stack)
        # print(Generator.dict_temp)
        return Return(temp, tipo, True)
