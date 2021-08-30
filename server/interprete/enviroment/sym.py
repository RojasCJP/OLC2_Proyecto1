from enum import Enum, auto


class Sym:
    def __init__(self, value, id_sym, types, symbol_obj=""):
        self.type = types
        self.value = value
        self.id = id_sym
        self.obj = symbol_obj
        self.attributes = {}
