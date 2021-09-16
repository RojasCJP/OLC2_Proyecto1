from ..abstracts.instruccion import *


class DeclareStruct(Instruction):

    def __init__(self, idd, types, line, column):
        Instruction.__init__(self, line, column)
        self.id = idd
        self.type = types

    def execute(self, env: Environment):
        struct = env.get_struct(self.type)
        if struct is None:
            print("no existe el type")
            return
        attributes = {}
        for attribute in struct:
            attributes.update({
                attribute: 0
            })
        env.save_var_struct(self.id, attributes, self.type)
