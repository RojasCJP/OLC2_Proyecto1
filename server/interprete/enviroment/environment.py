from .sym import *
from ..comandos.abstracts.returns import *


class Environment:
    def __init__(self, prev_env):
        self.variables = {}
        self.functions = {}
        self.structs = {}
        self.prev = prev_env

    def save_var(self, id_var, value, types):
        # todo tengo que cambiar el tipo de variable si es que cambia y verificar si es estricto, en ese caso no cambiarla
        env = self
        if isinstance(value, int) or isinstance(value, bool) or isinstance(value, str) or isinstance(value, float):
            new_sym = Sym(value, id_var, types)
        elif isinstance(value, Return):
            new_sym = Sym(value.value, id_var, types)
        else:
            value_value = value.execute(env)
            new_sym = Sym(value_value.value, id_var, value_value.type)

        while env is not None:
            if id_var in env.variables.keys():
                env.variables[id_var] = new_sym
                return
            env = env.prev
        self.variables[id_var] = new_sym

    def save_var_struct(self, id_var, attributes, types):
        env = self
        new_sym = Sym(None, id_var, Type.STRUCT, types)
        new_sym = attributes
        while env is not None:
            if id_var in env.variables.keys():
                env.variables[id_var] = new_sym
                return
            env = env.prev
        self.variables[id_var] = new_sym

    def save_func(self, id_func, function):
        if id_func in self.functions.keys():
            print("Función repetida")
        else:
            self.functions[id_func] = function

    def save_struct(self, id_struct, attr):
        if id_struct in self.structs.keys():
            print("Struct repetido")
        else:
            self.structs[id_struct] = attr

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
            end = end.prev
        return None

    def get_global_env(self):
        env = self
        while env.prev is not None:
            env = env.prev
        return env
