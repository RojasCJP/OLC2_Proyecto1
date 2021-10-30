from sym.Generator import *
from abstract.Expression import *
from abstract.Return import *


class ChangeArray(Expression):
    def __init__(self, id, indexs, value, line, column):
        Expression.__init__(self, line, column)
        self.id = id
        self.indexs = indexs
        self.value = value

    def compile(self, env):
        gen_aux = Generator()
        generator = gen_aux.get_instance()
        val = self.value.compile(env)
        generator.add_comment('cambiando el valor de arreglo')
        array = env.get_var(self.id)
        if array == None:
            print("error no existe el arreglo")
        temp = generator.add_temp()
        temp_pos = array.pos
        if not array.is_global:
            temp_pos = generator.add_temp()
            generator.add_expression(temp_pos, 'P', array.pos, '+')
        generator.get_stack(temp, temp_pos)
        sumado_aux = generator.add_temp()
        for element in self.indexs:
            elemento = element.compile(env)
            sumado = generator.add_temp()
            generator.add_expression(sumado, elemento.value, temp, '+')
            generator.get_heap(temp, sumado)
            generator.add_expression(sumado_aux, sumado, '', '')
        generator.set_heap(sumado_aux, val.value)
