from .Symbol import *
from ..abstract.Return import *


class Environment:
    variables = {}
    functions = {}
    structs = {}
    errores = []
    entrada = ""

    def __init__(self, prev_env):
        self.variables = {}
        self.functions = {}
        self.structs = {}
        self.prev = prev_env
        self.size = 0
        if prev_env is not None:
            self.size = self.prev.size
# TODO este metodo lo voy a tener que cambiar fijo
# TODO tambien voy a tener que poner cosas en los accesos de seguro

    def get_items_array(self, array):
        env = self
        array_return = []
        for element in array:
            if isinstance(element, int) or isinstance(element, bool) or isinstance(element, str) or isinstance(element, float):
                array_return.append(element)
            elif isinstance(element, Return):
                array_return.append(element.value)
            elif isinstance(element, list):
                array_return.append(self.get_items_array(element))
            else:
                element_value = element.execute(env)
                if element_value.type != Type.ARRAY:
                    array_return.append(element_value.value)
                else:
                    array_return.append(
                        self.get_items_array(element_value.value))
        return array_return

    def save_var(self, id_var, sym_type, in_heap):
        if id_var in self.variables.keys():
            print("Variable ya existe")
        else:
            newSymbol = Symbol(id_var, sym_type, self.size,
                               self.prev == None, in_heap)
            self.size += 1
            self.variables[id_var] = newSymbol
        return self.variables[id_var]

    def save_func(self, id_func, function):
        if id_func in self.functions.keys():
            print("Función repetida")
        else:
            self.functions[id_func] = function

    def save_struct(self, id_struct, attributes):
        if id_struct in self.structs.keys():
            print("Struct repetido")
        else:
            self.structs[id_struct] = attributes

    def get_var(self, id_var):
        env = self
        while env != None:
            if id_var in env.variables.keys():
                return env.variables[id_var]
            env = env.prev
        return None

    def get_func(self, id_func):
        if id_func in self.functions.keys():
            return self.functions[id_func]
        else:
            return None

    def get_struct(self, id_struct):
        env = self
        while env != None:
            if id_struct in env.structs.keys():
                return env.structs[id_struct]
            end = end.prev
        return None

    def get_global(self):
        env = self
        while env.prev != None:
            env = env.prev
        return env
