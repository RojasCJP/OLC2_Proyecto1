from .sym import *
from ..comandos.abstracts.returns import *


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

    def save_var(self, id_var, value, types):
        env = self
        if isinstance(value, int) or isinstance(value, bool) or isinstance(value, str) or isinstance(value, float) or isinstance(value, dict):
            new_sym = Sym(value, id_var, types)
        elif isinstance(value, list):
            arreglo = self.get_items_array(value)
            new_sym = Sym(arreglo, id_var, types)
        elif isinstance(value, Return):
            new_sym = Sym(value.value, id_var, types)
        else:
            value_value = value.execute(env)
            new_sym = Sym(value_value.value, id_var, value_value.type)

        while env is not None:
            if id_var in env.variables.keys():
                env.variables[id_var] = new_sym
                Environment.variables[id_var] = new_sym
                return
            env = env.prev
        self.variables[id_var] = new_sym
        Environment.variables[id_var] = new_sym

    def save_var_struct(self, id_var, attributes, types):
        env = self
        new_sym = Sym(None, id_var, Type.STRUCT, types)
        new_sym.value = attributes
        while env is not None:
            if id_var in env.variables.keys():
                env.variables[id_var] = new_sym
                Environment.variables[id_var] = new_sym
                return
            env = env.prev
        self.variables[id_var] = new_sym
        Environment.variables[id_var] = new_sym

    def save_func(self, id_func, function):
        if id_func in self.functions.keys():
            print("Función repetida")
        else:
            self.functions[id_func] = function
            Environment.functions[id_func] = function

    def save_struct(self, id_struct, attr):
        if id_struct in self.structs.keys():
            print("Struct repetido")
        else:
            self.structs[id_struct] = attr
            Environment.structs[id_struct] = attr

    def get_var(self, id_var):
        env = self
        while env is not None:
            if id_var in env.variables.keys():
                return env.variables[id_var]
            env = env.prev
        return None

    def get_function(self, id_func):
        env = self
        while env is not None:
            if id_func in env.functions.keys():
                return env.functions[id_func]
            env = env.prev
        return None

    def get_struct(self, id_struct):
        env = self
        while env is not None:
            if id_struct in env.structs.keys():
                return env.structs[id_struct]
            env = env.prev
        return None

    def get_global_env(self):
        env = self
        while env.prev is not None:
            env = env.prev
        return env
